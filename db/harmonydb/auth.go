package harmonydb

import "time"

func (db *HarmonyDB) GetCurrentStep(authID string) (string, error) {
	return db.rdb.Get(db, authID).Result()
}

func (db *HarmonyDB) SetStep(authID string, step string) error {
	return db.rdb.Set(db, authID, step, 10*time.Minute).Err()
}
