package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"

	"harmony-server/server/config"
	v1 "harmony-server/server/http/core/v1"
	"harmony-server/server/http/protocol"
)

var RegisterTests = RequestTestSet{
	Path: "/api/protocol/register",
	Type: http.MethodPost,
	Tests: []RequestTest{
		{
			Data: map[string]interface{}{
				"email":    "maho@amade.us",
				"username": "Maho Hiyajo",
				"password": "Ex@mpl3_p@ssw0rd",
			},
			Tester: func(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rec.Code)
				var response v1.RegisterResponse
				assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
				assert.NotEmpty(t, response.Session)
				assert.NotEmpty(t, response.UserID)
			},
		},
		{
			Data: map[string]interface{}{
				"email":    "maho",
				"username": "Maho Hiyajo",
				"password": "Ex@mpl3_p@ssw0rd",
			},
			Tester: func(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			},
		},
		{
			Data: map[string]interface{}{
				"email":    "maho@amade.us",
				"username": "_",
				"password": "Ex@mpl3_p@ssw0rd",
			},
			Tester: func(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotAcceptable, rec.Result().StatusCode)
				assert.NotEmpty(t, rec.Body.String())
			},
		},
		{
			Data: map[string]interface{}{
				"email":    "maho@amade.us",
				"username": "Maho Hiyajo",
				"password": "no lol",
			},
			Tester: func(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNotAcceptable, rec.Result().StatusCode)
				assert.NotEmpty(t, rec.Body.String())
			},
		},
	},
}

func TestRegister(t *testing.T) {
	e, g, _, mockDB, router := setupBoilerplate()
	protocol.New(&protocol.Dependencies{
		Router:    router,
		APIGroup:  g,
		Sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		Config:    &config.DefaultConf,
		DB:        mockDB,
	})
	for _, test := range RegisterTests.Tests {
		body, err := json.Marshal(test.Data)
		assert.NoError(t, err)
		req := httptest.NewRequest(RegisterTests.Type, RegisterTests.Path, bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		test.Tester(t, req, rec)
	}
}
