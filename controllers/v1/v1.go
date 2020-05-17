package v1

import (
	"uloom-api/providers"

	"github.com/labstack/echo/v4"
)

// Handler godoc
type Handler struct {
	FacebookProvider *providers.FacebookProvider
}

// NewHandler godoc
func NewHandler(facebookProvider *providers.FacebookProvider) *Handler {
	return &Handler{
		FacebookProvider: facebookProvider,
	}
}

// Register godoc
func (h *Handler) Register(g *echo.Group) {
	pingGroup := g.Group("/ping")
	pingGroup.GET("", h.GetPing)

	authGroup := g.Group("/auth")
	authGroup.GET("/facebook", h.GetAuthFacebook)
	authGroup.GET("/facebook/callback", h.GetAuthFacebookCallback)
}
