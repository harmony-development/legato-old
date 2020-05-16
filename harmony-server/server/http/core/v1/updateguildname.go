package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateGuildNameData is the data for a guild name update request
type UpdateGuildNameData struct {
	Guild int64  `validate:"required"`
	Name  string `validate:"name"`
}

// UpdateGuildName updates the guild name
func (h Handlers) UpdateGuildName(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data UpdateGuildNameData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild name updates")
	}
	owner, err := h.Deps.DB.GetOwner(data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to check perms")
	}
	if owner != ctx.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to change guild name")
	}
	if err := h.Deps.DB.UpdateGuildName(data.Guild, data.Name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to update guild name, please try again later")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] != nil {
		h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
			Type: "GuildNameUpdate",
			Data: map[string]interface{}{
				"guild": data.Guild,
				"name":  data.Name,
			},
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated guild name",
	})
}
