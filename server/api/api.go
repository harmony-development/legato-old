package api

import (
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
)

// API is a instance of the legato API.
type API struct {
	*echo.Echo
}

func New() *API {
	e := echo.New()

	server.NewHRPCServer(e).SetUnaryPre(
		server.ChainHandlerTransformers(),
	)

	return &API{
		e,
	}
}

func (i *API) Start(addr string) error {
	// Register the Chat service
	return i.Echo.Start(addr)
}

func (i *API) Shutdown() error {
	return i.Echo.Close()
}
