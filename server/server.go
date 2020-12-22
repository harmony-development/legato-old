package server

import (
	"errors"
	"fmt"
	"net"
	stdlibHTTP "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"

	"github.com/harmony-development/legato/server/api"
	"github.com/harmony-development/legato/server/api/core/v1/permissions"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	database_attachments_backend "github.com/harmony-development/legato/server/http/attachments/backend/database"
	"github.com/harmony-development/legato/server/http/attachments/backend/flatfile"
	"github.com/harmony-development/legato/server/intercom"
	"github.com/harmony-development/legato/server/logger"
	"github.com/soheilhy/cmux"
)

// Instance is an instance of the harmony server
type Instance struct {
	API             *api.API
	Sonyflake       *sonyflake.Sonyflake
	Config          *config.Config
	IntercomManager *intercom.Manager
	AuthManager     *auth.Manager
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
		inst.Logger.Fatal(errors.New("Config backend is not valid; must be 'PureFlatfile' or 'DatabaseFlatfile'."))
	}
	inst.API = api.New(api.Dependencies{
		Logger:         inst.Logger,
		DB:             inst.DB,
		AuthManager:    inst.AuthManager,
		Sonyflake:      inst.Sonyflake,
		Config:         inst.Config,
		Permissions:    permissions.NewManager(inst.DB),
		StorageBackend: storageBackend,
	})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", inst.Config.Server.Host, inst.Config.Server.Port))
	if err != nil {
		inst.Logger.Fatal(err)
	}

	muxer := cmux.New(listener)
	http2 := muxer.Match(cmux.HTTP2())
	http1 := muxer.Match(cmux.HTTP1())

	errChan := make(chan error)
	go func() {
		httpServer := http.New(http.Dependencies{
			DB:             inst.DB,
			Logger:         inst.Logger,
			Config:         inst.Config,
			StorageBackend: storageBackend,
		})
		err := (&stdlibHTTP.Server{
			Handler: stdlibHTTP.HandlerFunc(func(resp stdlibHTTP.ResponseWriter, req *stdlibHTTP.Request) {
				if strings.Contains(req.Header.Get("Access-Control-Request-Headers"), "x-grpc-web") || req.Header.Get("x-grpc-web") == "1" || req.Header.Get("Sec-Websocket-Protocol") == "grpc-websockets" {
					inst.API.GrpcWebServer.ServeHTTP(resp, req)
				} else if strings.HasPrefix(req.Header.Get("User-Agent"), "Prometheus") {
					inst.API.PrometheusServer.Handler.ServeHTTP(resp, req)
				} else {
					httpServer.ServeHTTP(resp, req)
				}
			}),
		}).Serve(http1)
		inst.Logger.CheckException(err)
		errChan <- err
	}()
	go func() {
		errChan <- inst.API.GrpcServer.Serve(http2)
	}()
	go func() {
		errChan <- muxer.Serve()
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
