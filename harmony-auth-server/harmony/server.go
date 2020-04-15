package harmony

import (
	"github.com/sirupsen/logrus"
	"harmony-auth-server/consts"
	"harmony-auth-server/harmony/auth"
	"harmony-auth-server/harmony/config"
	"harmony-auth-server/harmony/db"
	"harmony-auth-server/harmony/http"
	"harmony-auth-server/harmony/storage"
)

// Instance contains the server instance's variables.
type Instance struct {
	Server         *http.Server
	DB             *db.DB
	StorageManager *storage.Manager
	AuthHandler    *auth.Manager
	Consts         *consts.Constants
	Config         *config.Config
}

// Start begins the authentication server
func (inst Instance) Start() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Unable to load config", err)
	}
	inst.Config = cfg
	inst.Consts = consts.MakeConstants()
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
	inst.Server = http.New(inst.DB, inst.AuthHandler, inst.StorageManager, inst.Config, inst.Consts)
	logrus.Fatal(inst.Server.Start(inst.Config.Server.Port))
}

// Stop ends the authentication server gracefully
func (inst Instance) Stop() {
	inst.Server.Stop()
	if err := inst.DB.Close(); err != nil {
		logrus.Error("Error closing database", err)
	}
	if err := inst.DB.Close(); err != nil {
		logrus.Error("Error closing redis connection", err)
	}
}
