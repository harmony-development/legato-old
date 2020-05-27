package hm

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/responses"
)

func (m *Middlewares) Schema(schema interface{}) echo.MiddlewareFunc {
	schemaType := reflect.TypeOf(schema)
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			data := reflect.New(schemaType).Interface()
			ctx.Data = data
			if err := ctx.Bind(data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			if err := ctx.Validate(data); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			ctx.Data = reflect.ValueOf(data).Elem().Interface()
			return handler(ctx)
		}
	}
}
