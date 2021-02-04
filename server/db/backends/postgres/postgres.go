package postgres

import (
	"context"
	"database/sql"
	"fmt"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ztrue/tracerr"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"

	_ "github.com/lib/pq"
)

type postgresBackend struct {
}

func (p postgresBackend) New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	return New(cfg, logger, idgen)
}
func (p postgresBackend) Name() string {
	return "postgres"
}

func init() {
	db.RegisterBackend(postgresBackend{})
}

// DB is a wrapper for the SQL DB
type database struct {
	*sql.DB
	queries      *queries.Queries
	Logger       logger.ILogger
	Config       *config.Config
	OwnerCache   *lru.Cache
	SessionCache *lru.Cache
	Sonyflake    *sonyflake.Sonyflake
}

var ctx = context.Background()

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (*database, error) {
	db := &database{}
	db.Config = cfg
	db.Logger = logger
	db.Sonyflake = idgen
	var err error
	if db.DB, err = sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Host,
		cfg.Database.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.Database.SSL],
	)); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if err = db.Ping(); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if err = db.Migrate(); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.queries, err = queries.Prepare(context.Background(), db); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.OwnerCache, err = lru.New(cfg.Server.Policies.MaximumCacheSizes.Owner); err != nil {
		return nil, tracerr.Wrap(err)
	}
	if db.SessionCache, err = lru.New(cfg.Server.Policies.MaximumCacheSizes.Sessions); err != nil {
		return nil, tracerr.Wrap(err)
	}
	go db.SessionExpireRoutine()
	return db, nil
}
