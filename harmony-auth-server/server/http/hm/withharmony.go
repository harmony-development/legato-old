package hm

import "github.com/labstack/echo/v4"

func WithHarmony(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hctx := HarmonyContext{
			Context: ctx,
		}
		return handler(hctx)
	}
}
