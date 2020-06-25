package v1

import (
	"harmony-server/server/db"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"harmony-server/server/http/socket/client"
	"harmony-server/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusUpdateData struct {
	Status db.UserStatus `validate:"required"`
}

func (h Handlers) StatusUpdate(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(StatusUpdateData)
	if err := h.Deps.DB.SetStatus(ctx.UserID, data.Status); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	for c := range h.Deps.State.UserUpdateListeners {
		c.Send(&client.OutPacket{
			Type: UserUpdateEventType,
			Data: StatusUpdateEvent{
				UserID: util.U64TS(ctx.UserID),
				Status: data.Status,
			},
		})
	}
	return ctx.NoContent(http.StatusOK)
}
