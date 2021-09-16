// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package kv

import (
	"context"

	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/errwrap"
)

func (db *database) GetCurrentStep(ctx context.Context, authID string) (string, error) {
	var step string

	ok, err := db.store.Get(authID, &step)
	if err != nil {
		return "", errwrap.Wrapf(err, "failed to get step for ID %s", authID)
	}

	if !ok {
		return "", ephemeral.ErrStepNotFound
	}

	return step, nil
}

func (db *database) SetStep(ctx context.Context, authID string, step string) error {
	return errwrap.Wrapf(db.store.Set(authID, step), "failed to set step to %s for %s", step, authID)
}

func (db *database) DeleteAuthID(ctx context.Context, authID string) error {
	return errwrap.Wrapf(db.store.Delete(authID), "failed to delete auth ID %s", authID)
}
