package hotelsservice

import (
	"context"
	"log"

	"github.com/ruziba3vich/hotello-hotels/genprotos/hotels"
	"github.com/ruziba3vich/hotello-hotels/internal/items/storage"
)

type HotelsService struct {
	storage *storage.HotelsStorage
	logger  *log.Logger
	hotels.UnimplementedHotelsServiceServer
}

func NewHotelsService(storage *storage.HotelsStorage, logger *log.Logger) hotels.HotelsServiceServer {
	return &HotelsService{
		storage: storage,
		logger:  logger,
	}
}

func (s *HotelsService) CreateHotelService(ctx context.Context, req *hotels.CreateHotelRequest) (*hotels.Hotel, error) {
	s.logger.Printf("CreateHotelService called with request: %+v\n", req)
	return s.storage.CreateHotel(ctx, req)
}

func (s *HotelsService) CreateRoomService(ctx context.Context, req *hotels.CreateRoomRequest) (*hotels.Room, error) {
	s.logger.Printf("CreateRoomService called with request: %+v\n", req)
	return s.storage.CreateRoom(ctx, req)
}

func (s *HotelsService) GetAvailableRoomsByHotelService(ctx context.Context, req *hotels.GetAvailableRoomsByHotelRequest) (*hotels.GetAvailableRoomsByHotelResponse, error) {
	s.logger.Printf("GetAvailableRoomsByHotelService called with request: %+v\n", req)
	return s.storage.GetAvailableRoomsByHotel(ctx, req)
}

func (s *HotelsService) GetHotelByIdService(ctx context.Context, req *hotels.GetHotelByIdRequest) (*hotels.GetHotelByIdResponse, error) {
	s.logger.Printf("GetHotelByIdService called with request: %+v\n", req)
	return s.storage.GetHotelById(ctx, req)
}

func (s *HotelsService) GetAllHotelsService(ctx context.Context, req *hotels.GetAllHotelsRequest) (*hotels.GetAllHotelsResponse, error) {
	s.logger.Printf("GetAllHotelsService called with request: %+v\n", req)
	return s.storage.GetAllHotels(ctx, req)
}

func (s *HotelsService) AddRoomToHotelService(ctx context.Context, req *hotels.CreateRoomRequest) (*hotels.Room, error) {
	s.logger.Printf("AddRoomToHotelService called with request: %+v\n", req)
	return s.storage.AddRoomToHotel(ctx, req)
}

func (s *HotelsService) SetRoomToAvailableService(ctx context.Context, req *hotels.SetRoomToAvailableRequest) (*hotels.Room, error) {
	s.logger.Printf("SetRoomToAvailableService called with request: %+v\n", req)
	return s.storage.SetRoomToAvailable(ctx, req)
}

func (s *HotelsService) SetRoomToUnavailableService(ctx context.Context, req *hotels.SetRoomToUnavailableRequest) (*hotels.Room, error) {
	s.logger.Printf("SetRoomToUnavailableService called with request: %+v\n", req)
	return s.storage.SetRoomToUnavailable(ctx, req)
}
