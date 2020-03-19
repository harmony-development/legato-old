package v1

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"harmony-server/authentication"
	"harmony-server/rest/hm"
	"net/http"
	"os"
)

type User struct {
	ID       string `json:"userid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// Session is an endpoint for the auth server to send auth tokens to
func Session(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	session, rawUser := ctx.FormValue("session"), ctx.FormValue("user")
	if session == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid session")
	}
	token, err := authentication.ReadAuthToken(session)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid session")
	}
	if token.Identity != os.Getenv("HARMONY_IDENTITY") {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid identity")
	}
	var user User
	if err := json.Unmarshal([]byte(rawUser), user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad user")
	}
	authentication.SessionCache.Add(token.Session, user)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "session accepted",
	})
}