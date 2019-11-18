package handler

import (
	"github.com/bluskript/harmony-server/tablestructs"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"

	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/socket"
	"github.com/logrusorgru/aurora"
	"golang.org/x/crypto/bcrypt"
)

type loginData struct {
	email    string
	password string
}

// LoginHandler authenticates harmony users
func LoginHandler(raw interface{}, ws *socket.WebSocket) {
	if globals.HarmonyServer.DatabaseInstance == nil {
		log.Println(aurora.Red("Database Instance not available").Bold())
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
	var user = new(tablestructs.User)
	err := globals.HarmonyServer.DatabaseInstance.QueryRow("SELECT id, email, password, username FROM users WHERE email=?", data.email).Scan(&user.Id, &user.Email, &user.Password, &user.Username)
	if err != nil {
		log.Println(aurora.Yellow("Error querying for user " + err.Error()).Bold())
		loginErr("Invalid Email Or Password", ws)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.password))
	if err != nil {
		loginErr("Invalid Email Or Password", ws)
		return
	}

	expireTime := time.Now().Add(30 * 24 * time.Hour).UTC()
	claims := &AuthToken{
		Userid: user.Id,
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
