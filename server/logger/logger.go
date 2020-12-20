package logger

import (
	"context"
	"database/sql"
	"os"

	"github.com/harmony-development/legato/server/config"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	ErrorResponse(code codes.Code, err error, response string) error
	CheckException(err error)
	Exception(err error)
	Request(c context.Context, req interface{}, info *grpc.UnaryServerInfo)
	Fatal(err error)
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
	if err == nil || err == sql.ErrNoRows {
		return
	}
	l.Exception(err)
}

func (l Logger) ErrorResponse(code codes.Code, err error, response string) error {
	if l.Config.Server.Policies.Debug.RespondWithErrors {
		if l.Config.Server.Policies.Debug.ResponseErrorsIncludeTrace {
			return status.Error(code, tracerr.Sprint(err))
		}
		return status.Error(code, err.Error())
	}
	return status.Error(code, response)
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

// Fatal logs an exception and then aborts
func (l Logger) Fatal(err error) {
	l.CheckException(err)
	os.Exit(-1)
}

func (l Logger) Request(c context.Context, req interface{}, info *grpc.UnaryServerInfo) {
	if l.Config.Server.Policies.Debug.LogRequests {
		p, ok := peer.FromContext(c)
		ip := ""
		if ok {
			ip = p.Addr.String()
		}
		logrus.Info("[", info.FullMethod, "] FROM ", ip)
	}
}
