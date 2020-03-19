package v1

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-auth-server/authentication"
	"harmony-auth-server/rest/hm"
	"harmony-auth-server/types"
	"net/http"
)

// Authenticate takes in a user session and generates an instance-specific session
func Authenticate(c echo.Context) error {
	golog.Warn("hi")
	ctx := c.(hm.HarmonyContext)
	golog.Warn("hi1")
	host, session := ctx.FormValue("host"), ctx.FormValue("session")
	golog.Warn("hi2")
	if host == "" || session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid auth arguments")
	}
	golog.Warn("hi3")
	if !authentication.ValidateSession(session) {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	golog.Warn("hi4")
	serverSession := randstr.Hex(16)
	s := &types.Server{IP:host}
	identity, err := s.GetIdentity()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting server identity")
	}
	token, err := authentication.MakeServerSessionToken(serverSession, *identity)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating auth token")
	}
	types.Server{IP:host}.SendSession(token) // IMPORTANT : do not ever give the instance a user session!
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": serverSession,
	})
}