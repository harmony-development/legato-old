package rest

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func makeToken(id string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}