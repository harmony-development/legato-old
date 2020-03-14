package db

import (
	"harmony-auth-server/conf"
	"net/http"
	"net/url"
	"path"
	"time"
)

type User struct {
	ID       string `json:"userid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Server struct {
	IP string
}

func (s Server) SendSession(session string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion,"/session"), url.Values{"session": {session}, "expires": {string(time.Now().Add(24 * time.Hour).Unix())}})
	if err != nil {
		return
	}
}

func (s Server) SendUsernameUpdate(userID string, newUsername string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion,"/usernameupdate"), url.Values{"userid": {userID}, "username": {newUsername}})
	if err != nil {
		return
	}
}

func (s Server) SendAvatarUpdate(userID string, newAvatar string) {
	_, err := http.PostForm(path.Join(s.IP, "/api/", conf.InstanceAPIVersion, "/avatarupdate"), url.Values{"userid": {userID}, "avatar": {newAvatar}})
	if err != nil {
		return
	}
}