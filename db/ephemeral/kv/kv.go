// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package kv

import (
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/philippgille/gokv"
)

type database struct {
	store gokv.Store
}

func NewKVBackend(store gokv.Store) ephemeral.Database {
	return &database{
		store: store,
	}
}
