package handler

import (
	"fmt"
	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/socket"
	"github.com/dgrijalva/jwt-go"
)

// AuthToken is a structure for the authentication JWT
type AuthToken struct {
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

func deauth(ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "DEAUTH",
		Data: nil,
	}).Raw()
}

func verifyToken(rawToken string) (*AuthToken, error) {
	token, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return globals.HarmonyServer.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(AuthToken); ok && token.Valid {
		return &claims, nil
	}
	return nil, nil
}

func loginErr(reason string, ws *socket.WebSocket) {
	ws.Out <- (&socket.Event{
		Name: "LOGIN_ERROR",
		Data: reason,
	}).Raw()
}
