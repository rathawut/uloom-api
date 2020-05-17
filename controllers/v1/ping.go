package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPing godoc
func (h *Handler) GetPing(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Pong!",
	})
}
