package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/harmonydb"
	"harmony-server/socket"
	"regexp"
)

type registerData struct {
	Email    string `mapstructure:"email"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// compiling inside a function will slow it down a teensy bit
var emailMatch = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func verifyEmail(email string) bool {
	return emailMatch.MatchString(email)
}

func OnRegister(ws *socket.Client, rawMap map[string]interface{}) {
	var data registerData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		golog.Warnf("Error decoding register data : %v", err)
		return
	}
	if len(data.Password) < 5 || len(data.Password) > 64 {
		regErr(ws, "Password must be between 5 and 64 characters long")
		return
	}
	if len(data.Username) < 5 || len(data.Username) > 48 {
		regErr(ws, "Username must be between 5 and 48 characters long")
		return
	}
	if !verifyEmail(data.Email) {
		regErr(ws, "Invalid email")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		golog.Warnf("Error hashing password! Reason : %v", err)
		return
	}
	insertQuery, err := harmonydb.DBInst.Prepare("INSERT INTO users (id, email, username, avatar, password) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		golog.Errorf("Failed to prepare query. Reason : %v", err)
		return
	}
	userid := randstr.Hex(16)
	_, err = insertQuery.Exec(userid, data.Email, data.Username, "", string(hash))
	if err != nil {
		regErr(ws, "Email already registered")
		golog.Debugf("Error inserting user. Reason : %v", err)
		return
	}
	ws.Userid = userid
	sendToken(ws, userid)
}
