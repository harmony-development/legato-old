package v1

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

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
		_, err = jwt.ParseWithClaims(data.AuthToken, &auth.Token{}, func(_ *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
		// TODO : add session, fetch user, cache public key
	} else {
		// TODO : migrate user management from harmony-auth-server
	}
	return ctx.NoContent(http.StatusNotImplemented)
}
