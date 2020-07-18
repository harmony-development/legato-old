package v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB        db.IHarmonyDB
	Logger    logger.ILogger
	Sonyflake *sonyflake.Sonyflake
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies
}

func (v1 *V1) CreateGuild(c context.Context, r *corev1.CreateGuildRequest) (*corev1.CreateGuildResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	guildID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}
	str := ""
	if r.PictureUrl != nil {
		str = *r.PictureUrl
	}
	guild, err := v1.DB.CreateGuild(ctx.UserID, guildID, r.GuildName, str)
	if err != nil {
		return nil, err
	}
	return &corev1.CreateGuildResponse{
		GuildId: guild.GuildID,
	}, nil
}

func (v1 *V1) CreateInvite(c context.Context, r *corev1.CreateInviteRequest) (*corev1.CreateInviteResponse, error) {

}

func (v1 *V1) CreateChannel(c context.Context, r *corev1.CreateChannelRequest) (*corev1.CreateChannelResponse, error) {

}

func (v1 *V1) GetGuild(c context.Context, r *corev1.GetGuildRequest) (*corev1.GetGuildResponse, error) {

}

func (v1 *V1) GetGuildInvites(c context.Context, r *corev1.GetGuildInvitesRequest) (*corev1.GetGuildInvitesResponse, error) {

}

func (v1 *V1) GetGuildMembers(c context.Context, r *corev1.GetGuildMembersRequest) (*corev1.GetGuildMembersResponse, error) {

}

func (v1 *V1) GetGuildChannels(c context.Context, r *corev1.GetGuildChannelsRequest) (*corev1.GetGuildChannelsResponse, error) {

}

func (v1 *V1) GetChannelMessages(c context.Context, r *corev1.GetChannelMessagesRequest) (*corev1.GetChannelMessagesResponse, error) {

}

func (v1 *V1) UpdateGuildName(c context.Context, r *corev1.UpdateGuildNameRequest) (*corev1.UpdateGuildNameResponse, error) {

}

func (v1 *V1) UpdateMessage(c context.Context, r *corev1.UpdateMessageRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteGuild(c context.Context, r *corev1.DeleteGuildRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteInvite(c context.Context, r *corev1.DeleteInviteRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteChannel(c context.Context, r *corev1.DeleteChannelRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteMessage(c context.Context, r *corev1.DeleteMessageRequest) (*empty.Empty, error) {

}
