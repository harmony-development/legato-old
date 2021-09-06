// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package sqlite

import (
	"context"

	"github.com/harmony-development/legato/db/persist"
)

type user struct {
	ID      uint64 `gorm:"primarykey"`
	Local   *localuser
	Foreign *foreignuser
}

type localuser struct {
	Email    string `gorm:"unique"`
	Password []byte

	ID int `gorm:"primarykey"`
}

type foreignuser struct {
	ID int `gorm:"primarykey"`
}

type users struct {
	*database
}

func (db *users) Add(ctx context.Context, pers persist.UserInformation, ext persist.ExtendedUserInformation) error {
	switch data := ext.(type) {
	case persist.ForeignUserInformation:
		return db.db.Create(&user{
			ID:      pers.ID,
			Foreign: &foreignuser{},
		}).Error
	case persist.LocalUserInformation:
		return db.db.Create(&user{
			ID: pers.ID,
			Local: &localuser{
				Email:    data.Email,
				Password: data.Password,
			},
		}).Error
	default:
		panic("unhandled case")
	}
}

func (db *users) Get(ctx context.Context, id uint64) (ui persist.UserInformation, eui persist.ExtendedUserInformation, err error) {
	var user user

	err = db.db.Preload("Local").Preload("Foreign").First(&user, "id = ?", id).Error
	if err != nil {
		return persist.UserInformation{}, nil, err
	}

	ui.ID = user.ID

	if user.Local != nil {
		eui = persist.LocalUserInformation{
			Email:    user.Local.Email,
			Password: user.Local.Password,
		}
	} else if user.Foreign != nil {
		eui = persist.ForeignUserInformation{}
	} else {
		panic("unhandled / invalid db")
	}

	return
}

func (db *users) GetLocalByEmail(ctx context.Context, email string) (persist.UserInformation, persist.LocalUserInformation, error) {
	var user user

	err := db.db.Preload("Local").First(&user, "email = ?", email).Error
	if err != nil {
		return persist.UserInformation{}, persist.LocalUserInformation{}, err
	}

	return persist.UserInformation{
			ID: user.ID,
		}, persist.LocalUserInformation{
			Email:    user.Local.Email,
			Password: user.Local.Password,
		}, nil
}
