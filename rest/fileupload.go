package rest

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"io/ioutil"
	"net/http"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(30 << 20)
	if err != nil {
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		golog.Warnf("Error reading client file : %v", err)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		golog.Warnf("Error reading uploaded file : %v", err)
	}
	fname := randstr.Hex(16)
	err = ioutil.WriteFile(fmt.Sprintf("./filestore/%v", fname), fileBytes, 0666)
	if err != nil {
		golog.Warnf("Error saving file upload : %v", err)
		return
	}
	_, _ = w.Write([]byte(fname))
}