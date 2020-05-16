package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// CreateInviteData is the data that a CreateInvite request has
type CreateInviteData struct {
	Guild int64  `validate:"required"`
	Name  string `validate:"required"`
}

// CreateInvite : Create an invite for a given guild
func (h Handlers) CreateInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data CreateInviteData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}

	invite, err := h.Deps.DB.CreateInvite(data.Guild, -1, data.Name)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]int64{
		"invite": invite.InviteID,
	})
}
