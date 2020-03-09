package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/authentication"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// Authenticate is an API path that is meant to be triggered by a client, and initializes a session authorization
func Authenticate(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	host, session := ctx.FormValue("host"), ctx.FormValue("session")
	if host == "" || session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid auth arguments")
	}
	token, err := authentication.MakeToken(session, host)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "auth error, please try agian later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": *token,
	})
}