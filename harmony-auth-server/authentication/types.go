package authentication

import "github.com/dgrijalva/jwt-go"

type SessionClaims struct {
	Session string `json:"session"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}