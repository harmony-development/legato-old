package v1

import (
	"net/http"

	"harmony-server/server/http/hm"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// LeaveGuild unjoins a user from a guild
func (h Handlers) LeaveGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)

	if isOwner, err := h.Deps.DB.IsOwner(*ctx.Location.GuildID, ctx.UserID); err != nil || isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "you cannot leave a guild you own")
	}
	if err := h.Deps.DB.DeleteMember(*ctx.Location.GuildID, ctx.UserID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to leave guild, this issue has been reported")
	}
	h.Deps.State.GuildsLock.Lock()
	delete(h.Deps.State.Guilds[*ctx.Location.GuildID].Clients, ctx.UserID)
	h.Deps.State.GuildsLock.Unlock()
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully left guild",
	})
}
