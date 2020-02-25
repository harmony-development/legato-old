package rest

import "net/http"

func WithCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}