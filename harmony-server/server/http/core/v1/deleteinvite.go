package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteInviteData is the data for an invite delete request
type DeleteInviteData struct {
	Invite string `validate:"required"`
}

// DeleteInvite is the request to delete an invite
func (h Handlers) DeleteInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(DeleteInviteData)

	if err := h.Deps.DB.DeleteInvite(data.Invite); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted invite",
	})
}
