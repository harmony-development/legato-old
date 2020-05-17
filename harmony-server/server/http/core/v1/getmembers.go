package v1

import (
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMembersData is the data for a member list request
type GetMembersData struct {
	Guild uint64 `validate:"required"`
}

// GetMembers lists the members in a guild
func (h Handlers) GetMembers(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data GetMembersData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	inGuild, err := h.Deps.DB.UserInGuild(ctx.UserID, data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if !inGuild {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	res, err := h.Deps.DB.MembersInGuild(data.Guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"members": res,
	})
}
