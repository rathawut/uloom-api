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
	profile, err := h.FacebookProvider.GetProfile(c.QueryParam("code"))
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, profile)
}
