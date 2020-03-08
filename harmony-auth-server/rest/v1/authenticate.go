package v1

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-auth-server/rest/hm"
	"net/http"
)

func SendKeyToServer(server string, key string, ip string) {
	body, err := json.Marshal(map[string]string{
		"key": key,
		"ip": ip,
	})
	if err != nil {
		golog.Warnf("error marshalling auth request %v", err)
		return
	}

	_, err = http.Post(server, "application/json", bytes.NewBuffer(body))
	if err != nil {
		golog.Warnf("error POSTing target, %v", err)
		return
	}
}

// Authenticate is an API path that is meant to be triggered by a client, and initializes a session authorization
func Authenticate(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	host, session := ctx.FormValue("host"), ctx.FormValue("session")
	if host == "" || session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid auth arguments")
	}
	authKey := randstr.Hex(16)
	go SendKeyToServer(host, authKey, ctx.RealIP())
	return ctx.JSON(http.StatusOK, map[string]string{
		"key": authKey,
	})
}