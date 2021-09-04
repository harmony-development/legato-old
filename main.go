package main

import (
	"os"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/harmonydb"
	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server"
)

func main() {
	l := logger.New(os.Stdin)

	cfg, err := config.New(l, "configuration").ParseConfig()
	if err != nil {
		l.WithError(err).Fatal("Failed to read config")
	}

	db, err := harmonydb.New(l, cfg)
	if err != nil {
		l.WithError(err).Fatal("Failed to connect to database")
	}

	s, err := server.New(l, cfg, db)
	if err != nil {
		l.WithError(err).Fatal("Failed to setup server")
	}
	s.Start()
}
