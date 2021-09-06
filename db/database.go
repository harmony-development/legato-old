// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package db

import (
	"context"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
)

type sessionDB interface {
	GetSession(ctx context.Context, session string) (int64, error)
	SetSession(ctx context.Context, session string, userID int64) error
}

// Database handles access to long-lived data
type Database interface {
	sessionDB
}

type DatabaseFactory interface {
	NewDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)
}
