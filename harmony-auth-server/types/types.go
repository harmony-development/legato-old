package types

import (
	"encoding/json"
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
func (s Server) GetIdentity() (*string) {
	r, err := http.Get(path.Join(s.IP, "/api/", conf.InstanceAPIVersion, "/getidentity"))
	if err != nil {
		return nil
	}
	res := &identityResp{}
	err = getJson(r, res)
	return &res.Identity
}

// SendSession sends a POST request to a specific host to contain an IP
func (s Server) SendSession(session string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion,"/session"), url.Values{"session": {session}})
	if err != nil {
		return
	}
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