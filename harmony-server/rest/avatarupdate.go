package rest

import (
	"fmt"
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

func AvatarUpdate(limiter *rate.Limiter, w http.ResponseWriter, r *http.Request) {
	WithCors(w)
	err, userid, files := parseFileUpload(r)
	if err != nil || len(files) == 0 {
		golog.Debugf("Error updating avatar : %v", err)
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
		http.Error(w, "too many avatar updates", http.StatusTooManyRequests)
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
		http.Error(w, "error saving file", http.StatusInternalServerError)
		return
	}
	var oldAvatarID string
	err = harmonydb.DBInst.QueryRow("SELECT avatar FROM users WHERE id=$1", userid).Scan(&oldAvatarID)
	_, err = harmonydb.DBInst.Exec("UPDATE users SET avatar=$1 WHERE id=$2", fname, userid)
	if err != nil {
		http.Error(w, "error linking avatar to profile", http.StatusInternalServerError)
		go deleteFromFilestore(fname)
		return
	}
	go deleteFromFilestore(path.Base(oldAvatarID))
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", userid)
	if err != nil {
		golog.Warnf("Error selecting guilds for avatarupdate : %v", err)
		http.Error(w, "error propagating avatar update", http.StatusInternalServerError)
		return
	}
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			golog.Warnf("Error getting guildid from result on avatarupdate : %v", err)
			http.Error(w, "error propagating avatar update", http.StatusInternalServerError)
			return
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				for _, conn := range client {
					conn.Send(&globals.Packet{
						Type: "avatarupdate",
						Data: map[string]interface{}{
							"userid": userid,
							"avatar": fname,
						},
					})
				}
			}
		}
	}
}