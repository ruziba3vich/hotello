package kafkaconsumer

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/IBM/sarama"
	"github.com/ruziba3vich/hotello-users/genprotos/users"
)

type KafkaConsumer struct {
	service users.UsersServiceServer
	logger  *log.Logger
	wg      *sync.WaitGroup
	sig     <-chan os.Signal
}

func NewKafkaConsumer(service users.UsersServiceServer, wg *sync.WaitGroup, sig <-chan os.Signal, logger *log.Logger) *KafkaConsumer {
	return &KafkaConsumer{
		service: service,
		wg:      wg,
		sig:     sig,
		logger:  logger,
	}
}

func (kc *KafkaConsumer) StartConsumers(ctx context.Context, cancel context.CancelFunc, brokers []string) {
	go kc.consumeMessages(ctx, cancel, "user.registration", brokers, kc.handleRegistration)
	go kc.consumeMessages(ctx, cancel, "user.update.username", brokers, kc.handleUsernameUpdate)
	go kc.consumeMessages(ctx, cancel, "user.update.password", brokers, kc.handlePasswordUpdate)
	go kc.consumeMessages(ctx, cancel, "user.delete", brokers, kc.handleUserDelete)
}

func (kc *KafkaConsumer) consumeMessages(ctx context.Context, cancel context.CancelFunc, topic string, brokers []string, handler func(context.Context, *sarama.ConsumerMessage)) {
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

	go func() {
		<-kc.sig
		kc.logger.Println("Received shutdown signal, waiting for current tasks to complete...")
		cancel()
		kc.wg.Wait()
		kc.logger.Println("Shutting down consumer.")
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			kc.wg.Add(1)
			func(m *sarama.ConsumerMessage) {
				defer kc.wg.Done()
				handler(ctx, msg)
			}(msg)
		case <-ctx.Done():
			return
		}
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
