package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
)

type usernameUpdateData struct {
	APIVersion string `validate:"required"`
	Session    string `validate:"required"`
	Username   string `validate:"required"`
}

// UsernameUpdate handles requests from the client to update names
func (h Handlers) UsernameUpdate(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(usernameUpdateData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many username updates, please try again later")
	}

	session, exists := h.AuthManager.Sessions.GetSession(data.Session)
	if !exists {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
	}
	_, err := h.DB.Exec("UPDATE users SET username=$1 WHERE userid=$2", data.Username, session.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update username, please try again later")
	}

	servers, err := h.DB.GetInstanceList(session.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to broadcast username update, please try again later")
	}

	for _, server := range servers {
		go server.SendUsernameUpdate(session.UserID, data.Username, data.APIVersion)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated username",
	})
}
