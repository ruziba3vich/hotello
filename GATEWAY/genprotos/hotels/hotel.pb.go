// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: protos/hotels_proto/hotel.proto

package hotels

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hotel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId  string  `protobuf:"bytes,1,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Rating   float32 `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Address  string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Rooms    []*Room `protobuf:"bytes,6,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *Hotel) Reset() {
	*x = Hotel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hotel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotel) ProtoMessage() {}

func (x *Hotel) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotel.ProtoReflect.Descriptor instead.
func (*Hotel) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *Hotel) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *Hotel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hotel) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Hotel) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Hotel) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Hotel) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type GetAllHotelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllHotelsRequest) Reset() {
	*x = GetAllHotelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHotelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHotelsRequest) ProtoMessage() {}

func (x *GetAllHotelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHotelsRequest.ProtoReflect.Descriptor instead.
func (*GetAllHotelsRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllHotelsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllHotelsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllHotelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotels []*GetHotelByIdResponse `protobuf:"bytes,1,rep,name=hotels,proto3" json:"hotels,omitempty"`
}

func (x *GetAllHotelsResponse) Reset() {
	*x = GetAllHotelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllHotelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllHotelsResponse) ProtoMessage() {}

func (x *GetAllHotelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllHotelsResponse.ProtoReflect.Descriptor instead.
func (*GetAllHotelsResponse) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllHotelsResponse) GetHotels() []*GetHotelByIdResponse {
	if x != nil {
		return x.Hotels
	}
	return nil
}

type GetHotelByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId string `protobuf:"bytes,1,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
}

func (x *GetHotelByIdRequest) Reset() {
	*x = GetHotelByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelByIdRequest) ProtoMessage() {}

func (x *GetHotelByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelByIdRequest.ProtoReflect.Descriptor instead.
func (*GetHotelByIdRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{3}
}

func (x *GetHotelByIdRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

type GetHotelByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId  string  `protobuf:"bytes,1,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Rating   float32 `protobuf:"fixed32,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Address  string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetHotelByIdResponse) Reset() {
	*x = GetHotelByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHotelByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHotelByIdResponse) ProtoMessage() {}

func (x *GetHotelByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHotelByIdResponse.ProtoReflect.Descriptor instead.
func (*GetHotelByIdResponse) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{4}
}

func (x *GetHotelByIdResponse) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *GetHotelByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetHotelByIdResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GetHotelByIdResponse) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *GetHotelByIdResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetAvailableRoomsByHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId string `protobuf:"bytes,1,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit   int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAvailableRoomsByHotelRequest) Reset() {
	*x = GetAvailableRoomsByHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableRoomsByHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableRoomsByHotelRequest) ProtoMessage() {}

func (x *GetAvailableRoomsByHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableRoomsByHotelRequest.ProtoReflect.Descriptor instead.
func (*GetAvailableRoomsByHotelRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{5}
}

func (x *GetAvailableRoomsByHotelRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *GetAvailableRoomsByHotelRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAvailableRoomsByHotelRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAvailableRoomsByHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetAvailableRoomsByHotelResponse) Reset() {
	*x = GetAvailableRoomsByHotelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableRoomsByHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableRoomsByHotelResponse) ProtoMessage() {}

func (x *GetAvailableRoomsByHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableRoomsByHotelResponse.ProtoReflect.Descriptor instead.
func (*GetAvailableRoomsByHotelResponse) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{6}
}

func (x *GetAvailableRoomsByHotelResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId       string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	RoomType     string `protobuf:"bytes,2,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	NumberOfBeds int32  `protobuf:"varint,3,opt,name=number_of_beds,json=numberOfBeds,proto3" json:"number_of_beds,omitempty"`
	Available    bool   `protobuf:"varint,4,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{7}
}

func (x *Room) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *Room) GetNumberOfBeds() int32 {
	if x != nil {
		return x.NumberOfBeds
	}
	return 0
}

func (x *Room) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type CreateHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string  `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Rating   float32 `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Address  string  `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateHotelRequest) Reset() {
	*x = CreateHotelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHotelRequest) ProtoMessage() {}

func (x *CreateHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHotelRequest.ProtoReflect.Descriptor instead.
func (*CreateHotelRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{8}
}

func (x *CreateHotelRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHotelRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateHotelRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateHotelRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType     string `protobuf:"bytes,1,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	NumberOfBeds int32  `protobuf:"varint,2,opt,name=number_of_beds,json=numberOfBeds,proto3" json:"number_of_beds,omitempty"`
	Available    bool   `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
	HotelId      string `protobuf:"bytes,4,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{9}
}

func (x *CreateRoomRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *CreateRoomRequest) GetNumberOfBeds() int32 {
	if x != nil {
		return x.NumberOfBeds
	}
	return 0
}

func (x *CreateRoomRequest) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *CreateRoomRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

type SetRoomToUnavailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *SetRoomToUnavailableRequest) Reset() {
	*x = SetRoomToUnavailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoomToUnavailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoomToUnavailableRequest) ProtoMessage() {}

func (x *SetRoomToUnavailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoomToUnavailableRequest.ProtoReflect.Descriptor instead.
func (*SetRoomToUnavailableRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{10}
}

func (x *SetRoomToUnavailableRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type SetRoomToAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *SetRoomToAvailableRequest) Reset() {
	*x = SetRoomToAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hotels_proto_hotel_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRoomToAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoomToAvailableRequest) ProtoMessage() {}

func (x *SetRoomToAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hotels_proto_hotel_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoomToAvailableRequest.ProtoReflect.Descriptor instead.
func (*SetRoomToAvailableRequest) Descriptor() ([]byte, []int) {
	return file_protos_hotels_proto_hotel_proto_rawDescGZIP(), []int{11}
}

func (x *SetRoomToAvailableRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

var File_protos_hotels_proto_hotel_proto protoreflect.FileDescriptor

var file_protos_hotels_proto_hotel_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x05, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x45, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x22, 0x30, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x22,
	0x93, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x66, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3f, 0x0a,
	0x20, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x80,
	0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x65, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x42,
	0x65, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x22, 0x76, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x65, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x42, 0x65,
	0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x1b, 0x53,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x32, 0x9a, 0x04, 0x0a, 0x0d, 0x48, 0x6f,
	0x74, 0x65, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x2e,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x66,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x15, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x3e, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x42, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x55,
	0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x6f, 0x55, 0x6e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_hotels_proto_hotel_proto_rawDescOnce sync.Once
	file_protos_hotels_proto_hotel_proto_rawDescData = file_protos_hotels_proto_hotel_proto_rawDesc
)

func file_protos_hotels_proto_hotel_proto_rawDescGZIP() []byte {
	file_protos_hotels_proto_hotel_proto_rawDescOnce.Do(func() {
		file_protos_hotels_proto_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_hotels_proto_hotel_proto_rawDescData)
	})
	return file_protos_hotels_proto_hotel_proto_rawDescData
}

var file_protos_hotels_proto_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protos_hotels_proto_hotel_proto_goTypes = []any{
	(*Hotel)(nil),                            // 0: Hotel
	(*GetAllHotelsRequest)(nil),              // 1: GetAllHotelsRequest
	(*GetAllHotelsResponse)(nil),             // 2: GetAllHotelsResponse
	(*GetHotelByIdRequest)(nil),              // 3: GetHotelByIdRequest
	(*GetHotelByIdResponse)(nil),             // 4: GetHotelByIdResponse
	(*GetAvailableRoomsByHotelRequest)(nil),  // 5: GetAvailableRoomsByHotelRequest
	(*GetAvailableRoomsByHotelResponse)(nil), // 6: GetAvailableRoomsByHotelResponse
	(*Room)(nil),                             // 7: Room
	(*CreateHotelRequest)(nil),               // 8: CreateHotelRequest
	(*CreateRoomRequest)(nil),                // 9: CreateRoomRequest
	(*SetRoomToUnavailableRequest)(nil),      // 10: SetRoomToUnavailableRequest
	(*SetRoomToAvailableRequest)(nil),        // 11: SetRoomToAvailableRequest
}
var file_protos_hotels_proto_hotel_proto_depIdxs = []int32{
	7,  // 0: Hotel.rooms:type_name -> Room
	4,  // 1: GetAllHotelsResponse.hotels:type_name -> GetHotelByIdResponse
	7,  // 2: GetAvailableRoomsByHotelResponse.rooms:type_name -> Room
	8,  // 3: HotelsService.CreateHotelService:input_type -> CreateHotelRequest
	9,  // 4: HotelsService.CreateRoomService:input_type -> CreateRoomRequest
	5,  // 5: HotelsService.GetAvailableRoomsByHotelService:input_type -> GetAvailableRoomsByHotelRequest
	3,  // 6: HotelsService.GetHotelByIdService:input_type -> GetHotelByIdRequest
	1,  // 7: HotelsService.GetAllHotelsService:input_type -> GetAllHotelsRequest
	9,  // 8: HotelsService.AddRoomToHotelService:input_type -> CreateRoomRequest
	11, // 9: HotelsService.SetRoomToAvailableService:input_type -> SetRoomToAvailableRequest
	10, // 10: HotelsService.SetRoomToUnavailableService:input_type -> SetRoomToUnavailableRequest
	0,  // 11: HotelsService.CreateHotelService:output_type -> Hotel
	7,  // 12: HotelsService.CreateRoomService:output_type -> Room
	6,  // 13: HotelsService.GetAvailableRoomsByHotelService:output_type -> GetAvailableRoomsByHotelResponse
	4,  // 14: HotelsService.GetHotelByIdService:output_type -> GetHotelByIdResponse
	2,  // 15: HotelsService.GetAllHotelsService:output_type -> GetAllHotelsResponse
	7,  // 16: HotelsService.AddRoomToHotelService:output_type -> Room
	7,  // 17: HotelsService.SetRoomToAvailableService:output_type -> Room
	7,  // 18: HotelsService.SetRoomToUnavailableService:output_type -> Room
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_hotels_proto_hotel_proto_init() }
func file_protos_hotels_proto_hotel_proto_init() {
	if File_protos_hotels_proto_hotel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_hotels_proto_hotel_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Hotel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllHotelsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllHotelsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetHotelByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetHotelByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableRoomsByHotelRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetAvailableRoomsByHotelResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Room); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CreateHotelRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRoomRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SetRoomToUnavailableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_hotels_proto_hotel_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*SetRoomToAvailableRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_hotels_proto_hotel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_hotels_proto_hotel_proto_goTypes,
		DependencyIndexes: file_protos_hotels_proto_hotel_proto_depIdxs,
		MessageInfos:      file_protos_hotels_proto_hotel_proto_msgTypes,
	}.Build()
	File_protos_hotels_proto_hotel_proto = out.File
	file_protos_hotels_proto_hotel_proto_rawDesc = nil
	file_protos_hotels_proto_hotel_proto_goTypes = nil
	file_protos_hotels_proto_hotel_proto_depIdxs = nil
}
