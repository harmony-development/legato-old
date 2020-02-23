package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type getSelfData struct {
	Token string `mapstructure:"token"`
}

func OnGetSelf(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data getSelfData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		sendErr(ws, "Somethings wrong with your request dude")
		return
	}
	var username string
	var avatar string
	err := harmonydb.DBInst.QueryRow("SELECT username, avatar FROM users WHERE id=$1", ws.Userid).Scan(&username, &avatar)
	if err != nil {
		sendErr(ws, "We were unable to get your info, try again later")
		return
	}
	ws.Send(&globals.Packet{
		Type: "getself",
		Data: map[string]interface{}{
			"userid": ws.Userid,
			"username": username,
			"avatar": avatar,
		},
	})
}