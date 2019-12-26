package main

import (
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/socket"
	"harmony-server/socket/event"
	"net/http"
	"time"
)

// limit adds a ratelimit to an API path
// event is the event handler, duration is the time between requests, and burst is the amount of requests allowed to be done in an instant
func limit(event func(ws *globals.Client, data map[string]interface{}, limiter *rate.Limiter), duration time.Duration, burst int) globals.Event {
	limiter := rate.NewLimiter(rate.Every(duration), burst)
	return func(ws *globals.Client, data map[string]interface{}) {
		event(ws, data, limiter)
	}
}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	ws := socket.NewSocket(w, r)
	ws.Bind("login", limit(event.OnLogin, 10 * time.Second, 1))
	ws.Bind("register", limit(event.OnRegister, 10 * time.Minute, 1))
	ws.Bind("getguilds", limit(event.OnGetGuilds, 500 * time.Millisecond, 10))
	ws.Bind("message", limit(event.OnMessage, 100 * time.Millisecond, 10))
	ws.Bind("getmessages", limit(event.OnGetMessages, 100 * time.Millisecond, 5))
	ws.Bind("getchannels", limit(event.OnGetChannels, 100 * time.Millisecond, 5))
	ws.Bind("joinguild", limit(event.OnJoinGuild, 3 * time.Second, 1))
	ws.Bind("createguild", limit(event.OnCreateGuild, 20 * time.Second, 1))
	ws.Bind("leaveguild", limit(event.OnLeaveGuild, 3 * time.Second, 1))
	ws.Bind("updateguildpicture", limit(event.OnUpdateGuildPicture, 3 * time.Second, 1))
	ws.Bind("updateguildname", limit(event.OnUpdateGuildName, 3 * time.Second, 1))
	ws.Bind("getinvites", limit(event.OnGetInvites, 500 * time.Millisecond, 1))
	ws.Bind("addchannel", limit(event.OnAddChannel, 1 * time.Second, 5))
	ws.Bind("deletechannel", limit(event.OnDeleteChannel, 1 * time.Second, 5))
	ws.Bind("deleteinvite", limit(event.OnDeleteInvite, 200 * time.Millisecond, 5))
	ws.Bind("createinvite", limit(event.OnCreateInvite, 200 * time.Millisecond, 5))
	ws.Bind("getuser", limit(event.OnGetUser, 500 * time.Millisecond, 50))
}