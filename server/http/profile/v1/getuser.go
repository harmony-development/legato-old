package v1

import (
	"harmony-server/server/db"
	"harmony-server/server/http/hm"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser is the handler for a user info request
func (h Handlers) GetUser(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	user := *ctx.Location.User
	response := UserInfoResponse{
		UserName:   user.Username,
		UserAvatar: user.Avatar.String,
		UserStatus: db.UserStatus(user.Status),
	}

	if ctx.UserID == user.UserID {
		response.GuildList = user.Guildlist
	}
	return ctx.JSON(http.StatusOK, response)
}
