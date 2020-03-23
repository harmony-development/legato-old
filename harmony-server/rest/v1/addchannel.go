package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-server/db"
	"harmony-server/globals"
	"harmony-server/rest/hm"
	"net/http"
)

func AddChannel(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	guild, channelname := ctx.FormValue("guildid"), ctx.FormValue("channelname")
	if channelname == "" || guild == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameters")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channels being added, please wait a few seconds")
	}
	var channelID = randstr.Hex(16)

	if !db.IsUserOwner(guild, ctx.User.ID) {
		return echo.NewHTTPError(http.StatusUnauthorized, "insufficient permissions to add channel")
	}

	_, err := db.DBInst.Exec("INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3)", channelID, guild, channelname)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error occurred adding the guild. Do you have sufficient permission?")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients == nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "successfully added channel",
		})
	} else {
		for _, client := range globals.Guilds[guild].Clients {
			for _, conn := range client {
				conn.Send(&globals.Packet{
					Type: "addchannel",
					Data: map[string]interface{}{
						"guild":       guild,
						"channelname": guild,
						"channelid":   channelID,
					},
				})
			}
		}
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "successfully added channel",
		})
	}
}