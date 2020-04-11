package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"harmony-auth-server/harmony/config"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// Manager wraps logic for authentication
type Manager struct {
	SignKey    *rsa.PrivateKey
	Expiration time.Duration
	Sessions   Sessions
}

// SessionClaims is the structure for the JWT signed session token
type SessionClaims struct {
	Session  string `json:"session"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

// New creates a new authenticator
func New(conf *config.Config) *Manager {
	handler := &Manager{}
	handler.Expiration = conf.Server.SessionExpire
	privBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		logrus.Fatal("error reading private key!", err)
		os.Exit(-1)
	}
	handler.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		logrus.Fatal("error parsing RSA!", err)
		os.Exit(-1)
	}
	handler.Sessions = Sessions{
		SessionMap: make(map[string]Session),
		Mut:        &sync.RWMutex{},
	}
	go handler.Sessions.ExpireSessions()
	return handler
}

// MakeServerToken creates a signed session token for instance servers to use
func (m Manager) MakeServerToken(session string, identity string) (string, error) {
	claims := &SessionClaims{
		session,
		identity,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.Expiration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.SignKey)
}
