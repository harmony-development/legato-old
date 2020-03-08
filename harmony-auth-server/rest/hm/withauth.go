package hm

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/db"
	"net/http"
)

func WithAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		session := ctx.FormValue("session")
		userid, err := db.GetUserBySession(session)
		hctx, ok := ctx.(*HarmonyContext)
		if err != nil || !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
		}
		hctx.UserID = userid
		return handler(hctx)
	}
}