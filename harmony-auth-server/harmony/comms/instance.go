package comms

import (
	"net/http"
	"net/url"
	"path"
)

// Instance is an entry a user's list
type Instance struct {
	Label string
	Host  string
}

// SendAvatarUpdate announces avatar updates to all the hosts on the user's list
func (i Instance) SendAvatarUpdate(userID string, newAvatar string, apiVersion string) {
	_, err := http.PostForm(path.Join(i.Host, "/api/", apiVersion, "/avatarupdate"), url.Values{"userid": {userID}, "avatar": {newAvatar}})
	if err != nil {
		return
	}
}

// SendUsernameUpdate announces username updates to all the hosts on the user's list
func (i Instance) SendUsernameUpdate(userID string, newUsername string, apiVersion string) {
	_, err := http.PostForm(path.Join(i.Host, "/api/", apiVersion, "/usernameupdate"), url.Values{"userid": {userID}, "username": {newUsername}})
	if err != nil {
		return
	}
}
