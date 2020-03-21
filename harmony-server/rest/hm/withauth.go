package hm

import (
	"github.com/labstack/echo/v4"
	"harmony-server/authentication"
	"net/http"
)

func WithAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.FormValue("token")
		user, err := authentication.GetUserBySession(token)
		hctx, ok := ctx.(HarmonyContext)
		if err != nil || !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		hctx.User = user
		return handler(hctx)
	}
}