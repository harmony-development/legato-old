package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

func ListServers(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	servers, err := db.ListServersTransaction(*ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list servers, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"servers": servers,
	})
}