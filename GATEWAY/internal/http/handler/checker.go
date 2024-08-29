package handler

import (
	"context"
	"errors"

	"github.com/ruziba3vich/hotello-gateway/genprotos/users"
)

func (h *HotelloHandler) checkCreateUserRequest(c context.Context, req *users.CreateUserRequest) error {
	if len(req.Email) == 0 {
		return errors.New("email has not been provided")
	}
	if len(req.Username) == 0 {
		return errors.New("username has not been provided")
	}
	if len(req.Password) < 8 {
		return errors.New("password should be at least 8 chars long")
	}
	if len(req.FullName) == 0 {
		return errors.New("fullname has not been provided")
	}
	if h.service.CheckUserByEmail(c, &users.GetUserByFieldRequest{
		FieldName: "email",
		Value:     req.Email,
	}) {
		return errors.New("this email is already registered in the server")
	}

	if h.service.CheckUserByUsername(c, &users.GetUserByFieldRequest{
		FieldName: "username",
		Value:     req.Username,
	}) {
		return errors.New("this username is already registered in the server")
	}
	return nil
}

func (h *HotelloHandler) checkUsernameIsAvailable(c context.Context, req *users.UpdateUsernameRequest) bool {
	user, _ := h.service.GetUserByUsernameService(c, &users.GetUserByFieldRequest{
		FieldName: "username",
		Value:     req.NewUsername,
	})
	return user == nil
}

func (h *HotelloHandler) checkUpdatePasswordRequest(c context.Context, req *users.UpdatePasswordRequest) error {
	user
}
