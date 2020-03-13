package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// GetUser gets data for a certain user given a userid
func GetUser(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	userid := ctx.FormValue("userid")
	if userid == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "userid field required")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many get user requests, please try again in a few moments")
	}
	returnUser, err := db.GetUser(userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "user not found")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": *returnUser,
	})
}