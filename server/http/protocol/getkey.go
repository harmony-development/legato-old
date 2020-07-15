package protocol

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/responses"
)

func (h API) GetKey(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	keyBytes, err := x509.MarshalPKIXPublicKey(h.Deps.AuthManager.PubKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
	}
	pemData := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: keyBytes,
		},
	)
	return ctx.String(http.StatusOK, string(pemData))
}
