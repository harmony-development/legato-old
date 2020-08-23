package middleware

import (
	"context"

	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) LocationInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)

	if err := LocationHandler(m.DB, ctx, req, info.FullMethod, ctx.UserID); err != nil {
		return nil, err
	}
	return handler(c, req)
}

func LocationHandler(database db.IHarmonyDB, c context.Context, req interface{}, fullMethod string, userID uint64) error {
	if GetRPCConfig(fullMethod).Location == NoLocation {
		return nil
	}

	location, ok := req.(interface{ GetLocation() *corev1.Location })
	if !ok {
		panic("location middleware used on message without a location")
	}
	loc := location.GetLocation()
	locFlags := rpcConfigs[fullMethod].Location
	if !locFlags.Has(NoLocation) && loc == nil {
		return status.Error(codes.FailedPrecondition, responses.MissingLocation)
	}
	if locFlags.Has(GuildLocation) {
		if loc.GuildId == 0 {
			return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		ok, err := database.HasGuildWithID(loc.GuildId)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
		}
	}
	if locFlags.Has(ChannelLocation) {
		if loc.ChannelId == 0 {
			return status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		ok, err := database.HasChannelWithID(loc.GuildId, loc.ChannelId)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return status.Error(codes.FailedPrecondition, responses.BadLocationChannel)
		}
	}
	if locFlags.Has(MessageLocation) {
		if loc.GuildId == 0 {
			return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		if loc.ChannelId == 0 {
			return status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		if loc.MessageId == 0 {
			return status.Error(codes.InvalidArgument, responses.BadLocationMessage)
		}
		ok, err := database.HasMessageWithID(loc.GuildId, loc.ChannelId, loc.MessageId)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return status.Error(codes.FailedPrecondition, responses.BadLocationMessage)
		}
	}
	if locFlags.Has(JoinedLocation) {
		if loc.GuildId == 0 {
			return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		ok, err := database.UserInGuild(userID, loc.GuildId)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
		}
	}
	if locFlags.Has(AuthorLocation) {
		if loc.GuildId == 0 {
			return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}
		if loc.ChannelId == 0 {
			return status.Error(codes.InvalidArgument, responses.BadLocationChannel)
		}
		if loc.MessageId == 0 {
			return status.Error(codes.InvalidArgument, responses.BadLocationMessage)
		}
		owner, err := database.GetMessageOwner(loc.MessageId)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if owner != userID {
			return status.Error(codes.PermissionDenied, responses.BadLocationMessage)
		}
	}
	return nil
}
