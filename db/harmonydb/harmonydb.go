// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package harmonydb

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/go-redis/redis/v8"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	"github.com/harmony-development/legato/db/sql/gen"
	"github.com/jackc/pgx/v4/pgxpool"
)

// The default Harmony DB implementation. This uses sqlc.
type HarmonyDB struct {
	db.InitNothing

	*HarmonySessionDB
	*HarmonyAuthDB
}

type HarmonySessionDB struct {
	db.InitNothing

	// embed a context for easy use in queries
	ctx     context.Context
	queries *gen.Queries
}

type HarmonyAuthDB struct {
	db.InitNothing

	// embed a context for easy use in queries
	ctx context.Context
	rdb *redis.ClusterClient
}

func NewAuth(l log.Interface, cfg *config.Config) (*HarmonyAuthDB, error) {
	ctx := context.TODO()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Redis.Hosts,
		Password: cfg.Redis.Password,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &HarmonyAuthDB{
		ctx: ctx,
		rdb: rdb,
	}, nil
}

func NewSession(l log.Interface, cfg *config.Config) (*HarmonySessionDB, error) {
	ctx := context.TODO()

	username, password, host, port, db :=
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		username,
		password,
		host,
		port,
		db,
	)

	conn, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	q := gen.New(conn)

	return &HarmonySessionDB{
		ctx:     ctx,
		queries: q,
	}, nil
}

func New(l log.Interface, cfg *config.Config) (*HarmonyDB, error) {
	auth, err := NewAuth(l, cfg)
	if err != nil {
		return nil, err
	}
	sesh, err := NewSession(l, cfg)
	if err != nil {
		return nil, err
	}

	return &HarmonyDB{
		HarmonySessionDB: sesh,
		HarmonyAuthDB:    auth,
	}, nil
}
