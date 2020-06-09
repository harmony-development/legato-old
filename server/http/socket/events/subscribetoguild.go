package events

import (
	"encoding/json"
	"harmony-server/server/http/responses"
	"harmony-server/server/http/socket/client"
)

type subscribeToGuildData struct {
	Session string
	Guild   uint64
}

// SubscribeToGuild handles requests to subscribe to a guilds events
func (e Events) SubscribeToGuild(ws client.Client, event *client.Event, raw *json.RawMessage) {
	var data subscribeToGuildData
	if err := json.Unmarshal(*raw, &data); err != nil {
		ws.SendError("bad request")
		return
	}
	userID, err := e.DB.SessionToUserID(data.Session)
	if err != nil {
		ws.SendError(responses.InvalidSession)
		return
	}
	if !event.Limiter.Allow() {
		ws.SendError(responses.TooManyRequests)
		return
	}
	var count int
	inGuild, err := e.DB.UserInGuild(userID, data.Guild)
	if err != nil {
		ws.SendError(err.Error())
	}
	if !inGuild {
		ws.SendError(responses.NotInGuild)
	}
	if count == 1 {
		e.State.Guilds[data.Guild].AddClient(&userID, &ws)
	}
}
