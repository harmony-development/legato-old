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
	token, err := verifyToken(data.token)

	if err != nil {
		deauth(ws)
		return
	}
	servers, err := globals.HarmonyServer.DatabaseInstance.Query("SELECT server_id FROM user_servers WHERE user_id=?", token.Userid)
	if err != nil {
		return
	}
	defer func() {
		err := servers.Close()
		if err != nil {
			log.Printf("Error closing row %v", err)
		}
	}()
	err = servers.Err()
	if err != nil {
		log.Printf(err.Error())
	}
	for servers.Next() {
		
	}

}
