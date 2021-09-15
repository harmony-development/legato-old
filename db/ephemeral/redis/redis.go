// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package redis

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/ephemeral/kv"
	"github.com/philippgille/gokv/redis"
)

type backend struct{}

func Backend() ephemeral.Backend {
	return backend{}
}

func (backend) Name() string {
	return "redis"
}

// New creates a new ephemeral backend using redis.
func (backend) New(ctx context.Context, l log.Interface, cfg *config.Config) (ephemeral.Database, error) {
	rdb, err := redis.NewClient(redis.Options{
		Address:  cfg.Epheremal.Redis.Hosts[0],
		Password: cfg.Epheremal.Redis.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis %w", err)
	}

	return kv.NewKVBackend(rdb), nil
}
