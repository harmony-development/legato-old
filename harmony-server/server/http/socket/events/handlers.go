package events

import (
	"harmony-server/server/db"
	"harmony-server/server/logger"
	"harmony-server/server/state"
)

// Events contains the events + their dependencies
type Events struct {
	DB     *db.HarmonyDB
	State  *state.State
	Logger *logger.Logger
}
