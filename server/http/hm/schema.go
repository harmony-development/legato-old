package hm

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"

	"github.com/harmony-development/legato/server/responses"
)

func (m *Middlewares) Schema(schema interface{}) echo.MiddlewareFunc {
	schemaType := reflect.TypeOf(schema)

	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			data := reflect.New(schemaType).Interface()
			ctx.Data = data
			if err := ctx.Bind(data); err != nil {
				if m.Config.Server.Policies.Debug.RespondWithErrors {
					return echo.NewHTTPError(http.StatusBadRequest, err)
				}
				return echo.NewHTTPError(http.StatusBadRequest, responses.BadRequest)
			}
			ctx.Data = reflect.ValueOf(data).Elem().Interface()
			return handler(ctx)
		}
	}
}
