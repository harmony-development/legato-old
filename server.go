package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bluskript/harmony-server/globals"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/bluskript/harmony-server/rest"
	"github.com/bluskript/harmony-server/socket"
	"github.com/bluskript/harmony-server/socket/handler"
	"github.com/gorilla/mux"
	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongoServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	_, err = globals.HarmonyServer.Collections["users"].Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bsonx.Doc{{Key: "userid", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		log.Fatal(Red("Unable to create indexes : " + err.Error()).Bold())
	}
	globals.HarmonyServer.MongoInstance = client
	cancel()
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

	globals.HarmonyServer.JwtSecret = os.Getenv("JWT_SECRET")
	globals.HarmonyServer.Router = mux.NewRouter()
	globals.HarmonyServer.Router.Handle("/", http.FileServer(http.Dir("public/")))
	globals.HarmonyServer.Router.HandleFunc("/api/ping", rest.Ping)
	globals.HarmonyServer.Router.HandleFunc("/api/socket", websocketHandler)

	startMongoServer()

	log.Println(Green("Server successfully started on port " + strconv.Itoa(port)).Bold())
	callback(http.ListenAndServe(":"+strconv.Itoa(port), globals.HarmonyServer.Router))
}
