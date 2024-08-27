package usersservice

import (
	"context"
	"log"

	"github.com/ruziba3vich/hotello-users/genprotos/users"
	"github.com/ruziba3vich/hotello-users/internal/items/storage"
)

type UsersService struct {
	storage *storage.UsersStorage
	logger  *log.Logger
	users.UnimplementedUsersServiceServer
}

func NewUsersService(storage *storage.UsersStorage, logger *log.Logger) users.UsersServiceServer {
	return &UsersService{
		storage: storage,
		logger:  logger,
	}
}

func (s *UsersService) CreateUserService(ctx context.Context, req *users.CreateUserRequest) (*users.User, error) {
	s.logger.Printf("CreateUserService called with request: %+v\n", req)
	return s.storage.RegisterUser(ctx, req)
}

func (s *UsersService) LoginUserService(ctx context.Context, req *users.LoginUserRequest) (*users.RawResponse, error) {
	s.logger.Printf("LoginUserService called with request: %+v\n", req)
	return s.storage.LoginUser(ctx, req)
}

func (s *UsersService) GetUserByIdService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	s.logger.Printf("GetUserByIdService called with request: %+v\n", req)
	return s.storage.GetUserById(ctx, req)
}

func (s *UsersService) GetUserByUsernameService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	s.logger.Printf("GetUserByUsernameService called with request: %+v\n", req)
	return s.storage.GetUserByUsername(ctx, req)
}

func (s *UsersService) GetUserByEmailService(ctx context.Context, req *users.GetUserByFieldRequest) (*users.User, error) {
	s.logger.Printf("GetUserByEmailService called with request: %+v\n", req)
	return s.storage.GetUserByEmail(ctx, req)
}

func (s *UsersService) UpdateUsernameService(ctx context.Context, req *users.UpdateUsernameRequest) (*users.RawResponse, error) {
	s.logger.Printf("UpdateUsernameService called with request: %+v\n", req)
	return s.storage.UpdateUsername(ctx, req)
}

func (s *UsersService) UpdatePasswordService(ctx context.Context, req *users.UpdatePasswordRequest) (*users.RawResponse, error) {
	s.logger.Printf("UpdatePasswordService called with request: %+v\n", req)
	return s.storage.UpdatePasswordService(ctx, req)
}

func (s *UsersService) DeleteUserService(ctx context.Context, req *users.DeleteUserRequest) (*users.RawResponse, error) {
	s.logger.Printf("DeleteUserService called with request: %+v\n", req)
	return s.storage.SoftDeleteUser(ctx, req)
}

func (s *UsersService) VerifyCodeService(ctx context.Context, req *users.VerifyCodeRequest) (*users.LoginUserResponse, error) {
	s.logger.Printf("VerifyCodeService called with request: %+v\n", req)
	return s.storage.VerifyCodeFromEmail(ctx, req)
}
