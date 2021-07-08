package main

import (
	"flag"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/log/impl"
	"github.com/harmony-development/legato/server"
	"github.com/ztrue/tracerr"
)

func flags() (configPath string) {
	flag.StringVar(&configPath, "config", "config.yml", "the path to load the config from")
	flag.Parse()
	return
}

func startup() error {
	configPath := flags()

	cfg, err := config.Load(configPath)
	if err != nil {
		return err
	}

	logger := impl.New(cfg)

	srv := server.New(cfg, logger)
	return srv.Start()
}

func main() {
	if err := startup(); err != nil {
		tracerr.Print(err)
	}
}
