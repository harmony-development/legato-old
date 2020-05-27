package hm

import "github.com/labstack/echo/v4"

// WithHarmony adds a HarmonyContext to the echo.Context
func (m *Middlewares) WithHarmony(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hctx := HarmonyContext{
			Context: ctx,
		}
		return handler(hctx)
	}
}
