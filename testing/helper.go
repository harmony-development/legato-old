package testing

import (
	"harmony-server/server/http"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/routing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func setupBoilerplate() (*echo.Echo, *echo.Group, *hm.Middlewares, *MockDB, *routing.Router) {
	mockDB := &MockDB{}
	m := hm.New(mockDB)
	e := echo.New()
	e.Validator = &http.HarmonyValidator{
		Validator: validator.New(),
	}
	e.Use(m.WithHarmony)
	apiGroup := e.Group("/api")

	return e, apiGroup, m, mockDB, &routing.Router{Middlewares: m}
}
