package authentication

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"io/ioutil"
	"os"
)

var pubKey *rsa.PublicKey

type SessionData struct {
	ID       string `json:"userid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	ExpiresAt int64 `json:"expiresat"`
}

var SessionStore map[string]*SessionData

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
	SessionStore = make(map[string]*SessionData)
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

// GetUserBySession returns a user from the session cache or an error if the session is not valid
func GetUserBySession(session string) (*SessionData, error) {
	if SessionStore[session] == nil {
		golog.Warnf("bruhhh")
		return nil, errors.New("session is invalid")
	}
	return SessionStore[session], nil
}