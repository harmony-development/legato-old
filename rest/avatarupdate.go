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

func AvatarUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err, userid, files := parseFileUpload(r)
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
			"error": err.Error(),
		})
		return
	}

	if !globals.GetRESTClient(*userid).Allow() {
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
	var oldAvatarID string
	err = harmonydb.DBInst.QueryRow("SELECT avatar FROM users WHERE id=$1", userid).Scan(&oldAvatarID)
	_, err = harmonydb.DBInst.Exec("UPDATE users SET avatar=$1 WHERE id=$2", fname, userid)
	if err != nil {
		sendResp(w, map[string]string{
			"error": "Something weird happened on our end and we weren't able to set your avatar. Please try again in a few moments",
		})
		return
	}
	go deleteFromFilestore(path.Base(oldAvatarID))
	res, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", userid)
	if err != nil {
		golog.Warnf("Error selecting guilds for avatarupdate : %v", err)
		return
	}
	for res.Next() {
		var guildid string
		err = res.Scan(&guildid)
		if err != nil {
			golog.Warnf("Error getting guildid from result on avatarupdate : %v", err)
			return
		}
		if globals.Guilds[guildid] != nil {
			for _, client := range globals.Guilds[guildid].Clients  {
				golog.Debugf("Client conns : %v", client)
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