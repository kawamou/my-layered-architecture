package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawamou/my-layered-architecture/application"
)

// CreateUserHandler ...
type CreateUserHandler struct {
	s application.CreateUserService
}

type CreateUserHandlerRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// NewCreateUserHandler ...
func NewCreateUserHandler(service application.CreateUserService) *CreateUserHandler {
	return &CreateUserHandler{
		s: service,
	}
}

// Handle ...
func (u *CreateUserHandler) Handle(c *gin.Context) {
	withContext(c, func(ctx context.Context) {
		var req CreateUserHandlerRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, "bad request")
			return
		}
		if err := u.s.Handle(ctx, application.CreateUserServiceInput{
			ID:   req.ID,
			Name: req.Name,
			Age:  req.Age}); err != nil {

		}
		c.JSON(http.StatusOK, "ok")
	})
}
