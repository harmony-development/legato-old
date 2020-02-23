package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
)

type authData struct {
	Token string `mapstructure:"token"`
}

func OnAuth(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data authData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Token == "" {
		sendErr(ws, "Token is missing")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "Too many authentications, try again in a few seconds")
		return
	}
	if _, err := authentication.VerifyToken(data.Token); err != nil {
		sendErr(ws, "invalid token")
		return
	}
	ws.Authed = true
}