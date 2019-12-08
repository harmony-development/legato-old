package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type GetMessagesData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

type Message struct {
	Guild     string `json:"guild"`
	Channel   string `json:"channel"`
	Userid    string `json:"userid"`
	Createdat int    `json:"createdat"`
	Message   string `json:"message"`
	Messageid string `json:"messageid"`
}

func OnGetMessages(ws *socket.Client, rawMap map[string]interface{}) {
	var data GetMessagesData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT messageid, author, guildid, channelid, createdat, message FROM messages ORDER BY createdat DESC LIMIT 30")
	if err != nil {
		golog.Warnf("Error getting recent messages : %v", err)
		return
	}
	var returnMsgs [] Message
	for res.Next() {
		var msg Message
		err := res.Scan(&msg.Messageid, &msg.Userid, &msg.Guild, &msg.Channel, &msg.Createdat, &msg.Message)
		if err != nil {
			golog.Warnf("Error scanning next row getting messages. Reason: %v", err)
			return
		}
		returnMsgs = append(returnMsgs, msg)
	}
	ws.Send(&socket.Packet{
		Type: "getmessages",
		Data: map[string]interface{}{
			"messages": returnMsgs,
		},
	})
}
