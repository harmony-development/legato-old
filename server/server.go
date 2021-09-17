// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package server

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harmony-development/legato/api"
	authv1impl "github.com/harmony-development/legato/api/authv1"
	"github.com/harmony-development/legato/build"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/ephemeral/bigcache"
	"github.com/harmony-development/legato/db/ephemeral/redis"
	"github.com/harmony-development/legato/db/persist"
	"github.com/harmony-development/legato/db/persist/postgres"
	"github.com/harmony-development/legato/db/persist/sqlite"
	"github.com/harmony-development/legato/errwrap"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/key"
	"github.com/harmony-development/legato/logger"
	"github.com/sony/sonyflake"
)

const startupMessage = `Version %s
   __                  __      
  / /___  ___ _ ___ _ / /_ ___ 
 / // -_)/ _ ` + "`" + `// _ ` + "`" + `// __// _ \
/_/ \__/ \_, / \_,_/ \__/ \___/
        /___/ Commit %s
`

// nolint
var persistFactory = persist.NewFactory(
	postgres.Backend(),
	sqlite.Backend(),
)

// nolint
var ephemeralFactory = ephemeral.NewFactory(
	bigcache.Backend(),
	redis.Backend(),
)

// Server is an instance of Legato.
type Server struct {
	*fiber.App
	cfg       *config.Config
	l         log.Interface
	sonyflake *sonyflake.Sonyflake
}

// New creates a new server.
func New(l log.Interface) (*Server, error) {
	cfg, err := newConfig(l, "configuration")
	if err != nil {
		return nil, err
	}

	persist, err := persistFactory.New(context.TODO(), cfg.Database.Backend, l, cfg)
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to create persist backend")
	}

	ephemeral, err := ephemeralFactory.New(context.TODO(), cfg.Epheremal.Backend, l, cfg)
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to create ephemeral backend")
	}

	keyManager, err := tryMakeKeyManager(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.UnixMilli(cfg.IDStart),
	})
	s := newFiber(l, cfg)
	api.RegisterHandlers(s, l, cfg,
		authv1.NewAuthServiceHandler(
			authv1impl.New(keyManager, ephemeral, persist, sonyflake, cfg),
		),
	)

	return &Server{s, cfg, l, sonyflake}, nil
}

// Listen begins listening to the configured port.
func (s *Server) Listen() error {
	s.l.Info(formatStartup(s.cfg.Address, s.cfg.Port))

	return fmt.Errorf("error occurred while listening %w", s.App.Listen(s.cfg.Address+":"+strconv.Itoa(s.cfg.Port)))
}

func newConfig(l log.Interface, name string) (*config.Config, error) {
	l.Info("Reading config...")

	configReader := config.New(name)

	cfg, err := configReader.ParseConfig()
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to read config")
	}

	if err := configReader.WatchConfig(func(ev fsnotify.Event) {
		l.Info("Config change detected, reloading...")
		newConfig, err := configReader.ParseConfig()
		if err != nil {
			l.WithError(err).Error("Failed to reload config")

			return
		}
		*cfg = *newConfig
	}, func(error) {}); err != nil {
		return nil, errwrap.Wrap(err, "failed to watch config")
	}

	return cfg, nil
}

func newFiber(l log.Interface, cfg *config.Config) *fiber.App {
	s := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          api.FiberErrorHandler(l, cfg),
	})

	// TODO: move to config variable or smth but make this line shorter
	// nolint
	logFormat := "[${time}] ${status} |${green}${method}${white}|  ${path}  ↑${bytesSent} bytes  ↓${bytesReceived} bytes ${reqHeader:Authorization}\n"

	s.Use(fiberLogger.New(fiberLogger.Config{
		Format: logFormat,
	}))

	return s
}

func formatStartup(address string, port int) string {
	listenText := color.HiMagentaString(fmt.Sprintf("Listening on %s:%d", address, port))
	display := logger.Indent(
		log.InfoLevel,
		startupMessage,
		&listenText,
	)
	versionString := color.GreenString(build.Version)
	gitString := color.GreenString(
		fmt.Sprintf("%.7s", build.GitCommit),
	)

	return fmt.Sprintf(display, versionString, gitString)
}

func tryMakeKeyManager(privKeyPath string, pubKeyPath string) (key.Manager, error) {
	keyManager, err := key.NewEd25519KeyManagerFromFile(privKeyPath, pubKeyPath)
	if err == nil {
		return keyManager, nil
	}

	if !os.IsNotExist(err) {
		return nil, errwrap.Wrap(err, "unknown error creating key manager")
	}

	if err := key.WriteEd25519KeysToFile(privKeyPath, pubKeyPath); err != nil {
		return nil, fmt.Errorf("failed to save keys: %w", err)
	}

	keyManager, err = key.NewEd25519KeyManagerFromFile(privKeyPath, pubKeyPath)

	return keyManager, errwrap.Wrap(err, "failed to make key manager")
}
