package middleware

import (
	"context"

	"google.golang.org/grpc"
)

// LocalInterceptor ensures a user is a local user
func (m Middlewares) LocalInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !GetRPCConfig(info.FullMethod).Local {
		return handler(c, req)
	}

	ctx := c.(HarmonyContext)
	err := m.DB.UserIsLocal(ctx.UserID)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
