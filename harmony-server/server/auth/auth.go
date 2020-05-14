package auth

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"harmony-server/server/config"
)

// Manager wraps logic for authentication
type Manager struct {
	Dependencies       *Dependencies
}

type Dependencies struct {
	Config *config.Config
}

// New creates a new authenticator
func New(d *Dependencies) (*Manager, error) {
	m := &Manager{
		Dependencies: d,
	}

	return m, nil
}

func (m Manager) GetPublicKey(domain string) ([]byte, error) {
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/pubkey")
	resp, err := http.Get(u.RawQuery)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}