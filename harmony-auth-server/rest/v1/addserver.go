package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// AddServer adds a new server to a user's list
func AddServer(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	session, host := ctx.FormValue("session"), ctx.FormValue("host")
	if session == "" || host == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "session and host required")
	}
	user, err := db.GetUserBySession(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid ")
	}

	if err := db.AddServerTransaction(user.ID, host); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error adding server to list")
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully added server to list!",
	})
}