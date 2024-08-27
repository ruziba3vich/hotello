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

	"github.com/ruziba3vich/hotello-users/genprotos/users"
	kafkaconsumer "github.com/ruziba3vich/hotello-users/internal/items/kafka_imple"
	"github.com/ruziba3vich/hotello-users/internal/items/redisservice"
	usersservice "github.com/ruziba3vich/hotello-users/internal/items/service"
	"github.com/ruziba3vich/hotello-users/internal/items/storage"
	"github.com/ruziba3vich/hotello-users/internal/pkg/config"
	"github.com/ruziba3vich/hotello-users/internal/pkg/utils"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config, logger *log.Logger) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	mongoDb, err := storage.ConnectDB(cfg, ctx)
	if err != nil {
		logger.Println(err)
		return err
	}
	service := usersservice.NewUsersService(
		storage.New(
			redisservice.New(redisservice.NewRedisClient(cfg), logger),
			mongoDb,
			logger,
			utils.NewTokenGenerator(cfg),
			utils.NewPasswordHasher(),
			cfg,
		),
		logger,
	)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen(cfg.Protocol, cfg.Port)
	if err != nil {
		logger.Println(err)
		return err
	}

	server := grpc.NewServer()
	users.RegisterUsersServiceServer(server, service)

	kafkaconsumer := kafkaconsumer.NewKafkaConsumer(service, &sync.WaitGroup{}, sigChan, logger)
	kafkaconsumer.StartConsumers(ctx, cancel, cfg.GetKafkaBrokers())

	return server.Serve(listener)
}
