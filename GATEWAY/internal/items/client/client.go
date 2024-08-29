package client

import (
	"context"

	"github.com/ruziba3vich/hotello-gateway/genprotos/booking"
	"github.com/ruziba3vich/hotello-gateway/genprotos/hotels"
	"github.com/ruziba3vich/hotello-gateway/genprotos/users"
)

type (
	HotelServiceClient struct {
		hotelServiceClient   hotels.HotelsServiceClient
		usersServiceClient   users.UsersServiceClient
		bookingServiceClient booking.BookingServiceClient
	}
)

func New(hotelServiceClient hotels.HotelsServiceClient, usersServiceClient users.UsersServiceClient, bookingServiceClient booking.BookingServiceClient) *HotelServiceClient {
	return &HotelServiceClient{
		hotelServiceClient:   hotelServiceClient,
		usersServiceClient:   usersServiceClient,
		bookingServiceClient: bookingServiceClient,
	}
}

func (h *HotelServiceClient) CheckUserByEmail(ctx context.Context, req *users.GetUserByFieldRequest) bool {
	_, err := h.usersServiceClient.GetUserByEmailService(ctx, req)
	return err != nil
}

func (h *HotelServiceClient) CheckUserByUsername(ctx context.Context, req *users.GetUserByFieldRequest) bool {
	_, err := h.usersServiceClient.GetUserByEmailService(ctx, req)
	return err != nil
}

func (h *HotelServiceClient) LoginUserServer(ctx context.Context, req *users.LoginUserRequest) (*users.RawResponse, error) {
	return h.usersServiceClient.LoginUserService(ctx, req)
}

func (u *HotelServiceClient) GetUserByIdService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	return u.usersServiceClient.GetUserByIdService(ctx, req)
}

func (u *HotelServiceClient) GetUserByUsernameService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	return u.usersServiceClient.GetUserByUsernameService(ctx, req)
}

func (u *HotelServiceClient) GetUserByEmailService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	return u.usersServiceClient.GetUserByEmailService(ctx, req)
}

func (u *HotelServiceClient) VerifyCodeService(ctx context.Context, req *users.VerifyCodeRequest) (*users.LoginUserResponse, error) {
	return u.usersServiceClient.VerifyCodeService(ctx, req)
}
