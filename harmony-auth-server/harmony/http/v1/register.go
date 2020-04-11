package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
)

type registerData struct {
	Email    string `validate:"required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

// Register takes in an email, username, and password and tries to create an account on the server
func (h Handlers) Register(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(registerData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if len(data.Username) < h.Config.Server.UsernameLenMin || len(data.Username) > h.Config.Server.UsernameLenMax {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf("Username must be between %v and %v characters long",
				h.Config.Server.UsernameLenMin,
				h.Config.Server.UsernameLenMax,
			),
		)
	}
	if len(data.Password) < h.Config.Server.PassLenMin || len(data.Password) > h.Config.Server.PassLenMax {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Password must be between %v and %v characters long",
			h.Config.Server.PassLenMin,
			h.Config.Server.PassLenMax,
		), )
	}
	if !h.Consts.EmailRegex.MatchString(data.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, "email is invalid")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many registration requests, please try again later")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to securely save account, please try again later")
	}
	err = h.DB.RegisterUser(randstr.Hex(h.Config.Server.UserIDLength), data.Email, data.Username, string(hash))
	if err != nil {
		logrus.Warn(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to register account, email might already be registered")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully registered account, you may now log in",
	})
}
