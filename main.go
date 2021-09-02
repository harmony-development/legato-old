package main

import (
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/logger"
	"github.com/harmony-development/legato/server"
)

func main() {
	l := logger.New()
	cfg, err := config.New(l, "configuration").ParseConfig()
	if err != nil {
		l.WithError(err).Fatal("Failed to read config")
	}
	s := server.New(l, cfg)
	s.Start()
}
