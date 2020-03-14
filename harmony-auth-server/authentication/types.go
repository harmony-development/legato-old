package authentication

import "github.com/dgrijalva/jwt-go"

type SessionClaims struct {
	Session string `json:"session"`
	jwt.StandardClaims
}