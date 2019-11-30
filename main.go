package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"harmony-server/harmonydb"
	"net/http"
	"os"
)

const (
	PORT = ":8080"
)

func main() {
	harmonydb.DBInst = harmonydb.OpenDB()
	err := godotenv.Load()
	if err != nil {
		golog.Fatalf("Error loading .env! Reason : %v", err)
	}
	golog.SetLevel(os.Getenv("VERBOSITY_LEVEL"))
	golog.Infof("Started Harmony Server On Port %v", PORT)
	router := mux.NewRouter()
	router.HandleFunc("/api/socket", handleSocket)
	golog.Fatalf("Fatal error caused server to crash! %v", http.ListenAndServe(PORT, router))
}
