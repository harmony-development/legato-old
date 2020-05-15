package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
)

// CreateInviteData is the data that a CreateInvite request has
type CreateInviteData struct {
	Guild string `validate:"required"`
}

// CreateInvite : Create an invite for a given guild
func (h Handlers) CreateInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(CreateInviteData)


	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invites created, try again in a few seconds")
	}
	var inviteID = randstr.Hex(5)
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID)
	if h.Deps.State.Guilds[data.Guild] == nil || !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to create an invite")
	}
	if err := h.Deps.DB.AddInvite(inviteID, data.Guild); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create invite")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"invite": inviteID,
	})
}
