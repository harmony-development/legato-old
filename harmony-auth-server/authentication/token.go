package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// MakeSessionToken returns a signed session token, and an error if it fails
func MakeSessionToken(session string) (string, error) {
	claims := SessionClaims{
		session,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}