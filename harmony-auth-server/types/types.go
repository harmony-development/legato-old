package types

import (
	"net/http"
	"net/url"
	"path"
)

type User struct {
	Userid string `json:"userid"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
}

type Server struct {
	IP string
}

func (s Server) SendUsernameUpdate(userID string, newUsername string) {
	target, err := url.Parse(s.IP)
	if err != nil {
		return
	}
	target.Path = path.Join(target.Path, "/api/v1/usernameupdate")
	_, err = http.PostForm(target.String(), url.Values{"userid": userID, "username": newUsername})
	if err != nil {
		return
	}
}