package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type getSelfData struct {
	Token string `mapstructure:"token"`
}

func GetSelf(limiter *rate.Limiter, c echo.Context) error {
	var data getSelfData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		event.sendErr(ws, "Somethings wrong with your request dude")
		return
	}
	var username string
	var avatar string
	err := harmonydb.DBInst.QueryRow("SELECT username, avatar FROM users WHERE id=$1", ws.Userid).Scan(&username, &avatar)
	if err != nil {
		event.sendErr(ws, "We were unable to get your info, try again later")
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