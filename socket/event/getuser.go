package event

import (
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type getUsernameData struct {
	Token  string `mapstructure:"token"`
	Userid string `mapstructure:"userid"`
}

func OnGetUser(ws *globals.Client, rawMap map[string]interface{}) {
	var data getUsernameData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	var username string
	var avatar string
	err := harmonydb.DBInst.QueryRow("SELECT username, avatar FROM users WHERE id=$1", data.Userid).Scan(&username, &avatar)
	if err != nil {
		return
	}
	ws.Send(&globals.Packet{
		Type: "getuser",
		Data: map[string]interface{}{
			"userid": data.Userid,
			"username": username,
			"avatar": avatar,
		},
	})
}
