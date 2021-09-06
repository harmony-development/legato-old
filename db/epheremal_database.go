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

type authDB interface {
	GetCurrentStep(ctx context.Context, authID string) (string, error)
	SetStep(ctx context.Context, authID string, step string) error
	DeleteSession(ctx context.Context, authID string) error
}

// EpheremalDatabase handles access to short-lived data and pubsub
type EpheremalDatabase interface {
	authDB
}

type EpheremalDatabaseFactory interface {
	NewEpheremalDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (EpheremalDatabase, error)
}
