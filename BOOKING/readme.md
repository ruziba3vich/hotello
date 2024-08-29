## HotelsService Overview

The `HotelsService` is a gRPC-based service designed to manage hotels and rooms within the Hotello booking system. It provides a set of RPC methods to handle various hotel and room operations, ensuring efficient management and availability tracking. Below is a summary of the available RPC methods, including details on Kafka integration and notification handling.

### RPC Methodsservice HotelsService {
    rpc CreateHotelService              (CreateHotelRequest)                returns (Hotel);
    rpc CreateRoomService               (CreateRoomRequest)                 returns (Room);
    rpc GetAvailableRoomsByHotelService (GetAvailableRoomsByHotelRequest)   returns (GetAvailableRoomsByHotelResponse);
    rpc GetHotelByIdService             (GetHotelByIdRequest)               returns (GetHotelByIdResponse);
    rpc GetAllHotelsService             (GetAllHotelsRequest)               returns (GetAllHotelsResponse);
    rpc AddRoomToHotelService           (CreateRoomRequest)                 returns (Room);
    rpc SetRoomToAvailableService       (SetRoomToAvailableRequest)         returns (Room);
    rpc SetRoomToUnavailableService     (SetRoomToUnavailableRequest)       returns (Room);
}

- **`CreateHotelService(CreateHotelRequest) returns (Hotel)`**
  - Creates a new hotel with the details provided in the request.
  - **Request:** `CreateHotelRequest` (contains hotel details).
  - **Response:** `Hotel` (the created hotel object).
  - **Kafka Integration:** Not used.

- **`CreateRoomService(CreateRoomRequest) returns (Room)`**
  - Creates a new room with the details provided in the request.
  - **Request:** `CreateRoomRequest` (contains room details).
  - **Response:** `Room` (the created room object).
  - **Kafka Integration:** Not used.

- **`GetAvailableRoomsByHotelService(GetAvailableRoomsByHotelRequest) returns (GetAvailableRoomsByHotelResponse)`**
  - Retrieves a list of available rooms for a specific hotel.
  - **Request:** `GetAvailableRoomsByHotelRequest` (contains hotel ID).
  - **Response:** `GetAvailableRoomsByHotelResponse` (list of available rooms).
  - **Kafka Integration:** Not used.

- **`GetHotelByIdService(GetHotelByIdRequest) returns (GetHotelByIdResponse)`**
  - Fetches the details of a hotel based on its ID.
  - **Request:** `GetHotelByIdRequest` (contains hotel ID).
  - **Response:** `GetHotelByIdResponse` (hotel details).
  - **Kafka Integration:** Not used.

- **`GetAllHotelsService(GetAllHotelsRequest) returns (GetAllHotelsResponse)`**
  - Retrieves a list of all hotels.
  - **Request:** `GetAllHotelsRequest` (optional filters).
  - **Response:** `GetAllHotelsResponse` (list of all hotels).
  - **Kafka Integration:** Not used.

- **`AddRoomToHotelService(CreateRoomRequest) returns (Room)`**
  - Adds a new room to an existing hotel.
  - **Request:** `CreateRoomRequest` (contains room details and hotel ID).
  - **Response:** `Room` (the added room object).
  - **Kafka Integration:** Kafka is used to publish an event indicating that a new room has been added to a hotel. This event is consumed by other services to update their records or send notifications.

- **`SetRoomToAvailableService(SetRoomToAvailableRequest) returns (Room)`**
  - Marks a room as available for booking.
  - **Request:** `SetRoomToAvailableRequest` (contains room ID).
  - **Response:** `Room` (the updated room object).
  - **Kafka Integration:** Kafka is used to publish an event indicating that a room's status has been updated to available. Consumers of this event may trigger notifications to users about room availability.

- **`SetRoomToUnavailableService(SetRoomToUnavailableRequest) returns (Room)`**
  - Marks a room as unavailable for booking.
  - **Request:** `SetRoomToUnavailableRequest` (contains room ID).
  - **Response:** `Room` (the updated room object).
  - **Kafka Integration:** Kafka is used to publish an event indicating that a room's status has been updated to unavailable. This event can be consumed to notify users or update other system components.

### Notifications and Kafka Integration

- **Kafka Usage:** The `HotelsService` integrates with Kafka to handle real-time updates and notifications. Key events related to room status changes and new room additions are published to Kafka topics.
  
- **Notification Flow:**
  1. **Event Generation:** When a room is added or its status is changed, an event is generated and published to a Kafka topic.
  2. **Kafka Consumers:** Separate microservices or consumers subscribed to these Kafka topics process the events. They may update databases, trigger further processing, or generate notifications.
  3. **Real-Time Notifications:** Notifications are sent to users in real-time via WebSockets or other messaging channels, ensuring users receive up-to-date information about room availability and status changes.

### Summary

The `HotelsService` provides comprehensive functionality for managing hotels and their rooms, including creation, retrieval, and status updates. Kafka is utilized to ensure real-time updates and notifications are effectively communicated across the system, keeping users informed about changes in room availability and new room additions.

