package authentication

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func MakeToken(session string, host string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"session":  session,
		"host": host,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func VerifyToken(tokenstr string) (*string, *string, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var session string
		var host string
		if session, ok = claims["id"].(string); !ok {
			return nil, nil, errors.New("weird token")
		}
		if host, ok = claims["host"].(string); !ok {
			return nil, nil, errors.New("weird token")
		}
		return &session, &host, nil
	}
	return nil, nil, errors.New("weird token")
}
