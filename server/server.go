package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/log"
	"github.com/harmony-development/legato/server/api"
)

// Instance is a harmony server instance.
type Instance struct {
	c   *config.Config
	l   log.Logger
	api *api.API
}

func New(c *config.Config, l log.Logger) *Instance {
	api := api.New()

	return &Instance{c, l, api}
}

func (i *Instance) Serve(errChan chan error) {
	i.l.Log("Legato started")
	addr := fmt.Sprintf("%s:%d", i.c.Server.Host, i.c.Server.Port)
	errChan <- i.api.Start(addr)
}

// Run starts the harmony server taking a channel for errors and a channel for termination signals.
func (i *Instance) Run(errChan chan error, terminateChan chan os.Signal) error {
	go i.Serve(errChan)

	signal.Notify(terminateChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		i.l.Log("Fatal error occurred, exiting")

		return err
	case sig := <-terminateChan:
		i.l.Log("Received " + sig.String() + ", exiting")

		return nil
	}
}

// Start initializes the harmony server with default channels.
func (i *Instance) Start() error {
	return i.Run(
		make(chan error),
		make(chan os.Signal, 1),
	)
}
