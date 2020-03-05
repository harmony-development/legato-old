package v1

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type usernameUpdateData struct {
	Token string `mapstructure:"token"`
	Username string `mapstructure:"username"`
}

func UsernameUpdate(limiter *rate.Limiter, ctx echo.Context) error {
	var data usernameUpdateData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Username == "" {
		event.sendErr(ws, "Oops! We weren't able to set your username because of a weird bug")
		return
	}
	if !limiter.Allow() {
		event.sendErr(ws, "You're updating your username too fast, try again in a few moments")
		return
	}
	_, err := harmonydb.DBInst.Exec("UPDATE users SET username=$1 WHERE id=$2", data.Username, ws.Userid)
	if err != nil {
		event.sendErr(ws, "Something weird happened on our end and we weren't able to set your avatar. Please try again in a few moments")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", ws.Userid)
	if err != nil {
		golog.Warnf("Error selecting guilds for avatarupdate : %v", err)
		return
	}
	golog.Debugf("Username update. New Username : %v", data.Username)
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			golog.Warnf("Error getting guildid from result on avatarupdate : %v", err)
			return
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				for _, conn := range client {
					conn.Send(&globals.Packet{
						Type: "usernameupdate",
						Data: map[string]interface{}{
							"userid":   ws.Userid,
							"username": data.Username,
						},
					})
				}
			}
		}
	}
}