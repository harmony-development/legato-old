package db

import (
	"context"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
)

// Inittable is an interface for things that need an initialisation
// function to be called, such as to start a cleanup routine
type Inittable interface {
	Init()
}

type InitNothing struct{}

func (InitNothing) Init() {}

type sessionDB interface {
	GetSession(ctx context.Context, session string) (int64, error)
	SetSession(ctx context.Context, session string, userID int64) error
}

// Database handles access to long-lived data
type Database interface {
	Inittable

	sessionDB
}

type authDB interface {
	GetCurrentStep(ctx context.Context, authID string) (string, error)
	SetStep(ctx context.Context, authID string, step string) error
	DeleteSession(ctx context.Context, authID string) error
}

// EpheremalDatabase handles access to short-lived data and pubsub
type EpheremalDatabase interface {
	Inittable

	authDB
}

type DatabaseFactory interface {
	NewDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (Database, error)
}

type EpheremalDatabaseFactory interface {
	NewEpheremalDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (EpheremalDatabase, error)
}
