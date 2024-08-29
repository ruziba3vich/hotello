package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		secretKey string
	}
)

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables if set.")
		return nil, err
	}
	return &Config{
		secretKey: os.Getenv("SECRET_KEY"),
	}, nil
}

func (c *Config) GetSecretKey() string {
	return c.secretKey
}
