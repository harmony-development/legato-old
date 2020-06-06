package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"harmony-server/server/http/socket/client"
	"harmony-server/util"
)

type BioUpdateData struct {
	Bio string `validate:"required"`
}

func (h Handlers) BioUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(BioUpdateData)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	for c := range h.Deps.State.UserUpdateListeners {
		c.Send(&client.OutPacket{
			Type: UserUpdateEventType,
			Data: BioUpdateEvent{
				UserID:   util.U64TS(ctx.UserID),
				Bio: data.Bio,
			},
		})
	}
	if err := h.Deps.DB.UpdateBio(ctx.UserID, data.Bio); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.NoContent(http.StatusOK)
}
