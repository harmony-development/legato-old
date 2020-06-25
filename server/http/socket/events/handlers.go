package events

import (
	"harmony-server/server/db"
	"harmony-server/server/logger"
	"harmony-server/server/state"
)

// Events contains the events + their dependencies
type Events struct {
	DB     db.IHarmonyDB
	State  *state.State
	Logger logger.ILogger
}
