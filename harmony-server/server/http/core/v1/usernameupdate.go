package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"harmony-server/server/http/socket/client"
)

type UsernameUpdateData struct {
	Username string `validate:"required"`
}

func (h Handlers) UsernameUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(UsernameUpdateData)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	for c := range h.Deps.State.UserUpdateListeners {
		c.Send(&client.OutPacket{
			Type: UserUpdateEventType,
			Data: UserUpdateEvent{
				UserID:   ctx.UserID,
				Username: data.Username,
			},
		})
	}
	if err := h.Deps.DB.UpdateUsername(ctx.UserID, data.Username); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.NoContent(http.StatusOK)
}
