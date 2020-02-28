package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/rest"
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
	v1 := g.Group("/v1")
	v1.Use(middleware.CORS())
	v1.POST("/login/", rest.WithRateLimit(rest.Login, 10 * time.Second, 1))
	v1.POST("/avatarupdate/", rest.WithRateLimit(rest.AvatarUpdate, 3 * time.Second, 1))
	v1.POST("/updateguildpicture/:guildid/", rest.WithRateLimit(rest.UpdateGuildPicture, 3 * time.Second, 1))
	v1.POST("/message/:guildid/:channelid/*", rest.WithRateLimit(rest.Message, 500 * time.Millisecond, 20))
}

func makeEventBus() *globals.EventBus {
	bus := &globals.EventBus{}
	bus.Bind("login", limit(event.OnLogin, 10 * time.Second, 1))
	bus.Bind("register", limit(event.OnRegister, 1 * time.Hour, 1))
	bus.Bind("ping", limit(event.OnPing, 500 * time.Millisecond, 5))

	bus.Bind("getguilds", limit(event.OnGetGuilds, 500 * time.Millisecond, 10))
	bus.Bind("getchannels", limit(event.OnGetChannels, 100 * time.Millisecond, 5))
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