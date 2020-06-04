package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/util"

	"net/http"

	"github.com/labstack/echo/v4"
)

// GetGuild is the handler for a guild info request
func (h Handlers) GetGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	guild, err := h.Deps.DB.GetGuildByID(*ctx.Location.GuildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, GuildInfoResponse{
		GuildName:    guild.GuildName,
		GuildOwner:   util.U64TS(guild.OwnerID),
		GuildPicture: guild.PictureUrl,
	})
}
