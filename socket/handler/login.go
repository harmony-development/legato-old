package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"

	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/mongodocs"
	"github.com/bluskript/harmony-server/socket"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	if !ok {
		loginErr("Invalid arguments", ws)
		return
	}
	var data loginData
	data.email, ok = rawMap["email"].(string)
	if !ok {
		loginErr("Invalid Email", ws)
		return
	}
	data.password, ok = rawMap["password"].(string)
	if !ok {
		loginErr("Invalid Password", ws)
		return
	}
	var user *mongodocs.User
	err := globals.HarmonyServer.Collections["users"].FindOne(context.TODO(), bson.D{
		{"email", data.email},
	}).Decode(&user)
	if err != nil {
		loginErr("Invalid Email Or Password", ws)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.password))
	if err != nil {
		loginErr("Invalid Email Or Password", ws)
		return
	}

	expireTime := time.Now().Add(30 * 24 * time.Hour).UTC()
	claims := &Claims{
		Userid: user.Userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(globals.HarmonyServer.JwtSecret))
	if err != nil {
		whoops("LOGIN", ws)
	}
	login(tokenString, ws)
}
