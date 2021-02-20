package logger

import (
	"database/sql"
	"errors"
	"os"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/harmony-development/legato/server/config"
	"github.com/ztrue/tracerr"

	"github.com/harmony-development/sentry-go"
	"github.com/sirupsen/logrus"
)

// DebugScope is enums for types of debug
type DebugScope int

const (
	Streams DebugScope = iota
	Startup
)

type ILogger interface {
	CheckException(err error)
	Exception(err error)
	Debug(d DebugScope, v ...interface{})
	Verbose(d DebugScope, format string, v ...interface{})
	Fatal(err error)
	Warn(s string, v ...interface{})
}

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

// CheckException logs an exception if it's defined
func (l Logger) CheckException(err error) {
	if err == nil || errors.Is(err, sql.ErrNoRows) {
		return
	}
	l.Exception(err)
}

func (l Logger) Verbose(d DebugScope, format string, v ...interface{}) {
	switch d {
	case Streams:
		if !l.Config.Server.Policies.Debug.VerboseStreamHandling {
			return
		}
	}

	logrus.Infof(format, v...)
}

// Exception logs an exception
func (l Logger) Exception(err error) {
	if l.Config.Sentry.Enabled {
		sentry.CaptureException(err)
	}
	if l.Config.Server.Policies.Debug.LogErrors {
		logrus.Warnf("%s", tracerr.SprintSourceColor(err))
	}
}

// Warn warns
func (l Logger) Warn(fm string, v ...interface{}) {
	if l.Config.Server.Policies.Debug.LogErrors {
		logrus.Warnf("%s", v...)
	}
}

// Fatal logs an exception and then aborts
func (l Logger) Fatal(err error) {
	if err == nil {
		panic("fatal called with nil error")
	}
	println(tracerr.SprintSourceColor(err))
	os.Exit(-1)
}

func (l Logger) Debug(d DebugScope, v ...interface{}) {
	switch d {
	case Streams:
		if !l.Config.Server.Policies.Debug.VerboseStreamHandling {
			return
		}
	}

	var message strings.Builder

	for _, i := range v {
		switch value := i.(type) {
		case string:
			message.WriteString(value)
		default:
			message.WriteString(repr.String(value))
		}
		message.WriteString(" ")
	}

	logrus.Debug("[STREAMS] ", message.String())

}
