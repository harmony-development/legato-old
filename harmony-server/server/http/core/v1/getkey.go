package v1

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/responses"
)

func (h Handlers) GetKey(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, responses.TooManyRequests)
	}
	pemData := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(h.Deps.AuthManager.PubKey),
		},
	)
	return ctx.String(http.StatusOK, string(pemData))
}
