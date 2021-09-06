// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package sqlite

import (
	"context"

	"github.com/harmony-development/legato/db/persist"
)

type user struct {
	ID       uint64 `gorm:"primarykey"`
	Email    string `gorm:"unique"`
	Password []byte
}

type users struct {
	*database
}

func (db *users) Add(ctx context.Context, persist persist.UserInformation) error {
	return db.db.Create(&user{
		ID:       persist.ID,
		Email:    persist.Email,
		Password: persist.Password,
	}).Error
}

func (db *users) Get(ctx context.Context, id uint64) (persist.UserInformation, error) {
	var user user
	err := db.db.First(&user, "id = ?", id).Error
	if err != nil {
		return persist.UserInformation{}, err
	}
	return persist.UserInformation{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (db *users) GetByEmail(ctx context.Context, email string) (persist.UserInformation, error) {
	var user user
	err := db.db.First(&user, "email = ?", email).Error
	if err != nil {
		return persist.UserInformation{}, err
	}
	return persist.UserInformation{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
