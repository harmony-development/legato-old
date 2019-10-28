package handler

import (
	"context"
	"log"

	"github.com/bluskript/harmony-server/mongodocs"
	"github.com/bluskript/harmony-server/socket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type loginData struct {
	email    string
	password string
}

// MongoInstance is an instance of the mongo client to do queries with
var MongoInstance *mongo.Client

// LoginHandler authenticates harmony users
func LoginHandler(raw interface{}, ws *socket.WebSocket) {
	rawmap, ok := raw.(map[string]interface{})
	if ok {
		var data loginData
		if data.email, ok = rawmap["email"].(string); ok {
			if data.password, ok = rawmap["password"].(string); ok {
				userCollection := MongoInstance.Database("harmony").Collection("users")
				var user mongodocs.User
				err := userCollection.FindOne(context.TODO(), bson.D{
					{"email", data.email},
				}).Decode(&user)
				if err != nil {
					log.Println(err.Error())
					return
				}
				log.Println(user)
			} else {
				return
			}
		} else {
			return
		}
	} else {
		return
	}
}
