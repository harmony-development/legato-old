// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package redis

import (
	"context"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/ephemeral/kv"
	"github.com/philippgille/gokv/redis"
)

type factory struct{}

var Factory ephemeral.Factory = factory{}

func init() {
	ephemeral.RegisterBackend("redis", Factory)
}

func (factory) NewEpheremalDatabase(ctx context.Context, l log.Interface, cfg *config.Config) (ephemeral.Database, error) {
	rdb, err := redis.NewClient(redis.Options{
		Address:  cfg.Epheremal.Redis.Hosts[0],
		Password: cfg.Epheremal.Redis.Password,
	})
	if err != nil {
		return nil, err
	}

	return kv.NewKVBackend(rdb), nil
}
