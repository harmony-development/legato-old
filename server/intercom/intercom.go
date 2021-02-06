package intercom

import (
	"github.com/harmony-development/legato/server/logger"
	lru "github.com/hashicorp/golang-lru"
)

type Dependencies struct {
	Logger logger.ILogger
}

type Manager struct {
	Dependencies
	ForeignConnections *lru.Cache
}

func New(deps Dependencies) (*Manager, error) {
	manager := &Manager{
		Dependencies: deps,
	}

	cache, err := lru.NewWithEvict(4096, manager.OnConnectionEvict)
	if err != nil {
		return nil, err
	}
	manager.ForeignConnections = cache
	return manager, err
}

func (im Manager) Connect(host string) (interface{}, error) {
	panic("unimplemented")
}

func (im Manager) GetOrConnect(host string) (interface{}, error) {
	panic("unimplemented")
}

func (im Manager) OnConnectionEvict(key, value interface{}) {
	panic("unimplemented")
}
