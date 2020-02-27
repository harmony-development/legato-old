package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"golang.org/x/time/rate"
	"gopkg.in/h2non/bimg.v1"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"io/ioutil"
	"net/http"
	"path"
)

func UpdateGuildPicture(limiter *rate.Limiter, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err, _, files := parseFileUpload(r)
	var guild = mux.Vars(r)["guildid"]
	if err != nil || len(files) == 0 || guild == "" {
		golog.Debugf("Error updating guild picture : %v", err)
		http.Error(w, "invalid parameters", http.StatusBadRequest)
		return
	}

	file, err := files[0].Open()
	if err != nil {
		golog.Debugf("Error opening file : %v", err)
		http.Error(w, "error opening file", http.StatusInternalServerError)
		return
	}

	if !limiter.Allow() {
		http.Error(w, "too many requests", http.StatusTooManyRequests)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			golog.Warnf("Error closing file : %v", err)
		}
	}()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}
	scaled, err := bimg.NewImage(fileBytes).Process(bimg.Options{
		Height: 128,
		Width: 128,
		Quality: 60,
		Crop: true,
	})
	if err != nil {
		golog.Warnf("Error scaling image : %v", err)
		http.Error(w, "error resizing image", http.StatusInternalServerError)
		return
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), scaled, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		http.Error(w, "error saving image", http.StatusInternalServerError)
		return
	}

	var oldPictureID string
	err = harmonydb.DBInst.QueryRow("SELECT picture FROM guilds WHERE guildid=$1", guild).Scan(&oldPictureID)
	_, err = harmonydb.DBInst.Exec("UPDATE guilds SET picture=$1 WHERE guildid=$2", fname, guild)
	if err != nil {
		golog.Warnf("Error updating picture. %v", err)
		http.Error(w, "error linking picture to guild", http.StatusInternalServerError)
		go deleteFromFilestore(fname)
		return
	}
	go deleteFromFilestore(path.Base(oldPictureID))
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "updateguildpicture",
				Data: map[string]interface{}{
					"guild":   guild,
					"picture": fname,
					"success": true,
				},
			})
		}
	}
}