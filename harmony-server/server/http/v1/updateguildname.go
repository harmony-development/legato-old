package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/handling"
	"net/http"
)

type UpdateGuildName struct {
	Guild string `validate:"required"`
	Name string `validate:"name"`
}

func (h Handlers) UpdateGuildName(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data UpdateGuildName
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
	if *owner != ctx.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to change guild name")
	}
	if err := h.Deps.DB.UpdateGuildName(data.Guild, data.Name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to update guild name, please try again later")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] != nil {
		h.Deps.State.Guilds[data.Guild].Broadcast(&handling.OutPacket{
			Type: "GuildNameUpdate",
			Data: map[string]string{
				"guild": data.Guild,
				"name": data.Name,
			},
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated guild name",
	})
}