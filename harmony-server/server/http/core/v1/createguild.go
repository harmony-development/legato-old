package v1

import (
	"net/http"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"harmony-server/util"

	"github.com/labstack/echo/v4"
)

// CreateGuildData is the data for a guild creation request
type CreateGuildData struct {
	GuildName string `json:"guild_name" validate:"required"`
}

// CreateGuild creates a guild for a user
func (h Handlers) CreateGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(CreateGuildData)
	guildID, err := h.Deps.Sonyflake.NextID()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	guild, err := h.Deps.DB.CreateGuild(ctx.UserID, guildID, data.GuildName, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, GuildCreateResponse{
		GuildID: util.U64TS(guild.GuildID),
	})
}
