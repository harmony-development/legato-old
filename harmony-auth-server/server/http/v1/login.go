package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"harmony-auth-server/server/db"
	"harmony-auth-server/server/http/hm"
	"net/http"
)

type loginData struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

// Login takes in an email and password and returns a session token for connecting to instances
func (h Handlers) Login(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(loginData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	var user db.User
	if err := h.DB.QueryRow(
		"SELECT userid, password FROM users WHERE email=$1", data.Email,
	).Scan(&user.UserID, &user.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting user details")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	session := h.AuthManager.Sessions.MakeSession(user.UserID, h.Config.Server.SessionExpire)
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": session,
	})
}
