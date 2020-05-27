package state

import (
	"sync"

	"harmony-server/server/http/socket/client"
	"harmony-server/server/state/guild"
)

// State contains the variables related to application state
type State struct {
	Guilds              map[uint64]*guild.Guild
	GuildsLock          *sync.RWMutex
	UserUpdateListeners map[*client.Client]struct{}
}
