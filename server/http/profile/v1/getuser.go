package v1

import (
	"net/http"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/http/hm"

	"github.com/labstack/echo/v4"
)

// GetUser is the handler for a user info request
func (h Handlers) GetUser(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	user := *ctx.Location.User
	response := UserInfoResponse{
		UserName:   user.Username,
		UserAvatar: user.Avatar.String,
		UserStatus: profilev1.UserStatus(user.Status),
	}
	return ctx.JSON(http.StatusOK, response)
}
