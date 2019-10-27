package main

import (
	"context"
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
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Print(Red(err.Error()).Bold())
	}
	rest.MongoInstance = client
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := socket.CreateServer(w, r)
	if err != nil {
		log.Printf("Error creating websocket server: %v", err)
		return
	}
	ws.On("ping", handler.PingHandler)
}

func startServer(port int, callback func(error)) {
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("public/")))
	router.HandleFunc("/api/ping", rest.Ping)
	router.HandleFunc("/api/socket", websocketHandler)

	startMongoServer()

	log.Println(Green("Server successfully started on port " + strconv.Itoa(port)).Bold())
	callback(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
