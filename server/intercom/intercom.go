package intercom

import (
	"context"
	"time"

	"github.com/harmony-development/legato/server/logger"
	lru "github.com/hashicorp/golang-lru"
	"google.golang.org/grpc"
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

func (im Manager) Connect(host string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := grpc.DialContext(ctx, host)
	if err != nil {
		return nil, err
	}
	im.ForeignConnections.Add(host, client)
	return client, nil
}

func (im Manager) GetOrConnect(host string) (*grpc.ClientConn, error) {
	conn, exists := im.ForeignConnections.Get(host)
	if exists {
		return conn.(*grpc.ClientConn), nil
	} else {
		return im.Connect(host)
	}
}

func (im Manager) OnConnectionEvict(key, value interface{}) {
	conn := value.(*grpc.ClientConn)
	im.Logger.CheckException(conn.Close())
}
