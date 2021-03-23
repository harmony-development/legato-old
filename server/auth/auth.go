package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/protobuf/types/known/emptypb"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/intercom"
)

type IManager interface {
	MakeAuthToken(userID uint64, target, username, avatar string) (string, error)
	GetPublicKey(host string) (string, error)
	GetOwnPublicKey() *rsa.PublicKey
}

// Manager wraps logic for authentication
type Manager struct {
	*Dependencies
	PrivKey *rsa.PrivateKey
	PubKey  *rsa.PublicKey
}

// Dependencies are items that an authentication manager needs
type Dependencies struct {
	Config          *config.Config
	IntercomManager *intercom.Manager
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

func (m Manager) GetOwnPublicKey() *rsa.PublicKey {
	return m.PubKey
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

// GetPublicKey gets the public key from a host
func (m Manager) GetPublicKey(host string) (string, error) {
	foundationClient := authv1.NewAuthServiceClient(host)
	reply, err := foundationClient.Key(&emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return reply.GetKey(), nil
}
