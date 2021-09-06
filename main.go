// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/apex/log"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harmony-development/legato/api"
	authv1impl "github.com/harmony-development/legato/api/authv1"
	"github.com/harmony-development/legato/build"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/persist"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/key"
	"github.com/harmony-development/legato/logger"

	// DATABASE BACKENDS
	_ "github.com/harmony-development/legato/db/ephemeral/bigcache"
	_ "github.com/harmony-development/legato/db/ephemeral/redis"
	_ "github.com/harmony-development/legato/db/persist/postgres"
)

var startupMessage = `Version %s
   __                  __      
  / /___  ___ _ ___ _ / /_ ___ 
 / // -_)/ _ ` + "`" + `// _ ` + "`" + `// __// _ \
/_/ \__/ \_, / \_,_/ \__/ \___/
        /___/ Commit %s
`

func main() {
	l := logger.New(os.Stdin)

	cfg, err := config.New(l, "configuration").ParseConfig()
	if err != nil {
		l.WithError(err).Fatal("Failed to read config")
	}

	keyManager, err := tryMakeKeyManager(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	if err != nil {
		l.WithError(err).Fatal("Failed to initialize key manager")
	}
	persistFactory, err := persist.GetBackend(string(cfg.Database.Backend))
	if err != nil {
		l.WithError(err).Fatal("Failed to initialize persistent database")
	}

	persist, err := persistFactory.NewDatabase(context.TODO(), l, cfg)
	if err != nil {
		l.WithError(err).Fatal("Failed to connect to database")
	}

	_ = persist

	ephemeralFactory, err := ephemeral.GetBackend(string(cfg.Epheremal.Backend))
	if err != nil {
		l.WithError(err).Fatal("Failed to initialize ephemeral database")
	}

	ephemeral, err := ephemeralFactory.NewEpheremalDatabase(context.TODO(), l, cfg)
	if err != nil {
		l.WithError(err).Fatal("Failed to connect to epheremal database")
	}

	s := newServer(l, cfg)
	registerServices := api.Setup(l, s)

	registerServices(
		authv1.NewAuthServiceHandler(
			authv1impl.New(keyManager, ephemeral),
		),
	)

	l.Info(formatStartup(cfg.Address, cfg.Port))
	s.Listen(cfg.Address + ":" + strconv.Itoa(cfg.Port))
}

func newServer(l log.Interface, cfg *config.Config) *fiber.App {
	s := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          api.FiberErrorHandler(l, cfg),
	})

	s.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} |${green}${method}${white}|  ${path}  ↑${bytesSent} bytes  ↓${bytesReceived} bytes ${reqHeader:Authorization}\n",
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

func tryMakeKeyManager(privKeyPath string, pubKeyPath string) (key.KeyManager, error) {
	keyManager, err := key.NewEd25519KeyManagerFromFile(privKeyPath, pubKeyPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := key.WriteEd25519KeysToFile(privKeyPath, pubKeyPath); err != nil {
				return nil, err
			}
			keyManager, err = key.NewEd25519KeyManagerFromFile(privKeyPath, pubKeyPath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return keyManager, nil
}
