package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"path"
)

type avatarUpdateData struct {
	Token string `mapstructure:"token"`
	Avatar string `mapstructure:"avatar"`
}

func OnAvatarUpdate(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data avatarUpdateData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Token == "" {
		deauth(ws)
		return
	}
	if data.Avatar == "" {
		sendErr(ws, "Oops! We weren't able to set the avatar because of a weird bug")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're updating your avatar too fast, try again in a few moments")
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	var oldAvatarID string
	err = harmonydb.DBInst.QueryRow("SELECT avatar FROM users WHERE id=$1", userid).Scan(&oldAvatarID)
	_, err = harmonydb.DBInst.Exec("UPDATE users SET avatar=$1 WHERE id=$2", data.Avatar, userid)
	if err != nil {
		sendErr(ws, "Something weird happened on our end and we weren't able to set your avatar. Please try again in a few moments")
		return
	}
	go DeleteFromFilestore(path.Base(oldAvatarID))
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", userid)
	if err != nil {
		golog.Warnf("Error selecting guilds for avatarupdate : %v", err)
		return
	}
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			golog.Warnf("Error getting guildid from result on avatarupdate : %v", err)
			return
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				golog.Debugf("Client conns : %v", client)
				for _, conn := range client {
					conn.Send(&globals.Packet{
						Type: "avatarupdate",
						Data: map[string]interface{}{
							"userid": userid,
							"avatar": data.Avatar,
						},
					})
				}
			}
		}
	}
}