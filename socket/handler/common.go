package handler

import (
	"github.com/bluskript/harmony-server/socket"
	"github.com/dgrijalva/jwt-go"
)

// Claims is a structure for the authentication JWT
type Claims struct {
	Userid string
	jwt.StandardClaims
}

func whoops(name string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: name,
		Data: "Whoops! Seems like something went wrong on our end! Please try again later!",
	}).Raw()
}

func register(token string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "REGISTER",
		Data: token,
	}).Raw()
}

func regErr(reason string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "REGISTER_ERROR",
		Data: reason,
	}).Raw()
}

func login(token string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "LOGIN",
		Data: token,
	}).Raw()
}

func loginErr(reason string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "LOGIN_ERROR",
		Data: reason,
	}).Raw()
}
