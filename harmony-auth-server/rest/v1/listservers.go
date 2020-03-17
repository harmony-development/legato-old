package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/authentication"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// ListServers returns an array of servers saved in the DB
func ListServers(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	session := ctx.FormValue("session")
	user, err := authentication.GetUserBySession(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	servers, err := db.ListServersTransaction(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list servers, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"servers": servers,
	})
}