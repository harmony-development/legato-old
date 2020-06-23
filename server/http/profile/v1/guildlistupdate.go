package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GuildListUpdateData struct {
	NewList string `validate:"required"`
}

func (h Handlers) GuildListUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(GuildListUpdateData)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	if err := h.Deps.DB.UpdateGuildList(ctx.UserID, data.NewList); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.NoContent(http.StatusOK)
}
