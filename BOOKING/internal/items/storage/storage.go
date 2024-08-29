package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/ruziba3vich/hotello-booking/genprotos/booking"
	"github.com/ruziba3vich/hotello-booking/genprotos/hotels"
	"github.com/ruziba3vich/hotello-booking/internal/items/models"
	"github.com/ruziba3vich/hotello-booking/internal/pkg/config"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Storage struct {
		database     *DB
		writer       *kafka.Writer
		hotelsClient hotels.HotelsServiceClient
		logger       *log.Logger
	}

	DB struct {
		Client                  *mongo.Client
		BookingsCollection      *mongo.Collection
		NotificationsCollection *mongo.Collection
	}
)

func New(database *DB, writer *kafka.Writer, hotelsClient hotels.HotelsServiceClient, logger *log.Logger) *Storage {
	return &Storage{
		database:     database,
		writer:       writer,
		hotelsClient: hotelsClient,
		logger:       logger,
	}
}

func (s *Storage) GetRoomAvailabilityInInterval(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
	filter := bson.M{
		"room_id": req.RoomId,
		"status":  string(models.AVAILABLE),
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
	var bookingOrder models.BookEntity

	bookingOrder.FromBookRoomRequest(req)

	_, err := s.database.BookingsCollection.InsertOne(ctx, &bookingOrder)
	if err != nil {
		return nil, fmt.Errorf("error creating booking: %v", err)
	}

	err = s.sendNotification(ctx, "Room booked successfully", req.UserId, req.RoomId)
	if err != nil {
		return nil, fmt.Errorf("error sending notification: %v", err)
	}

	if _, err := s.hotelsClient.SetRoomToUnavailableService(ctx, &hotels.SetRoomToUnavailableRequest{
		RoomId: req.RoomId,
	}); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &booking.RawResponse{Message: "Room booked successfully"}, nil
}

func (s *Storage) RevokeOrder(ctx context.Context, req *booking.RevokeOrderRequest) (*booking.RawResponse, error) {
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

	if _, err := s.hotelsClient.SetRoomToAvailableService(ctx, &hotels.SetRoomToAvailableRequest{
		RoomId: req.RoomId,
	}); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &booking.RawResponse{Message: "Booking revoked successfully"}, nil
}

func (s *Storage) GetBookingsByUserId(ctx context.Context, req *booking.GetBookingsByUserIdRequest) (*booking.GetBookingsByUserIdResponse, error) {
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

func (s *Storage) GetNotificationsByUserId(ctx context.Context, req *booking.GetNotificationsByUserIdRequest) (*booking.GetNotificationsResponse, error) {
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

func (s *Storage) GetRoomAvailability(ctx context.Context, req *booking.GetRoomAvailabilityRequest) (*booking.GetRoomAvailabilityResponse, error) {
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

func (s *Storage) GetBookingById(ctx context.Context, req *booking.GetBookingByIdRequest) (*booking.BookEntity, error) {
	id, err := primitive.ObjectIDFromHex(req.BookingId)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	filter := bson.M{
		"_id": id,
	}

	var response models.BookEntity

	if err := s.database.BookingsCollection.FindOne(ctx, filter).Decode(&response); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return response.ToProto(), nil
}

func (s *Storage) CompleteBooking(ctx context.Context, req *booking.CompleteBookingRequest) (*booking.RawResponse, error) {
	id, err := primitive.ObjectIDFromHex(req.BookingId)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	filter := bson.M{"booking_id": id}
	update := bson.M{"$set": bson.M{"status": string(models.AVAILABLE)}}

	result, err := s.database.BookingsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("no matching booking found")
	}

	message := fmt.Sprintf("you have successfully completed your booking %s", req.BookingId)

	bookingEnt, err := s.GetBookingById(ctx, &booking.GetBookingByIdRequest{BookingId: req.BookingId})
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	if err := s.sendNotification(ctx, message, bookingEnt.UserId, bookingEnt.RoomId); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	if _, err := s.hotelsClient.SetRoomToAvailableService(ctx, &hotels.SetRoomToAvailableRequest{
		RoomId: bookingEnt.RoomId,
	}); err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return &booking.RawResponse{
		Message: message,
	}, nil
}

func (s *Storage) sendNotification(ctx context.Context, message, userId, roomId string) error {
	protonotification := booking.Notification{
		UserId:  userId,
		RoomId:  roomId,
		Message: message,
	}

	var notification models.Notification
	notification.FromProto(&protonotification)

	msgBytes, err := json.Marshal(&notification)
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

	if _, err := s.database.NotificationsCollection.InsertOne(ctx, &notification); err != nil {
		s.logger.Println(err)
		return err
	}

	s.logger.Printf("Notification sent for user: %s", userId)
	return nil
}

func (s *Storage) GetNotificationById(ctx context.Context, req *booking.GetNotificationByIdRequest) (*booking.Notification, error) {
	id, err := primitive.ObjectIDFromHex(req.NotificationId)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	filter := bson.M{"notification_id": id}
	update := bson.M{"$set": bson.M{"read": true}}

	var notification models.Notification

	if err := s.database.NotificationsCollection.FindOneAndUpdate(ctx, filter, update).Decode(&notification); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("notification with ID %s not found", req.NotificationId)
		}
		return nil, fmt.Errorf("error retrieving notification: %v", err)
	}

	return notification.ToProto(), nil
}

/*
   rpc BookRoom                    (BookRoomRequest)                 returns (RawResponse);
   rpc RevokeOrder                 (RevokeOrderRequest)              returns (RawResponse);
   rpc GetBookingsByUserId         (GetBookingsByUserIdRequest)      returns (GetBookingsByUserIdResponse);
   // rpc SendNotification            (Notification)                    returns (google.protobuf.Empty);
   rpc GetNotificationsByUserId    (GetNotificationsByUserIdRequest) returns (GetNotificationsResponse);
*/
// =======================================================================================================================

func ConnectDB(cfg *config.Config, ctx context.Context) (*DB, error) {
	clientOptions := options.Client().ApplyURI(cfg.DbConfig.MongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %s", err.Error())
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %s", err.Error())
	}

	return &DB{
		Client:                  client,
		BookingsCollection:      client.Database(cfg.DbConfig.MongoDB).Collection(cfg.DbConfig.BookingsCollection),
		NotificationsCollection: client.Database(cfg.DbConfig.MongoDB).Collection(cfg.DbConfig.NotificationsCollection),
	}, nil
}

func (db *DB) DisconnectDB(ctx context.Context) error {
	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %s", err.Error())
	}
	return nil
}
