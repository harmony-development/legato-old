package harmonydb

import (
	"context"
	"time"
)

func (db *HarmonyAuthDB) GetCurrentStep(ctx context.Context, authID string) (string, error) {
	return db.rdb.Get(ctx, authID).Result()
}

func (db *HarmonyAuthDB) SetStep(ctx context.Context, authID string, step string) error {
	return db.rdb.Set(ctx, authID, step, 10*time.Minute).Err()
}

func (db *HarmonyAuthDB) DeleteSession(ctx context.Context, authID string) error {
	return db.rdb.Del(ctx, authID).Err()
}
