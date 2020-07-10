package v1

import (
	"net/http"
	"strconv"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
	"github.com/labstack/echo/v4"
)

type MoveGuildData struct {
	TargetGuild           string `validate:"required"`
	TargetHomeserver      string `validate:"required"`
	BeforeGuild           string `validate:"required"`
	BeforeGuildHomeserver string `validate:"required"`
	AfterGuild            string `validate:"required"`
	AfterGuildHomeserver  string `validate:"required"`
}

func (h Handlers) MoveGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	data := ctx.Data.(MoveGuildData)
	targetGuild, err := strconv.ParseUint(data.TargetGuild, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
	}
	beforeGuild, err := strconv.ParseUint(data.BeforeGuild, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
	}
	afterGuild, err := strconv.ParseUint(data.AfterGuild, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
	}

	if err := h.DB.MoveGuild(ctx.UserID, targetGuild, data.TargetHomeserver, beforeGuild, afterGuild, data.BeforeGuildHomeserver, data.AfterGuildHomeserver); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	return ctx.NoContent(http.StatusOK)
}
