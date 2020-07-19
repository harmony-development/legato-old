package middleware

import (
	"context"

	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) LocationInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if GetRPCConfig(info.FullMethod).Location == NoLocation {
		return handler(c, req)
	}

	ctx := c.(HarmonyContext)
	location, ok := req.(interface{ GetLocation() *corev1.Location })
	if !ok {
		panic("location middleware used on message without a location")
	}
	loc := location.GetLocation()
	locFlags := rpcConfigs[info.FullMethod].Location
	if locFlags.Has(GuildLocation) {
		if loc.GuildId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		ok, err := m.DB.HasGuildWithID(loc.GuildId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return nil, status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
		}
	}
	if locFlags.Has(ChannelLocation) {
		if loc.ChannelId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		ok, err := m.DB.HasChannelWithID(loc.GuildId, loc.ChannelId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return nil, status.Error(codes.FailedPrecondition, responses.BadLocationChannel)
		}
	}
	if locFlags.Has(MessageLocation) {
		if loc.GuildId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		if loc.ChannelId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		if loc.MessageId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.BadLocationMessage)
		}
		ok, err := m.DB.HasMessageWithID(loc.GuildId, loc.ChannelId, loc.MessageId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return nil, status.Error(codes.FailedPrecondition, responses.BadLocationMessage)
		}
	}
	if locFlags.Has(JoinedLocation) {
		if loc.GuildId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		ok, err := m.DB.UserInGuild(ctx.UserID, loc.GuildId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return nil, status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
		}
	}
	if locFlags.Has(AuthorLocation) {
		if loc.GuildId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		if loc.ChannelId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		if loc.MessageId == 0 {
			return nil, status.Error(codes.InvalidArgument, responses.BadLocationMessage)
		}
		owner, err := m.DB.GetMessageOwner(loc.MessageId)
		if err != nil {
			return nil, status.Error(codes.Internal, responses.InternalServerError)
		}
		if owner != ctx.UserID {
			return nil, status.Error(codes.PermissionDenied, responses.BadLocationMessage)
		}
	}
	return handler(ctx, req)
}
