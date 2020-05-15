package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

// AddChannelData represents data received from client on AddChannel
type AddChannelData struct {
	Guild       string `validate:"required"`
	ChannelName string `validate:"required"`
}

// AddChannel is a request to add a channel to a guild
func (h Handlers) AddChannel(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(AddChannelData)


	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channels being added, please wait a few seconds")
	}
	var channelID = randstr.Hex(16)
	owner, err := h.Deps.DB.GetOwner(data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to verify ownership, please try again later")
	}
	if !(*owner == ctx.UserID) {
		return echo.NewHTTPError(http.StatusUnauthorized, "insufficient permissions to add channel")
	}
	if err := h.Deps.DB.AddChannelToGuild(channelID, data.Guild, data.ChannelName); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil || h.Deps.State.Guilds[data.Guild].Clients == nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "successfully added channel",
		})
	} else {
		h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
			Type: "AddChannel",
			Data: map[string]interface{}{
				"guild":       data.Guild,
				"channelName": data.ChannelName,
				"channelID":   channelID,
			},
		})
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "successfully added channel",
		})
	}
}
