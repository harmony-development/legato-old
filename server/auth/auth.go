package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/dgrijalva/jwt-go"

	"github.com/harmony-development/legato/server/config"
)

// Manager wraps logic for authentication
type Manager struct {
	Dependencies *Dependencies
	PrivKey      *rsa.PrivateKey
	PubKey       *rsa.PublicKey
}

// Dependencies are items that an authentication manager needs
type Dependencies struct {
	Config *config.Config
}

// Token is the structure for an authentication JWT
type Token struct {
	jwt.StandardClaims
	UserID   uint64
	Target   string
	Username string
	Avatar   string
}

// New creates a new authenticator
func New(d *Dependencies) (*Manager, error) {
	m := &Manager{
		Dependencies: d,
	}
	priv, err := ioutil.ReadFile(d.Config.Server.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading private key : %v", err)
	}
	m.PrivKey, err = jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, err
	}
	pub, err := ioutil.ReadFile(d.Config.Server.PublicKeyPath)
	if err != nil {
		return nil, err
	}
	m.PubKey, err = jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m Manager) MakeAuthToken(userID uint64, target, username, avatar string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &Token{
		UserID:   userID,
		Target:   target,
		Username: username,
		Avatar:   avatar,
	})
	return token.SignedString(m.PrivKey)
}

// GetPublicKey gets the public key from a domain
func (m Manager) GetPublicKey(domain string) ([]byte, error) {
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/protocol/key")
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	return body, nil
}
