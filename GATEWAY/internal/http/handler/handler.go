package handler

import (
	"log"

	"github.com/ruziba3vich/hotello-gateway/internal/items/client"
	"github.com/ruziba3vich/hotello-gateway/internal/items/kafkasrv"
	"github.com/ruziba3vich/hotello-gateway/internal/pkg/utils"
)

type (
	HotelloHandler struct {
		service           *client.HotelServiceClient
		logger            *log.Logger
		usersPublisher    *kafkasrv.KafkaPublisher
		bookingsPublisher *kafkasrv.KafkaPublisher
		hotelsPublisher   *kafkasrv.KafkaPublisher
		utils             utils.TokenGenerator
	}
)

func New(service *client.HotelServiceClient, usersPublisher *kafkasrv.KafkaPublisher, bookingsPublisher *kafkasrv.KafkaPublisher, hotelsPublisher *kafkasrv.KafkaPublisher, utils utils.TokenGenerator, logger *log.Logger) *HotelloHandler {
	return &HotelloHandler{
		service:           service,
		usersPublisher:    usersPublisher,
		hotelsPublisher:   hotelsPublisher,
		bookingsPublisher: bookingsPublisher,
		utils:             utils,
		logger:            logger,
	}
}
