package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
	"github.com/ruziba3vich/hotello-booking/genprotos/hotels"
	kafkaconsumer "github.com/ruziba3vich/hotello-booking/internal/items/kafkasrv"
	bookingservice "github.com/ruziba3vich/hotello-booking/internal/items/service"
	"github.com/ruziba3vich/hotello-booking/internal/items/storage"
	"github.com/ruziba3vich/hotello-booking/internal/pkg/config"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run(cfg *config.Config, logger *log.Logger) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	mongoDb, err := storage.ConnectDB(cfg, ctx)
	if err != nil {
		logger.Println("Failed to connect to MongoDB:", err)
		return err
	}

	hotelsClient, err := getHotelServiceClient(cfg)
	if err != nil {
		return err
	}

	service := bookingservice.NewBookingService(
		storage.New(
			mongoDb,
			createKafkaWriter(cfg),
			hotelsClient,
			logger,
		),
		logger,
	)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen(cfg.Protocol, cfg.Port)
	if err != nil {
		logger.Println("Failed to create gRPC listener:", err)
		return err
	}

	server := grpc.NewServer()
	booking.RegisterBookingServiceServer(server, service)

	kafkaconsumer := kafkaconsumer.NewBookingKafkaConsumer(service, &sync.WaitGroup{}, sigChan, logger)
	kafkaconsumer.StartConsumers(ctx, cancel, cfg.GetKafkaBrokers())

	logger.Printf("Booking server has started on %s\n", cfg.Port)

	return server.Serve(listener)
}

func createKafkaWriter(cfg *config.Config) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(cfg.GetKafkaBrokers()...),
		Topic:    cfg.KafkaConfig.NotificationsTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getHotelServiceClient(cfg *config.Config) (hotels.HotelsServiceClient, error) {
	conn, err := grpc.NewClient(cfg.GetHotelsService(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return hotels.NewHotelsServiceClient(conn), nil
}
