package rest

import (
	"fmt"
	"harmony-server/authentication"
	"io/ioutil"
	"net/http"

	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
)

// FileUpload handles all file upload requests sent to the server
func FileUpload(w http.ResponseWriter, r *http.Request) {
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
		sendResp(w, map[string]string{
			"error": "Your file failed the vibe check, please try again later",
		})
		return
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
		return
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		return
	}
	_, _ = w.Write([]byte(fname))
}
