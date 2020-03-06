package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type getUsernameData struct {
	Token  string `mapstructure:"token"`
	Userid string `mapstructure:"userid"`
}

func GetUser(limiter *rate.Limiter, c echo.Context) error {
	var data getUsernameData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		event.sendErr(ws, "Something's wrong with your request dude")
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
