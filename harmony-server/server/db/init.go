package db

import (
	"database/sql"
	"fmt"
	"harmony-server/server/config"

	lru "github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
)

// DB is a wrapper for the SQL DB
type DB struct {
	*sql.DB
	Config       *config.Config
	OwnerCache   *lru.Cache
	SessionCache *lru.Cache
}

// New creates a new DB connection
func New(cfg *config.Config) (*DB, error) {
	db := &DB{}
	var err error
	if db.DB, err = sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.DB.User,
		cfg.DB.Password,
		"harmony",
		cfg.DB.Host,
		cfg.DB.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.DB.SSL],
	)); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err = db.Migrate(); err != nil {
		return nil, err
	}
	if err = db.AddSampleData(); err != nil {
		return nil, err
	}
	if db.OwnerCache, err = lru.New(cfg.Server.OwnerCacheMax); err != nil {
		return nil, err
	}
	if db.SessionCache, err = lru.New(cfg.Server.SessionCacheMax); err != nil {
		return nil, err
	}
	return db, nil
}
