package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteInvite is the request to delete an invite
func (h Handlers) DeleteInvite(c echo.Context) error {
	if err := h.Deps.DB.DeleteInvite(c.Param("invite_id")); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete invite, please try again later")
	}
	return nil
}
