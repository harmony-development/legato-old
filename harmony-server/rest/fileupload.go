package rest

import (
	"fmt"
	"gopkg.in/h2non/bimg.v1"
	"harmony-server/authentication"
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
)

func sendVibeCheck(w http.ResponseWriter) {
	sendResp(w, map[string]string{
		"error": "Your file failed the vibe check, please try again later",
	})
}

// FileUpload handles all file upload requests sent to the server
func FileUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		return
	}
	token := r.FormValue("token")
	if token == "" {
		golog.Debugf("Invalid token received during file upload : %v", token)
		sendResp(w, map[string]string{
			"error": "Invalid token",
		})
		return
	}
	if _, err := authentication.VerifyToken(token); err != nil {
		sendResp(w, map[string]string{
			"error": "Invalid token",
		})
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		golog.Warnf("Error reading client file : %v", err)
		sendVibeCheck(w)
		return
	}

	defer file.Close()
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
	_, _ = w.Write([]byte(fname))
}
