package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
)

type removeSessionData struct {
	Session string `validate:"required"`
	Host    string `validate:"required"`
}

// RemoveServer removes a server from a user's list
func (h Handlers) RemoveServer(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(removeSessionData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	session, exists := h.AuthManager.Sessions.GetSession(data.Session)
	if !exists {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid session")
	}
	if err := h.DB.RemoveInstance(data.Host, session.UserID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error removing server from list")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully added server to list!",
	})
}
