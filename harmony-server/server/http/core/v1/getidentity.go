package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetIdentity returns the identity of this instance (used to prevent authentication vulnerabilities
func (h Handlers) GetIdentity(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"identity": h.Deps.Config.Server.Identity,
	})
}
