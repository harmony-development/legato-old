package server

import (
	"github.com/sirupsen/logrus"
	"harmony-server/server/config"
	"harmony-server/server/db"
)

// Instance is an instance of the harmony server
type Instance struct {
	Config *config.Config
	DB *db.DB
}

// Start begins the instance server
func (inst Instance) Start()  {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Unable to load config", err)
	}
	inst.Config = cfg
	if err := ConnectSentry(cfg); err != nil {
		logrus.Fatal("Error connecting to sentry", err)
	}
	inst.DB, err = db.New(inst.Config)
	if err != nil {
		logrus.Fatal("Error initializing DB", err)
	}
}
