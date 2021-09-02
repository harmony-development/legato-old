package harmonydb

import "github.com/harmony-development/legato/db/sql/gen"

func (db *HarmonyDB) GetSession(session string) (int64, error) {
	return db.queries.GetSession(db, session)
}

func (db *HarmonyDB) SetSession(session string, userID int64) error {
	return db.queries.SetSession(db, gen.SetSessionParams{
		Userid:    userID,
		Sessionid: session,
	})
}
