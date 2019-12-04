package event

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type GetMessagesData struct {
	token string
	targetGuild string
}

type Message struct {
	Guild string `json:"guild"`
	Userid string `json:"userid"`
	Createdat int `json:"createdat"`
	Message string `json:"message"`
	Messageid string `json:"messageid"`
}

func OnGetMessages(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data GetMessagesData
	if data.token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.targetGuild, ok = rawMap["guild"].(string); !ok {
		return
	}
	userid := verifyToken(data.token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	if globals.Guilds[data.targetGuild] == nil || globals.Guilds[data.targetGuild].Clients[userid] == nil {
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT messageid, author, guildid, createdat, message FROM messages ORDER BY createdat DESC LIMIT 30")
	if err != nil {
		golog.Warnf("Error getting recent messages : %v", err)
		return
	}
	var returnMsgs[] Message
	for res.Next() {
		var msg Message
		err := res.Scan(&msg.Messageid, &msg.Userid, &msg.Guild, &msg.Createdat, &msg.Message)
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