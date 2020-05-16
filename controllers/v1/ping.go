package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Message godoc
type Message struct {
	Message string
}

// PingPong godoc
func (h *Handler) PingPong(c echo.Context) error {
	message := &Message{
		Message: "Pong!",
	}
	return c.JSON(http.StatusOK, message)
}
