package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
	"github.com/ruziba3vich/hotello-booking/internal/items/models"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Storage struct {
		database *DB
		writer   *kafka.Writer
		logger   *log.Logger
	}

	DB struct {
		Client                  *mongo.Client
		BookingsCollection      *mongo.Collection
		NotificationsCollection *mongo.Collection
	}
)

func New(database *DB, writer *kafka.Writer, logger *log.Logger) *Storage {
	return &Storage{
		database: database,
		writer:   writer,
		logger:   logger,
	}
}

func (s *Storage) GetRoomAvailabilityInInterval(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
	filter := bson.M{
		"room_id": req.RoomId,
		"status":  models.CONFIRMED,
		"$or": []bson.M{
			{
				"$and": []bson.M{
					{"book_from": bson.M{"$lt": req.To.AsTime()}},
					{"book_till": bson.M{"$gt": req.From.AsTime()}},
				},
			},
		},
	}

	count, err := s.database.BookingsCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error checking room availability: %v", err)
	}

	return &booking.GetRoomAvailabilityResponse{
		Available: count == 0,
	}, nil
}

func (s *Storage) BookRoom(ctx context.Context, req *booking.BookRoomRequest) (*booking.RawResponse, error) {
	fromTime := req.BookFrom.AsTime()
	tillTime := req.BookTill.AsTime()

	bookingOrder := booking.BookEntity{
		UserId:   req.UserId,
		RoomId:   req.RoomId,
		BookFrom: timestamppb.New(fromTime),
		BookTill: timestamppb.New(tillTime),
		Status:   string(models.CONFIRMED),
		MadeAt:   timestamppb.Now(),
	}

	_, err := s.database.BookingsCollection.InsertOne(ctx, &bookingOrder)
	if err != nil {
		return nil, fmt.Errorf("error creating booking: %v", err)
	}

	err = s.sendNotification(ctx, "Room booked successfully", req.UserId, req.RoomId)
	if err != nil {
		return nil, fmt.Errorf("error sending notification: %v", err)
	}

	return &booking.RawResponse{Message: "Room booked successfully"}, nil
}

func (s *Storage) RevokeOrderService(ctx context.Context, req *booking.RevokeOrderRequest) (*booking.RawResponse, error) {
	filter := bson.M{"user_id": req.UserId, "room_id": req.RoomId, "status": "CONFIRMED"}

	bookingOrder := s.database.BookingsCollection.FindOne(ctx, filter)
	if bookingOrder.Err() != nil {
		return nil, fmt.Errorf("error finding booking: %v", bookingOrder.Err())
	}

	_, err := s.database.BookingsCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error revoking booking: %v", err)
	}

	err = s.sendNotification(ctx, "Booking revoked successfully", req.UserId, req.RoomId)
	if err != nil {
		return nil, fmt.Errorf("error sending notification: %v", err)
	}

	return &booking.RawResponse{Message: "Booking revoked successfully"}, nil
}

func (s *Storage) GetBookingsByUserIdService(ctx context.Context, req *booking.GetBookingsByUserIdRequest) (*booking.GetBookingsByUserIdResponse, error) {
	filter := bson.M{"user_id": req.UserId}

	findOptions := options.Find()
	findOptions.SetLimit(int64(req.Limit))
	findOptions.SetSkip(int64(req.Page * req.Limit))

	cursor, err := s.database.BookingsCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("error retrieving bookings: %v", err)
	}
	defer cursor.Close(ctx)

	var bookings []*booking.BookEntity
	if err := cursor.All(ctx, &bookings); err != nil {
		return nil, fmt.Errorf("error decoding bookings: %v", err)
	}

	return &booking.GetBookingsByUserIdResponse{Bookings: bookings}, nil
}

func (s *Storage) GetNotificationsByUserIdService(ctx context.Context, req *booking.GetNotificationsByUserIdRequest) (*booking.GetNotificationsResponse, error) {
	filter := bson.M{"user_id": req.UserId}

	cursor, err := s.database.NotificationsCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving notifications: %v", err)
	}
	defer cursor.Close(ctx)

	var notifications []*booking.Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, fmt.Errorf("error decoding notifications: %v", err)
	}

	return &booking.GetNotificationsResponse{Notifications: notifications}, nil
}

func (s *Storage) GetRoomAvailabilityService(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
	fromTime := req.From.AsTime()
	toTime := req.To.AsTime()

	filter := bson.M{
		"room_id": req.RoomId,
		"status":  "CONFIRMED",
		"$or": []bson.M{
			{
				"$and": []bson.M{
					{"book_from": bson.M{"$lt": toTime}},
					{"book_till": bson.M{"$gt": fromTime}},
				},
			},
		},
	}

	count, err := s.database.BookingsCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error checking room availability: %v", err)
	}

	return &booking.GetRoomAvailabilityResponse{Available: count == 0}, nil
}

func (s *Storage) sendNotification(ctx context.Context, message, userId, roomId string) error {
	notification := booking.Notification{
		UserId:  userId,
		RoomId:  roomId,
		Message: message,
	}

	msgBytes, err := proto.Marshal(&notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	kafkaMsg := kafka.Message{
		Key:   []byte(userId),
		Value: msgBytes,
	}

	err = s.writer.WriteMessages(ctx, kafkaMsg)
	if err != nil {
		s.logger.Printf("error sending Kafka message: %v", err)
		return err
	}

	s.logger.Printf("Notification sent for user: %s", userId)
	return nil
}

/*
   rpc BookRoom                    (BookRoomRequest)                 returns (RawResponse);
   rpc RevokeOrder                 (RevokeOrderRequest)              returns (RawResponse);
   rpc GetBookingsByUserId         (GetBookingsByUserIdRequest)      returns (GetBookingsByUserIdResponse);
   // rpc SendNotification            (Notification)                    returns (google.protobuf.Empty);
   rpc GetNotificationsByUserId    (GetNotificationsByUserIdRequest) returns (GetNotificationsResponse);
*/
