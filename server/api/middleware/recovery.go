package middleware

import (
	"errors"
	"runtime/debug"

	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) RecoveryFunc(p interface{}) error {
	if err, ok := p.(error); ok {
		m.Logger.Exception(err)
	}
	m.Logger.Exception(errors.New(string(debug.Stack())))
	return status.Error(codes.Internal, responses.UnknownError)
}
