package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	maxFiles = 3
)

func Message(limiter *rate.Limiter, w http.ResponseWriter, r *http.Request) {
	WithCors(w)
	err, userid, files := parseFileUpload(r)
	message := r.FormValue("message")
	vars := mux.Vars(r)
	channel, guild := vars["guildid"], vars["channelid"]
	if err != nil || message == "" || channel == "" || guild == "" {
		golog.Debugf("Error receiving message : %v", err)
		http.Error(w, "invalid parameters", http.StatusBadRequest)
		return
	}

	if len(files) > maxFiles {
		http.Error(w, "too many files uploaded", http.StatusBadRequest)
		return
	}

	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[*userid] == nil {
		return
	}
	if !limiter.Allow() {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}
	fileTransaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		golog.Warnf("error making file transaction : %v", err)
		http.Error(w, "error beginning file transaction", http.StatusInternalServerError)
		return
	}
	var messageID = randstr.Hex(16)
	var attachments = make([]string, len(files))
	for i, v := range files {
		file, err := v.Open()
		if err != nil {
			golog.Warnf("Failed to parse file : %v", err)
			http.Error(w, "error opening file", http.StatusInternalServerError)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			golog.Warnf("Error reading uploaded file : %v", err)
			http.Error(w, "error reading file", http.StatusInternalServerError)
		}
		err = file.Close()
		if err != nil {
			golog.Warnf("Failed to close file : %v", err)
			http.Error(w, "error closing file", http.StatusInternalServerError)
		}
		fname := randstr.Hex(16)
		err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
		if err != nil {
			golog.Warnf("Error saving file upload : %v", err)
			http.Error(w, "error saving file", http.StatusInternalServerError)
		}
		_, err = fileTransaction.Exec("INSERT INTO attachments(messageid, attachment) VALUES($1, $2)", messageID, fname)
		if err != nil {
			golog.Warnf("Error inserting into attachments : %v", err)
			http.Error(w, "error linking file to message", http.StatusInternalServerError)
			go deleteFromFilestore(fname)
		} else {
			attachments[i] = fname
		}
	}
	err = fileTransaction.Commit()
	if err != nil {
		golog.Warnf("Error committing attachment transaction  : %v", err)
		http.Error(w, "error committing attachment transaction", http.StatusInternalServerError)
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
