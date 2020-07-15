package events

import (
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/state"
)

// Events contains the events + their dependencies
type Events struct {
	DB     db.IHarmonyDB
	State  *state.State
	Logger logger.ILogger
}
