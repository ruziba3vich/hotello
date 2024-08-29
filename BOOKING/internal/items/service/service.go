package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
	"github.com/ruziba3vich/hotello-booking/internal/items/storage"
)

type BookingService struct {
	storage *storage.Storage
	logger  *log.Logger
	booking.UnimplementedBookingServiceServer
}

func NewBookingService(storage *storage.Storage, logger *log.Logger) *BookingService {
	return &BookingService{
		storage: storage,
		logger:  logger,
	}
}

func (s *BookingService) BookRoomService(ctx context.Context, req *booking.BookRoomRequest) (*booking.RawResponse, error) {
	s.logger.Printf("BookRoomService called with request: %+v\n", req)
	return s.storage.BookRoom(ctx, req)
}

func (s *BookingService) RevokeOrderService(ctx context.Context, req *booking.RevokeOrderRequest) (*booking.RawResponse, error) {
	s.logger.Printf("RevokeOrderService called with request: %+v\n", req)
	return s.storage.RevokeOrder(ctx, req)
}

func (s *BookingService) GetBookingsByUserIdService(ctx context.Context, req *booking.GetBookingsByUserIdRequest) (*booking.GetBookingsByUserIdResponse, error) {
	s.logger.Printf("GetBookingsByUserIdService called with request: %+v\n", req)
	return s.storage.GetBookingsByUserId(ctx, req)
}

func (s *BookingService) GetNotificationsByUserIdService(ctx context.Context, req *booking.GetNotificationsByUserIdRequest) (*booking.GetNotificationsResponse, error) {
	s.logger.Printf("GetNotificationsByUserIdService called with request: %+v\n", req)
	return s.storage.GetNotificationsByUserId(ctx, req)
}

func (s *BookingService) GetRoomAvailabilityService(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
	s.logger.Printf("GetRoomAvailabilityService called with request: %+v\n", req)
	return s.storage.GetRoomAvailability(ctx, req)
}

func (s *BookingService) GetNotificationByIdService(ctx context.Context, req *booking.GetNotificationByIdRequest) (*booking.Notification, error) {
	s.logger.Printf("GetNotificationByIdService called with request: %+v\n", req)
	return s.storage.GetNotificationById(ctx, req)
}

func (s *BookingService) CompleteBookingService(ctx context.Context, req *booking.CompleteBookingRequest) (*booking.RawResponse, error) {
	s.logger.Printf("CompleteBookingService called with request: %+v\n", req)
	return s.storage.CompleteBooking(ctx, req)
}

func (s *BookingService) GetBookingByIdService(ctx context.Context, req *booking.GetBookingByIdRequest) (*booking.BookEntity, error) {
	s.logger.Printf("GetBookingByIdService called with request: %+v\n", req)
	return s.storage.GetBookingById(ctx, req)
}
