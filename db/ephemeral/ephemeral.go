// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package ephemeral

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
)

type authDB interface {
	GetCurrentStep(ctx context.Context, authID string) (string, error)
	SetStep(ctx context.Context, authID string, step string) error
	DeleteSession(ctx context.Context, authID string) error
}

// Database handles access to short-lived data and pubsub
type Database interface {
	authDB
}

type Factory interface {
	NewEpheremalDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)
}

var backends = map[string]Factory{}

func RegisterBackend(name string, factory Factory) {
	backends[name] = factory
}

func GetBackend(name string) (Factory, error) {
	factory, ok := backends[name]
	if !ok {
		return nil, fmt.Errorf("ephemeral backend not found: %s", name)
	}
	return factory, nil
}
