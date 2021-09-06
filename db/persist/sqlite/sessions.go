// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package sqlite

import "context"

type session struct {
	ID     string `gorm:"primarykey"`
	UserID uint64 `gorm:"not null"`
	User   *user
}

type sessions struct {
	*database
}

func (db *sessions) Get(ctx context.Context, sessionID string) (uint64, error) {
	var ses session
	err := db.db.First(&ses, "id = ?", sessionID).Error
	if err != nil {
		return 0, err
	}
	return ses.UserID, nil
}

func (db *sessions) Add(ctx context.Context, sessionID string, userID uint64) error {
	return db.db.Create(&session{
		ID:     sessionID,
		UserID: userID,
	}).Error
}
