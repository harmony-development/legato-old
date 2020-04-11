package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
)

type addServerData struct {
	Session string `validate:"required"`
	Host    string `validate:"required"`
}

// AddInstance adds a new server to a user's list
func (h Handlers) AddInstance(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(addServerData)
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
	if err := h.DB.AddInstance("", data.Host, session.UserID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error adding instance to list")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully added server to list!",
	})
}
