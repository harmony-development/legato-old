package authentication

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"io/ioutil"
	"os"
)

var signKey *rsa.PrivateKey

// Init initializes all the things necessary for authentication
func Init() {
	makeSessionCache()
	makeUserIDCache()

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