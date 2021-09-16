// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"errors"
	"net/http"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/hrpc/server"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/errwrap"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
)

// FiberRPCHandler converts a RPC handler to a Fiber handler.
func FiberRPCHandler(handler server.RawHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp, err := handler(c.Context(), c.Request())
		if err != nil {
			return err
		}

		return errwrap.Wrap(c.Send(resp), "")
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

// Setup registers the Harmony protocol API.
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

		contentType := string(c.Request().Header.Peek("Content-Type"))

		var herr *Error
		if errors.As(e, &herr) {
			data, err := server.MarshalHRPC(&harmonytypesv1.Error{
				Identifier:   herr.Identifier,
				HumanMessage: herr.HumanMessage,
				MoreDetails:  herr.MoreDetails,
			}, contentType)
			if err != nil {
				return errwrap.Wrapf(err, "failed to wrap %v", data)
			}

			// nolint
			return c.Status(http.StatusBadRequest).Send(data)
		}

		err := &harmonytypesv1.Error{
			Identifier: ErrorInternalServerError,
		}
		if cfg.Debug.RespondWithErrors {
			err.HumanMessage = e.Error()
		}

		data, marshalErr := server.MarshalHRPC(err, contentType)
		if marshalErr != nil {
			return errwrap.Wrapf(marshalErr, "failed to marshal error: (%s, %v)", contentType, err)
		}

		// nolint
		return c.Status(http.StatusInternalServerError).Send(data)
	}
}
