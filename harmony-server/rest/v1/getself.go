package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)


func GetSelf(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	var username string
	var avatar string
	err := harmonydb.DBInst.QueryRow("SELECT username, avatar FROM users WHERE id=$1", ctx.UserID).Scan(&username, &avatar)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get your profile, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"userid": *ctx.UserID,
		"username": username,
		"avatar": avatar,
	})
}