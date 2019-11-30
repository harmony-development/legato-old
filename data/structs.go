package data

import "harmony-server/socket"

type (
	Guild struct {
		clients map[string]*socket.Client
	}
)