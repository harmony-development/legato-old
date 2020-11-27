package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func (m Middlewares) LoggingInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)

	m.Logger.Request(c, req, info)

	return handler(ctx, req)
}
