package event

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/socket"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func regErr(ws *socket.Client, msg string) {
	ws.Send(&socket.Packet{
		Type: "registererror",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func loginErr(ws *socket.Client, msg string) {
	ws.Send(&socket.Packet{
		Type: "loginerror",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func deauth(ws *socket.Client) {
	ws.Send(&socket.Packet{
		Type: "deauth",
		Data: map[string]interface{}{
			"message": "token is missing or invalid",
		},
	})
}

func sendToken(ws *socket.Client, id string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		golog.Warnf("Error signing token. Reason : %v", err) // pray to god this never happens
		return
	}

	ws.Send(&socket.Packet{
		Type: "token",
		Data: map[string]interface{}{
			"token": tokenString,
			"userid": id,
		},
	})
}

func VerifyToken(tokenstr string) string {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string)
	} else {
		return ""
	}
}

func registerSocket(guildid string, ws *socket.Client, userid string) {
	if globals.Guilds[guildid] != nil {
		globals.Guilds[guildid].Clients[userid] = ws
	} else {
		globals.Guilds[guildid] = &globals.Guild{
			Clients: map[string]*socket.Client{
				userid: ws,
			},
		}
	}
}

func DeleteFromFilestore(fileid string) {
	err := os.Remove(fmt.Sprintf("./filestore/%v", fileid))
	if err != nil {
		golog.Warnf("Error deleting from filestore : %v", err)
	}
}