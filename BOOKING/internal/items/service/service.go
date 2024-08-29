package bookingservice

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
	
func NewBookingService(storage *storage.Storage, logger *log.Logger) booking.BookingServiceServer {
	return &BookingService{
		storage: storage,
		logger:  logger,
	}
}

func (s *BookingService) BookRoomService(ctx context.Context, req *booking.BookRoomRequest) (*booking.RawResponse, error) {
	return s.storage.BookRoom(ctx, req)
}

func (s *BookingService) RevokeOrderService(ctx context.Context, req *booking.RevokeOrderRequest) (*booking.RawResponse, error) {
	return s.storage.RevokeOrderService(ctx, req)
}

func (s *BookingService) GetBookingsByUserIdService(ctx context.Context, req *booking.GetBookingsByUserIdRequest) (*booking.GetBookingsByUserIdResponse, error) {
	return s.storage.GetBookingsByUserIdService(ctx, req)
}

func (s *BookingService) GetNotificationsByUserIdService(ctx context.Context, req *booking.GetNotificationsByUserIdRequest) (*booking.GetNotificationsResponse, error) {
	return s.storage.GetNotificationsByUserIdService(ctx, req)
}

func (s *BookingService) GetRoomAvailabilityService(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
	return s.storage.GetRoomAvailabilityService(ctx, req)
}

func (s *BookingService) GetNotificationByUserIdService(ctx context.Context, req *booking.GetNotificationByIdRequest) (*booking.Notification, error) {
	return s.storage.GetNotificationById(ctx, req)
}
