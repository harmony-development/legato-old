package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Channel struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// GetChannels gets the channels for a given guild
func (h Handlers) GetChannels(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	res, err := h.Deps.DB.ChannelsForGuild(*ctx.Location.GuildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list channels, please try again later")
	}
	ret := []Channel{}
	for _, channel := range res {
		ret = append(ret, Channel{
			Name: channel.ChannelName,
			ID:   u64TS(channel.ChannelID),
		})
	}
	return ctx.JSON(http.StatusOK, ChannelListResponse{
		Channels: ret,
	})
}
