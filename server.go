package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/bluskript/harmony-server/rest"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongoServer() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(port int, callback func(error)) {
	router := httprouter.New()
	http.Handle("/", http.FileServer(http.Dir("public/")))
	router.GET("/ping", rest.Ping)

	go startMongoServer()

	log.Println("Server successfully started on port " + strconv.Itoa(port))
	callback(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
