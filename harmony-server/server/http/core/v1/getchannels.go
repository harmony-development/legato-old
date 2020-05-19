package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetChannels gets the channels for a given guild
func (h Handlers) GetChannels(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	res, err := h.Deps.DB.ChannelsForGuild(*ctx.Location.GuildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list channels, please try again later")
	}
	return ctx.JSON(http.StatusOK, res)
}
