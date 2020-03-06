package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/authentication"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func Login(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
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
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	token, err := authentication.MakeToken(userid)
	if err != nil || token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating token")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": *token,
	})
}