// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package persist

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/persist/sql/gen"
)

type sessionDB interface {
	GetSession(ctx context.Context, session string) (int64, error)
	AddSession(ctx context.Context, session string, userID int64) error
}

type userDB interface {
	GetUserByEmail(ctx context.Context, email string) (gen.GetUserByEmailRow, error)
}

// Database handles access to long-lived data
type Database interface {
	sessionDB
	userDB
}

type Factory interface {
	NewDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)
}

var backends = map[string]Factory{}

func RegisterBackend(name string, factory Factory) {
	backends[name] = factory
}

func GetBackend(name string) (Factory, error) {
	factory, ok := backends[name]
	if !ok {
		return nil, fmt.Errorf("database backend not found: %s", name)
	}
	return factory, nil
}
