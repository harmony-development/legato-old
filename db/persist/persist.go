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
)

// Database handles access to long-lived data.
type Database interface {
	Sessions() Sessions
	Users() Users
}

// FactoryFunc constructs a new Database.
type FactoryFunc = func(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)

var backends = map[string]FactoryFunc{}

// RegisterBackend registers a new database backend with a name and a factory function.
func RegisterBackend(name string, factory FactoryFunc) {
	backends[name] = factory
}

// GetBackend gets a factory function by name,
// or returns an error if there isn't one with that name.
func GetBackend(name string) (FactoryFunc, error) {
	factory, ok := backends[name]
	if !ok {
		return nil, fmt.Errorf("database backend not found: %s", name)
	}

	return factory, nil
}
