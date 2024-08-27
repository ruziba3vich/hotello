## HotelsService Overview

The `HotelsService` is a gRPC-based service that manages hotel-related operations, including hotel creation, room management, and querying available rooms. It provides several RPC methods to handle these tasks effectively.

### RPC Methods

1. **CreateHotelService**
   - **Request:** `CreateHotelRequest`
   - **Response:** `Hotel`
   - **Description:** Creates a new hotel in the system and returns the details of the newly created hotel.

2. **CreateRoomService**
   - **Request:** `CreateRoomRequest`
   - **Response:** `Room`
   - **Description:** Creates a new room within a specified hotel and returns the details of the newly created room.

3. **GetAvailableRoomsByHotelService**
   - **Request:** `GetAvailableRoomsByHotelRequest`
   - **Response:** `GetAvailableRoomsByHotelResponse`
   - **Description:** Retrieves a list of available rooms for a specific hotel based on the provided criteria.

4. **GetHotelByIdService**
   - **Request:** `GetHotelByIdRequest`
   - **Response:** `GetHotelByIdResponse`
   - **Description:** Fetches details of a specific hotel by its unique identifier.

5. **GetAllHotelsService**
   - **Request:** `GetAllHotelsRequest`
   - **Response:** `GetAllHotelsResponse`
   - **Description:** Returns a list of all hotels stored in the system.

6. **AddRoomToHotelService**
   - **Request:** `CreateRoomRequest`
   - **Response:** `Room`
   - **Description:** Adds a new room to an existing hotel and returns the details of the added room.

7. **SetRoomToAvailableService**
   - **Request:** `SetRoomToAvailableRequest`
   - **Response:** `Room`
   - **Description:** Marks a specific room as available and returns the updated room details.

8. **SetRoomToUnavailableService**
   - **Request:** `SetRoomToUnavailableRequest`
   - **Response:** `Room`
   - **Description:** Marks a specific room as unavailable and returns the updated room details.

### Usage

This service is intended for managing hotels and rooms within a hotel booking system. The methods provided allow for comprehensive hotel management, from creating hotels and rooms to handling room availability and querying available rooms. This makes it suitable for integration into larger hotel management systems or as a standalone service in a microservices architecture.
