package testing

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/routing"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func routingTestHandler(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	return ctx.String(http.StatusOK, strings.Join(ctx.ParamValues(), " "))
}

var routeTests = []struct {
	tester func(t *testing.T, e *echo.Echo)
	route  routing.Route
}{
	{
		func(t *testing.T, e *echo.Echo) {
			req := httptest.NewRequest(http.MethodPost, "/api/any", strings.NewReader(`{}`))
			req.Header.Set(echo.HeaderAuthorization, "let me in")
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		routing.Route{
			Method:      routing.ANY,
			Path:        "/any",
			Auth:        false,
			Schema:      struct{}{},
			Location:    routing.LocationNone,
			Handler:     routingTestHandler,
			Permissions: hm.NoPermission,
		},
	},
	{
		func(t *testing.T, e *echo.Echo) {
			req := httptest.NewRequest(http.MethodDelete, "/api/delete/42069/", strings.NewReader(`{}`))
			req.Header.Set(echo.HeaderAuthorization, "let me in")
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, len(strings.Split(rec.Body.String(), " ")), 1)
			}
		},
		routing.Route{
			Method:      routing.DELETE,
			Path:        "/delete/:guild_id/",
			Auth:        true,
			Schema:      struct{}{},
			Location:    routing.LocationGuild,
			Handler:     routingTestHandler,
			Permissions: hm.ModifyChannels,
		},
	},
	{
		func(t *testing.T, e *echo.Echo) {
			req := httptest.NewRequest(http.MethodGet, "/api/get/42069/3264/", strings.NewReader(`{}`))
			req.Header.Set(echo.HeaderAuthorization, "let me in")
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, len(strings.Split(rec.Body.String(), " ")), 2)
			}
		},
		routing.Route{
			Method:      routing.GET,
			Path:        "/get/:guild_id/:channel_id/",
			Auth:        true,
			Schema:      struct{}{},
			Location:    routing.LocationGuildAndChannel,
			Handler:     routingTestHandler,
			Permissions: hm.ModifyGuild,
		},
	},
	{
		func(t *testing.T, e *echo.Echo) {
			req := httptest.NewRequest(http.MethodPatch, "/api/patch/10239/", strings.NewReader(`{}`))
			req.Header.Set(echo.HeaderAuthorization, "let me in")
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, len(strings.Split(rec.Body.String(), " ")), 1)
			}
		},
		routing.Route{
			Method:      routing.PATCH,
			Path:        "/patch/:user_id/",
			Auth:        true,
			Schema:      struct{}{},
			Location:    routing.LocationUser,
			Handler:     routingTestHandler,
			Permissions: hm.ModifyInvites,
		},
	},
	{
		func(t *testing.T, e *echo.Echo) {
			req := httptest.NewRequest(http.MethodPost, "/api/post/10222/12314/43612/", strings.NewReader(`{}`))
			req.Header.Set(echo.HeaderAuthorization, "let me in")
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			if assert.Equal(t, http.StatusOK, rec.Code) {
				assert.Equal(t, len(strings.Split(rec.Body.String(), " ")), 3)
			}
		},
		routing.Route{
			Method:      routing.POST,
			Path:        "/post/:guild_id/:channel_id/:message_id/",
			Auth:        true,
			Schema:      struct{}{},
			Location:    routing.LocationGuildChannelAndMessage,
			Handler:     routingTestHandler,
			Permissions: hm.Owner,
		},
	},
}

func TestRouting(t *testing.T) {
	e, g, _, _, router := setupBoilerplate()
	for _, test := range routeTests {
		router.BindRoute(g, test.route)
	}
	for _, test := range routeTests {
		test.tester(t, e)
	}
}
