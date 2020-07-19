package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func (m Middlewares) GuildPermissionInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	if rpcConfigs[info.FullMethod].Permission != ModifyGuild {
		return handler(ctx, req)
	}
	return handler(ctx, req)
}
