package cassandra

import (
	"context"
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/ztrue/tracerr"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"

	_ "embed"

	_ "github.com/lib/pq"
)

type cassandraBackend struct {
}

func (b cassandraBackend) New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (*database, error) {
	return New(cfg, logger, idgen)
}
func (p cassandraBackend) Name() string {
	return "cassandra"
}

// func init() {
// 	db.RegisterBackend(cassandraBackend{})
// }

// DB is a wrapper for the DB
type database struct {
	gocqlx.Session
	Logger    logger.ILogger
	Config    *config.Config
	Sonyflake *sonyflake.Sonyflake
	userTable *table.Table
}

var ctx = context.Background()

//go:embed "setup.sql"
var setup string

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (*database, error) {
	if err := Migrate(); err != nil {
		return nil, err
	}
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "harmony"
	wrapped, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	db := &database{Session: wrapped, Config: cfg, Logger: logger, Sonyflake: idgen}
	db.setupUserTable()
	return db, nil
}

func Migrate() error {
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.Keyspace = "system"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return tracerr.Wrap(err)
	}

	for _, line := range strings.Split(setup, ";") {
		if len(line) == 0 {
			continue
		}
		if err := session.ExecStmt(line); err != nil {
			return tracerr.Wrap(err)
		}
	}
	return nil
}
