package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
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

func OnRegister(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data registerData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		golog.Warnf("Error decoding register data : %v", err)
		return
	}
	if len(data.Password) < 5 || len(data.Password) > 64 {
		sendErr(ws, "Password must be between 5 and 64 characters long")
		return
	}
	if len(data.Username) < 5 || len(data.Username) > 48 {
		sendErr(ws, "Username must be between 5 and 48 characters long")
		return
	}
	if !verifyEmail(data.Email) {
		sendErr(ws, "Invalid email")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're registering too fast. Try again in a few minutes")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		sendErr(ws, "Something went wrong during registration. Please submit again")
		golog.Warnf("Error hashing password! Reason : %v", err)
		return
	}
	insertQuery, err := harmonydb.DBInst.Prepare("INSERT INTO users (id, email, username, avatar, password) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		sendErr(ws, "Something went wrong during registration. Please submit again")
		golog.Errorf("Failed to prepare query. Reason : %v", err)
		return
	}
	userid := randstr.Hex(16)
	_, err = insertQuery.Exec(userid, data.Email, data.Username, "", string(hash))
	if err != nil {
		sendErr(ws, "Email already registered")
		golog.Debugf("Error inserting user. Reason : %v", err)
		return
	}
	ws.Userid = userid
	sendToken(ws, userid)
}
