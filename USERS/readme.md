# UsersService Microservice Overview

The `UsersService` microservice handles various user-related operations within our system. This service utilizes both gRPC and Kafka for communication to ensure efficient integration within our microservice architecture. Below is an overview of the available RPC services, their communication methods, and their specific roles.

## RPC Services

### 1. CreateUserService

- **Method**: `rpc CreateUserService(CreateUserRequest) returns (User);`
- **Communication**: Kafka and gRPC
- **Description**: Creates a new user. Receives user details in the `CreateUserRequest`. The request is first handled asynchronously via Kafka for background processing and then returns the created `User` object using gRPC for direct response.

### 2. LoginUserService

- **Method**: `rpc LoginUserService(LoginUserRequest) returns (RawResponse);`
- **Communication**: gRPC
- **Description**: Handles user login by verifying credentials and returning a response indicating success or failure. Uses gRPC for synchronous requests, providing immediate validation and response.

### 3. GetUserByIdService

- **Method**: `rpc GetUserByIdService(GetUserByFieldRequest) returns (User);`
- **Communication**: gRPC
- **Description**: Retrieves user details based on user ID. Invoked using gRPC for efficient querying and retrieval of user information.

### 4. GetUserByUsernameService

- **Method**: `rpc GetUserByUsernameService(GetUserByFieldRequest) returns (User);`
- **Communication**: gRPC
- **Description**: Fetches user details using the username. Uses gRPC for efficient querying by username.

### 5. GetUserByEmailService

- **Method**: `rpc GetUserByEmailService(GetUserByFieldRequest) returns (User);`
- **Communication**: gRPC
- **Description**: Retrieves user information based on email. Uses gRPC for synchronous communication, ensuring quick lookups.

### 6. UpdateUsernameService

- **Method**: `rpc UpdateUsernameService(UpdateUsernameRequest) returns (RawResponse);`
- **Communication**: Kafka
- **Description**: Updates the username of an existing user. Handled via Kafka for asynchronous processing, allowing updates to be handled in the background.

### 7. UpdatePasswordService

- **Method**: `rpc UpdatePasswordService(UpdatePasswordRequest) returns (RawResponse);`
- **Communication**: Kafka
- **Description**: Handles password updates. Uses Kafka for asynchronous processing, beneficial for operations involving additional verification or background tasks.

### 8. DeleteUserService

- **Method**: `rpc DeleteUserService(DeleteUserRequest) returns (RawResponse);`
- **Communication**: Kafka
- **Description**: Deletes a user from the system. Performed asynchronously via Kafka to handle data consistency checks and background processes.

### 9. VerifyCodeService

- **Method**: `rpc VerifyCodeService(VerifyCodeRequest) returns (LoginUserResponse);`
- **Communication**: gRPC
- **Description**: Verifies a code sent to the user’s email during login or other verification processes. Uses gRPC for synchronous communication to validate the code and respond with a login result.

## Summary

- **gRPC Services**: `LoginUserService`, `GetUserByIdService`, `GetUserByUsernameService`, `GetUserByEmailService`, and `VerifyCodeService` are used for direct, synchronous communication, providing immediate responses and facilitating real-time interactions.
- **Kafka Services**: `CreateUserService`, `UpdateUsernameService`, `UpdatePasswordService`, and `DeleteUserService` utilize Kafka for asynchronous processing, enabling efficient background operations and reducing the load on the primary service.

By leveraging both gRPC and Kafka, the `UsersService` ensures robust and scalable handling of user-related operations, catering to both real-time and background processing needs.
