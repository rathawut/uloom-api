package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAuthFacebook godoc
func (h *Handler) GetAuthFacebook(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, h.FacebookProvider.BuildURL())
}

// GetAuthFacebookCallback godoc
func (h *Handler) GetAuthFacebookCallback(c echo.Context) error {
	return c.JSON(http.StatusOK, c.QueryParams())
}
