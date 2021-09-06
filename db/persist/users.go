// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package persist

import "context"

type UserInformation struct {
	ID       uint64
	Email    string
	Password []byte
}

type Users interface {
	Add(ctx context.Context, user UserInformation) error

	Get(ctx context.Context, id uint64) (UserInformation, error)
	GetByEmail(ctx context.Context, email string) (UserInformation, error)
}
