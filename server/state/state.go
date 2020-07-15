package state

import (
	"sync"

	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/server/state/guild"
)

// State contains the variables related to application state
type State struct {
	Guilds              map[uint64]*guild.Guild
	GuildsLock          *sync.RWMutex
	UserUpdateListeners map[*client.Client]struct{}
}
