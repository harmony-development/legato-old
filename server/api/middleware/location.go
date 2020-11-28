package middleware

import (
	"context"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) LocationInterceptor(c context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx := c.(HarmonyContext)

	if err := LocationHandler(m.DB, req, info.FullMethod, ctx.UserID); err != nil {
		return nil, err
	}
	return handler(c, req)
}

func LocationHandler(database db.IHarmonyDB, req interface{}, fullMethod string, userID uint64) error {
	if GetRPCConfig(fullMethod).Location.Has(NoLocation) {
		return nil
	}

	locFlags := rpcConfigs[fullMethod].Location

	if locFlags.Has(GuildLocation) {
		location, ok := req.(interface {
			GetGuildId() uint64
		})
		if !ok {
			panic("guild location middleware used on message without a guild location")
		}
		guildID := location.GetGuildId()

		if guildID == 0 {
			return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
		}

		ok, err := database.HasGuildWithID(guildID)
		if err != nil {
			return status.Error(codes.Internal, responses.InternalServerError)
		}
		if !ok {
			return status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
		}

		if locFlags.Has(ChannelLocation) {
			location, ok := req.(interface {
				GetChannelId() uint64
			})

			if !ok {
				panic("channel location middleware used on message without a channel location")
			}

			channelID := location.GetChannelId()

			if channelID == 0 {
				return status.Error(codes.InvalidArgument, responses.BadLocationChannel)
			}
			ok, err := database.HasChannelWithID(guildID, channelID)
			if err != nil {
				return status.Error(codes.Internal, responses.InternalServerError)
			}
			if !ok {
				return status.Error(codes.FailedPrecondition, responses.BadLocationChannel)
			}

			if locFlags.Has(MessageLocation) {
				location, ok := req.(interface {
					GetMessageId() uint64
				})

				if !ok {
					panic("message location middleware used on message without a message location")
				}

				messageID := location.GetMessageId()

				if messageID == 0 {
					return status.Error(codes.InvalidArgument, responses.BadLocationMessage)
				}
				ok, err := database.HasMessageWithID(guildID, channelID, messageID)
				if err != nil {
					return status.Error(codes.Internal, responses.InternalServerError)
				}
				if !ok {
					return status.Error(codes.FailedPrecondition, responses.BadLocationMessage)
				}
				if locFlags.Has(AuthorLocation) {
					owner, err := database.GetMessageOwner(messageID)
					if err != nil {
						return status.Error(codes.Internal, responses.InternalServerError)
					}
					if owner != userID {
						return status.Error(codes.PermissionDenied, responses.BadLocationMessage)
					}
				}
			}
		}
		if locFlags.Has(JoinedLocation) {
			if guildID == 0 {
				return status.Error(codes.InvalidArgument, responses.MissingLocationGuild)
			}
			ok, err := database.UserInGuild(userID, guildID)
			if err != nil {
				return status.Error(codes.Internal, responses.InternalServerError)
			}
			if !ok {
				return status.Error(codes.FailedPrecondition, responses.BadLocationGuild)
			}
		}
	}
	return nil
}
