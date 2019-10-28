package handler

import (
	"context"
	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/mongodocs"
	"github.com/bluskript/harmony-server/socket"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type loginData struct {
	email    string
	password string
}

// MongoInstance is an instance of the mongo client to do queries with
var MongoInstance *mongo.Client

// LoginHandler authenticates harmony users
func LoginHandler(raw interface{}, ws *socket.WebSocket) {
	if globals.HarmonyServer.Collections["users"] == nil {
		log.Println(aurora.Red("Users collection not available").Bold())
		return
	}
	rawMap, ok := raw.(map[string]interface{})
	if ok {
		var data loginData
		if data.email, ok = rawMap["email"].(string); ok {
			if data.password, ok = rawMap["password"].(string); ok {
				var user mongodocs.User
				err := globals.HarmonyServer.Collections["users"].FindOne(context.TODO(), bson.D{
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
