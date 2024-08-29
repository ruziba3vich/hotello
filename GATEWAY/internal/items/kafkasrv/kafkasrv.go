package kafkasrv

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type KafkaPublisher struct {
	producer sarama.SyncProducer
	logger   *log.Logger
}

func NewKafkaPublisher(brokers []string, topic string, logger *log.Logger) (*KafkaPublisher, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaPublisher{
		producer: producer,
		logger:   logger,
	}, nil
}

func (kp *KafkaPublisher) Publish(message interface{}, topic string) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msgBytes),
	}

	partition, offset, err := kp.producer.SendMessage(kafkaMsg)
	if err != nil {
		kp.logger.Printf("Error publishing message to topic %s: %s", topic, err.Error())
		return err
	}

	kp.logger.Printf("Message published to topic %s, partition %d, offset %d", topic, partition, offset)
	return nil
}

func (kp *KafkaPublisher) Close() error {
	return kp.producer.Close()
}
