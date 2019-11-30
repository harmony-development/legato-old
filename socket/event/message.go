package event

import (
	"harmony-server/globals"
	"harmony-server/socket"
)

type messageData struct {
	token       string
	targetGuild string
	message     string
}

func OnMessage(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data messageData
	if data.token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.message, ok = rawMap["message"].(string); !ok || len(data.message) == 0 {
		return
	}
	if data.targetGuild, ok = rawMap["guild"].(string); !ok {
		return
	}
	userid := verifyToken(data.token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	// either the guild doesn't exist or the client isn't subbed to it - it doesn't matter.
	if globals.Guilds[data.targetGuild] == nil || globals.Guilds[data.targetGuild].Clients[userid] == nil {
		return
	}

	// unfortunately O(n) is the only way to do this, we just need to make n really smol (︶︹︶)
	for _, client := range globals.Guilds[data.targetGuild].Clients {
		client.Send(&socket.Packet{
			Type: "message",
			Data: map[string]interface{}{
				"guild": data.targetGuild,
				"message": data.message,
			},
		})
	}
}
