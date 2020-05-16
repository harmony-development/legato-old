package events

import (
	"harmony-server/server/db"
	"harmony-server/server/state"
)

// Events contains the event events + their dependencies
type Events struct {
	DB    *db.HarmonyDB
	State *state.State
}

// NewEvents creates a new events instance
func NewEvents(db *db.HarmonyDB, state *state.State) Events {
	return Events{
		DB:    db,
		State: state,
	}
}
