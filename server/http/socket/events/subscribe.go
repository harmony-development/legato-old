package events

import (
	"encoding/json"
	"sync"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/http/responses"

	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/server/state/guild"
)

type subscribeData struct {
	Session string
}

// Subscribe gets
func (e Events) Subscribe(ws client.Client, event *client.Event, raw *json.RawMessage) {
	var data subscribeData
	if err := json.Unmarshal(*raw, &data); err != nil {
		ws.SendError(responses.InvalidRequest)
		return
	}
	userID, err := e.DB.SessionToUserID(data.Session)
	if err != nil {
		ws.SendError(responses.InvalidSession)
		return
	}
	ws.UserID = &userID
	if !event.Limiter.Allow() {
		ws.SendError(responses.TooManyRequests)
		return
	}
	guildIDs, err := e.DB.GuildsForUser(userID)
	if err != nil {
		ws.SendError(responses.UnknownError)
		e.Logger.Exception(err)
		return
	}
	for _, id := range guildIDs {
		if e.State.Guilds[id] == nil {
			e.State.GuildsLock.Lock()
			e.State.Guilds[id] = &guild.Guild{
				Clients: make(map[uint64]*guild.ClientArray),
				RWMutex: &sync.RWMutex{},
			}
			e.State.Guilds[id].AddClient(&userID, &ws)
			e.State.GuildsLock.Unlock()
		} else {
			e.State.Guilds[id].AddClient(&userID, &ws)
		}
	}
	if err := e.DB.SetStatus(userID, profilev1.UserStatus_USER_STATUS_OFFLINE); err != nil {
		ws.SendError(responses.UnknownError)
		e.Logger.Exception(err)
	}
}
