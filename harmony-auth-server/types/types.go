package types

import (
	"encoding/json"
	"errors"
	"github.com/kataras/golog"
	"harmony-auth-server/conf"
	"net/http"
	"net/url"
	"path"
)

type User struct {
	ID       string `json:"userid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Server struct {
	IP string
}

type identityResp struct {
	Identity string `json:"identity"`
}

func getJson(res *http.Response, target interface{}) error {
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

// GetIdentity requests an instance server to identify itself
func (s Server) GetIdentity() (*string, error) {
	u, err := url.Parse(s.IP)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/", conf.InstanceAPIVersion, "/getidentity/")
	r, err := http.Get(u.String())
	if err != nil {
		golog.Warn("error GETing identity ", err)
		return nil, errors.New("error GETing")
	}
	res := &identityResp{}
	err = getJson(r, res)
	if err != nil {
		return nil, err
	}
	return &res.Identity, nil
}

// SendSession sends a POST request to a specific host to contain an IP
func (s Server) SendSession(session string, user *User) (*http.Response, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	userOut, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(s.IP)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/", conf.InstanceAPIVersion,"/session")
	res, err := http.PostForm(
		u.String(),
		url.Values{
			"session": {session},
			"user": {string(userOut)},
		},
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendUsernameUpdate sends a POST request to a specific host to notify a username update
func (s Server) SendUsernameUpdate(userID string, newUsername string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion,"/usernameupdate"), url.Values{"userid": {userID}, "username": {newUsername}})
	if err != nil {
		return
	}
}

// SendAvatarUpdate sends a POST request to a specific host to notify an avatar update
func (s Server) SendAvatarUpdate(userID string, newAvatar string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion, "/avatarupdate"), url.Values{"userid": {userID}, "avatar": {newAvatar}})
	if err != nil {
		return
	}
}