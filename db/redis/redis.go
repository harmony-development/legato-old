// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package redis

import (
	"context"

	"github.com/apex/log"
	"github.com/go-redis/redis/v8"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
)

type database struct {
	rdb *redis.ClusterClient
}

type factory struct {
}

var Factory db.EpheremalDatabaseFactory = factory{}

func (factory) NewEpheremalDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (db.EpheremalDatabase, error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Epheremal.Redis.Hosts,
		Password: cfg.Epheremal.Redis.Password,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &database{
		rdb: rdb,
	}, nil
}
