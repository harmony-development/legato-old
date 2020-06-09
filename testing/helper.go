package testing

import (
	harmony_http "harmony-server/server/http"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/routing"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestTestSet struct {
	Path  string
	Type  string
	Tests []RequestTest
}

type RequestTest struct {
	Data    map[string]interface{}
	DBFlags MockFlags
	Tester  func(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder)
}

func setupBoilerplate() (*echo.Echo, *echo.Group, *hm.Middlewares, *MockDB, *routing.Router) {
	mockDB := &MockDB{
		Flags: MockFlags{},
	}
	m := hm.New(mockDB)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = &harmony_http.HarmonyValidator{
		Validator: validator.New(),
	}
	e.Use(m.WithHarmony)
	apiGroup := e.Group("/api")

	return e, apiGroup, m, mockDB, &routing.Router{Middlewares: m}
}
