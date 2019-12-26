package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type GetMessagesData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Channel string `mapstructure:"channel"`
}

type Message struct {
	Guild     string `json:"guild"`
	Channel   string `json:"channel"`
	Userid    string `json:"userid"`
	Createdat int    `json:"createdat"`
	Message   string `json:"message"`
	Messageid string `json:"messageid"`
}

func OnGetMessages(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data GetMessagesData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		sendErr(ws, "")
		sendErr(ws, "Something was wrong with your request. Please try again")
		return
	}
	if data.Guild == "" || data.Token == "" || data.Channel == "" {
		sendErr(ws, "Something was wrong with your request. Please try again")
		golog.Warnf("Error decoding getmessages request")
		return
	}
	userid ,err := authentication.VerifyToken(data.Token)
	if err != nil { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT messageid, author, guildid, channelid, createdat, message FROM messages WHERE guildid=$1 ORDER BY createdat DESC LIMIT 30")
	if err != nil {
		sendErr(ws, "We weren't able to get a list of messages. Please try again")
		golog.Warnf("Error getting recent messages : %v", err)
		return
	}
	var returnMsgs [] Message
	for res.Next() {
		var msg Message
		err := res.Scan(&msg.Messageid, &msg.Userid, &msg.Guild, &msg.Channel, &msg.Createdat, &msg.Message)
		if err != nil {
			sendErr(ws, "We weren't able to get a list of messages. Please try again")
			golog.Warnf("Error scanning next row getting messages. Reason: %v", err)
			return
		}
		returnMsgs = append(returnMsgs, msg)
	}
	ws.Send(&globals.Packet{
		Type: "getmessages",
		Data: map[string]interface{}{
			"messages": returnMsgs,
		},
	})
}
