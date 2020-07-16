package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func (m Middlewares) HarmonyContextInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := HarmonyContext{
		Context: c,
	}
	return handler(ctx, req)
}
