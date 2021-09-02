package harmonydb

import (
	"context"

	"github.com/harmony-development/legato/db/sql/gen"
)

// The default Harmony DB implementation. This uses sqlc.
type HarmonyDB struct {
	// embed a context for easy use in queries
	context.Context
	queries gen.Queries
}

// var _ db.Auth = HarmonyDB{}
