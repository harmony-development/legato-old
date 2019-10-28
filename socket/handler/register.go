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
			regErr("Missing Email", ws)
			return
		}
		data.username, ok = rawMap["username"].(string)
		if !ok {
			regErr("Missing Username", ws)
			return
		}
		data.password, ok = rawMap["password"].(string)
		if !ok {
			regErr("Missing Password", ws)
			return
		}
		if !emailRegex.MatchString(data.email) {
			regErr("Invalid Email", ws)
			return
		}

		if len(data.username) >= 20 {
			regErr("Your name is too long", ws)
			return
		}

		if len(data.password) < 3 || len(data.password) > 30 {
			regErr("Password must be at least 3 characters and at most 30 characters long", ws)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(data.password), bcrypt.DefaultCost)

		if err != nil {
			whoops("REGISTER_ERROR", ws)
			return
		}

		// all inputs are valid here
		_, err = globals.HarmonyServer.Collections["users"].InsertOne(context.TODO(), bson.D{
			{Key: "email", Value: data.email},
			{Key: "username", Value: data.username},
			{Key: "password", Value: string(hash)},
		})
		if err != nil {
			log.Println(Red(err.Error()).Bold())
			whoops("REGISTER_ERROR", ws)
			return
		}
	}
}
