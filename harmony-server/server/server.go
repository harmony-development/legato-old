package server

import (
	"github.com/sirupsen/logrus"
	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"harmony-server/server/state/guild"
	"harmony-server/server/storage"
	"sync"
)

// Instance is an instance of the harmony server
type Instance struct {
	Server         *http.Server
	State          *state.State
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         *logger.Logger
	DB             *db.DB
}

// Start begins the instance server
func (inst Instance) Start() {
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Unable to load config", err)
	}
	inst.Config = cfg
	if err := ConnectSentry(cfg); err != nil {
		logrus.Fatal("Error connecting to sentry", err)
	}
	inst.Logger = logger.New(cfg)
	inst.DB, err = db.New(inst.Config)
	if err != nil {
		inst.Logger.Fatal("Error migrating database : ", err)
	}
	inst.AuthManager, err = auth.New()
	if err != nil {
		inst.Logger.Fatal(err)
	}
	inst.StorageManager = &storage.Manager{
		ImageDeleteQueue:        make(chan string, 512),
		GuildPictureDeleteQueue: make(chan string, 512),
		ImagePath:               inst.Config.Server.ImagePath,
		GuildPicturePath:        inst.Config.Server.GuildPicturePath,
	}
	go inst.StorageManager.DeleteRoutine()
	inst.State = &state.State{
		Guilds:     make(map[string]*guild.Guild),
		GuildsLock: &sync.RWMutex{},
	}
	inst.Server = http.New(&http.Dependencies{
		DB:             inst.DB,
		AuthManager:    inst.AuthManager,
		StorageManager: inst.StorageManager,
		State:          inst.State,
		Config:         inst.Config,
	})
	logrus.Fatal(inst.Server.Start(inst.Config.Server.Port))
}
