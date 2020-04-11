package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
)

type authenticateData struct {
	Host     string `validate:"required"`
	Identity string `validate:"required"`
	Session  string `validate:"required"`
}

// Authenticate takes in a user session and generates an instance-specific session
func (h Handlers) Authenticate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(authenticateData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	instanceSession := randstr.Hex(h.Config.Server.SessionLength)
	token, err := h.AuthManager.MakeServerToken(instanceSession, data.Identity)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating server session")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
