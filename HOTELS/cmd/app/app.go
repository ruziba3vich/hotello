package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/ruziba3vich/hotello-hotels/genprotos/hotels"
	redisservice "github.com/ruziba3vich/hotello-hotels/internal/items/redissrv"
	hotelservice "github.com/ruziba3vich/hotello-hotels/internal/items/service"
	"github.com/ruziba3vich/hotello-hotels/internal/items/storage"
	"github.com/ruziba3vich/hotello-hotels/internal/pkg/config"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config, logger *log.Logger) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	mongoDb, err := storage.ConnectDB(cfg, ctx)
	if err != nil {
		logger.Println("Failed to connect to database:", err)
		return err
	}

	redissrv := redisservice.NewRedisService(redisservice.NewRedisClient(cfg), logger)

	service := hotelservice.NewHotelsService(
		storage.NewHotelsStorage(redissrv, mongoDb, logger),
		logger,
	)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen(cfg.Protocol, cfg.Port)
	if err != nil {
		logger.Println("Failed to start listener:", err)
		return err
	}

	server := grpc.NewServer()
	hotels.RegisterHotelsServiceServer(server, service)

	g.Go(func() error {
		logger.Println("gRPC server starting...")
		if err := server.Serve(listener); err != nil {
			logger.Println("gRPC server error:", err)
			return err
		}
		return nil
	})

	g.Go(func() error {
		select {
		case sig := <-sigChan:
			logger.Println("Received signal:", sig)
			server.GracefulStop()
			logger.Println("gRPC server stopped gracefully")
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil && err != context.Canceled {
		logger.Println("Error during execution:", err)
		return err
	}

	return nil
}
