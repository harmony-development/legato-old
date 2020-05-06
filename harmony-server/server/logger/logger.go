package logger

import (
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"harmony-server/server/config"
	"os"
)

type Logger struct {
	Config *config.Config
}

func New(cfg *config.Config) *Logger {
	return &Logger{
		Config: cfg,
	}
}

func (l Logger) Exception(err error) {
	if l.Config.Sentry.Enabled {
		sentry.CaptureException(errors.WithStack(err))
	}
	if l.Config.Server.LogErrors {
		logrus.Warn(err)
	}
}

func (l Logger) Fatal(err error) {
	l.Exception(err)
	os.Exit(-1)
}