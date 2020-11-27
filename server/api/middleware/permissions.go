package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) GuildPermissionInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if GetRPCConfig(info.FullMethod).Permission == NoPermission {
		return handler(c, req)
	}

	ctx := c.(HarmonyContext)
	location, ok := req.(interface {
		GetGuildId() uint64
	})
	if !ok {
		panic("guild permission middleware used on message without a location")
	}
	guildID := location.GetGuildId()
	if rpcConfigs[info.FullMethod].Permission.HasAny(ModifyInvites, ModifyChannels, ModifyGuild, Owner) {
		owner, err := m.DB.GetOwner(guildID)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if owner != ctx.UserID {
			return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
		}
	}
	return handler(ctx, req)
}
