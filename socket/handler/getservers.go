package handler

import (
	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/socket"
	"github.com/logrusorgru/aurora"
	"log"
)

type getServerData struct {
	token string
}

func getServers(raw interface{}, ws *socket.WebSocket) {
	if globals.HarmonyServer.DatabaseInstance == nil {
		log.Println(aurora.Red("Database instance not available").Bold())
		return
	}
	rawMap, ok := raw.(map[string]interface{})
	if !ok {
		return
	}
	var data getServerData
	data.token, ok = rawMap["token"].(string)
	if !ok {
		return
	}
	
}
