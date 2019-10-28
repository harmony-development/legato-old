package handler

import (
	"regexp"
	"github.com/bluskript/harmony-server/socket"
)

type registerData struct {
	email string
	username string
	password string
}

// RegisterHandler handles user registration in harmony
func RegisterHandler(raw interface{}, ws *socket.WebSocket) {
	rawmap, ok := raw.(map[string]interface{})
	if ok {
		var data registerData
		if data.email, ok = rawmap["email"].(string); ok {
			if data.username, ok = rawmap["username"].(string); ok {
				if data.password, ok = rawmap["password"].(string); ok {
					if regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(data.email) {

					} else {
						ws.Out <- (&socket.Event{
							Name: "REGISTER_ERROR",
							Data: "Invalid email",
						}).Raw()
					}
				} else {
					ws.Out <- (&socket.Event{
							Name: "REGISTER_ERROR",
							Data: "Missing Password",
					}).Raw()
					return
				}
			} else {
				ws.Out <- (&socket.Event{
					Name: "REGISTER_ERROR",
					Data: "Invalid email",
				}).Raw()
				return
			}
		} else {
			ws.Out <- (&socket.Event{
							Name: "REGISTER_ERROR",
							Data: "Invalid email",
						}).Raw()
			return
		}
	} else {
		return
	}
}
