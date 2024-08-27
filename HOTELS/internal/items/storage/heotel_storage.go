// storage/hotels.go

package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/ruziba3vich/hotello-hotels/genprotos/hotels"
	"github.com/ruziba3vich/hotello-hotels/internal/items/models"
	redisservice "github.com/ruziba3vich/hotello-hotels/internal/items/redissrv"
	"github.com/ruziba3vich/hotello-hotels/internal/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	HotelsStorage struct {
		redisService *redisservice.RedisService
		database     *DB
		logger       *log.Logger
	}

	DB struct {
		Client           *mongo.Client
		HotelsCollection *mongo.Collection
		RoomsCollection  *mongo.Collection
	}
)

func NewHotelsStorage(redisService *redisservice.RedisService, database *DB, logger *log.Logger) *HotelsStorage {
	return &HotelsStorage{
		redisService: redisService,
		database:     database,
		logger:       logger,
	}
}

func (s *HotelsStorage) CreateHotel(ctx context.Context, req *hotels.CreateHotelRequest) (*hotels.Hotel, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		hotel := models.Hotel{
			Name:     req.Name,
			Location: req.Location,
			Rating:   float64(req.Rating),
			Address:  req.Address,
		}

		result, err := s.database.HotelsCollection.InsertOne(sessCtx, hotel)
		if err != nil {
			s.logger.Println("Error inserting hotel into MongoDB:", err)
			return nil, err
		}

		hotel.ID = result.InsertedID.(primitive.ObjectID)
		if err := s.redisService.StoreHotelInRedis(ctx, hotel.ToProto()); err != nil {
			s.logger.Println("Error storing hotel in Redis:", err)
			return nil, err
		}

		return hotel.ToProto(), nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	result, err := session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	hotel, ok := result.(*hotels.Hotel)
	if !ok {
		s.logger.Println("Unexpected result from transaction callback")
		return nil, fmt.Errorf("unexpected result from transaction callback")
	}

	return hotel, nil
}

func (s *HotelsStorage) CreateRoom(ctx context.Context, req *hotels.CreateRoomRequest) (*hotels.Room, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		room := models.Room{
			RoomType:     req.RoomType,
			NumberOfBeds: int(req.NumberOfBeds),
			Available:    req.Available,
		}

		hotelID, err := primitive.ObjectIDFromHex(req.HotelId)
		if err != nil {
			return nil, fmt.Errorf("invalid hotel ID format: %s", err.Error())
		}

		room.HotelID = hotelID

		_, err = s.database.RoomsCollection.InsertOne(sessCtx, room)
		if err != nil {
			s.logger.Println("Error inserting room into MongoDB:", err)
			return nil, err
		}

		if err := s.redisService.StoreRoomInRedis(ctx, room.ToProto()); err != nil {
			s.logger.Println("Error storing room in Redis:", err)
			return nil, err
		}

		return room.ToProto(), nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	result, err := session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	room, ok := result.(*hotels.Room)
	if !ok {
		s.logger.Println("Unexpected result from transaction callback")
		return nil, fmt.Errorf("unexpected result from transaction callback")
	}

	return room, nil
}

func (s *HotelsStorage) GetAvailableRoomsByHotel(ctx context.Context, req *hotels.GetAvailableRoomsByHotelRequest) (*hotels.GetAvailableRoomsByHotelResponse, error) {
	hotelID, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		s.logger.Println("Invalid hotel ID format:", err)
		return nil, err
	}

	filter := bson.M{
		"hotel_id":  hotelID,
		"available": true,
	}

	cursor, err := s.database.RoomsCollection.Find(ctx, filter)
	if err != nil {
		s.logger.Println("Error finding available rooms in MongoDB:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var rooms []models.Room
	if err := cursor.All(ctx, &rooms); err != nil {
		s.logger.Println("Error decoding rooms from MongoDB:", err)
		return nil, err
	}

	var protoRooms []*hotels.Room
	for _, r := range rooms {
		protoRooms = append(protoRooms, r.ToProto())
	}

	return &hotels.GetAvailableRoomsByHotelResponse{Rooms: protoRooms}, nil
}

func (s *HotelsStorage) GetHotelById(ctx context.Context, req *hotels.GetHotelByIdRequest) (*hotels.GetHotelByIdResponse, error) {
	hotelID, err := primitive.ObjectIDFromHex(req.HotelId)
	if err != nil {
		s.logger.Println("Invalid hotel ID format:", err)
		return nil, err
	}

	hotel, err := s.redisService.GetHotelFromRedis(ctx, req.HotelId)
	if err == nil {
		return &hotels.GetHotelByIdResponse{
			HotelId:  hotel.HotelId,
			Name:     hotel.Name,
			Location: hotel.Location,
			Rating:   hotel.Rating,
			Address:  hotel.Address,
		}, nil
	}

	filter := bson.M{"_id": hotelID}
	var mongoHotel models.Hotel
	if err := s.database.HotelsCollection.FindOne(ctx, filter).Decode(&mongoHotel); err != nil {
		s.logger.Println("Error finding hotel in MongoDB:", err)
		return nil, err
	}

	return &hotels.GetHotelByIdResponse{
		HotelId:  mongoHotel.ID.Hex(),
		Name:     mongoHotel.Name,
		Location: mongoHotel.Location,
		Rating:   float32(mongoHotel.Rating),
		Address:  mongoHotel.Address,
	}, nil
}

func (s *HotelsStorage) GetAllHotels(ctx context.Context, req *hotels.GetAllHotelsRequest) (*hotels.GetAllHotelsResponse, error) {
	page := req.Page
	limit := req.Limit

	filter := bson.M{}
	options := options.Find()
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))

	cursor, err := s.database.HotelsCollection.Find(ctx, filter, options)
	if err != nil {
		s.logger.Println("Error finding hotels in MongoDB:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var hotelsList []models.Hotel
	if err := cursor.All(ctx, &hotelsList); err != nil {
		s.logger.Println("Error decoding hotels from MongoDB:", err)
		return nil, err
	}

	var protoHotels []*hotels.GetHotelByIdResponse
	for _, h := range hotelsList {
		protoHotels = append(protoHotels, h.ToGetAllHotels())
	}

	return &hotels.GetAllHotelsResponse{Hotels: protoHotels}, nil
}

func (s *HotelsStorage) AddRoomToHotel(ctx context.Context, req *hotels.CreateRoomRequest) (*hotels.Room, error) {
	return s.CreateRoom(ctx, req)
}

func (s *HotelsStorage) SetRoomToAvailable(ctx context.Context, req *hotels.SetRoomToAvailableRequest) (*hotels.Room, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		roomID, err := primitive.ObjectIDFromHex(req.RoomId)
		if err != nil {
			return nil, fmt.Errorf("invalid room ID format: %s", err.Error())
		}

		update := bson.M{"$set": bson.M{"available": true}}
		result := s.database.RoomsCollection.FindOneAndUpdate(sessCtx, bson.M{"_id": roomID}, update)
		if result.Err() != nil {
			s.logger.Println("Error updating room availability in MongoDB:", result.Err())
			return nil, result.Err()
		}

		var updatedRoom models.Room
		if err := result.Decode(&updatedRoom); err != nil {
			s.logger.Println("Error decoding updated room:", err)
			return nil, err
		}

		if err := s.redisService.StoreRoomInRedis(ctx, updatedRoom.ToProto()); err != nil {
			s.logger.Println("Error storing room in Redis:", err)
			return nil, err
		}

		return updatedRoom.ToProto(), nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	result, err := session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	room, ok := result.(*hotels.Room)
	if !ok {
		s.logger.Println("Unexpected result from transaction callback")
		return nil, fmt.Errorf("unexpected result from transaction callback")
	}

	return room, nil
}

func (s *HotelsStorage) SetRoomToUnavailable(ctx context.Context, req *hotels.SetRoomToUnavailableRequest) (*hotels.Room, error) {
	session, err := s.database.Client.StartSession()
	if err != nil {
		s.logger.Println("Error starting MongoDB session:", err)
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		roomID, err := primitive.ObjectIDFromHex(req.RoomId)
		if err != nil {
			return nil, fmt.Errorf("invalid room ID format: %s", err.Error())
		}

		update := bson.M{"$set": bson.M{"available": false}}
		result := s.database.RoomsCollection.FindOneAndUpdate(sessCtx, bson.M{"_id": roomID}, update)
		if result.Err() != nil {
			s.logger.Println("Error updating room availability in MongoDB:", result.Err())
			return nil, result.Err()
		}

		var updatedRoom models.Room
		if err := result.Decode(&updatedRoom); err != nil {
			s.logger.Println("Error decoding updated room:", err)
			return nil, err
		}

		if err := s.redisService.StoreRoomInRedis(ctx, updatedRoom.ToProto()); err != nil {
			s.logger.Println("Error storing room in Redis:", err)
			return nil, err
		}

		return updatedRoom.ToProto(), nil
	}

	txnOpts := options.Transaction().SetWriteConcern(writeconcern.Majority())

	result, err := session.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		s.logger.Println("Transaction failed and was rolled back:", err)
		return nil, err
	}

	room, ok := result.(*hotels.Room)
	if !ok {
		s.logger.Println("Unexpected result from transaction callback")
		return nil, fmt.Errorf("unexpected result from transaction callback")
	}

	return room, nil
}

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
		Client:           client,
		HotelsCollection: client.Database(cfg.DbConfig.MongoDB).Collection(cfg.DbConfig.HotelsCollection),
		RoomsCollection:  client.Database(cfg.DbConfig.MongoDB).Collection(cfg.DbConfig.RoomsCollection),
	}, nil
}

func (db *DB) DisconnectDB(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %s", err.Error())
	}
	return nil
}
