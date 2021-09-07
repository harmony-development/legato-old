// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package config

type StringSet map[string]struct{}

func (set StringSet) Has(s string) bool {
	_, ok := set[s]

	return ok
}

func NewStringSet(vals ...string) StringSet {
	set := StringSet{}
	for _, v := range vals {
		set[v] = struct{}{}
	}

	return set
}

func (set StringSet) Values() []string {
	ret := []string{}
	for k := range set {
		ret = append(ret, k)
	}

	return ret
}
