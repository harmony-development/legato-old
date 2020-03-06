package event

import (
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"time"
)

func OnPing(ws *globals.Client, _ map[string]interface{}, limiter *rate.Limiter) {
	if !limiter.Allow() {
		return
	}
	ws.LastPong = time.Now()
}