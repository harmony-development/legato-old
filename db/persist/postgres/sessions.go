// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package postgres

import (
	"context"

	"github.com/harmony-development/legato/db/persist/sql/gen"
)

func (db *database) GetSession(ctx context.Context, session string) (int64, error) {
	return db.queries.GetSession(ctx, session)
}

func (db *database) AddSession(ctx context.Context, session string, userID int64) error {
	return db.queries.AddSession(ctx, gen.AddSessionParams{
		Userid:    userID,
		Sessionid: session,
	})
}
