package v1

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

	"harmony-server/server/auth"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
)

type LoginData struct {
	AuthToken string `validate:"required_without=Email"`
	Domain    string `validate:"required_without=Email"`
	Email     string `validate:"required_without=AuthToken"`
	Password  string `validate:"required_without=AuthToken"`
}

func (h Handlers) Login(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(*LoginData)
	if data.AuthToken != "" {
		pem, err := h.Deps.AuthManager.GetPublicKey(data.Domain)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pem)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		t, err := jwt.ParseWithClaims(data.AuthToken, &auth.Token{}, func(_ *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		token := t.Claims.(*auth.Token)
		session := randstr.Hex(16)
		if err := h.Deps.DB.AddForeignUser(token.UserID, data.Domain, token.Username, token.Avatar); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		if err := h.Deps.DB.AddForeignSession(token.UserID, data.Domain, session); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		return ctx.JSON(http.StatusOK, LoginResponse{Session: session})
	} else {
		user, err := h.Deps.DB.GetUser(data.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, responses.InvalidEmail)
		}
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, responses.InvalidPassword)
		}
		session := randstr.Hex(16)
		if err := h.Deps.DB.AddSession(user.UserID, session); err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		return ctx.JSON(http.StatusOK, LoginResponse{Session: session})
	}
}
