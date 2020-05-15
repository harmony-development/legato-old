package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMembersData is the data for a member list request
type GetMembersData struct {
	Guild string `validate:"required"`
}

// GetMembers lists the members in a guild
func (h Handlers) GetMembers(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(GetMembersData)
	h.Deps.State.GuildsLock.RLock()
	if h.Deps.State.Guilds[data.Guild] == nil || h.Deps.State.Guilds[data.Guild].Clients[ctx.UserID] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to list members")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many member listing requests, please try again later")
	}
	res, err := h.Deps.DB.Query("SELECT userid FROM guildmembers WHERE guildid=$1", data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
	}
	var returnMembers []string
	for res.Next() {
		var userid string
		err = res.Scan(&userid)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
		}
		returnMembers = append(returnMembers, userid)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"members": returnMembers,
	})
}