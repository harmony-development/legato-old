// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package postgres

import (
	"context"

	"github.com/harmony-development/legato/db/persist/sql/gen"
)

func (db *database) GetUserByEmail(ctx context.Context, email string) (gen.GetUserByEmailRow, error) {
	return db.queries.GetUserByEmail(ctx, email)
}
