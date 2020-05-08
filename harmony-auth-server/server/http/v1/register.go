package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

	"harmony-auth-server/server/http/hm"
)

type registerData struct {
	Email    string `validate:"required,email"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

// Register takes in an email, username, and password and tries to create an account on the server
func (h Handlers) Register(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := new(registerData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid")
	}
	if len(data.Username) < h.Config.Server.UsernameLenMin || len(data.Username) > h.Config.Server.UsernameLenMax {
		return echo.NewHTTPError(http.StatusBadRequest,
			map[string]interface{}{
				"message": "register.username-length",
				"fields": map[string]interface{}{
					"minLength": h.Config.Server.UsernameLenMin,
					"maxLength": h.Config.Server.UsernameLenMax,
				},
			},
		)
	}
	if len(data.Password) < h.Config.Server.PassLenMin || len(data.Password) > h.Config.Server.PassLenMax {
		return echo.NewHTTPError(http.StatusBadRequest,
			map[string]interface{}{
				"message": "register.password-length",
				"fields": map[string]interface{}{
					"minLength": h.Config.Server.PassLenMin,
					"maxLength": h.Config.Server.PassLenMax,
				},
			},
		)
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "rate-limit")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "register.error")
	}
	exists, err := h.DB.EmailRegistered(data.Email)
	if exists {
		return echo.NewHTTPError(http.StatusConflict, "register.already-registered")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "register.error-registering")
	}
	userid := randstr.Hex(h.Config.Server.UserIDLength)
	err = h.DB.AddUser(userid, data.Email, data.Username, string(hash))
	if err != nil {
		logrus.Warn(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "register.error-registering")
	}
	session := h.AuthManager.Sessions.MakeSession(userid, h.Config.Server.SessionExpire)
	return ctx.JSON(http.StatusOK, map[string]string{
		"session": session,
	})
}
