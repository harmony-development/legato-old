package main

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/v1"
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
	r := echo.New()
	r.Use(middleware.Recover())
	api := r.Group("/api")
	apiV1(*api)
	r.Any("/api/socket", handleSocket)
	r.Static("/filestore", "filestore")
	r.Static("/static", "static")
	r.File("/", "static/index.html")
	go v1.CleanupRoutine()
	golog.Fatalf("Fatal error caused server to crash! %v", http.ListenAndServe(PORT, r))
}
