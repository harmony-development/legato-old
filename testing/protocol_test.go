package testing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"

	"harmony-server/server/config"
	harmonyHttp "harmony-server/server/http"
	v1 "harmony-server/server/http/core/v1"
	"harmony-server/server/http/protocol"
	"harmony-server/server/http/routing"
)

func TestRegister(t *testing.T) {
	testData := `{
		"email": "maho@amade.us",
		"username": "Maho Hiyajo",
		"password": "Ex@mpl3_p@ssw0rd"
	}`
	e, g, m, mockDB := setupBoilerplate()
	protocol.New(&protocol.Dependencies{
		Router: &routing.Router{
			Middlewares: m,
		},
		APIGroup:  g,
		Sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		Config:    &config.DefaultConf,
		DB:        mockDB,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/protocol/register", strings.NewReader(testData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.Validator = &harmonyHttp.HarmonyValidator{
		Validator: validator.New(),
	}
	e.ServeHTTP(rec, req)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var response v1.RegisterResponse
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
		assert.NotEmpty(t, response.Session)
		assert.NotEmpty(t, response.UserID)
	} else {
		t.Error(rec.Body.String())
	}
}
