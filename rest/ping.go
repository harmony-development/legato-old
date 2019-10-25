package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoInstance *mongo.Client

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(MongoInstance.ListDatabaseNames(context.TODO(), bson.D{}))
}
