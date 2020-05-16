package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetGuilds lists the guilds a user is in
func (h Handlers) GetGuilds(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	res, err := h.Deps.DB.GuildsForUser(ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}
