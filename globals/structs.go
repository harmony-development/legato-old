package globals

import "harmony-server/socket"

type (
	Guild struct {
		Clients map[string]*socket.Client
	}
)

var Guilds = make(map[string]*Guild)