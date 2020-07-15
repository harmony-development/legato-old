package v1

import (
	"net/http"
	"strconv"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
)

type RemoveGuildData struct {
	GuildID    string `valdiate:"required"`
	Homeserver string `validate:"required"`
}

func (h Handlers) RemoveGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	data := ctx.Data.(RemoveGuildData)
	guildID, err := strconv.ParseUint(data.GuildID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, responses.InvalidRequest)
	}
	if err := h.DB.RemoveGuildFromList(ctx.UserID, guildID, data.Homeserver); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.NoContent(http.StatusOK)
}
