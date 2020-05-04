package events

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"harmony-server/server/http/socket/client"
	"harmony-server/server/state/guild"
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
	ws.UserID = userID
	if !event.Limiter.Allow() {
		ws.SendError("too many requests")
		return
	}
	res, err := e.DB.Query("SELECT guilds.guildid FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", userID)
	if err != nil {
		ws.SendError("Unable to get guilds list")
		logrus.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	for res.Next() {
		var guildID string
		err := res.Scan(&guildID)
		if err != nil {
			ws.SendError("Unable to subscribe to guilds")
			logrus.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		if e.State.Guilds[guildID] == nil {
			e.State.GuildsLock.Lock()
			e.State.Guilds[guildID] = &guild.Guild{}
			e.State.Guilds[guildID].AddClient(userID, &ws)
		} else {
			e.State.Guilds[guildID].AddClient(userID, &ws)
		}
	}
}
