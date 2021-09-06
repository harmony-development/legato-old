// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package postgres

import (
	"context"

	"github.com/harmony-development/legato/db/persist/sql/gen"
)

type sessions struct {
	*database
}

func (db *sessions) Get(ctx context.Context, session string) (uint64, error) {
	val, err := db.queries.GetSession(ctx, session)
	if err != nil {
		return 0, err
	}

	return uint64(val), nil
}

func (db *sessions) Add(ctx context.Context, session string, userID uint64) error {
	return db.queries.AddSession(ctx, gen.AddSessionParams{
		Userid:    int64(userID),
		Sessionid: session,
	})
}
