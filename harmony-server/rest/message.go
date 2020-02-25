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

const (
	maxFiles = 3
)

func Message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err, userid, files := parseFileUpload(r)

	message, channel, guild := r.FormValue("message"), r.FormValue("channel"), r.FormValue("guild")
	if err != nil || message == "" || channel == "" || guild == "" {
		golog.Debugf("Error receiving message : %v", err)
		sendResp(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if len(files) > maxFiles {
		sendResp(w, "you uploaded more files than you're allowed to upload")
		return
	}

	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[*userid] == nil {
		return
	}
	if rateLimits[getIP(r)] != nil && !rateLimits[getIP(r)].limiter.Allow() {
		sendResp(w, map[string]string{
			"error": "You're sending too many messages with attachments! Please try again later",
		})
		return
	}
	fileTransaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		golog.Warnf("error making file transaction : %v", err)
		sendVibeCheck(w)
		return
	}
	var messageID = randstr.Hex(16)
	var attachments = make([]string, len(files))
	for i, v := range files {
		file, err := v.Open()
		if err != nil {
			golog.Warnf("Failed to parse file : %v", err)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			golog.Warnf("Error reading uploaded file : %v", err)
			return
		}
		err = file.Close()
		if err != nil {
			golog.Warnf("Failed to close file : %v", err)
			return
		}
		fname := randstr.Hex(16)
		err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
		if err != nil {
			golog.Warnf("Error saving file upload : %v", err)
			sendVibeCheck(w)
			return
		}
		_, err = fileTransaction.Exec("INSERT INTO attachments(messageid, attachment) VALUES($1, $2)", messageID, fname)
		if err != nil {
			golog.Warnf("Error inserting into attachments : %v", err)
			sendVibeCheck(w)
			return
		}
		attachments[i] = fname
	}
	err = fileTransaction.Commit()
	if err != nil {
		golog.Warnf("Error committing attachment transaction  : %v", err)
		sendVibeCheck(w)
		return
	}

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
					"messageid": messageID,
				},
			})
		}
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO messages(messageid, guildid, channelid, author, createdat, message) VALUES($1, $2, $3, $4, $5, $6)", messageID, guild, channel, userid, time.Now().UTC().Unix(), message)
}
