package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetChannelsData is the data for GetChannels
type GetChannelsData struct {
	Guild int64 `validate:"required"`
}

// GetChannels gets the channels for a given guild
func (h Handlers) GetChannels(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data GetChannelsData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	res, err := h.Deps.DB.ChannelsForGuild(data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list channels, please try again later")
	}
	return ctx.JSON(http.StatusOK, res)
}
