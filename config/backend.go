// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package config

import "fmt"

// TODO: get this to use the actual registered backends list.
var (
	persistBackends   = StringSet{}
	ephemeralBackends = StringSet{}
)

func init() {
	persistBackends.Add(
		"postgres",
		"sqlite",
	)
	ephemeralBackends.Add(
		"bigcache",
		"redis",
	)
}

type PersistBackend string

func (e *PersistBackend) UnmarshalText(text []byte) error {
	ok := persistBackends.Has(string(text))
	if !ok {
		return fmt.Errorf("persist backend must be one of: %v", persistBackends.Values())
	}
	*e = PersistBackend(text)
	return nil
}

type EpheremalBackend string

func (e *EpheremalBackend) UnmarshalText(text []byte) error {
	ok := ephemeralBackends.Has(string(text))
	if !ok {
		return fmt.Errorf("ephemeral backend must be one of: %v", ephemeralBackends.Values())
	}
	*e = EpheremalBackend(text)
	return nil
}
