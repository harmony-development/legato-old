package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/rest/v1"
	"harmony-server/socket"
	"harmony-server/socket/event"
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

func apiV1(g echo.Group) {
	r := g.Group("/v1")
	r.Use(middleware.CORS())
	r.POST("/login*", v1.WithRateLimit(v1.Login, 5 * time.Second, 1))
	r.POST("/register*", v1.WithRateLimit(v1.Register, 10 * time.Minute, 10))
	r.POST("/getguilds*", v1.WithRateLimit(v1.GetGuilds, 5 * time.Second, 3))
	r.POST("/getchannels/:guildid", v1.WithRateLimit(v1.GetChannels, 500 * time.Second, 5))
	r.POST("/avatarupdate*", v1.WithRateLimit(v1.AvatarUpdate, 3 * time.Second, 1))
	r.POST("/updateguildpicture/:guildid/", v1.WithRateLimit(v1.UpdateGuildPicture, 3 * time.Second, 1))
	r.POST("/message/:guildid/:channelid/*", v1.WithRateLimit(v1.Message, 500 * time.Millisecond, 20))
}

func makeEventBus() *globals.EventBus {
	bus := &globals.EventBus{}
	bus.Bind("ping", limit(event.OnPing, 500 * time.Millisecond, 5))
	bus.Bind("subscribe", limit(event.OnSubscribe, 500 * time.Millisecond, 10))
	bus.Bind("getmembers", limit(event.OnGetMembers, 5 * time.Second, 3))
	bus.Bind("getmessages", limit(event.OnGetMessages, 100 * time.Millisecond, 5))
	bus.Bind("getinvites", limit(event.OnGetInvites, 500 * time.Millisecond, 1))

	bus.Bind("message", limit(event.OnMessage, 100 * time.Millisecond, 10))
	bus.Bind("deletemessage", limit(event.OnDeleteMessage, 1 * time.Second, 8))

	bus.Bind("joinguild", limit(event.OnJoinGuild, 3 * time.Second, 1))
	bus.Bind("leaveguild", limit(event.OnLeaveGuild, 3 * time.Second, 1))

	bus.Bind("createguild", limit(event.OnCreateGuild, 20 * time.Second, 1))
	bus.Bind("deleteguild", limit(event.OnDeleteGuild, 10 * time.Second, 1))
	
	bus.Bind("updateguildname", limit(event.OnUpdateGuildName, 3 * time.Second, 1))

	bus.Bind("addchannel", limit(event.OnAddChannel, 1 * time.Second, 5))
	bus.Bind("createinvite", limit(event.OnCreateInvite, 200 * time.Millisecond, 5))

	bus.Bind("deletechannel", limit(event.OnDeleteChannel, 1 * time.Second, 5))
	bus.Bind("deleteinvite", limit(event.OnDeleteInvite, 200 * time.Millisecond, 5))

	bus.Bind("getuser", limit(event.OnGetUser, 500 * time.Millisecond, 50))
	bus.Bind("getself", limit(event.OnGetSelf, 500 * time.Millisecond, 20))

	bus.Bind("usernameupdate", limit(event.OnUsernameUpdate, 10 * time.Second, 1))
	return bus
}


func handleSocket(ctx echo.Context) error {
	ws := socket.NewSocket(ctx.Response(), ctx.Request())
	ws.EventBus = globals.Bus
	return nil
}