package testing

import (
	"harmony-server/server/http/hm"

	"github.com/labstack/echo/v4"
)

func setupBoilerplate() (*echo.Echo, *echo.Group, *hm.Middlewares, *MockDB) {
	mockDB := &MockDB{}
	m := hm.New(mockDB)
	e := echo.New()
	e.Use(m.WithHarmony)
	apiGroup := e.Group("/api")

	return e, apiGroup, m, mockDB
}
