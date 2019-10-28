package handler

import (
	"context"
	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/socket"
	. "github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type registerData struct {
	email    string
	username string
	password string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// RegisterHandler handles user registration in harmony
func RegisterHandler(raw interface{}, ws *socket.WebSocket) {
	if globals.HarmonyServer.Collections["users"] == nil {
		log.Println(Red("Users collection not available").Bold())
		return
	}

	rawMap, ok := raw.(map[string]interface{})
	if ok {
		var data registerData
		data.email, ok = rawMap["email"].(string)
		if !ok {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Missing email",
			}).Raw()
			return
		}
		data.username, ok = rawMap["username"].(string)
		if !ok {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Invalid username",
			}).Raw()
			return
		}
		data.password, ok = rawMap["password"].(string)
		if !ok {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Missing Password",
			}).Raw()
			return
		}
		if !emailRegex.MatchString(data.email) {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Invalid email",
			}).Raw()
			return
		}

		if len(data.username) >= 20 {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Username is too long",
			}).Raw()
			return
		}

		if len(data.password) < 3 || len(data.password) >= 30 {
			ws.Out <- (&socket.Event{
				Name: "REGISTER_ERROR",
				Data: "Password must be between 3 and 30 characters long",
			}).Raw()
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(data.password), bcrypt.DefaultCost)

		if err != nil {
			whoops("REGISTER_ERROR", ws)
			return
		}

		// all inputs are valid here
		_, err = globals.HarmonyServer.Collections["users"].InsertOne(context.TODO(), bson.D{
			{Key: "email", Value: string(hash)},
			{Key: "username", Value: data.username},
			{Key: "password", Value: data.password},
		})
		if err != nil {
			log.Println(Red(err.Error()).Bold())
			whoops("REGISTER_ERROR", ws)
			return
		}
	}
}
