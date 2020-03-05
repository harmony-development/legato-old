package main

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/globals"
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

func makeEventBus() *globals.EventBus {
	bus := &globals.EventBus{}
	bus.Bind("ping", limit(event.OnPing, 500 * time.Millisecond, 5))
	bus.Bind("subscribe", limit(event.OnSubscribe, 500 * time.Millisecond, 10))
	return bus
}


func handleSocket(ctx echo.Context) error {
	ws := socket.NewSocket(ctx.Response(), ctx.Request())
	ws.EventBus = globals.Bus
	return nil
}