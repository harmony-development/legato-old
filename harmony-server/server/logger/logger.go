package logger

import (
	"fmt"
	"github.com/getsentry/sentry-go"
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

func (l Logger) Exception(err ...interface{}) {
	if l.Config.Sentry.Enabled {
		sentry.CaptureMessage(fmt.Sprint(err))
	}
	if l.Config.Server.LogErrors {
		logrus.Warn(err)
	}
}

func (l Logger) Fatal(err ...interface{}) {
	l.Exception(err)
	os.Exit(-1)
}