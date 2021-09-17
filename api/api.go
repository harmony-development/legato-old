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

		return errwrap.Wrap(c.Send(resp), "failed to send response")
	}
}

func RegisterHandlers(app *fiber.App, l log.Interface, cfg *config.Config, all ...server.HRPCServiceHandler) {
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

func newHarmonyError(cfg *config.Config, e error) (int, *harmonytypesv1.Error) {
	var herr *Error
	if errors.As(e, &herr) {
		return http.StatusBadRequest, &harmonytypesv1.Error{
			Identifier:   herr.Identifier,
			HumanMessage: herr.HumanMessage,
			MoreDetails:  herr.MoreDetails,
		}
	}

	err := &harmonytypesv1.Error{
		Identifier: ErrorInternalServerError,
	}
	if cfg.Debug.RespondWithErrors {
		err.HumanMessage = e.Error()
	}

	return http.StatusInternalServerError, err
}

func FiberErrorHandler(l log.Interface, cfg *config.Config) fiber.ErrorHandler {
	return func(c *fiber.Ctx, e error) error {
		if cfg.Debug.LogErrors && e != nil {
			l.WithError(e).WithFields(log.Fields{
				"path":  c.OriginalURL(),
				"error": e,
			}).Error("error in http handler")
		}

		contentType := string(c.Request().Header.Peek("Content-Type"))
		statusCode, err := newHarmonyError(cfg, e)

		data, marshalErr := server.MarshalHRPC(err, contentType)
		if marshalErr != nil {
			return errwrap.Wrapf(marshalErr, "failed to marshal error: (%s, %v)", contentType, e)
		}

		// nolint
		return c.Status(statusCode).Send(data)
	}
}
