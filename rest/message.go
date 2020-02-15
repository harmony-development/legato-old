package rest

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"io/ioutil"
	"net/http"
	"time"
)

func Message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err, userid, file := parseFileUpload(r)
	message, channel, guild := r.FormValue("message"), r.FormValue("channel"), r.FormValue("guild")
	if err != nil || message == "" || channel == "" || guild == "" {
		golog.Debugf("Error receiving message : %v", err)
		sendResp(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[*userid] == nil {
		return
	}
	if !globals.GetRESTClient(*userid).Allow() {
		sendResp(w, map[string]string{
			"error": "You're sending too many messages with attachments! Please try again later",
		})
		return
	}
	defer (*file).Close()
	fileBytes, err := ioutil.ReadAll(*file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		sendVibeCheck(w)
		return
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		sendVibeCheck(w)
		return
	}
	var messageID = randstr.Hex(16)
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "message",
				Data: map[string]interface{}{
					"guild":     guild,
					"channel":   channel,
					"userid":    userid,
					"createdat": time.Now().UTC().Unix(),
					"message":   message,
					"attachment": fname,
					"messageid": messageID,
				},
			})
		}
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO messages(messageid, guildid, channelid, author, createdat, message, attachment) VALUES($1, $2, $3, $4, $5, $6, $7)", messageID, guild, channel, userid, time.Now().UTC().Unix(), message, fname)
}
