package api

import (
	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/hrpc/server"
	authv1impl "github.com/harmony-development/legato/api/authv1"
	chatv1impl "github.com/harmony-development/legato/api/chatv1"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
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
		l.Infof("Registering %s", handler.Name())
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
func Setup(l log.Interface, app *fiber.App) {
	RegisterHandlers(l, app,
		authv1.NewAuthServiceHandler(authv1impl.AuthV1{}),
		chatv1.NewChatServiceHandler(chatv1impl.ChatV1{}),
	)
}
