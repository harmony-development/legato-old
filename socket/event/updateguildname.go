package event

import "harmony-server/socket"

type updateGuildName struct {
	Token string
	Guild string
	Name string
}

func OnUpdateGuildName(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data updateGuildName
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
		return
	}
	if data.Name, ok = rawMap["name"].(string); !ok {
		return
	}
}