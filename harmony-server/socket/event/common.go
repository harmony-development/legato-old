package event

import (
	"harmony-server/globals"
	"os"
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

func Deauth(ws *globals.Client) {
	ws.Send(&globals.Packet{
		Type: "deauth",
		Data: map[string]interface{}{
			"message": "token is missing or invalid",
		},
	})
}

func RegisterSocket(guildid string, ws *globals.Client, userid string) {
	if globals.Guilds[guildid] != nil {
		globals.Guilds[guildid].Lock.Lock()
		if globals.Guilds[guildid].Clients[userid] == nil {
			globals.Guilds[guildid].Clients[userid] = []*globals.Client{ws}
		} else {
			globals.Guilds[guildid].Clients[userid] = append(globals.Guilds[guildid].Clients[userid], ws)
		}
		globals.Guilds[guildid].Lock.Unlock()
	} else {
		globals.GuildsLock.Lock()
		globals.Guilds[guildid] = &globals.Guild{
			Clients: map[string][]*globals.Client{
				userid: {ws},
			},
		}
		globals.GuildsLock.Unlock()
	}
}
