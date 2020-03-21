package v1

import (
	"database/sql"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

type MessageData struct {
	Userid     string  `json:"userid"`
	Createdat  int     `json:"createdat"`
	Message    string  `json:"message"`
	Attachments []string `json:"attachment"`
	Messageid  string  `json:"messageid"`
}

func GetMessages(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	guild, channel, lastmessage := ctx.FormValue("guild"), ctx.FormValue("channel"), ctx.FormValue("lastmessage")
	if guild == "" || channel == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameters")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[ctx.User.ID] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to list messages")
	}
	var res *sql.Rows
	var err error
	if lastmessage == "" {
		res, err = db.DBInst.Query("SELECT messageid, author, createdat, message FROM messages WHERE guildid=$1 AND channelid=$2 ORDER BY createdat DESC LIMIT 30", guild, channel)
	} else {
		res, err = db.DBInst.Query("SELECT messageid, author, createdat, message FROM messages WHERE guildid=$1 AND channelid=$2 AND createdat < (SELECT createdat FROM messages WHERE messageid=$3) ORDER BY createdat DESC LIMIT 30", guild, channel, lastmessage)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameters")
	}
	var returnMsgs []MessageData
	for res.Next() {
		var msg MessageData
		err := res.Scan(&msg.Messageid, &msg.Userid, &msg.Createdat, &msg.Message)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to list messages, please try again later")
		}

		attachments, err := db.DBInst.Query("SELECT attachment FROM attachments WHERE messageid=$1", msg.Messageid)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to list messages, please try again later")
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
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": returnMsgs,
	})
}
