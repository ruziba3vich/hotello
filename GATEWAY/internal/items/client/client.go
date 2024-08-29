package client

import (
	"github.com/ruziba3vich/hotello-gateway/genprotos/hotels"
	"github.com/ruziba3vich/hotello-gateway/genprotos/users"
)

type (
	HotelServiceClient struct {
		hotelServiceClient hotels.HotelsServiceClient
		usersServiceClient users.UsersServiceClient
	}
)

func New(hotelServiceClient hotels.HotelsServiceClient, usersServiceClient users.UsersServiceClient) *HotelServiceClient {
	
}
