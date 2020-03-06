package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func UsernameUpdate(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	username := ctx.FormValue("username")
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many username updates, please try again later")
	}
	_, err := harmonydb.DBInst.Exec("UPDATE users SET username=$1 WHERE id=$2", username, *ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update username, please try again later")
	}
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", *ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update")
	}
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update")
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				for _, conn := range client {
					conn.Send(&globals.Packet{
						Type: "usernameupdate",
						Data: map[string]interface{}{
							"userid":   *ctx.UserID,
							"username": username,
						},
					})
				}
			}
		}
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated username",
	})
}