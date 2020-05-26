package v1

import (
	"net/http"

	"harmony-server/server/http/hm"

	"github.com/labstack/echo/v4"
)

// CreateGuildData is the data for a guild creation request
type CreateGuildData struct {
	GuildName string `validate:"required"`
}

// CreateGuild creates a guild for a user
func (h Handlers) CreateGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(CreateGuildData)

	guild, err := h.Deps.DB.CreateGuild(ctx.UserID, data.GuildName, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]uint64{
		"guild": guild.GuildID,
	})
}
