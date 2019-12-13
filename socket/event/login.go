package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type loginData struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

func OnLogin(ws *globals.Client, rawMap map[string]interface{}) {
	var data loginData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Email == "" || data.Password == "" {
		loginErr(ws, "Missing Email Or Password")
		return
	}
	var passwd string
	var userid string
	err := harmonydb.DBInst.QueryRow("SELECT password, id from users WHERE email=$1", data.Email).Scan(&passwd, &userid)
	if err != nil {
		loginErr(ws, "Invalid email or password") // either something weird happened or the email doesn't exist
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(data.Password))
	if err != nil {
		loginErr(ws, "Invalid email or password")
		return
	}
	ws.Userid = userid
	sendToken(ws, userid)
}
