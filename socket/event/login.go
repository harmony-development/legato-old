package event

import (
	"golang.org/x/crypto/bcrypt"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type loginData struct {
	Email string
	Password string
}

func OnLogin(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data loginData
	if data.Email, ok = rawMap["email"].(string); !ok {
		loginErr(ws, "Email is required")
		return
	}
	if data.Password, ok = rawMap["password"].(string); !ok {
		loginErr(ws, "Password is required")
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