package intercom

import (
	"time"

	"github.com/harmony-development/legato/server/logger"
	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/grpc"
)

type Dependencies struct {
	Logger logger.Logger
}

type IntercomManager struct {
	Dependencies
	ForeignConnections *lru.Cache
}

func New(deps Dependencies) (*IntercomManager, error) {
	manager := &IntercomManager{
		Dependencies: deps,
	}

	cache, err := lru.NewWithEvict(4096, manager.OnConnectionEvict)
	if err != nil {
		return nil, err
	}
	manager.ForeignConnections = cache
	return manager, err
}

func (im IntercomManager) Connect(host string) (*grpc.ClientConn, error) {
	client, err := grpc.Dial(host, grpc.WithTimeout(15*time.Second))
	if err != nil {
		return nil, err
	}
	im.ForeignConnections.Add(host, client)
	return client, nil
}

func (im IntercomManager) OnConnectionEvict(key interface{}, value interface{}) {
	conn := value.(*grpc.ClientConn)
	im.Logger.CheckException(conn.Close())
}
