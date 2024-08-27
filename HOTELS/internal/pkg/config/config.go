package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	MongoURI   string
	MongoDB    string
	Collection string
}

type Config struct {
	DbConfig     DbConfig
	Port         string
	Protocol     string
	redisUri     string
	kafkaBrokers []string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables if set.")
	}

	return &Config{
		DbConfig: DbConfig{
			MongoURI:   os.Getenv("MONGO_URI"),
			MongoDB:    os.Getenv("MONGO_DB"),
			Collection: os.Getenv("MONGO_COLLECTION"),
		},
		Port:         os.Getenv("PORT"),
		Protocol:     os.Getenv("PROTOCOL"),
		redisUri:     os.Getenv("REDIS_URI"),
		kafkaBrokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
	}, nil
}

func (c *Config) GetRedisURI() string {
	return c.redisUri
}

func (c *Config) GetKafkaBrokers() []string {
	return c.kafkaBrokers
}
