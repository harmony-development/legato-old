package authentication

import (
	"errors"
	"github.com/hashicorp/golang-lru"
	"harmony-auth-server/db"
	"harmony-auth-server/types"
)

var UserSessionCache *lru.ARCCache
var UserIDCache *lru.ARCCache

func makeSessionCache() {
	UserSessionCache, _ = lru.NewARC(500000)
	return
}

func makeUserIDCache() {
	UserIDCache, _ = lru.NewARC(500000)
	return
}

func ValidateSession(session string) bool {
	_, ok := UserSessionCache.Get(session)
	if !ok {
		if err := db.VerifySession(session); err != nil {
			return false
		}
	}
	return true
}

// GetUserBySession returns user details given a session string
func GetUserBySession(session string) (*types.User, error) {
	entry, ok := UserSessionCache.Get(session)
	if !ok {
		user, err := db.GetUserFromDB(session)
		if err != nil {
			return nil, err
		}
		UserSessionCache.Add(session, *user)
		UserIDCache.Add(user.ID, *user)
		return user, nil
	}
	user, ok := entry.(types.User)
	if !ok {
		return nil, errors.New("not a user")
	}
	return &user, nil
}

// GetUserByID returns user details given a user's ID
func GetUserByID(userid string) (*types.User, error) {
	entry, ok := UserIDCache.Get(userid)
	if !ok {
		user, err := db.GetUser(userid)
		if err != nil {
			return nil, err
		}
		UserIDCache.Add(userid, *user)
		return user, nil
	}
	user, ok := entry.(types.User)
	if !ok {
		return nil, errors.New("not a user")
	}
	return &user, nil
}