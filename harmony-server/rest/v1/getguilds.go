package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

type returnGuild struct {
	Guildname string `json:"guildname"`
	Picture   string `json:"picture"`
	Owner     string `json:"owner"`
}

func GetGuilds(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "you're getting guilds too fast, please wait a moment")
	}
	res, err := db.DBInst.Query("SELECT guilds.guildid, guilds.guildname, guilds.owner, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", ctx.User.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get guilds at this time, please try again later")
	}
	var returnGuilds = make(map[string]returnGuild)
	for res.Next() {
		var guildID string
		var fetchedGuild returnGuild
		err := res.Scan(&guildID, &fetchedGuild.Guildname, &fetchedGuild.Owner, &fetchedGuild.Picture)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "we were unable to list one of the guilds, please try again later")
		}
		returnGuilds[guildID] = fetchedGuild
	}
	return ctx.JSON(http.StatusOK, returnGuilds)
}
