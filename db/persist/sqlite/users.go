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
	Username string `gorm:"unique"`

	Local   *localuser
	Foreign *foreignuser
}

type localuser struct {
	Email    string `gorm:"unique"`
	Password []byte

	UserID uint64
	ID     int `gorm:"primarykey"`
}

type foreignuser struct {
	UserID uint64
	ID     int `gorm:"primarykey"`
}

type users struct {
	*database
}

func (db *users) Add(ctx context.Context, pers persist.UserInformation, ext persist.ExtendedUserInformation) error {
	switch data := ext.(type) {
	case persist.ForeignUserInformation:
		return db.db.Create(&user{
			ID:       pers.ID,
			Username: pers.Username,
			Foreign:  &foreignuser{},
		}).Error
	case persist.LocalUserInformation:
		return db.db.Create(&user{
			ID:       pers.ID,
			Username: pers.Username,
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
	ui.Username = user.Username

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
	var luser localuser
	var user user

	err := db.db.First(&luser, "email = ?", email).Error
	if err != nil {
		return persist.UserInformation{}, persist.LocalUserInformation{}, err
	}

	err = db.db.First(&user, "id = ?", luser.UserID).Error
	if err != nil {
		return persist.UserInformation{}, persist.LocalUserInformation{}, err
	}

	user.Local = &luser

	return persist.UserInformation{
			ID: user.ID,
		}, persist.LocalUserInformation{
			Email:    user.Local.Email,
			Password: user.Local.Password,
		}, nil
}
