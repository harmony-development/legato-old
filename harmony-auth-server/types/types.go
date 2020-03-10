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
	_, err := http.PostForm(path.Join(s.IP, "/api/v1/usernameupdate"), url.Values{"userid": {userID}, "username": {newUsername}})
	if err != nil {
		return
	}
}