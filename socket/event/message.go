package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"time"
)

type messageData struct {
	Token   string `mapstructure:"token"`
	Guild   string `mapstructure:"guild"`
	Channel string `mapstructure:"channel"`
	Message string `mapstructure:"message"`
}

func OnMessage(ws *globals.Client, rawMap map[string]interface{}) {
	var data messageData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	var messageID = randstr.Hex(16)

	// unfortunately O(n) is the only way to do this, we just need to make n really smol (︶︹︶)
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&globals.Packet{
			Type: "message",
			Data: map[string]interface{}{
				"guild":     data.Guild,
				"channel":   data.Channel,
				"userid":    userid,
				"createdat": time.Now().UTC().Unix(),
				"message":   data.Message,
				"messageid": messageID,
			},
		})
	}
	_, err := harmonydb.DBInst.Exec("INSERT INTO messages(messageid, guildid, channelid, createdat, author, message) VALUES($1, $2, $3, $4, $5, $6)", messageID, data.Guild, data.Channel, time.Now().UTC().Unix(), userid, data.Message)
	if err != nil {
		golog.Warnf("error saving message to database : %v", err)
		return
	}
}
