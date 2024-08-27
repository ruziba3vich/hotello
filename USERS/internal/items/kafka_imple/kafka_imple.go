package kafkaconsumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/ruziba3vich/hotello-users/genprotos/users"
)

type KafkaConsumer struct {
	service users.UsersServiceServer
	logger  *log.Logger
}

func NewKafkaConsumer(service users.UsersServiceServer, logger *log.Logger) *KafkaConsumer {
	return &KafkaConsumer{
		service: service,
		logger:  logger,
	}
}

func (kc *KafkaConsumer) StartConsumers(brokers []string) {
	go kc.consumeMessages("user.registration", brokers, kc.handleRegistration)
	go kc.consumeMessages("user.update.username", brokers, kc.handleUsernameUpdate)
	go kc.consumeMessages("user.update.password", brokers, kc.handlePasswordUpdate)
	go kc.consumeMessages("user.delete", brokers, kc.handleUserDelete)
}

func (kc *KafkaConsumer) consumeMessages(topic string, brokers []string, handler func(context.Context, *sarama.ConsumerMessage)) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		kc.logger.Fatalf("Error creating Kafka consumer: %s", err.Error())
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		kc.logger.Fatalf("Error creating partition consumer: %s", err.Error())
	}
	defer partitionConsumer.Close()

	ctx := context.Background()

	for msg := range partitionConsumer.Messages() {
		kc.logger.Printf("Received message from topic %s: %s", topic, string(msg.Value))
		handler(ctx, msg)
	}
}

func (kc *KafkaConsumer) handleRegistration(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req users.CreateUserRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling registration message: %s", err.Error())
		return
	}
	resp, err := kc.service.CreateUserService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error registering user: %s", err.Error())
	} else {
		kc.logger.Println(resp)
	}
}

func (kc *KafkaConsumer) handleUsernameUpdate(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req users.UpdateUsernameRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling username update message: %s", err.Error())
		return
	}
	resp, err := kc.service.UpdateUsernameService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error updating username: %s", err.Error())
	} else {
		kc.logger.Println(resp)
	}
}

func (kc *KafkaConsumer) handlePasswordUpdate(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req users.UpdatePasswordRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling password update message: %s", err.Error())
		return
	}
	resp, err := kc.service.UpdatePasswordService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error updating password: %s", err.Error())
	} else {
		kc.logger.Println(resp)
	}
}

func (kc *KafkaConsumer) handleUserDelete(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req users.DeleteUserRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling delete user message: %s", err.Error())
		return
	}
	resp, err := kc.service.DeleteUserService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error deleting user: %s", err.Error())
	} else {
		kc.logger.Println(resp)
	}
}
