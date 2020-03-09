package db

import (
	"github.com/thanhpk/randstr"
	"harmony-auth-server/types"
	"time"
)

func MakeSessionTransaction(userid string) (*string, error) {
	expiration := time.Now().Add(48 * time.Hour)
	sessionid := randstr.Hex(16)
	_, err := DB.Exec("INSERT INTO sessions(sessionid, expiration, userid) VALUES($1. $2. $30)", sessionid, expiration.Unix(), userid)
	if err != nil {
		return nil, err
	}
	return &sessionid, nil
}

func GetUserBySession(session string) (*types.User, error) {
	res, err := DB.Query("SELECT sessions.userid, users.username, users.avatar FROM sessions INNER JOIN users ON sessions.userid = users.userid WHERE sessionid=$1 AND expiration>$2", session, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := res.Scan(&user.Userid, &user.Username, &user.Avatar); err != nil {
		return nil, err
	}
	return &user, nil
}

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

func ListServersTransaction(userid string) ([]string, error) {
	res, err := DB.Query("SELECT host FROM servers WHERE userid=$1", userid)
	if err != nil {
		return nil, err
	}
	var servers []string

	for res.Next() {
		var host string
		servers = append(servers, host)
	}

	return servers, nil
}