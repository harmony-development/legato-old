package postgres

import (
	"context"
	"fmt"

	"github.com/ztrue/tracerr"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/backends/ent_shared"
	"github.com/harmony-development/legato/server/db/ent/entgen"
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

// DB is a wrapper for the ent DB
type database struct {
	*ent_shared.DB
}

var ctx = context.Background()

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (*database, error) {
	client, err := entgen.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Host,
		cfg.Database.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.Database.SSL],
	))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	db, err := ent_shared.New(client, logger)
	if err != nil {
		return nil, err
	}
	return &database{
		DB: db,
	}, nil
}
