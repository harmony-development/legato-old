package api

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/hrpc/server"
)

// FiberRPCHandler converts a RPC handler to a Fiber handler
func FiberRPCHandler(handler server.RawHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp, err := handler(c.Context(), c.Request())
		if err != nil {
			return err
		}
		return c.Send(resp)
	}
}

func RegisterHandlers(l log.Interface, app *fiber.App, all ...server.HRPCServiceHandler) {
	l.Info("Registering services...")
	for _, handler := range all {
		serviceLog := l.WithFields(log.Fields{
			"service": handler.Name(),
		})
		for path, handler := range handler.Routes() {
			serviceLog.Infof("Registered %s", path)
			app.All(path, FiberRPCHandler(handler))
		}
	}
}

// Setup registers the Harmony protocol API
func Setup(l log.Interface, app *fiber.App) func(all ...server.HRPCServiceHandler) {
	return func(all ...server.HRPCServiceHandler) {
		RegisterHandlers(l, app, all...)
	}
}
