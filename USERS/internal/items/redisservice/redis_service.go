package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/hotello-users/genprotos/users"
	"github.com/ruziba3vich/hotello-users/internal/pkg/config"
)

type (
	RedisService struct {
		redisDb *redis.Client
		logger  *log.Logger
	}
)

func New(redisDb *redis.Client, logger *log.Logger) *RedisService {
	return &RedisService{
		logger:  logger,
		redisDb: redisDb,
	}
}

func NewRedisClient(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.GetRedisURI(),
		Password: "",
		DB:       0,
	})

	return rdb
}

func (r *RedisService) StoreUserInRedis(ctx context.Context, user *users.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = r.redisDb.Set(ctx, user.Id, userJSON, time.Hour*24).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisService) GetUserFromRedis(ctx context.Context, userId string) (*users.User, error) {
	userJSON, err := r.redisDb.Get(ctx, userId).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		r.logger.Printf("ERROR WHILE GETTING DATA FROM REDIS : %s\n", err.Error())
		return nil, err
	}

	var user users.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		r.logger.Printf("ERROR WHILE UNMARSHALING DATA : %s\n", err.Error())
		return nil, err
	}
	return &user, nil
}

func (r *RedisService) DeleteUserFromRedis(ctx context.Context, userID string) error {
	result, err := r.redisDb.Del(ctx, userID).Result()
	if err != nil {
		return err
	}

	if result == 0 {
		r.logger.Printf("User with ID %s does not exist in Redis", userID)
	} else {
		r.logger.Printf("User with ID %s has been deleted from Redis", userID)
	}

	return nil
}

func (r *RedisService) StoreEmailAndCode(ctx context.Context, email string, code int) error {
	codeKey := "verification_code:" + email
	err := r.redisDb.Set(ctx, codeKey, code, time.Minute*15).Err()
	if err != nil {
		r.logger.Printf("ERROR WHILE STORING VERIFICATION CODE: %s\n", err.Error())
		return err
	}
	return nil
}

func (r *RedisService) GetCodeByEmail(ctx context.Context, email string) (int, error) {
	codeKey := "verification_code:" + email
	codeStr, err := r.redisDb.Get(ctx, codeKey).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		r.logger.Printf("ERROR WHILE GETTING VERIFICATION CODE: %s\n", err.Error())
		return 0, err
	}

	var code int
	_, err = fmt.Sscanf(codeStr, "%d", &code)
	if err != nil {
		r.logger.Printf("ERROR WHILE PARSING VERIFICATION CODE: %s\n", err.Error())
		return 0, err
	}

	return code, nil
}
