package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

//TODO make the channel list ordered
type returnChannel struct {
	Name string `json:"name"`
}

// GetChannelsData is the data for GetChannels
type GetChannelsData struct {
	Guild string `validate:"required"`
}

// GetChannels gets the channels for a given guild
func (h Handlers) GetChannels(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "you're getting channels too often! Please try again in a few seconds.")
	}
	data := ctx.Data.(GetChannelsData)
	res, err := h.Deps.DB.Query("SELECT channelid, channelname FROM channels WHERE guildid=$1", data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list channels, please try again later")
	}
	var returnChannels = make(map[string]returnChannel)
	for res.Next() {
		var channelID string
		var channel returnChannel
		err = res.Scan(&channelID, &channel.Name)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to get channel, please try again later")
		}
		returnChannels[channelID] = channel
	}
	return ctx.JSON(http.StatusOK, returnChannels)
}