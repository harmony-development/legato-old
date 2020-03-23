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
	ctx := c.(hm.HarmonyContext)
	host, identity, session := ctx.FormValue("host"), ctx.FormValue("identity"), ctx.FormValue("session")
	if host == "" || session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid auth arguments")
	}
	user, err := authentication.GetUserBySession(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	serverSession := randstr.Hex(16)
	token, err := authentication.MakeServerSessionToken(serverSession, identity)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating server session")
	}
	resp, err := types.Server{IP: host}.SendSession(token, user) // IMPORTANT : do not ever give the instance a user session!
	if err != nil || resp.StatusCode != http.StatusOK {
		golog.Warn(resp.StatusCode)
		return echo.NewHTTPError(http.StatusInternalServerError, "error sending session")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": serverSession,
	})
}
