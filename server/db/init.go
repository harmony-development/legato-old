package db

import "github.com/harmony-development/legato/server/db/types"

var backends = map[string]types.IBackend{}

func RegisterBackend(b types.IBackend) {
	backends[b.Name()] = b
}

func GetBackend(s string) types.IBackend {
	return backends[s]
}

func Backends() (r []string) {
	for k := range backends {
		r = append(r, k)
	}
	return
}
