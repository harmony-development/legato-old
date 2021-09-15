// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package bigcache

import (
	"context"

	"github.com/apex/log"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/ephemeral/kv"
	"github.com/philippgille/gokv/bigcache"
	"github.com/philippgille/gokv/encoding"
)

type backend struct{}

func Backend() ephemeral.Backend {
	return backend{}
}

func (backend) Name() string {
	return "bigcache"
}

// New creates a new ephemeral backend using bigcache.
func (backend) New(ctx context.Context, l log.Interface, cfg *config.Config) (ephemeral.Database, error) {
	cache, err := bigcache.NewStore(bigcache.Options{
		Codec: encoding.Gob,
	})
	if err != nil {
		return nil, err
	}

	return kv.NewKVBackend(cache), nil
}
