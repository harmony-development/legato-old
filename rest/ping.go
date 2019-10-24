package rest

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Fatal(err.Error())
	}
}