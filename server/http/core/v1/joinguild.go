package v1

import (
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/util"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// JoinGuildData is the data for a guild join request
type JoinGuildData struct {
	InviteCode string `json:"invite_id" validate:"required"`
}

// JoinGuild is the request to join a guild
func (h Handlers) JoinGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(JoinGuildData)
	guildID, err := h.Deps.DB.ResolveGuildID(data.InviteCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "error joining guild, invite code may not exist")
	}
	if err := h.Deps.DB.AddMemberToGuild(ctx.UserID, guildID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to join guild, please try again later")
	}
	if err := h.Deps.DB.IncrementInvite(data.InviteCode); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error updating invite counter")
	}
	return ctx.JSON(http.StatusOK, JoinGuildResponse{
		GuildID: util.U64TS(guildID),
	})
}
