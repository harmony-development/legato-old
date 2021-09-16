// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package ephemeral

import (
	"context"
	"errors"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/errwrap"
)

type authDB interface {
	GetCurrentStep(ctx context.Context, authID string) (string, error)
	SetStep(ctx context.Context, authID string, step string) error
	DeleteAuthID(ctx context.Context, authID string) error
}

// Database handles access to short-lived data and pubsub.
type Database interface {
	authDB
}

var (
	ErrDatabaseNotFound = errors.New("database not found")
	ErrStepNotFound     = errors.New("step not found")
)

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
func (b Factory) New(ctx context.Context, name string, l log.Interface, cfg *config.Config) (Database, error) {
	backend, ok := b[name]
	if !ok {
		return nil, errwrap.Wrap(ErrDatabaseNotFound, name)
	}

	db, err := backend.New(ctx, l, cfg)

	return db, errwrap.Wrap(err, "failed to create backend")
}
