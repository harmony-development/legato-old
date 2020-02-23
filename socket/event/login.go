package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type loginData struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

func OnLogin(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data loginData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Email == "" || data.Password == "" {
		sendErr(ws, "You need both an email and password to login")
		return
	}
	var passwd string
	var userid string
	err := harmonydb.DBInst.QueryRow("SELECT password, id from users WHERE email=$1", data.Email).Scan(&passwd, &userid)
	if err != nil {
		sendErr(ws, "Invalid email or password") // either something weird happened or the email doesn't exist
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(data.Password))
	if err != nil {
		sendErr(ws, "Invalid email or password")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're logging in too fast. Try again in a few seconds")
		return
	}
	ws.Userid = userid
	sendToken(ws, userid)
}
