package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type usernameUpdateData struct {
	Token string `mapstructure:"token"`
	Username string `mapstructure:"username"`
}

func OnUsernameUpdate(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data usernameUpdateData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Token == "" {
		deauth(ws)
		return
	}
	if data.Username == "" {
		sendErr(ws, "Oops! We weren't able to set your username because of a weird bug")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're updating your username too fast, try again in a few moments")
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	_, err = harmonydb.DBInst.Exec("UPDATE users SET username=$1 WHERE id=$2", data.Username, userid)
	if err != nil {
		sendErr(ws, "Something weird happened on our end and we weren't able to set your avatar. Please try again in a few moments")
		return
	}
	ws.Send(&globals.Packet{
		Type: "usernameupdate",
		Data: map[string]interface{}{
			"success": "true",
		},
	})
}