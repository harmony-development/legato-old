package authentication

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	lru "github.com/hashicorp/golang-lru"
	"github.com/kataras/golog"
	"harmony-auth-server/db"
	"harmony-auth-server/types"
	"io/ioutil"
	"os"
	"time"
)

var signKey *rsa.PrivateKey
var UserSessionCache *lru.ARCCache
var UserIDCache *lru.ARCCache

type SessionClaims struct {
	Session string `json:"session"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

// Init initializes all the things necessary for authentication
func Init() {
	UserIDCache, _ = lru.NewARC(500000)
	UserSessionCache, _ = lru.NewARC(500000)

	privBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		golog.Fatal("error reading private key!", err)
		os.Exit(-1)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		golog.Fatal("error parsing RSA!", err)
		os.Exit(-1)
	}
}

// ValidateSession checks if a specific user session is valid
func ValidateSession(session string) bool {
	_, ok := UserSessionCache.Get(session)
	if !ok {
		if err := db.VerifySession(session); err != nil {
			return false
		}
	}
	return true
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

// MakeServerSessionToken returns a signed session token for instances, and an error if it fails
func MakeServerSessionToken(session string, identity string) (string, error) {
	claims := SessionClaims{
		session,
		identity,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(signKey)
}