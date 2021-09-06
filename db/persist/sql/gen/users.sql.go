// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package gen

import (
	"context"
)

const getUserByEmail = `-- name: GetUserByEmail :one

SELECT
    UserID,
    Passwd
FROM Users WHERE Email=$1 LIMIT 1
`

type GetUserByEmailRow struct {
	Userid int64
	Passwd []byte
}

// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(&i.Userid, &i.Passwd)
	return i, err
}
