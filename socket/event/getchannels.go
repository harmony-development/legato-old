package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type getChannelsData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnGetChannels(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data getChannelsData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid ,err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're getting channel listings too fast, try again soon")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT channelid, channelname FROM channels WHERE guildid=$1", data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to get a list of channels for this guild. You should try again")
		golog.Warnf("Error selecting channels : %v", err)
		return
	}

	var returnChannels = make(map[string]string)
	for res.Next() {
		var channelid, channelname string
		err = res.Scan(&channelid, &channelname)
		if err != nil {
			sendErr(ws, "We weren't able to get a full listing of guild channels. You should try again")
			golog.Warnf("Error scanning channels : %v", err)
			return
		}
		returnChannels[channelid] = channelname
	}
	ws.Send(&globals.Packet{
		Type: "getchannels",
		Data: map[string]interface{}{
			"channels": returnChannels,
		},
	})
}
