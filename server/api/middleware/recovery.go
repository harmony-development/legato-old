package middleware

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryRecoveryFunc recovers unary requests
func (m Middlewares) UnaryRecoveryFunc(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			m.Logger.Exception(fmt.Errorf("%+v", r))
			m.Logger.Exception(errors.New(string(debug.Stack())))
			err = status.Error(codes.Internal, responses.UnknownError)
		}
	}()
	return handler(ctx, req)
}

// StreamRecoveryFunc recovers streams
func (m Middlewares) StreamRecoveryFunc(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			m.Logger.Exception(fmt.Errorf("%+v", r))
			m.Logger.Exception(errors.New(string(debug.Stack())))
			err = status.Error(codes.Internal, responses.UnknownError)
		}
	}()

	return handler(srv, ss)
}
