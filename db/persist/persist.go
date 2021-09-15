// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package persist

import (
	"context"
	"errors"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
)

const DatabaseNotFoundError = "database not found"

// Database handles access to long-lived data.
type Database interface {
	Sessions() Sessions
	Users() Users
}

type Backend interface {
	Name() string
	New(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)
}

type Factory map[string]Backend

func NewFactory(backends ...Backend) Factory {
	res := make(map[string]Backend)
	for _, backend := range backends {
		res[backend.Name()] = backend
	}

	return res
}

// New creates a new backend by name,
// or returns an error if there isn't one with that name or it fails to construct.
func (b Factory) New(name string, ctx context.Context, l log.Interface, cfg *config.Config) (Database, error) {
	backend, ok := b[name]
	if !ok {
		return nil, fmt.Errorf("%w: %s", errors.New(DatabaseNotFoundError), name)
	}

	return backend.New(ctx, l, cfg)
}
