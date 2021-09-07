// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package persist

import "context"

type UserInformation struct {
	ID       uint64
	Username string
}

type (
	ExtendedUserInformation interface{ IsUserInfo() }
	isUserInfo              struct{}
)

func (isUserInfo) IsUserInfo() {}

type LocalUserInformation struct {
	Email    string
	Password []byte

	isUserInfo
}

type ForeignUserInformation struct {
	isUserInfo
}

type Users interface {
	Add(ctx context.Context, user UserInformation, info ExtendedUserInformation) error

	Get(ctx context.Context, id uint64) (UserInformation, ExtendedUserInformation, error)
	GetLocalByEmail(ctx context.Context, email string) (UserInformation, LocalUserInformation, error)
}
