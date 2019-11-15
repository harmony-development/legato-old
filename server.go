package main

import (
	"database/sql"
	"github.com/bluskript/harmony-server/globals"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bluskript/harmony-server/rest"
	"github.com/bluskript/harmony-server/socket"
	"github.com/bluskript/harmony-server/socket/handler"
	"github.com/gorilla/mux"
	. "github.com/logrusorgru/aurora"
	_ "github.com/mattn/go-sqlite3"
)

func openDatabase() {
	database, err := sql.Open("sqlite3", "harmony.db")
	if err != nil {
		log.Fatal(Red("unable to open harmony.db, harmony cannot continue " + err.Error()).Bold())
	}
		statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id TEXT NOT NULL PRIMARY KEY UNIQUE,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		username TEXT)`)
	if err != nil {
		log.Fatal(Red("Cannot prepare database initialization statements, harmony cannot continue. " + err.Error()).Bold())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(Red("Cannot execute initialization statements, harmony cannot continue " + err.Error()).Bold())
	}

	globals.HarmonyServer.DatabaseInstance = database
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	globals.HarmonyServer.SocketServer, err = socket.CreateServer(w, r)
	if err != nil {
		log.Printf("Error creating websocket server: %v", err)
		return
	}
	globals.HarmonyServer.SocketServer.On("ping", handler.PingHandler)
	globals.HarmonyServer.SocketServer.On("login", handler.LoginHandler)
	globals.HarmonyServer.SocketServer.On("register", handler.RegisterHandler)
}

func startServer(port int, callback func(error)) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(Red("Unable to read .env file! Please assign JWT_SECRET variable in .env!").Bold())
	}

	if _, err := os.Stat("./storage/"); os.IsNotExist(err) {
		err = os.Mkdir("./storage/", 0700)

		if err != nil {
			log.Fatal(Red("Error making storage folder").Bold(), err)
		}
	}

	globals.HarmonyServer.JwtSecret = os.Getenv("JWT_SECRET")
	globals.HarmonyServer.Router = mux.NewRouter()
	globals.HarmonyServer.Router.PathPrefix("/storage/").Handler(http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))
	globals.HarmonyServer.Router.HandleFunc("/api/attach", rest.HandleAttachment)
	globals.HarmonyServer.Router.HandleFunc("/api/socket", websocketHandler)
	openDatabase()
	log.Println(Green("Server successfully started on port " + strconv.Itoa(port)).Bold())
	callback(http.ListenAndServe(":"+strconv.Itoa(port), globals.HarmonyServer.Router))
}
