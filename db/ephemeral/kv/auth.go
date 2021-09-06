// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package kv

import (
	"context"
)

func (db *database) GetCurrentStep(ctx context.Context, authID string) (string, error) {
	var step string
	_, err := db.store.Get(authID, &step)
	if err != nil {
		return "", err
	}
	return step, nil
}

func (db *database) SetStep(ctx context.Context, authID string, step string) error {
	return db.store.Set(authID, step)
}

func (db *database) DeleteSession(ctx context.Context, authID string) error {
	return db.store.Delete(authID)
}
