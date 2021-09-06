// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package postgres

import (
	"context"

	"github.com/harmony-development/legato/db/persist"
)

type users struct {
	*database
}

func (db *users) Add(ctx context.Context, persist persist.UserInformation) error {
	panic("unimplemented")
}

func (db *users) Get(ctx context.Context, id uint64) (persist.UserInformation, error) {
	panic("unimplemented")
}

func (db *users) GetByEmail(ctx context.Context, email string) (persist.UserInformation, error) {
	it, err := db.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return persist.UserInformation{}, err
	}

	return persist.UserInformation{
		ID:       uint64(it.Userid),
		Password: it.Passwd,
	}, nil
}
