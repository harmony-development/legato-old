package main

import (
	"log"
	"net/http"
	"strconv"

	"./rest"
	"github.com/julienschmidt/httprouter"
)

func startServer(port int, callback func(error)) {
	router := httprouter.New()
	http.Handle("/", http.FileServer(http.Dir("public/")))
	router.GET("/ping", rest.Ping)
	log.Println("Server successfully started on port " + strconv.Itoa(port))
	callback(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
