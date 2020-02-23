package authentication

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = os.Getenv("JWT_SECRET")

// VerifyToken verifies if a token has a valid signature
func VerifyToken(tokenstr string) (string, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var id string
		if id, ok = claims["id"].(string); !ok {
			return "", fmt.Errorf("weird token")
		}
		return id, nil
	}
	return "", fmt.Errorf("weird token")
}
