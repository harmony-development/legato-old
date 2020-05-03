package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-server/server/http/hm"
	"net/http"
)

type CreateGuildData struct {
	GuildName string `validate:"requried"`
}

func (h Handlers) CreateGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	var data CreateGuildData
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "you're creating too many guilds, please try again in a minute or two")
	}
	guildID := randstr.Hex(16)
	if err := h.Deps.DB.AddGuild(guildID, ctx.UserID, data.GuildName, ""); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": guildID,
	})
}