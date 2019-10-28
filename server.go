package main

import (
	"context"
	"github.com/bluskript/harmony-server/globals"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bluskript/harmony-server/rest"
	"github.com/bluskript/harmony-server/socket"
	"github.com/bluskript/harmony-server/socket/handler"
	"github.com/gorilla/mux"
	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongoServer() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Print(Red(err.Error()).Bold())
		return
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Print(Red(err.Error()).Bold())
	}
	globals.HarmonyServer.Collections["users"] = client.Database("harmony").Collection("users")
	globals.HarmonyServer.MongoInstance = client
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
	globals.HarmonyServer.Router = mux.NewRouter()
	globals.HarmonyServer.Router.Handle("/", http.FileServer(http.Dir("public/")))
	globals.HarmonyServer.Router.HandleFunc("/api/ping", rest.Ping)
	globals.HarmonyServer.Router.HandleFunc("/api/socket", websocketHandler)

	startMongoServer()

	log.Println(Green("Server successfully started on port " + strconv.Itoa(port)).Bold())
	callback(http.ListenAndServe(":"+strconv.Itoa(port), globals.HarmonyServer.Router))
}
