package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) GuildPermissionInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if GetRPCConfig(info.FullMethod).Permission == "" {
		ctx := c.(HarmonyContext)
		if GetRPCConfig(info.FullMethod).WantsRoles {
			location, ok := req.(interface {
				GetGuildId() uint64
			})
			if !ok {
				panic("wants roles middleware used on message without a location")
			}
			roles, err := m.DB.RolesForUser(location.GetGuildId(), ctx.UserID)
			if err != nil {
				return nil, status.Error(codes.Internal, responses.InternalServerError)
			}
			ctx.UserRoles = roles
			owner, err := m.DB.GetOwner(location.GetGuildId())
			if err != nil {
				return nil, status.Error(codes.Internal, responses.InternalServerError)
			}
			ctx.IsOwner = owner == ctx.UserID
		}
		return handler(ctx, req)
	}

	ctx := c.(HarmonyContext)
	location, ok := req.(interface {
		GetGuildId() uint64
	})
	if !ok {
		panic("guild permission middleware used on message without a location")
	}
	guildID := location.GetGuildId()
	owner, err := m.DB.GetOwner(guildID)
	if err != nil {
		return nil, status.Error(codes.Internal, responses.InternalServerError)
	}
	if owner == ctx.UserID {
		return handler(c, req)
	}

	channelID := uint64(0)
	channelLocation, ok := req.(interface {
		GetChannelId() uint64
	})
	if ok {
		channelID = channelLocation.GetChannelId()
	}

	roles, err := m.DB.RolesForUser(guildID, ctx.UserID)
	if err != nil {
		return nil, status.Error(codes.Internal, responses.InternalServerError)
	}
	ctx.UserRoles = roles

	if !m.Perms.Check(GetRPCConfig(info.FullMethod).Permission, roles, guildID, channelID) {
		return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
	}

	return handler(ctx, req)
}
