package rest

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/golog"
	"harmony-server/authentication"
	"mime/multipart"
	"net/http"
	"os"
)

func sendResp(w http.ResponseWriter, data interface{}) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		golog.Warnf("Error sending JSON Response : %v", err)
		return
	}
	_, _ = w.Write(marshalled)
}

func deleteFromFilestore(fileid string) {
	err := os.Remove(fmt.Sprintf("./filestore/%v", fileid))
	if err != nil {
		golog.Warnf("Error deleting from filestore : %v", err)
	}
}

func parseFileUpload(r *http.Request) (error, *string, []*multipart.FileHeader) {
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		return fmt.Errorf("error parsing form"), nil, nil
	}
	token := r.FormValue("token")
	if token == "" {
		golog.Debugf("Invalid token received during file upload : %v", token)
		return fmt.Errorf("invalid token"), nil, nil
	}
	var userid string
	if userid, err = authentication.VerifyToken(token); err != nil {
		return fmt.Errorf("invalid token"), nil, nil
	}

	files := r.MultipartForm.File["files"]

	return nil, &userid, files
}