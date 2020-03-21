package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

func DeleteChannel(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guild, channelid := ctx.FormValue("guild"), ctx.FormValue("channelid")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != ctx.User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete channel")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channel deletions, please try again in a few seconds")
	}
	err := db.DeleteChannelTransaction(guild, channelid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting channel, please try again later")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deletechannel",
				Data: map[string]interface{}{
					"guild":     guild,
					"channelid": channelid,
				},
			})
		}
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted channel",
	})
}
