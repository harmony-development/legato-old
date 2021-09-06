// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package postgres

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/persist"
	"github.com/harmony-development/legato/db/persist/sql/gen"
	"github.com/jackc/pgx/v4/pgxpool"
)

type database struct {
	queries *gen.Queries

	s *sessions
	u *users
}

func (d *database) Sessions() persist.Sessions {
	return d.s
}

func (d *database) Users() persist.Users {
	return d.u
}

type factory struct{}

var Factory persist.Factory = factory{}

func init() {
	persist.RegisterBackend("postgres", Factory)
}

func (factory) NewDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (persist.Database, error) {
	username, password, host, port, db :=
		cfg.Database.Postgres.Username,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.DB

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		username,
		password,
		host,
		port,
		db,
	)

	println(connString)

	conn, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	q := gen.New(conn)

	return &database{
		queries: q,
	}, nil
}
