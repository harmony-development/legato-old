package sqlite

import (
	"fmt"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/backends/ent_shared"
	"github.com/harmony-development/legato/server/db/ent/entgen"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"

	// backend
	_ "github.com/mattn/go-sqlite3"
	"github.com/sony/sonyflake"
)

type sqliteBackend struct {
}

func (p sqliteBackend) New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	return New(cfg, logger, idgen)
}
func (p sqliteBackend) Name() string {
	return "sqlite"
}

func init() {
	db.RegisterBackend(sqliteBackend{})
}

type database struct {
	*ent_shared.DB
}

// New creates a new DB connection
func New(cfg *config.Config, logger logger.ILogger, idgen *sonyflake.Sonyflake) (types.IHarmonyDB, error) {
	c, err := entgen.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", cfg.Database.Filename))
	if err != nil {
		return nil, err
	}
	db, err := ent_shared.New(c, logger)
	if err != nil {
		return nil, err
	}
	return &database{
		DB: db,
	}, nil
}
