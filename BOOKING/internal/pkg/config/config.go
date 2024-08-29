package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DbConfig      DbConfig
		Port          string
		Protocol      string
		KafkaConfig   KafkaConfig
		hotelsService string
	}
	DbConfig struct {
		MongoURI                string
		MongoDB                 string
		BookingsCollection      string
		NotificationsCollection string
	}

	KafkaConfig struct {
		NotificationsTopic string
		kafkaBrokers       []string
	}
)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables if set.")
	}

	return &Config{
		DbConfig: DbConfig{
			MongoURI:                os.Getenv("MONGO_URI"),
			MongoDB:                 os.Getenv("MONGO_DB"),
			BookingsCollection:      os.Getenv("BOOKINGS_COLLECTION"),
			NotificationsCollection: os.Getenv("NOTIFICATIONS_COLLECTION"),
		},
		Port:     os.Getenv("PORT"),
		Protocol: os.Getenv("PROTOCOL"),
		KafkaConfig: KafkaConfig{
			kafkaBrokers:       strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
			NotificationsTopic: os.Getenv("NOTIFICATIONS_TPIC"),
		},
		hotelsService: os.Getenv("HOTELS_SERVICE"),
	}, nil
}

func (c *Config) GetKafkaBrokers() []string {
	return c.KafkaConfig.kafkaBrokers
}

func (c *Config) GetHotelsService() string {
	return c.hotelsService
}
