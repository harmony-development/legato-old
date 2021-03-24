package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/creasty/defaults"
	"github.com/harmony-development/legato/server/config"
	"github.com/labstack/echo/v4"
)

func DefaultConf() *config.Config {
	var cfg config.Config
	defaults.MustSet(&cfg)
	return &cfg
}

func DummyContext(e *echo.Echo) echo.Context {
	return e.NewContext(httptest.NewRequest(http.MethodGet, "https://127.0.0.1", nil), httptest.NewRecorder())
}
