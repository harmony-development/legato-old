package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"harmony-server/harmonydb"
	"harmony-server/rest"
	"net/http"
	"os"
)

const (
	PORT = ":2288"
)

func main() {
	harmonydb.DBInst = harmonydb.OpenDB()
	err := godotenv.Load()
	if err != nil {
		golog.Fatalf("Error loading .env! Reason : %v", err)
	}
	golog.SetLevel(os.Getenv("VERBOSITY_LEVEL"))
	_ = os.Mkdir("./filestore", 0777)
	golog.Infof("Started Harmony Server On Port %v", PORT)
	router := mux.NewRouter()
	router.HandleFunc("/api/socket", handleSocket)
	router.HandleFunc("/api/rest/fileupload", rest.FileUpload)
	golog.Fatalf("Fatal error caused server to crash! %v", http.ListenAndServe(PORT, router))
}
