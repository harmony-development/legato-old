package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
)

type ConnectData struct {
	Target string `validate:"required"`
}

func (h Handlers) Connect(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	user, err := h.Deps.DB.GetUserByID(ctx.UserID)
	if err != nil {
		h.Deps.Logger.Exception(err)
		return echo.NewHTTPError(http.StatusInternalServerError, responses.InvalidRequest)
	}
	data := ctx.Data.(ConnectData)
	token, err := h.Deps.AuthManager.MakeAuthToken(ctx.UserID, data.Target, user.Username, user.Avatar.String)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
