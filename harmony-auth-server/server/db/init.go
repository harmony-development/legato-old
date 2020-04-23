package db

import (
	"database/sql"
	"fmt"
	"harmony-auth-server/server/config"
	// postgres support for gorm
	_ "github.com/lib/pq"
)

// DB is an wrapper for the SQL DB
type DB struct {
	*sql.DB
}

// New initializes the connection to the database
func New(cfg *config.Config) (*DB, error) {
	db := &DB{}
	var err error
	if db.DB, err = sql.Open("postgres", fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v",
		cfg.DB.User,
		cfg.DB.Password,
		"harmonyauth",
		cfg.DB.Host,
		cfg.DB.Port,
		map[bool]string{true: "enable", false: "disable"}[cfg.DB.SSL],
	)); err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
