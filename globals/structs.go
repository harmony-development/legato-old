package globals

import "harmony-server/socket"

type (
	Guild struct {
		Clients map[string]*socket.Client
		Owner string
	}
)

var Guilds = make(map[string]*Guild)