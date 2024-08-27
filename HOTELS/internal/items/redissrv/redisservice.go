package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/hotello-hotels/genprotos/hotels"
	"github.com/ruziba3vich/hotello-hotels/internal/pkg/config"
)

type RedisService struct {
	client *redis.Client
	logger *log.Logger
}

func NewRedisService(client *redis.Client, logger *log.Logger) *RedisService {
	return &RedisService{
		client: client,
		logger: logger,
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

func (s *RedisService) StoreHotelInRedis(ctx context.Context, hotel *hotels.Hotel) error {
	hotelID := hotel.HotelId
	data, err := json.Marshal(hotel)
	if err != nil {
		s.logger.Println("Error marshalling hotel to JSON:", err)
		return err
	}

	err = s.client.Set(ctx, hotelID, data, 0).Err()
	if err != nil {
		s.logger.Println("Error storing hotel in Redis:", err)
		return err
	}

	return nil
}

func (s *RedisService) GetHotelFromRedis(ctx context.Context, hotelID string) (*hotels.Hotel, error) {
	data, err := s.client.Get(ctx, hotelID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("hotel with ID %s not found in Redis", hotelID)
	}
	if err != nil {
		s.logger.Println("Error retrieving hotel from Redis:", err)
		return nil, err
	}

	var hotel hotels.Hotel
	err = json.Unmarshal([]byte(data), &hotel)
	if err != nil {
		s.logger.Println("Error unmarshalling hotel JSON:", err)
		return nil, err
	}

	return &hotel, nil
}

func (s *RedisService) StoreRoomInRedis(ctx context.Context, room *hotels.Room) error {
	roomID := room.RoomId
	data, err := json.Marshal(room)
	if err != nil {
		s.logger.Println("Error marshalling room to JSON:", err)
		return err
	}

	err = s.client.Set(ctx, roomID, data, 0).Err()
	if err != nil {
		s.logger.Println("Error storing room in Redis:", err)
		return err
	}

	return nil
}

func (s *RedisService) GetRoomFromRedis(ctx context.Context, roomID string) (*hotels.Room, error) {
	data, err := s.client.Get(ctx, roomID).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("room with ID %s not found in Redis", roomID)
	}
	if err != nil {
		s.logger.Println("Error retrieving room from Redis:", err)
		return nil, err
	}

	var room hotels.Room
	err = json.Unmarshal([]byte(data), &room)
	if err != nil {
		s.logger.Println("Error unmarshalling room JSON:", err)
		return nil, err
	}

	return &room, nil
}
