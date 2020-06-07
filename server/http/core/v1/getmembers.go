package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/util"

	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMembers lists the members in a guild
func (h Handlers) GetMembers(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)

	inGuild, err := h.Deps.DB.UserInGuild(ctx.UserID, *ctx.Location.GuildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !inGuild {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	res, err := h.Deps.DB.MembersInGuild(*ctx.Location.GuildID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
	}

	return ctx.JSON(http.StatusOK, MemberListResponse{
		Members: util.U64TSA(res),
	})
}
