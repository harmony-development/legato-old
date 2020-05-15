package db

import (
	"database/sql"
	"fmt"
	"harmony-server/server/config"

	lru "github.com/hashicorp/golang-lru"
	_ "github.com/lib/pq"
)

// HarmonyDB is a wrapper for the SQL HarmonyDB
type HarmonyDB struct {
	*sql.DB
	Queries      *Queries
	Config       *config.Config
	OwnerCache   *lru.Cache
	SessionCache *lru.Cache
}

// New creates a new DB connection
func New(cfg *config.Config) (*HarmonyDB, error) {
	db := &HarmonyDB{}
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
