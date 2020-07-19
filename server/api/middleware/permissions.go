package middleware

import (
	"context"

	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) GuildPermissionInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)
	location, ok := req.(interface{ GetLocation() *corev1.Location })
	if !ok {
		panic("guild permission middleware used on message without a location")
	}
	loc := location.GetLocation()
	if rpcConfigs[info.FullMethod].Permission.HasAny(ModifyInvites, ModifyChannels, ModifyGuild, Owner) {
		owner, err := m.DB.GetOwner(loc.GuildId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if owner != ctx.UserID {
			return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
		}
	}
	return handler(ctx, req)
}
