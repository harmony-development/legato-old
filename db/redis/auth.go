// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package redis

import (
	"context"
	"time"
)

func (db *database) GetCurrentStep(ctx context.Context, authID string) (string, error) {
	return db.rdb.Get(ctx, authID).Result()
}

func (db *database) SetStep(ctx context.Context, authID string, step string) error {
	return db.rdb.Set(ctx, authID, step, 10*time.Minute).Err()
}

func (db *database) DeleteSession(ctx context.Context, authID string) error {
	return db.rdb.Del(ctx, authID).Err()
}
