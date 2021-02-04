package server

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"

	"github.com/harmony-development/legato/server/api"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/attachments/backend"

	"github.com/harmony-development/legato/server/db/types"
	database_attachments_backend "github.com/harmony-development/legato/server/http/attachments/backend/database"
	"github.com/harmony-development/legato/server/http/attachments/backend/flatfile"
	"github.com/harmony-development/legato/server/intercom"
	"github.com/harmony-development/legato/server/logger"
)

// Instance is an instance of the harmony server
type Instance struct {
	API             *api.API
	Sonyflake       *sonyflake.Sonyflake
	Config          *config.Config
	IntercomManager *intercom.Manager
	AuthManager     *auth.Manager
	Logger          logger.ILogger
	DB              types.IHarmonyDB
}

// Start begins the instance server
func (inst Instance) Start() {
	_ = os.Mkdir("./filestore", 0o777)
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
	bk := db.GetBackend(inst.Config.Database.Backend)
	if bk == nil {
		inst.Logger.Fatal(fmt.Errorf("'%s' is not a known backend! Known backends are: %+v", inst.Config.Database.Backend, db.Backends()))
	}
	inst.DB, err = bk.New(inst.Config, inst.Logger, inst.Sonyflake)
	if err != nil {
		inst.Logger.Fatal(err)
	}
	inst.IntercomManager, err = intercom.New(intercom.Dependencies{
		Logger: inst.Logger,
	})
	if err != nil {
		inst.Logger.Fatal(err)
	}
	inst.AuthManager, err = auth.New(&auth.Dependencies{Config: cfg, IntercomManager: inst.IntercomManager})
	if err != nil {
		inst.Logger.Fatal(err)
	}
	var storageBackend backend.AttachmentBackend

	switch inst.Config.Server.StorageBackend {
	case "PureFlatfile":
		storageBackend = &flatfile.Backend{
			Dependencies: flatfile.Dependencies{
				Config: inst.Config,
			},
		}
	case "DatabaseFlatfile":
		storageBackend = &database_attachments_backend.Backend{
			Dependencies: database_attachments_backend.Dependencies{
				Config: inst.Config,
				DB:     inst.DB,
			},
		}
	default:
		inst.Logger.Fatal(errors.New("config backend is not valid; must be 'PureFlatfile' or 'DatabaseFlatfile'"))
	}
	inst.API = api.New(api.Dependencies{
		Logger:         inst.Logger,
		DB:             inst.DB,
		AuthManager:    inst.AuthManager,
		Sonyflake:      inst.Sonyflake,
		Config:         inst.Config,
		Permissions:    permissions.NewManager(inst.DB, inst.Logger),
		StorageBackend: storageBackend,
	})

	errChan := make(chan error)

	go func() {
		inst.Logger.Debug(logger.Startup, "API routes:", inst.API.Routes())
		errChan <- inst.API.Start(fmt.Sprintf("%s:%d", inst.Config.Server.Host, inst.Config.Server.Port))
	}()

	terminateChan := make(chan os.Signal, 1)
	signal.Notify(terminateChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logrus.Info("Legato started")
		if err := <-errChan; err != nil {
			logrus.Error(err)
		}
		terminateChan <- os.Interrupt
	}()

	<-terminateChan

	logrus.Info("Legato ended")
}
