package authentication

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hashicorp/golang-lru"
	"github.com/kataras/golog"
	"io/ioutil"
	"os"
)

var pubKey *rsa.PublicKey

var SessionCache *lru.ARCCache

// SessionClaims is the token's structure when received from the auth server
type SessionClaims struct {
	Session string `json:"session"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

// Init initializes variables necessary for authentication
func Init() {
	pubBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		golog.Fatalf("error reading public key! ", err)
		os.Exit(-1)
	}
	pubKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		golog.Fatalf("invalid public key! ", err)
		os.Exit(-1)
	}
	SessionCache, err = lru.NewARC(250000)
	if err != nil {
		golog.Fatalf("error making session cache! ", err)
		os.Exit(-1)
	}
}

// ReadAuthToken takes in a token string and returns a session structure or an error
func ReadAuthToken(raw string) (*SessionClaims, error) {
	token, err := jwt.ParseWithClaims(raw, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		return nil, err
	}
	if session, ok := token.Claims.(*SessionClaims); ok {
		return session, nil
	}
	return nil, errors.New("invalid session claims")
}

// IsValid checks if a session exists in its list
func IsValid(session string) bool {
	return SessionCache.Contains(session)
}