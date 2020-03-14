package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

// GetIdentity returns the identity of this instance (used to prevent authentication vulnerabilities
func GetIdentity(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"identity": os.Getenv("HARMONY_IDENTITY"),
	})
}