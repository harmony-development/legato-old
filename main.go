package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest"
	"net/http"
	"os"
)

const (
	PORT = ":2288"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		golog.Fatalf("Error loading .env! Reason : %v", err)
	}
	harmonydb.DBInst = harmonydb.OpenDB()
	golog.SetLevel(os.Getenv("VERBOSITY_LEVEL"))
	_ = os.Mkdir("./filestore", 0777)
	globals.Bus = *makeEventBus()
	golog.Infof("Started Harmony Server On Port %v", PORT)
	router := mux.NewRouter()
	router.HandleFunc("/api/socket", handleSocket)
	router.HandleFunc("/api/rest/fileupload", rest.FileUpload)
	router.PathPrefix("/filestore/").Handler(http.StripPrefix("/filestore/", http.FileServer(http.Dir("./filestore"))))
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("./static")))
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	golog.Fatalf("Fatal error caused server to crash! %v", http.ListenAndServe(PORT, router))
}
