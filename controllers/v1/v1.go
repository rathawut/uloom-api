package v1

import (
	"github.com/labstack/echo/v4"
)

// Handler godoc
type Handler struct{}

// NewHandler godoc
func NewHandler() *Handler {
	return &Handler{}
}

// Register godoc
func (h *Handler) Register(g *echo.Group) {
	ping := g.Group("/ping")
	ping.GET("", h.PingPong)
}
