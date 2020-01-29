package event

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"harmony-server/globals"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func sendErr(ws *globals.Client, msg string) {
	ws.Send(&globals.Packet{
		Type: "error",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func deauth(ws *globals.Client) {
	ws.Send(&globals.Packet{
		Type: "deauth",
		Data: map[string]interface{}{
			"message": "token is missing or invalid",
		},
	})
}

func sendToken(ws *globals.Client, id string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		golog.Warnf("Error signing token. Reason : %v", err) // pray to god this never happens
		return
	}

	ws.Send(&globals.Packet{
		Type: "token",
		Data: map[string]interface{}{
			"token":  tokenString,
			"userid": id,
		},
	})
}

func registerSocket(guildid string, ws *globals.Client, userid string) {
	if globals.Guilds[guildid] != nil {
		if globals.Guilds[guildid].Clients[userid] == nil {
			globals.Guilds[guildid].Clients[userid] = []*globals.Client{ws}
		} else {
			globals.Guilds[guildid].Clients[userid] = append(globals.Guilds[guildid].Clients[userid], ws)
		}
	} else {
		globals.Guilds[guildid] = &globals.Guild{
			Clients: map[string][]*globals.Client{
				userid: {ws},
			},
		}
	}
}

// DeleteFromFilestore deletes a file from the storage
func DeleteFromFilestore(fileid string) {
	err := os.Remove(fmt.Sprintf("./filestore/%v", fileid))
	if err != nil {
		golog.Warnf("Error deleting from filestore : %v", err)
	}
}
