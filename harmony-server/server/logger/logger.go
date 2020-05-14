package logger

import (
	"harmony-server/server/config"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Logger is the Harmony logger
type Logger struct {
	Config *config.Config
}

// New creates a Logger with a given configuration
func New(cfg *config.Config) *Logger {
	return &Logger{
		Config: cfg,
	}
}

// Exception logs an exception
func (l Logger) Exception(err error) {
	if l.Config.Sentry.Enabled {
		sentry.CaptureException(errors.WithStack(err))
	}
	if l.Config.Server.LogErrors {
		logrus.Warn(err)
	}
}

// Fatal logs an exception and then aborts
func (l Logger) Fatal(err error) {
	l.Exception(err)
	os.Exit(-1)
}
