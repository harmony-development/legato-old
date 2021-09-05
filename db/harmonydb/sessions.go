package harmonydb

import (
	"context"

	"github.com/harmony-development/legato/db/sql/gen"
)

func (db *HarmonySessionDB) GetSession(ctx context.Context, session string) (int64, error) {
	return db.queries.GetSession(ctx, session)
}

func (db *HarmonySessionDB) SetSession(ctx context.Context, session string, userID int64) error {
	return db.queries.SetSession(ctx, gen.SetSessionParams{
		Userid:    userID,
		Sessionid: session,
	})
}
