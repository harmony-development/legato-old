package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	harmony_http "github.com/harmony-development/legato/server/http"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/routing"

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

func setupBoilerplate() (*echo.Echo, *echo.Group, *hm.Middlewares, *MockDB, *routing.Router, MockLogger) {
	mockDB := &MockDB{
		Flags: MockFlags{},
	}
	mockLogger := MockLogger{}
	m := hm.New(mockDB, mockLogger)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Validator = &harmony_http.HarmonyValidator{
		Validator: validator.New(),
	}
	e.Use(m.WithHarmony)
	apiGroup := e.Group("/api")

	return e, apiGroup, m, mockDB, &routing.Router{Middlewares: m}, mockLogger
}
