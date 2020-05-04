package logger

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"harmony-server/server/config"
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
		sentry.CaptureException(err)
	}
	if l.Config.Server.LogErrors {
		logrus.Warn(err.Error())
	}
}
