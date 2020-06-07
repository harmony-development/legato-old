package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// CreateInviteData is the data that a CreateInvite request has
type CreateInviteData struct {
	Name string `json:"invite_name" validate:"required"`
	Uses *int32 `json:"invite_uses"`
}

func (c CreateInviteData) uses() int32 {
	if c.Uses == nil {
		return -1
	}
	return int32(*c.Uses)
}

// CreateInvite : Create an invite for a given guild
func (h Handlers) CreateInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(CreateInviteData)

	invite, err := h.Deps.DB.CreateInvite(*ctx.Location.GuildID, data.uses(), data.Name)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, InviteCreateResponse{
		ID:   invite.InviteID,
		Name: data.Name,
		Uses: invite.PossibleUses.Int32,
	})
}
