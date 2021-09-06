// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package sqlite

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/persist"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

type factory struct{}

var Factory persist.Factory = factory{}

func init() {
	persist.RegisterBackend("sqlite", Factory)
}

func (factory) NewDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (persist.Database, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.SQLite.File), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %+w", err)
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, fmt.Errorf("database migration failed: %+w", err)
	}

	return &database{
		db: db,
	}, nil
}

func (d *database) Sessions() persist.Sessions {
	return &sessions{d}
}

func (d *database) Users() persist.Users {
	return &users{d}
}
