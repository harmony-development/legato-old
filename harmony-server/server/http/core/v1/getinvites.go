package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

type GetInvitesData struct {
	Guild uint64 `validate:"required"`
}

func (h Handlers) GetInvites(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data GetInvitesData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}

	invites, err := h.Deps.DB.GetInvites(data.Guild)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get invites")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"invites": invites,
	})
}
