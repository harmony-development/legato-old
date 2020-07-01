package protocol

import (
	"database/sql"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GuildLeaveData struct {
	Nonce   string `validate:"required"`
	GuildID uint64 `validate:"required"`
}

func (h API) GuildLeave(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(GuildLeaveData)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	nonceInfo, err := h.Deps.DB.GetNonceInfo(data.Nonce)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, responses.NonceNotFound)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
	}
	if err := h.Deps.DB.RemoveGuildFromList(nonceInfo.UserID, data.GuildID, nonceInfo.HomeServer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}

	return ctx.NoContent(http.StatusOK)
}
