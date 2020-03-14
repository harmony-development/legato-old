package main

import (
	"github.com/joho/godotenv"
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"harmony-auth-server/authentication"
	"harmony-auth-server/db"
	"harmony-auth-server/rest"
	"net/http"
	"os"
)

const (
	PORT = ":2289"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		golog.Fatalf("Error loading .env! %v", err)
	}
	golog.SetLevel(os.Getenv("VERBOSITY_LEVEL"))
	authentication.Init()
	db.DB = *db.OpenDB()
	_ = os.Mkdir("./avatars", 0777)
	r := echo.New()
	r.Pre(middleware.AddTrailingSlash())
	r.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisableStackAll:   true,
	}))
	api := r.Group("/api")
	rest.Setup(*api)
	r.Static("/avatars", "avatars")
	golog.Infof("Started Harmony AUTHENTICATION Server On Port %v", PORT)
	golog.Fatalf("Fatal error caused server to crash! %v", http.ListenAndServe(PORT, r))
}
