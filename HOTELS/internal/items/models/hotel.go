package models

import (
	"github.com/ruziba3vich/hotello-hotels/genprotos/hotels"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Location string             `bson:"location"`
	Rating   float64            `bson:"rating"`
	Address  string             `bson:"address"`
	Rooms    []Room             `bson:"rooms"`
}

type Room struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	RoomType     string             `bson:"room_type"`
	NumberOfBeds int                `bson:"number_of_beds"`
	Available    bool               `bson:"available"`
	HotelID      primitive.ObjectID `bson:"hotel_id"`
}

func (h *Hotel) ToProto() *hotels.Hotel {
	var rooms []*hotels.Room
	for _, r := range h.Rooms {
		rooms = append(rooms, &hotels.Room{
			RoomId:       r.ID.Hex(),
			RoomType:     r.RoomType,
			NumberOfBeds: int32(r.NumberOfBeds),
			Available:    r.Available,
		})
	}
	return &hotels.Hotel{
		HotelId:  h.ID.Hex(),
		Name:     h.Name,
		Location: h.Location,
		Rating:   float32(h.Rating),
		Address:  h.Address,
		Rooms:    rooms,
	}
}

func (r *Room) ToProto() *hotels.Room {
	return &hotels.Room{
		RoomId:       r.ID.Hex(),
		RoomType:     r.RoomType,
		NumberOfBeds: int32(r.NumberOfBeds),
		Available:    r.Available,
	}
}

func (h *Hotel) ToGetAllHotels() *hotels.GetHotelByIdResponse {
	return &hotels.GetHotelByIdResponse{
		HotelId:  h.ID.Hex(),
		Name:     h.Name,
		Location: h.Location,
		Rating:   float32(h.Rating),
		Address:  h.Address,
	}
}
