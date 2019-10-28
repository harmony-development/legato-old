package handler

import "github.com/bluskript/harmony-server/socket"

func whoops(name string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: name,
		Data: "Whoops! Seems like something went wrong on our end! Please try again later!",
	}).Raw()
}