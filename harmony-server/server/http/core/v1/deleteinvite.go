package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// DeleteInviteData is the data for an invite delete request
type DeleteInviteData struct {
	Guild  string `validate:"required"`
	Invite string `validate:"required"`
}

// DeleteInvite is the request to delete an invite
func (h Handlers) DeleteInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(DeleteInviteData)

	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many message deletions, please try again later")
	}
	isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to check permissions")
	}
	if !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete invite")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invite deletions, please wait a few moments")
	}
	if err := h.Deps.DB.DeleteInvite(data.Invite, data.Guild); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted invite",
	})
}
