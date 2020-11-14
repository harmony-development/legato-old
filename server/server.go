package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
	"golang.org/x/sync/errgroup"

	"github.com/harmony-development/legato/server/api"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http"
	"github.com/harmony-development/legato/server/intercom"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/storage"
	"github.com/soheilhy/cmux"
)

// Instance is an instance of the harmony server
type Instance struct {
	API             *api.API
	Sonyflake       *sonyflake.Sonyflake
	Config          *config.Config
	IntercomManager *intercom.Manager
	AuthManager     *auth.Manager
	StorageManager  *storage.Manager
	Logger          logger.ILogger
	DB              db.IHarmonyDB
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
	inst.DB, err = db.New(inst.Config, inst.Logger, inst.Sonyflake)
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
	inst.API = api.New(api.Dependencies{
		Logger:      inst.Logger,
		DB:          inst.DB,
		AuthManager: inst.AuthManager,
		Sonyflake:   inst.Sonyflake,
		Config:      inst.Config,
	})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", inst.Config.Server.Port))
	if err != nil {
		inst.Logger.Fatal(err)
	}

	multiplexer := cmux.New(listener)

	grpcListener := multiplexer.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcWebListener := multiplexer.Match(
		cmux.HTTP1HeaderField("content-type", "application/grpc-web"),
		cmux.HTTP1HeaderField("Sec-Websocket-Protocol", "grpc-websockets"),
	)
	prometheusListener := multiplexer.Match(cmux.HTTP1HeaderFieldPrefix("User-Agent", "Prometheus"))
	httpListener := multiplexer.Match(cmux.HTTP1Fast())

	grp := new(errgroup.Group)
	grp.Go(func() error {
		httpServer := http.New(http.Dependencies{
			DB:     inst.DB,
			Logger: inst.Logger,
			Config: inst.Config,
		})
		err := httpServer.Server.Serve(httpListener)
		inst.Logger.CheckException(err)
		return err
	})
	grp.Go(func() error {
		return inst.API.Start(grpcListener, grpcWebListener, prometheusListener)
	})
	grp.Go(func() error {
		return multiplexer.Serve()
	})

	logrus.Info("Legato started")
	grp.Wait()
}
