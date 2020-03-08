package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

func Login(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	email, password := ctx.FormValue("email"), ctx.FormValue("password")
	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	var passwd, userid string
	if err := db.DB.QueryRow("SELECT password, id from users WHERE email=$1", email).Scan(&passwd, &userid); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	session, err := db.MakeSessionTransaction(userid)
	if err != nil || session == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create session, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": *session,
	})
}