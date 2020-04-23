package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/server/http/hm"
	"net/http"
)

type listServersData struct {
	Session string `validate:"required"`
}

// ListInstances returns an array of servers saved in the DB
func (h Handlers) ListInstances(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(listServersData)

	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	session, exists := h.AuthManager.Sessions.GetSession(data.Session)

	if !exists {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}

	if !ctx.Limiter.Reserve().OK() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many instance listing requests")
	}

	servers, err := h.DB.GetInstanceList(session.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list servers, please try again later")
	}
	ctx.Limiter.Allow()
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"servers": servers,
	})
}
