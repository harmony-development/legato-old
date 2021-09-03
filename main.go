package main

import (
	"os"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server"
)

func main() {
	l := logger.New(os.Stdin)
	cfg, err := config.New(l, "configuration").ParseConfig()
	if err != nil {
		l.WithError(err).Fatal("Failed to read config")
	}
	s, err := server.New(l, cfg)
	if err != nil {
		l.WithError(err).Fatal("Failed to setup server")
	}
	s.Start()
}
