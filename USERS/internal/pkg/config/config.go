package config

import (
	"log"
	"os"
	"strconv"
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
	secretKey    string
	redisUri     string
	smtpHost     string
	smtpPort     int
	smtpUser     string
	smtpPass     string
	kafkaBrokers []string
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
		Port:         os.Getenv("PORT"),
		Protocol:     os.Getenv("PROTOCOL"),
		secretKey:    os.Getenv("SECRET_KEY"),
		redisUri:     os.Getenv("REDIS_URI"),
		smtpHost:     os.Getenv("SMTP_HOST"),
		smtpUser:     os.Getenv("SMTP_USER"),
		smtpPass:     os.Getenv("SMTP_PASS"),
		kafkaBrokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		smtpPort:     smtpPort,
	}, nil
}

func (c *Config) GetSecretKey() string {
	return c.secretKey
}

func (c *Config) GetRedisURI() string {
	return c.redisUri
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

func (c *Config) GetKafkaBrokers() []string {
	return c.kafkaBrokers
}
