package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

// Manager wraps logic for authentication
type Manager struct {
	PublicKey *rsa.PublicKey
}

// TokenClaims is the structure for the JWT signed session token
type TokenClaims struct {
	Session  string `json:"session"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

// New creates a new authenticator
func New() (*Manager, error) {
	handler := &Manager{}
	pubBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		return nil, err
	}
	handler.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

// ReadAuthToken takes in a token string and returns a session structure or an error
func (m Manager) ReadAuthToken(raw string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(raw, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return m.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if session, ok := token.Claims.(*TokenClaims); ok {
		return session, nil
	}
	return nil, errors.New("invalid session claims")
}
