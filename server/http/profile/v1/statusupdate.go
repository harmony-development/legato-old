package v1

import (
	"net/http"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/util"

	"github.com/labstack/echo/v4"
)

type StatusUpdateData struct {
	Status profilev1.UserStatus `validate:"required"`
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
