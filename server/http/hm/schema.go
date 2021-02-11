package hm

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
	"github.com/harmony-development/legato/server/http/responses"
)

func (m *Middlewares) Schema(schema interface{}) echo.MiddlewareFunc {
	schemaType := reflect.TypeOf(schema)
	v := &validator.Validate{}

	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			data := reflect.New(schemaType).Interface()
			ctx.Data = data
			if err := ctx.Bind(data); err != nil {
				if m.Config.Server.Policies.Debug.RespondWithErrors {
					return echo.NewHTTPError(http.StatusBadRequest, err)
				}
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			if err := v.Struct(data); err != nil {
				if m.Config.Server.Policies.Debug.RespondWithErrors {
					return echo.NewHTTPError(http.StatusBadRequest, err)
				}
				return echo.NewHTTPError(http.StatusBadRequest, responses.InvalidRequest)
			}
			ctx.Data = reflect.ValueOf(data).Elem().Interface()
			return handler(ctx)
		}
	}
}
