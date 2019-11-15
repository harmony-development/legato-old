package rest

import (
	"github.com/thanhpk/randstr"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func HandleAttachment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(16 << 20)
	if err != nil {
		return
	}
	file, handler, err := r.FormFile("attachment")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}

	err = ioutil.WriteFile("./storage/" + randstr.Hex(16) + filepath.Ext(handler.Filename), fileBytes, 0644)
	if err != nil {
		log.Println(err)
		return
	}

}