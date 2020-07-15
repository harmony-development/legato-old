package v1

import (
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/util"

	"net/http"

	"github.com/labstack/echo/v4"
)

// GetGuilds lists the guilds a user is in
func (h Handlers) GetGuilds(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	res, err := h.Deps.DB.GuildsForUserWithData(ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, JoinedGuildsResponse{
		Guilds: func() []JoinedGuildsResponseGuild {
			var ret []JoinedGuildsResponseGuild
			for _, guild := range res {
				ret = append(ret, JoinedGuildsResponseGuild{
					GuildID:      util.U64TS(guild.GuildID),
					GuildName:    guild.GuildName,
					GuildPicture: guild.PictureUrl,
					GuildOwner:   util.U64TS(guild.OwnerID),
				})
			}
			return ret
		}(),
	})
}
