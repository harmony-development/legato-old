package server

import (
	"github.com/harmony-development/legato/server/config"

	"github.com/getsentry/sentry-go"
)

// This file shouldn't exist, but sentry's package is garbage so I had to deviate from the standard package layout.
// Lord forgive me for my sins.

// ConnectSentry connects to sentry
func ConnectSentry(cfg *config.Config) error {
	if !cfg.Sentry.Enabled {
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.DSN,
		AttachStacktrace: cfg.Sentry.AttachStacktraces,
	})
	if err != nil {
		return err
	}

	sentry.CaptureMessage("Harmony Server Started")

	return nil
}
