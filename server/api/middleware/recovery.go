package middleware

import (
	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) RecoveryFunc(p interface{}) error {
	if err, ok := p.(error); ok {
		m.Logger.Exception(err)
	}
	return status.Error(codes.Internal, responses.UnknownError)
}
