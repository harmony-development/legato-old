package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

type returnGuild struct {
	GuildName string
	Picture   string
	Owner     string
}

// GetGuilds lists the guilds a user is in
func (h Handlers) GetGuilds(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	res, err := h.Deps.DB.Query("SELECT guilds.guildid, guilds.guildname, guilds.owner, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get guilds at this time, please try again later")
	}
	var returnGuilds = make(map[string]returnGuild)
	for res.Next() {
		var guildID string
		var fetchedGuild returnGuild
		err := res.Scan(&guildID, &fetchedGuild.GuildName, &fetchedGuild.Owner, &fetchedGuild.Picture)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "we were unable to list one of the guilds, please try again later")
		}
		returnGuilds[guildID] = fetchedGuild
	}
	return ctx.JSON(http.StatusOK, returnGuilds)
}
