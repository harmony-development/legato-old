package v1

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"harmony-auth-server/conf"
	"harmony-auth-server/db"
	"harmony-auth-server/rest/hm"
	"net/http"
)

// Register takes in an email, username, and password and tries to create an account on the server
func Register(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	email, username, password := ctx.FormValue("email"), ctx.FormValue("username"), ctx.FormValue("password")
	if len(username) < conf.UsernameLenMin || len(username) > conf.UsernameLenMax {
		return echo.NewHTTPError(http.StatusBadRequest, conf.UsernameLenMessage)
	}
	if len(password) < conf.PassLenMin || len(password) > conf.PassLenMax {
		return echo.NewHTTPError(http.StatusBadRequest, conf.PasswordLenMessage)
	}
	if !conf.EmailValidation.MatchString(email) {
		return echo.NewHTTPError(http.StatusBadRequest, "email is invalid")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many registration requests, please try again later")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to securely save account, please try again later")
	}
	userid := randstr.Hex(16)
	err = db.RegisterUser(userid, email, username, string(hash))
	if err != nil {
		golog.Warn(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to register account, email might already be registered")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully registered account, you may now log in",
	})
}