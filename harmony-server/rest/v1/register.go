package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/authentication"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
	"regexp"
)

const (
	passwordMin = 5
	passwordMax = 64
	usernameMin = 3
	usernameMax = 48
)

// top tier micro-optimization
var emailMatch = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func verifyEmail(email string) bool {
	return emailMatch.MatchString(email)
}

func Register(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	email, username, password := ctx.FormValue("email"), ctx.FormValue("username"), ctx.FormValue("password")
	if len(username) < usernameMin || len(username) > usernameMax {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("username must be between %v and %v characters long", usernameMin, usernameMax))
	}
	if len(password) < passwordMin || len(password) > passwordMax {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("password must be between %v and %v characters long", usernameMin, usernameMax))
	}
	if !verifyEmail(email) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid email")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many register requests, please try again in a few minutes")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error hashing password, please try again later")
	}
	insertQuery, err := harmonydb.DBInst.Prepare("INSERT INTO users (id, email, username, avatar, password) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating account, the email/username may already be registered")
	}
	userid := randstr.Hex(16)
	_, err = insertQuery.Exec(userid, email, username, "", string(hash))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating account, the email/username may already be registered")
	}
	token, err := authentication.MakeToken(userid)
	if err != nil || token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating token")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": *token,
	})
}