package db

import (
	"github.com/thanhpk/randstr"
	"harmony-auth-server/types"
	"time"
)

// RegisterUser registers a new account in the database given some arguments
func RegisterUser(id string, email string, username string, hash string) error {
	_, err := DB.Exec("INSERT INTO users (userid, email, username, avatar, password) VALUES ($1, $2, $3, $4, $5)", id, email, username, "", hash)
	return err
}

// MakeSessionTransaction generates a session corresponding to a user
func MakeSessionTransaction(userid string) (*string, error) {
	expiration := time.Now().Add(48 * time.Hour)
	sessionid := randstr.Hex(16)
	_, err := DB.Exec("INSERT INTO sessions(sessionid, expiration, userid) VALUES($1, $2, $3)", sessionid, expiration.Unix(), userid)
	if err != nil {
		return nil, err
	}
	return &sessionid, nil
}

// VerifySession checks if a certain session is valid
func VerifySession(session string) error {
	_, err := DB.Query("SELECT 1 FROM sessions WHERE sessionid=$1 AND expiration>$2", session, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

// GetUserBySession fetches a user from database using a session
func GetUserBySession(session string) (*types.User, error) {
	res, err := DB.Query("SELECT sessions.userid, users.username, users.avatar FROM sessions INNER JOIN users ON sessions.userid = users.userid WHERE sessionid=$1 AND expiration>$2", session, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := res.Scan(&user.ID, &user.Username, &user.Avatar); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser fetches a user by their ID
func GetUser(userid string) (*types.User, error) {
	var user = types.User{}
	res, err := DB.Query("SELECT username, avatar FROM users WHERE userid=$1", userid)
	if err != nil {
		return nil, err
	}
	if err := res.Scan(&user.Username, &user.Avatar); err != nil {
		return nil, err
	}
	return &user, nil
}

// ListServersTransaction returns an array of servers from SQL given a User ID
func ListServersTransaction(userid string) ([]types.Server, error) {
	res, err := DB.Query("SELECT host FROM servers WHERE userid=$1", userid)
	if err != nil {
		return nil, err
	}
	var servers []types.Server

	for res.Next() {
		var host string
		servers = append(servers, types.Server{IP:host})
	}

	return servers, nil
}

// AddServerTransaction adds a new server to a user's list
func AddServerTransaction(userid string, host string) error {
	_, err := DB.Query("INSERT INTO servers (userid, host) VALUES ($1, $2)", userid, host)
	return err
}