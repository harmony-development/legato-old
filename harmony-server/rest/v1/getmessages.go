package v1

import (
	"database/sql"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type GetMessagesData struct {
	Token       string `mapstructure:"token"`
	Guild       string `mapstructure:"guild"`
	Channel     string `mapstructure:"channel"`
	LastMessage string `mapstructure:"lastmessage"`
}

type Message struct {
	Userid     string  `json:"userid"`
	Createdat  int     `json:"createdat"`
	Message    string  `json:"message"`
	Attachments []string `json:"attachment"`
	Messageid  string  `json:"messageid"`
}

func GetMessages(limiter *rate.Limiter, ctx echo.Context) error {
	var data GetMessagesData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		event.sendErr(ws, "Something was wrong with your request. Please try again")
		return
	}
	if data.Guild == "" || data.Token == "" || data.Channel == "" {
		event.sendErr(ws, "Something was wrong with your request. Please try again")
		golog.Warnf("Error decoding getmessages request")
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil {
		return
	}
	var res *sql.Rows
	var err error
	if data.LastMessage == "" {
		res, err = harmonydb.DBInst.Query("SELECT messageid, author, createdat, message FROM messages WHERE guildid=$1 AND channelid=$2 ORDER BY createdat DESC LIMIT 30", data.Guild, data.Channel)
	} else {
		res, err = harmonydb.DBInst.Query("SELECT messageid, author, createdat, message FROM messages WHERE guildid=$1 AND channelid=$2 AND createdat < (SELECT createdat FROM messages WHERE messageid=$3) ORDER BY createdat DESC LIMIT 30", data.Guild, data.Channel, data.LastMessage)
	}
	if err != nil {
		event.sendErr(ws, "We weren't able to get a list of messages. Please try again")
		golog.Warnf("Error getting recent messages : %v", ws.Userid)
		return
	}
	var returnMsgs []Message
	for res.Next() {
		var msg Message
		err := res.Scan(&msg.Messageid, &msg.Userid, &msg.Createdat, &msg.Message)
		if err != nil {
			event.sendErr(ws, "We weren't able to get a list of messages. Please try again")
			golog.Warnf("Error scanning next row getting messages. Reason: %v", err)
			return
		}

		attachments, err := harmonydb.DBInst.Query("SELECT attachment FROM attachments WHERE messageid=$1", msg.Messageid)
		if err != nil {
			golog.Warnf("Error scanning for attachments : %v", err)
			return
		}

		for attachments.Next() {
			var attachment string
			err := attachments.Scan(&attachment)
			if err != nil {
				golog.Warnf("Error scanning attachment : %v", err)
			}
			msg.Attachments = append(msg.Attachments, attachment)
		}
		
		returnMsgs = append(returnMsgs, msg)
	}
	if data.LastMessage == "" {
		ws.Send(&globals.Packet{
			Type: "getmessages",
			Data: map[string]interface{}{
				"messages": returnMsgs,
			},
		})
	} else {
		ws.Send(&globals.Packet{
			Type: "getmessages-old",
			Data: map[string]interface{}{
				"messages": returnMsgs,
			},
		})
	}
}
