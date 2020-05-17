package state

import (
	"harmony-server/server/state/guild"
	"sync"
)

// State contains the variables related to application state
type State struct {
	Guilds     map[uint64]*guild.Guild
	GuildsLock *sync.RWMutex
}
