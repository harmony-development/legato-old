package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/server/http/hm"
	"net/http"
)

type getUserData struct {
	UserID string `validate:"required"`
}

// GetUser gets data for a certain user given a userid
func (h Handlers) GetUser(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(getUserData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many get user requests, please try again in a few moments")
	}

	returnUser, err := h.DB.GetUser(data.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "user not found")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user": *returnUser,
	})
}
