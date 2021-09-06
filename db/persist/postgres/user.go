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

func (db *users) Add(ctx context.Context, pers persist.UserInformation, ext persist.ExtendedUserInformation) error {
	panic("unimplemented")
}

func (db *users) Get(ctx context.Context, id uint64) (ui persist.UserInformation, eui persist.ExtendedUserInformation, err error) {
	panic("unimplemented")
}

func (db *users) GetLocalByEmail(ctx context.Context, email string) (persist.UserInformation, persist.LocalUserInformation, error) {
	panic("unimplemented")
}
