package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"harmony-server/server/config"
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

// New creates a new authenticator
func New(d *Dependencies) (*Manager, error) {
	m := &Manager{
		Dependencies: d,
	}
	priv, err := ioutil.ReadFile(d.Config.Server.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading private key : %v", err)
	}
	privPem, _ := pem.Decode(priv)
	m.PrivKey, err = x509.ParsePKCS1PrivateKey(privPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key : %v", err)
	}
	pub, err := ioutil.ReadFile(d.Config.Server.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading public key : %v", err)
	}
	pubPem, _ := pem.Decode(pub)
	pubKey, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key : %v", err)
	}
	m.PubKey = pubKey.(*rsa.PublicKey)
	return m, nil
}

// GetPublicKey gets the public key from a domain
func (m Manager) GetPublicKey(domain string) ([]byte, error) {
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/key")
	resp, err := http.Get(u.RawQuery)
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
