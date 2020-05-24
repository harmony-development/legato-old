package events

import (
	"encoding/json"

	"harmony-server/server/http/socket/client"
	"harmony-server/server/state/guild"

	"github.com/sirupsen/logrus"
)

type subscribeData struct {
	Session string
}

// Subscribe gets
func (e Events) Subscribe(ws client.Client, event *client.Event, raw *json.RawMessage) {
	var data subscribeData
	if err := json.Unmarshal(*raw, &data); err != nil {
		ws.SendError("bad request")
		return
	}
	userID, err := e.DB.SessionToUserID(data.Session)
	if err != nil {
		ws.SendError("invalid session")
		return
	}
	ws.UserID = &userID
	if !event.Limiter.Allow() {
		ws.SendError("too many requests")
		return
	}
	guildIDs, err := e.DB.GuildsForUser(userID)
	if err != nil {
		ws.SendError("Unable to get guilds list")
		logrus.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	for _, id := range guildIDs {
		if e.State.Guilds[id] == nil {
			e.State.GuildsLock.Lock()
			e.State.Guilds[id] = &guild.Guild{}
			e.State.Guilds[id].AddClient(&userID, &ws)
		} else {
			e.State.Guilds[id].AddClient(&userID, &ws)
		}
	}
}
