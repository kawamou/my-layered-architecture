package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawamou/my-layered-architecture/application"
)

// PingHandler ...
type PingHandler struct {
	s application.PingService
}

// PingHandlerRequest ...
type PingHandlerRequest struct{}

// PingHandlerResponse ...
type PingHandlerResponse struct {
	Message string `json:"message"`
}

// NewPingHandler ...
func NewPingHandler(service application.PingService) *PingHandler {
	return &PingHandler{
		s: service,
	}
}

// Handle ...
func (p *PingHandler) Handle(c *gin.Context) {
	withContext(c, func(ctx context.Context) {
		log.Println(p.s.Handle(ctx))
		response := PingHandlerResponse{
			Message: "pong",
		}
		c.JSON(http.StatusOK, response)
	})
}
