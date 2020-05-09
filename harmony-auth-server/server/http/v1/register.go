package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

	"harmony-auth-server/server/http/hm"
	"harmony-auth-server/util"
)

type registerData struct {
	Email         string `validate:"required,email"`
	Username      string `validate:"required"`
	Password      string `validate:"required"`
	PasswordStats struct {
		Capital   int
		Lowercase int
		Numbers   int
		Special   int
	} `validate:"required"`
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
	if !util.InRange(len(data.Username), h.Config.Server.UsernamePolicy.MinLength, h.Config.Server.UsernamePolicy.MaxLength) {
		return echo.NewHTTPError(http.StatusBadRequest,
			map[string]interface{}{
				"message": "register.username-length",
				"fields": map[string]interface{}{
					"minLength": h.Config.Server.UsernamePolicy.MinLength,
					"maxLength": h.Config.Server.UsernamePolicy.MaxLength,
				},
			},
		)
	}
	if !util.InRange(len(data.Password), h.Config.Server.PasswordPolicy.MinLength, h.Config.Server.PasswordPolicy.MaxLength) {
		return echo.NewHTTPError(http.StatusBadRequest,
			map[string]interface{}{
				"message": "register.password-length",
				"fields": map[string]interface{}{
					"minLength": h.Config.Server.PasswordPolicy.MinLength,
					"maxLength": h.Config.Server.PasswordPolicy.MaxLength,
				},
			},
		)
	}
	if data.PasswordStats.Capital < h.Config.Server.PasswordPolicy.MinCapital ||
		data.PasswordStats.Lowercase < h.Config.Server.PasswordPolicy.MinLowercase ||
		data.PasswordStats.Numbers < h.Config.Server.PasswordPolicy.MinNumbers ||
		data.PasswordStats.Special < h.Config.Server.PasswordPolicy.MinSpecial {
		return echo.NewHTTPError(http.StatusBadRequest,
			map[string]interface{}{
				"message": "register.password-policy",
				"fields": map[string]interface{}{
					"minCapital":   h.Config.Server.PasswordPolicy.MinCapital,
					"minLowercase": h.Config.Server.PasswordPolicy.MinLowercase,
					"minNumbers":   h.Config.Server.PasswordPolicy.MinNumbers,
					"minSpecial":   h.Config.Server.PasswordPolicy.MinSpecial,
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
