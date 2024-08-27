package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	MongoURI         string
	MongoDB          string
	HotelsCollection string
	RoomsCollection  string
}

type Config struct {
	DbConfig DbConfig
	Port     string
	Protocol string
	redisUri string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables if set.")
	}

	return &Config{
		DbConfig: DbConfig{
			MongoURI:         os.Getenv("MONGO_URI"),
			MongoDB:          os.Getenv("MONGO_DB"),
			HotelsCollection: os.Getenv("MONGO_HOTELS_COLLECTION"),
			RoomsCollection:  os.Getenv("MONGO_ROOMS_COLLECTION"),
		},
		Port:     os.Getenv("PORT"),
		Protocol: os.Getenv("PROTOCOL"),
		redisUri: os.Getenv("REDIS_URI"),
	}, nil
}

func (c *Config) GetRedisURI() string {
	return c.redisUri
}
