package event

import (
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
	"time"
)

type messageData struct {
	token         string
	targetGuild   string
	targetChannel string
	message       string
}

func OnMessage(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data messageData
	if data.token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.message, ok = rawMap["message"].(string); !ok || len(data.message) == 0 {
		return
	}
	if data.targetGuild, ok = rawMap["guild"].(string); !ok {
		return
	}
	if data.targetChannel, ok = rawMap["channel"].(string); !ok {
		return
	}
	userid := VerifyToken(data.token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[data.targetGuild] == nil || globals.Guilds[data.targetGuild].Clients[userid] == nil {
		return
	}
	var messageID = randstr.Hex(16)

	// unfortunately O(n) is the only way to do this, we just need to make n really smol (︶︹︶)
	for _, client := range globals.Guilds[data.targetGuild].Clients {
		client.Send(&socket.Packet{
			Type: "message",
			Data: map[string]interface{}{
				"guild":     data.targetGuild,
				"channel":   data.targetChannel,
				"userid":    userid,
				"createdat": time.Now().UTC().Unix(),
				"message":   data.message,
				"messageid": messageID,
			},
		})
	}
	_, err := harmonydb.DBInst.Exec("INSERT INTO messages(messageid, guildid, channelid, createdat, author, message) VALUES(?, ?, ?, ?, ?, ?)", messageID, data.targetGuild, data.targetChannel, time.Now().UTC().Unix(), userid, data.message)
	if err != nil {
		golog.Warnf("error saving message to database : %v", err)
		return
	}
}
