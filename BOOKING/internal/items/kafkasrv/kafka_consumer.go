package kafkaconsumer

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/IBM/sarama"
	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
)

type BookingKafkaConsumer struct {
	service booking.BookingServiceServer
	logger  *log.Logger
	wg      *sync.WaitGroup
	sig     <-chan os.Signal
}

func NewBookingKafkaConsumer(service booking.BookingServiceServer, wg *sync.WaitGroup, sig <-chan os.Signal, logger *log.Logger) *BookingKafkaConsumer {
	return &BookingKafkaConsumer{
		service: service,
		wg:      wg,
		sig:     sig,
		logger:  logger,
	}
}

func (kc *BookingKafkaConsumer) StartConsumers(ctx context.Context, cancel context.CancelFunc, brokers []string) {
	go kc.consumeMessages(ctx, cancel, "booking.room.book", brokers, kc.handleBookRoom)
	go kc.consumeMessages(ctx, cancel, "booking.order.revoke", brokers, kc.handleRevokeOrder)
}

func (kc *BookingKafkaConsumer) handleBookRoom(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req booking.BookRoomRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling book room message: %s", err.Error())
		return
	}
	resp, err := kc.service.BookRoomService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error booking room: %s", err.Error())
	} else {
		kc.logger.Printf("Successfully booked room: %+v", resp)
	}
}

func (kc *BookingKafkaConsumer) handleRevokeOrder(ctx context.Context, msg *sarama.ConsumerMessage) {
	var req booking.RevokeOrderRequest
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		kc.logger.Printf("Error unmarshalling revoke order message: %s", err.Error())
		return
	}
	resp, err := kc.service.RevokeOrderService(ctx, &req)
	if err != nil {
		kc.logger.Printf("Error revoking order: %s", err.Error())
	} else {
		kc.logger.Printf("Successfully revoked order: %+v", resp)
	}
}

func (kc *BookingKafkaConsumer) consumeMessages(ctx context.Context, cancel context.CancelFunc, topic string, brokers []string, handler func(context.Context, *sarama.ConsumerMessage)) {
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
