package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/authentication"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

func Verify(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	token := ctx.FormValue("token")
	if token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing token")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many token verification requests, please try again later")
	}
	session, host, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if ctx.RealIP() != *host { // CRITICAL : if the token doesn't have the same host as the server requesting, then it's INVALID!
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	user, err := db.GetUserBySession(*session)
	if err != nil { // if session does not exist
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}