package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/ztrue/tracerr"
)

type DB struct {
	session *gocql.Session
	conf    *config.Config
}

var _ db.IHarmonyDB = &DB{}

func New(conf *config.Config) (*DB, error) {
	db := &DB{}

	cfg := gocql.NewCluster()
	cfg.PoolConfig.HostSelectionPolicy =
		gocql.TokenAwareHostPolicy(
			gocql.RoundRobinHostPolicy(), // Load balance by sequentially trying hosts for queries
		)

	cfg.Compressor = &gocql.SnappyCompressor{} // Optimized for performance but not that much compression
	cfg.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{
		NumRetries: conf.Server.Cassandra.NumRetries,
	}

	cfg.Consistency = gocql.LocalQuorum

	session, err := cfg.CreateSession()
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	db.session = session

	return db, nil
}
