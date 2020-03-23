package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/authentication"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// UsernameUpdate handles requests from the client to update names
func UsernameUpdate(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	session, username := ctx.FormValue("session"), ctx.FormValue("username")
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many username updates, please try again later")
	}
	user, err := authentication.GetUserBySession(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	_, err = db.DB.Exec("UPDATE users SET username=$1 WHERE userid=$2", username, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update username, please try again later")
	}
	servers, err := db.ListServersTransaction(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update, please try again later")
	}

	for _, server := range servers {
		go server.SendUsernameUpdate(user.ID, username)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated username",
	})
}
