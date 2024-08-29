package models

import (
	"time"

	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	CONFIRMED  = Status("CONFIRMED")
	REJECTED   = Status("REJECTED")
	INPROGRESS = Status("INPROGRESS")
	AVAILABLE  = Status("AVAILABLE")
)

func (n *Notification) FromProto(obj *booking.Notification) {
	n.NotificationID = primitive.NewObjectID()
	n.UserID = obj.UserId
	n.RoomID = obj.RoomId
	n.Message = obj.Message
	n.Read = false
}

func (n *Notification) ToProto() *booking.Notification {
	return &booking.Notification{
		NotificationId: n.NotificationID.Hex(),
		UserId:         n.UserID,
		RoomId:         n.RoomID,
		Message:        n.Message,
		Read:           n.Read,
	}
}

type (
	Status string

	Notification struct {
		NotificationID primitive.ObjectID `bson:"_id,omitempty"`
		UserID         string             `bson:"user_id"`
		RoomID         string             `bson:"room_id"`
		Message        string             `bson:"message"`
		Read           bool               `bson:"read"`
	}

	BookEntity struct {
		Id       primitive.ObjectID `bson:"_id"`
		UserID   string             `bson:"user_id"`
		RoomID   string             `bson:"room_id"`
		BookFrom time.Time          `bson:"book_from"`
		BookTill time.Time          `bson:"book_till"`
		Status   string             `bson:"status"`
		MadeAt   time.Time          `bson:"made_at"`
	}
)

func (b *BookEntity) FromBookRoomRequest(obj *booking.BookRoomRequest) {
	b.Id = primitive.NewObjectID()
	b.UserID = obj.UserId
	b.RoomID = obj.RoomId
	b.BookFrom = obj.BookFrom.AsTime()
	b.BookTill = obj.BookTill.AsTime()
	b.Status = "INPROGRESS"
	b.MadeAt = time.Now()
}

func (b *BookEntity) ToProto() *booking.BookEntity {
	return &booking.BookEntity{
		BookinId: b.Id.Hex(),
		UserId:   b.UserID,
		RoomId:   b.RoomID,
		BookFrom: convertToTimestampProto(b.BookFrom),
		BookTill: convertToTimestampProto(b.BookTill),
		Status:   b.Status,
		MadeAt:   convertToTimestampProto(b.MadeAt),
	}
}

func convertToTimestampProto(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
