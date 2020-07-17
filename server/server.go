package server

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"

	"github.com/harmony-development/legato/server/api"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/state"
	"github.com/harmony-development/legato/server/state/guild"
	"github.com/harmony-development/legato/server/storage"
)

// Instance is an instance of the harmony server
type Instance struct {
	API            *api.API
	Sonyflake      *sonyflake.Sonyflake
	State          *state.State
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         logger.ILogger
	DB             db.IHarmonyDB
}

// Start begins the instance server
func (inst Instance) Start() {
	_ = os.Mkdir("./filestore", 0777)
	cfg, err := config.Load()
	if err != nil {
		logrus.Fatal("Unable to load config", err)
	}
	inst.Config = cfg
	inst.Sonyflake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Unix(cfg.Server.SnowflakeStart, 0),
	})
	if err := ConnectSentry(cfg); err != nil {
		logrus.Fatal("Error connecting to sentry", err)
	}
	inst.Logger = logger.New(cfg)
	inst.DB, err = db.New(inst.Config, inst.Logger)
	if err != nil {
		inst.Logger.Fatal(err)
	}
	inst.AuthManager, err = auth.New(&auth.Dependencies{Config: cfg})
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
		Guilds:     make(map[uint64]*guild.Guild),
		GuildsLock: &sync.RWMutex{},
	}
	inst.API = api.New(api.Dependencies{
		Logger: inst.Logger,
		DB:     inst.DB,
	})
	logrus.Info("Legato started")
	inst.Logger.CheckException(inst.API.Start(inst.Config.Server.Port))
}
