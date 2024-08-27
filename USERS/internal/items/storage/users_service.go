package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ruziba3vich/hotello-users/genprotos/users"
	"github.com/ruziba3vich/hotello-users/internal/items/models"
	"github.com/ruziba3vich/hotello-users/internal/items/redisservice"
	"github.com/ruziba3vich/hotello-users/internal/pkg/config"
	"github.com/ruziba3vich/hotello-users/internal/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	DB struct {
		Client          *mongo.Client
		UsersCollection *mongo.Collection
	}

	UsersStorage struct {
		redisService   *redisservice.RedisService
		database       *DB
		logger         *log.Logger
		tokenGenerator *utils.TokenGenerator
		pwdHashsher    *utils.PasswordHasher
		config         config.Config
	}
)

func New(redisService *redisservice.RedisService, database *DB, logger *log.Logger, tokenGenerator *utils.TokenGenerator, pwdHashsher *utils.PasswordHasher, config config.Config) *UsersStorage {
	return &UsersStorage{
		redisService:   redisService,
		database:       database,
		logger:         logger,
		tokenGenerator: tokenGenerator,
		pwdHashsher:    pwdHashsher,
		config:         config,
	}
}

func (s *UsersStorage) RegisterUser(ctx context.Context, req *users.CreateUserRequest) (*users.User, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	hashshedPwd, err := s.pwdHashsher.HashPassword(req.Password)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	req.Password = hashshedPwd

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		var user models.User
		user.FromCreateUserRequest(req)

		_, err := s.database.UsersCollection.InsertOne(sessCtx, user)
		if err != nil {
			s.logger.Println("Error inserting user into MongoDB:", err)
			return nil, fmt.Errorf("failed to insert user into MongoDB: %s", err.Error())
		}

		if err := s.redisService.StoreUserInRedis(ctx, user.ToProto()); err != nil {
			s.logger.Println("Error storing user in Redis:", err)
			return nil, fmt.Errorf("failed to store user in Redis: %s", err.Error())
		}

		return user.ToProto(), nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	result, err := session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	user, ok := result.(*users.User)
	if !ok {
		s.logger.Println("Unexpected result from transaction callback")
		return nil, fmt.Errorf("unexpected result from transaction callback")
	}

	return user, nil
}

func (s *UsersStorage) LoginUser(ctx context.Context, req *users.LoginUserRequest) (*users.RawResponse, error) {
	user, err := s.GetUserByEmail(ctx, &users.GetUserByFieldRequest{
		FieldName: "email",
		Value:     req.Email,
	})
	if err != nil {
		return nil, err
	}
	if !s.pwdHashsher.CheckPasswordHash(req.Password, user.Password) {
		s.logger.Println("-- Missmatch in password --")
		return nil, errors.New("password missmatch")
	}
	if err := s.sendVerificationCode(ctx, req.Email); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &users.RawResponse{
		Message: fmt.Sprintf("Verification code has been sent to %s****%s", req.Email[:3], req.Email[len(req.Email)-5:]),
	}, nil
}

func (s *UsersStorage) GetUserById(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	user, err := s.redisService.GetUserFromRedis(ctx, req.Value)
	if err == nil {
		return user, nil
	}

	objectID, err := primitive.ObjectIDFromHex(req.Value)
	if err != nil {
		s.logger.Println("Invalid ObjectID format:", err)
		return nil, err
	}

	var mongoUser models.User
	filter := bson.M{"_id": objectID}

	if err := s.database.UsersCollection.FindOne(ctx, filter).Decode(&mongoUser); err != nil {
		s.logger.Println("Error finding user in MongoDB:", err)
		return nil, err
	}

	if err := s.redisService.StoreUserInRedis(ctx, mongoUser.ToProto()); err != nil {
		s.logger.Println("Error storing user in Redis:", err)
		return nil, err
	}

	return mongoUser.ToProto(), nil
}

func (s *UsersStorage) GetUserByUsername(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	return s.getUserByField(ctx, req)
}

func (s *UsersStorage) GetUserByEmail(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	return s.getUserByField(ctx, req)
}

func (s *UsersStorage) UpdatePasswordService(ctx context.Context, req *users.UpdatePasswordRequest) (*users.RawResponse, error) {
	user, err := s.GetUserById(ctx, &users.GetUserByFieldRequest{
		Value: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if !s.pwdHashsher.CheckPasswordHash(req.PreviousPassword, user.Password) {
		return nil, errors.New("password missmatch")
	}
	hashedPwd, err := s.pwdHashsher.HashPassword(req.NewPassword)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return s.updateUserByField(ctx, &models.UpdateUserRequest{FieldName: "password", Value: hashedPwd}, user)
}

func (s *UsersStorage) UpdateUsername(ctx context.Context, req *users.UpdateUsernameRequest) (*users.RawResponse, error) {
	user, err := s.GetUserByUsername(ctx, &users.GetUserByFieldRequest{
		FieldName: "username",
		Value:     req.NewUsername,
	})
	if err == nil {
		return nil, fmt.Errorf("%s has already been taken", req.NewUsername)
	}
	return s.updateUserByField(ctx, &models.UpdateUserRequest{FieldName: "username", Value: req.NewUsername}, user)
}

func (s *UsersStorage) SoftDeleteUser(ctx context.Context, req *users.DeleteUserRequest) (*users.RawResponse, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		update := bson.M{"$set": bson.M{"deleted": true}}

		_, err := s.database.UsersCollection.UpdateOne(sessCtx, bson.M{"_id": req.Id}, update)
		if err != nil {
			s.logger.Println("Error updating user in MongoDB:", err)
			return nil, fmt.Errorf("failed to update user in MongoDB: %s", err.Error())
		}

		if err := s.redisService.DeleteUserFromRedis(sessCtx, req.Id); err != nil {
			s.logger.Println("Error deleting user from Redis:", err)
			return nil, fmt.Errorf("failed to delete user from Redis: %s", err.Error())
		}

		return nil, nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	_, err = session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	return &users.RawResponse{
		Message: "user has successfully been deleted",
	}, nil
}

func (s *UsersStorage) VerifyCodeFromEmail(ctx context.Context, req *users.VerifyCodeRequest) (*users.LoginUserResponse, error) {
	if err := s.verifyEmail(ctx, req.Email, int(req.Code)); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	user, err := s.GetUserByEmail(ctx, &users.GetUserByFieldRequest{
		FieldName: "email",
		Value:     req.Email,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.tokenGenerator.GenerateToken(user.Id)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return &users.LoginUserResponse{
		Username:  user.Username,
		Token:     token,
		ExpiresIn: 24 * 60 * 60,
	}, nil
}

func (s *UsersStorage) getUserByField(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	filter := bson.M{req.FieldName: req.Value}

	var mongoUser models.User

	if err := s.database.UsersCollection.FindOne(ctx, filter).Decode(&mongoUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found: %s", err.Error())
		}
		s.logger.Println("Error finding user in MongoDB:", err)
		return nil, err
	}

	return mongoUser.ToProto(), nil
}

func (s *UsersStorage) updateUserByField(ctx context.Context, req *models.UpdateUserRequest, user *users.User) (*users.RawResponse, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err.Error())
		return nil, err
	}
	defer session.EndSession(ctx)

	if req.FieldName == "username" {
		user.Username = req.Value
	} else {
		user.Password = req.Value
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		filter := bson.M{"_id": user.Id}
		update := bson.M{"$set": bson.M{req.FieldName: req.Value}}
		_, err := s.database.UsersCollection.UpdateOne(sessCtx, filter, update)
		if err != nil {
			s.logger.Println("Error updating username in MongoDB:", err)
			return nil, err
		}

		if err = s.redisService.StoreUserInRedis(ctx, user); err != nil {
			s.logger.Println("Error updating username in Redis:", err)
			return nil, err
		}

		return &users.RawResponse{
			Message: "username has successfully been updated",
		}, nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	_, err = session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	return &users.RawResponse{
		Message: fmt.Sprintf("%s has successfully been updated", req.FieldName),
	}, nil
}

// =======================================================================================================================

func ConnectDB(cfg *config.Config, ctx context.Context) (*DB, error) {
	clientOptions := options.Client().ApplyURI(cfg.DbConfig.MongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %s", err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %s", err.Error())
	}

	return &DB{
		Client:          client,
		UsersCollection: client.Database(cfg.DbConfig.MongoDB).Collection(cfg.DbConfig.Collection),
	}, nil
}

func (db *DB) DisconnectDB(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %s", err.Error())
	}
	return nil
}
