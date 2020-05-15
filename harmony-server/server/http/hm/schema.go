package hm

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/responses"
)

func (m *Middlewares) Schema(schema interface{}) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			if err := ctx.Bind(schema); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			if err := ctx.Validate(schema); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			ctx.Data = schema
		}
	}
}