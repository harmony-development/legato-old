package middleware

import (
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (m Middlewares) LocationHandler(req interface{}, fullMethod string, userID uint64) error {
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

		ok, err := m.DB.HasGuildWithID(guildID)
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
			ok, err := m.DB.HasChannelWithID(guildID, channelID)
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
				ok, err := m.DB.HasMessageWithID(guildID, channelID, messageID)
				if err != nil {
					return status.Error(codes.Internal, responses.InternalServerError)
				}
				if !ok {
					return status.Error(codes.FailedPrecondition, responses.BadLocationMessage)
				}
				if locFlags.Has(AuthorLocation) {
					owner, err := m.DB.GetMessageOwner(messageID)
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
			ok, err := m.DB.UserInGuild(userID, guildID)
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
