package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	MongoURI   string
	MongoDB    string
	Collection string
}

type Config struct {
	DbConfig    DbConfig
	Port        string
	Protocol    string
	secretKey   string
	redisUri    string
	rabbitMqUri string
	smtpHost    string
	smtpPort    int
	smtpUser    string
	smtpPass    string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables if set.")
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		DbConfig: DbConfig{
			MongoURI:   os.Getenv("MONGO_URI"),
			MongoDB:    os.Getenv("MONGO_DB"),
			Collection: os.Getenv("MONGO_COLLECTION"),
		},
		Port:        os.Getenv("PORT"),
		Protocol:    os.Getenv("PROTOCOL"),
		secretKey:   os.Getenv("SECRET_KEY"),
		redisUri:    os.Getenv("REDIS_URI"),
		rabbitMqUri: os.Getenv("RABBITMQ_URI"),
		smtpHost:    os.Getenv("SMTP_HOST"),
		smtpUser:    os.Getenv("SMTP_USER"),
		smtpPass:    os.Getenv("SMTP_PASS"),
		smtpPort:    smtpPort,
	}, nil
}

func (c *Config) GetSecretKey() string {
	return c.secretKey
}

func (c *Config) GetRedisURI() string {
	return c.redisUri
}

func (c *Config) GetRabbitMqURI() string {
	return c.rabbitMqUri
}

func (c *Config) GetSmptpPass() string {
	return c.smtpPass
}

func (c *Config) GetSmptpUser() string {
	return c.smtpUser
}

func (c *Config) GetSmptpHost() string {
	return c.smtpHost
}

func (c *Config) GetSmptpPort() int {
	return c.smtpPort
}
