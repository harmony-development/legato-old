package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type loginData struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

func OnLogin(ws *socket.Client, rawMap map[string]interface{}) {
	var data loginData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Email == "" || data.Password == "" {
		loginErr(ws, "Missing Email Or Password")
		return
	}
	var passwd string
	var id string
	err := harmonydb.DBInst.QueryRow("SELECT password, id from users WHERE email=?", data.Email).Scan(&passwd, &id)
	if err != nil {
		loginErr(ws, "Invalid email or password") // either something weird happened or the email doesn't exist
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(data.Password))
	if err != nil {
		loginErr(ws, "Invalid email or password")
		return
	}
	sendToken(ws, id)
}
