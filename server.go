package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bluskript/harmony-server/rest"
	"github.com/julienschmidt/httprouter"
	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongoServer() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(Red(err.Error()).Bold())
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(Red(err.Error()).Bold())
	}
	rest.MongoInstance = client
}

func startServer(port int, callback func(error)) {
	router := httprouter.New()
	http.Handle("/", http.FileServer(http.Dir("public/")))
	router.GET("/ping", rest.Ping)

	startMongoServer()
	startSocketServer()

	log.Println(Green("Server successfully started on port " + strconv.Itoa(port)).Bold())
	callback(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
