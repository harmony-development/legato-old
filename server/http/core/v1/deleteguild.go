package v1

import (
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/util"

	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteGuild is the handler for a delete guild request
func (h Handlers) DeleteGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	if h.Deps.DB.DeleteGuild(*ctx.Location.GuildID) != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting guild, please try again later or report to the administrator")
	}

	if h.Deps.State.Guilds[*ctx.Location.GuildID] != nil {
		h.Deps.State.GuildsLock.Lock()
		h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
			Type: GuildDeleteEventType,
			Data: GuildDeleteEvent{
				GuildID: util.U64TS(*ctx.Location.GuildID),
			},
		})
		delete(h.Deps.State.Guilds, *ctx.Location.GuildID)
		h.Deps.State.GuildsLock.RUnlock()
	}

	return nil
}
