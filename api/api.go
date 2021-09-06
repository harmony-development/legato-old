// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/hrpc/server"
	"github.com/harmony-development/legato/config"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"google.golang.org/protobuf/proto"
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

func FiberErrorHandler(l log.Interface, cfg *config.Config) fiber.ErrorHandler {
	return func(c *fiber.Ctx, e error) error {
		if cfg.Debug.LogErrors && e != nil {
			l.WithError(e).WithFields(log.Fields{
				"path": c.OriginalURL(),
			}).Error("error in http handler")
		}

		switch v := e.(type) {
		case *Error:
			data, err := proto.Marshal(&harmonytypesv1.Error{
				Identifier:   v.Identifier,
				HumanMessage: v.HumanMessage,
				MoreDetails:  v.MoreDetails,
			})
			if err != nil {
				return err
			}
			return c.Status(http.StatusBadRequest).Send(data)
		default:
			err := &harmonytypesv1.Error{
				Identifier: InternalServerError,
			}
			if cfg.Debug.RespondWithErrors {
				err.HumanMessage = v.Error()
			}
			data, marshalErr := proto.Marshal(err)
			if marshalErr != nil {
				return marshalErr
			}
			return c.Status(http.StatusInternalServerError).Send(data)
		}
	}
}
