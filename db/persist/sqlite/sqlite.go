// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package sqlite

import (
	"context"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/persist"
	"github.com/harmony-development/legato/errwrap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

type backend struct{}

func Backend() persist.Backend {
	return backend{}
}

func (b backend) Name() string {
	return "sqlite"
}

// New creates a new persistent backend using sqlite.
func (b backend) New(ctx context.Context, l log.Interface, cfg *config.Config) (persist.Database, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.SQLite.File), &gorm.Config{})
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to open sqlite database")
	}

	err = db.AutoMigrate(
		&user{},
		&session{},
		&foreignuser{},
		&localuser{},
	)
	if err != nil {
		return nil, errwrap.Wrap(err, "database migration failed for sqlite")
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
