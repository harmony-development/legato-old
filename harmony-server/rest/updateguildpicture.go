package rest

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"gopkg.in/h2non/bimg.v1"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"io/ioutil"
	"net/http"
	"path"
)

func UpdateGuildPicture(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err, _, files := parseFileUpload(r)
	var guild = r.FormValue("guild")
	if err != nil || len(files) == 0 {
		golog.Debugf("Error updating avatar : %v", err)
		sendResp(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	file, err := files[0].Open()
	if err != nil {
		golog.Debugf("Error opening file : %v", err)
		sendResp(w, map[string]string{
			"error": "we were unable to read your file",
		})
		return
	}

	if !getVisitor("updateguildpicture", getIP(r)).Allow() {
		sendResp(w, map[string]string{
			"error": "You're sending too many files! Wait a bit and try again",
		})
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
		sendVibeCheck(w)
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
		sendVibeCheck(w)
		return
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), scaled, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		sendVibeCheck(w)
		return
	}

	var oldPictureID string
	err = harmonydb.DBInst.QueryRow("SELECT picture FROM guilds WHERE guildid=$1", guild).Scan(&oldPictureID)
	_, err = harmonydb.DBInst.Exec("UPDATE guilds SET picture=$1 WHERE guildid=$2", fname, guild)
	if err != nil {
		golog.Warnf("Error updating picture. %v", err)
		sendVibeCheck(w)
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