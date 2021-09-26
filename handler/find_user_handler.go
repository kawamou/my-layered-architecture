package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawamou/my-layered-architecture/application"
)

type FindUserHandler struct {
	s application.FindUserService
}

func NewFindUserHandler(s application.FindUserService) *FindUserHandler {
	return &FindUserHandler{
		s: s,
	}
}

type FindUserHandlerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *FindUserHandler) Handle(c *gin.Context) {
	withContext(c, func(ctx context.Context) {
		id := c.Param("id")
		output, err := u.s.Handle(ctx, application.FindUserServiceInput{ID: id})
		if err != nil {
			log.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}
		c.JSON(http.StatusOK, output)
	})
}
