package rest

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
	"harmony-server/harmonydb"
	"net/http"
)

func Login(limiter *rate.Limiter, ctx echo.Context) error {
	email, password := ctx.FormValue("email"), ctx.FormValue("password")
	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form")
	}
	var passwd, userid string
	if err := harmonydb.DBInst.QueryRow("SELECT password, id from users WHERE email=$1", email).Scan(&passwd, &userid); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}
	if !limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	token, err := makeToken(userid)
	if err != nil || token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating token")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": *token,
	})
}