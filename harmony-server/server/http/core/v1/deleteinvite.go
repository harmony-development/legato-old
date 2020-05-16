package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteInviteData is the data for an invite delete request
type DeleteInviteData struct {
	Guild  int64 `validate:"required"`
	Invite int64 `validate:"required"`
}

// DeleteInvite is the request to delete an invite
func (h Handlers) DeleteInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data DeleteInviteData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}

	if err := h.Deps.DB.DeleteInvite(data.Invite); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted invite",
	})
}
