package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/hotello-gateway/genprotos/users"
	errorss "github.com/ruziba3vich/hotello-gateway/internal/items/errors"
)

func (h *HotelloHandler) RegisterUserHandler(c *gin.Context) {
	var req users.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}

	if err := h.checkCreateUserRequest(c, &req); err != nil {
		c.IndentedJSON(http.StatusConflict, errorss.ERROR{Error: err.Error()})
		return
	}

	if err := h.usersPublisher.Publish(&req, "user.registration"); err != nil { // get this into config
		h.logger.Printf("-- ERROR WHLE PUBLISHING MESSAGE -- %s\n", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, errorss.Response{Message: "You have successfully been registerd"})
}

func (h *HotelloHandler) LoginUserHandler(c *gin.Context) {
	var req users.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}

	response, err := h.service.LoginUserServer(c, &req)
	if err != nil {
		h.logger.Printf("-- ERROR HAS BEEN RETURNED FROM THE SERVER %s -- ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, errorss.Response{Message: response.Message})
}

func (h *HotelloHandler) GetUserByIdHandler(c *gin.Context) {
	var req users.GetUserByFieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}

	response, err := h.service.GetUserByIdService(c, &req)
	if err != nil {
		h.logger.Printf("-- ERROR HAS BEEN RETURNED FROM THE SERVER %s -- ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *HotelloHandler) GetUserByUsernameHandler(c *gin.Context) {
	var req users.GetUserByFieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}

	response, err := h.service.GetUserByUsernameService(c, &req)
	if err != nil {
		h.logger.Printf("-- ERROR HAS BEEN RETURNED FROM THE SERVER %s -- ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *HotelloHandler) GetUserByEmailHandler(c *gin.Context) {
	var req users.GetUserByFieldRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}

	response, err := h.service.GetUserByEmailService(c, &req)
	if err != nil {
		h.logger.Printf("-- ERROR HAS BEEN RETURNED FROM THE SERVER %s -- ", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func (h *HotelloHandler) UpdateUsernameHandler(c *gin.Context) {
	/*
		Id will be set in the header in middleware
	*/
	id := c.GetHeader("id")
	var req users.UpdateUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}
	req.Id = id
	if h.checkUsernameIsAvailable(c, &req) {
		if err := h.usersPublisher.Publish(&req, "user.registration"); err != nil { // get this into config
			h.logger.Printf("-- ERROR WHLE PUBLISHING MESSAGE -- %s\n", err.Error())
			c.IndentedJSON(http.StatusInternalServerError, errorss.ERROR{Error: err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, errorss.Response{Message: "username successfully updated"})
	}
	c.IndentedJSON(http.StatusConflict, errorss.ERROR{Error: "username is already taken"})
}

func (h *HotelloHandler) UpdatePasswordHandler(c *gin.Context) {
	id := c.GetHeader("id")
	var req users.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Printf("-- ERROR WHILE BINDING DATA  %s--\n", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errorss.ERROR{Error: err.Error()})
		return
	}
	req.Id = id

	
}

/*
	CreateUserService(ctx context.Context, in *CreateUserRequest) (*User, error)
    LoginUserService(ctx context.Context, in *LoginUserRequest) (*RawResponse, error)
    GetUserByIdService(ctx context.Context, in *GetUserByFieldRequest) (*User, error)
    GetUserByUsernameService(ctx context.Context, in *GetUserByFieldRequest) (*User, error)
    GetUserByEmailService(ctx context.Context, in *GetUserByFieldRequest) (*User, error)
    UpdateUsernameService(ctx context.Context, in *UpdateUsernameRequest) (*RawResponse, error)
    UpdatePasswordService(ctx context.Context, in *UpdatePasswordRequest) (*RawResponse, error)
    DeleteUserService(ctx context.Context, in *DeleteUserRequest) (*RawResponse, error)
    VerifyCodeService(ctx context.Context, in *VerifyCodeRequest) (*LoginUserResponse, error)
*/
