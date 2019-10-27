package handler

import (
	"github.com/bluskript/harmony-server/socket"
)

// PingHandler handles the ping socket event
func PingHandler(data interface{}, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "response",
		Data: "pong!",
	}).Raw()
}
