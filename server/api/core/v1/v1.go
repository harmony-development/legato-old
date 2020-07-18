package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	NoPermissionsError = errors.New("No permissions")
	NotInGuild         = errors.New("Not in guild")
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

func (v1 *V1) EnsureOwner(guildID, ownerID uint64) error {
	owner, err := v1.DB.GetOwner(guildID)
	if err != nil {
		return err
	}
	if owner == ownerID {
		return nil
	}
	return NoPermissionsError
}

func (v1 *V1) EnsureInGuild(guildID, userID uint64) error {
	ok, err := v1.DB.UserInGuild(userID, guildID)
	if err != nil {
		return err
	}
	if !ok {
		return NotInGuild
	}
	return nil
}

func (v1 *V1) ActionsToProto(msgs []json.RawMessage) (ret []*corev1.Action) {
	for _, msg := range msgs {
		var action *corev1.Action
		json.Unmarshal([]byte(msg), &action)
		ret = append(ret, action)
	}
	return
}

func (v1 *V1) ProtoToActions(msgs []*corev1.Action) (ret [][]byte) {
	for _, msg := range msgs {
		data, _ := json.Marshal(msg)
		ret = append(ret, json.RawMessage(data))
	}
	return
}

func (v1 *V1) EmbedsToProto(embeds []json.RawMessage) (ret []*corev1.Embed) {
	for _, embed := range embeds {
		var action *corev1.Embed
		json.Unmarshal([]byte(embed), &action)
		ret = append(ret, action)
	}
	return
}

func (v1 *V1) ProtoToEmbeds(embeds []*corev1.Embed) (ret [][]byte) {
	for _, embed := range embeds {
		data, _ := json.Marshal(embed)
		ret = append(ret, json.RawMessage(data))
	}
	return
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
	ctx := c.(middleware.HarmonyContext)
	if err := v1.EnsureOwner(r.ForGuild, ctx.UserID); err != nil {
		return nil, err
	}
	inv := int32(-1)
	if r.PossibleUses != nil {
		inv = *r.PossibleUses
	}
	invite, err := v1.DB.CreateInvite(r.ForGuild, inv, r.Name)
	if err != nil {
		return nil, err
	}
	return &corev1.CreateInviteResponse{
		Name: invite.InviteID,
	}, nil
}

func (v1 *V1) CreateChannel(c context.Context, r *corev1.CreateChannelRequest) (*corev1.CreateChannelResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := v1.EnsureOwner(r.GuildId, ctx.UserID); err != nil {
		return nil, err
	}
	channel, err := v1.DB.AddChannelToGuild(r.GuildId, r.ChannelName)
	if err != nil {
		return nil, err
	}
	return &corev1.CreateChannelResponse{
		ChannelId: channel.ChannelID,
	}, nil
}

func (v1 *V1) GetGuild(c context.Context, r *corev1.GetGuildRequest) (*corev1.GetGuildResponse, error) {
	guild, err := v1.DB.GetGuildByID(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &corev1.GetGuildResponse{
		GuildName:    guild.GuildName,
		GuildOwner:   guild.OwnerID,
		GuildPicture: guild.PictureUrl,
	}, nil
}

func (v1 *V1) GetGuildInvites(c context.Context, r *corev1.GetGuildInvitesRequest) (*corev1.GetGuildInvitesResponse, error) {
	err := v1.EnsureOwner(r.GuildId, c.(middleware.HarmonyContext).UserID)
	if err != nil {
		return nil, err
	}
	invites, err := v1.DB.GetInvites(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &corev1.GetGuildInvitesResponse{
		Invites: func() (ret []*corev1.GetGuildInvitesResponse_Invite) {
			for _, inv := range invites {
				ret = append(ret, &corev1.GetGuildInvitesResponse_Invite{
					InviteId: inv.InviteID,
					PossibleUses: func() int32 {
						if inv.PossibleUses.Valid {
							return inv.PossibleUses.Int32
						}
						return -1
					}(),
					UseCount: inv.Uses,
				})
			}
			return
		}(),
	}, nil
}

func (v1 *V1) GetGuildMembers(c context.Context, r *corev1.GetGuildMembersRequest) (*corev1.GetGuildMembersResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.EnsureInGuild(r.GuildId, ctx.UserID)
	if err != nil {
		return nil, err
	}
	members, err := v1.DB.MembersInGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &corev1.GetGuildMembersResponse{
		Members: members,
	}, nil
}

func (v1 *V1) GetGuildChannels(c context.Context, r *corev1.GetGuildChannelsRequest) (*corev1.GetGuildChannelsResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.EnsureInGuild(r.GuildId, ctx.UserID)
	if err != nil {
		return nil, err
	}
	chans, err := v1.DB.ChannelsForGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &corev1.GetGuildChannelsResponse{
		Channels: func() (ret []*corev1.GetGuildChannelsResponse_Channel) {
			for _, channel := range chans {
				ret = append(ret, &corev1.GetGuildChannelsResponse_Channel{
					ChannelId:   channel.ChannelID,
					ChannelName: channel.ChannelName,
				})
			}
			return
		}(),
	}, nil
}

func (v1 *V1) GetChannelMessages(c context.Context, r *corev1.GetChannelMessagesRequest) (*corev1.GetChannelMessagesResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.EnsureInGuild(r.GuildId, ctx.UserID)
	if err != nil {
		return nil, err
	}
	var messages []queries.Message
	if r.BeforeMessage != nil {
		time, err := v1.DB.GetMessageDate(*r.BeforeMessage)
		if err != nil {
			return nil, err
		}
		messages, err = v1.DB.GetMessagesBefore(r.GuildId, r.ChannelId, time)
	} else {
		messages, err = v1.DB.GetMessages(r.GuildId, r.ChannelId)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &corev1.GetChannelMessagesResponse{
		Messages: func() (ret []*corev1.GetChannelMessagesResponse_Message) {
			for _, message := range messages {
				createdAt, _ := ptypes.TimestampProto(message.CreatedAt.UTC())
				var editedAt *timestamppb.Timestamp
				if message.EditedAt.Valid {
					editedAt, _ = ptypes.TimestampProto(editedAt.AsTime().UTC())
				}
				var embeds []*corev1.Embed
				var actions []*corev1.Action
				for _, rawEmbed := range message.Embeds {
					var embed corev1.Embed
					if err := json.Unmarshal(rawEmbed, &embed); err != nil {
						continue
					}
					embeds = append(embeds, &embed)
				}
				for _, rawAction := range message.Actions {
					var action corev1.Action
					if err := json.Unmarshal(rawAction, &action); err != nil {
						continue
					}
					actions = append(actions, &action)
				}
				ret = append(ret, &corev1.GetChannelMessagesResponse_Message{
					MessageId: message.MessageID,
					GuildId:   message.GuildID,
					ChannelId: message.ChannelID,
					AuthorId:  message.UserID,
					CreatedAt: createdAt,
					EditedAt:  editedAt,
					Content:   message.Content,
					Embeds:    embeds,
					Actions:   actions,
				})
			}
			return
		}(),
	}, nil
}

func (v1 *V1) UpdateGuildName(c context.Context, r *corev1.UpdateGuildNameRequest) (*corev1.UpdateGuildNameResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.EnsureOwner(r.GuildId, ctx.UserID)
	if err != nil {
		return nil, err
	}
	if err := v1.DB.UpdateGuildName(r.GuildId, r.NewGuildName); err != nil {
		return nil, err
	}
	return &corev1.UpdateGuildNameResponse{}, nil
}

func (v1 *V1) UpdateMessage(c context.Context, r *corev1.UpdateMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if !r.UpdateActions && !r.UpdateEmbeds && !r.UpdateContent {
		return nil, errors.New("bad request; nothing is being edited")
	}

	owner, err := v1.DB.GetMessageOwner(r.MessageId)
	if err != nil {
		return nil, err
	}
	if owner != ctx.UserID {
		return nil, NoPermissionsError
	}

	var actions *[][]byte
	var embeds *[][]byte
	if r.UpdateActions {
		val := v1.ProtoToActions(r.Actions)
		actions = &val
	}
	if r.UpdateEmbeds {
		val := v1.ProtoToEmbeds(r.Embeds)
		embeds = &val
	}
	_, err = v1.DB.UpdateMessage(r.MessageId, r.Content, embeds, actions)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (v1 *V1) DeleteGuild(c context.Context, r *corev1.DeleteGuildRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteInvite(c context.Context, r *corev1.DeleteInviteRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteChannel(c context.Context, r *corev1.DeleteChannelRequest) (*empty.Empty, error) {

}

func (v1 *V1) DeleteMessage(c context.Context, r *corev1.DeleteMessageRequest) (*empty.Empty, error) {

}
