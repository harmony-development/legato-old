package events

import (
	"encoding/json"
	"harmony-server/server/http/socket/client"
)

type subscribeToGuildData struct {
	Session string
	Guild   string
}

// OnSubscribeToGuild handles requests to subscribe to a guilds events
func (e Events) OnSubscribeToGuild(ws client.Client, event *client.Event, raw *json.RawMessage) {
	var data subscribeToGuildData
	if err := json.Unmarshal(*raw, &data); err != nil {
		ws.SendError("bad request")
		return
	}
	userID, err := e.DB.SessionToUserID(data.Session)
	if err != nil {
		ws.SendError("invalid session")
		return
	}
	if !event.Limiter.Allow() {
		ws.SendError("too many subscription attempts, please try later")
		return
	}
	var count int
	res, err := e.DB.Query("SELECT COUNT(*) FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1 AND guilds.guildid=$2", userID, data.Guild)
	if err != nil {
		ws.SendError("unable to subscribe to guild")
		return
	}
	err = res.Scan(&count)
	if err != nil {
		ws.SendError("unable to subscribe to guild")
		return
	}
	if count == 1 {
		e.State.Guilds[data.Guild].AddClient(userID, &ws)
	}
	return
}
