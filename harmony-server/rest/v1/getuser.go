package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func GetUser(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	target := ctx.FormValue("target")
	var username string
	var avatar string
	err := harmonydb.DBInst.QueryRow("SELECT username, avatar FROM users WHERE id=$1", target).Scan(&username, &avatar)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get user profile, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"username": username,
		"avatar": avatar,
	})
}
