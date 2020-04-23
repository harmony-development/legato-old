package server

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"harmony-auth-server/server/auth"
	"harmony-auth-server/server/config"
	"harmony-auth-server/server/db"
	"harmony-auth-server/server/http"
	"harmony-auth-server/server/storage"
	"time"
)

// Instance is an instance of the harmony auth server
type Instance struct {
	Server         *http.Server
	DB             *db.DB
	StorageManager *storage.Manager
	AuthHandler    *auth.Manager
	Config         *config.Config
}

// Start begins the authentication server
func (inst Instance) Start() {
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
		logrus.Fatal("Error connecting to database", err)
	}
	if err := inst.DB.Migrate(); err != nil {
		logrus.Fatal("Error persisting database schema", err)
	}
	inst.AuthHandler = auth.New(cfg)
	inst.StorageManager = storage.New()
	go inst.StorageManager.DeleteRoutine(cfg.Server.AvatarPath)
	inst.Server = http.New(inst.DB, inst.AuthHandler, inst.StorageManager, inst.Config)
	logrus.Fatal(inst.Server.Start(inst.Config.Server.Port))
}

// Stop ends the authentication server gracefully
func (inst Instance) Stop() {
	inst.Server.Stop()
	if inst.Config.Sentry.Enabled {
		sentry.Flush(2 * time.Second)
	}

	if err := inst.DB.Close(); err != nil {
		logrus.Error("Error closing database", err)
	}

	if err := inst.DB.Close(); err != nil {
		logrus.Error("Error closing redis connection", err)
	}
}
