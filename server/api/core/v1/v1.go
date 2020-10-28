package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/responses"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	PubSub    SubscriptionManager
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies
}

func (v1 *V1) ActionsToProto(msgs json.RawMessage) (ret []*corev1.Action) {
	json.Unmarshal([]byte(msgs), &ret)
	return
}

func (v1 *V1) ProtoToActions(msgs []*corev1.Action) (ret []byte) {
	ret, _ = json.Marshal(msgs)
	return
}

func (v1 *V1) EmbedsToProto(embeds json.RawMessage) (ret []*corev1.Embed) {
	json.Unmarshal([]byte(embeds), &ret)
	return
}

func (v1 *V1) ProtoToEmbeds(embeds []*corev1.Embed) (ret []byte) {
	ret, _ = json.Marshal(embeds)
	return
}

func (v1 *V1) OverridesToProto(overrides []byte) (ret *corev1.Override) {
	if len(overrides) == 0 {
		return
	}
	ret = new(corev1.Override)
	proto.Unmarshal(overrides, ret)
	return
}

func (v1 *V1) ProtoToOverrides(overrides *corev1.Override) (ret []byte) {
	ret, _ = proto.Marshal(overrides)
	return
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    1,
		},
		Auth:       true,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/CreateGuild")
}

func (v1 *V1) CreateGuild(c context.Context, r *corev1.CreateGuildRequest) (*corev1.CreateGuildResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	guildID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}
	channelID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}
	guild, err := v1.DB.CreateGuild(ctx.UserID, guildID, channelID, r.GuildName, r.PictureUrl)
	if err != nil {
		return nil, err
	}
	return &corev1.CreateGuildResponse{
		GuildId: guild.GuildID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyInvites,
	}, "/protocol.core.v1.CoreService/CreateInvite")
}

func (v1 *V1) CreateInvite(c context.Context, r *corev1.CreateInviteRequest) (*corev1.CreateInviteResponse, error) {
	inv := int32(-1)
	if r.PossibleUses != 0 {
		inv = r.PossibleUses
	}
	invite, err := v1.DB.CreateInvite(r.Location.GuildId, inv, r.Name)
	if err != nil {
		return nil, err
	}
	return &corev1.CreateInviteResponse{
		Name: invite.InviteID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyChannels,
	}, "/protocol.core.v1.CoreService/CreateChannel")
}

func (v1 *V1) CreateChannel(c context.Context, r *corev1.CreateChannelRequest) (*corev1.CreateChannelResponse, error) {
	channel, err := v1.DB.AddChannelToGuild(r.Location.GuildId, r.ChannelName, r.PreviousId, r.NextId, r.IsCategory)
	if err != nil {
		return nil, err
	}
	r.Location.ChannelId = channel.ChannelID
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_CreatedChannel{
			CreatedChannel: &corev1.GuildEvent_ChannelCreated{
				Location:   r.Location,
				Name:       r.ChannelName,
				PreviousId: r.PreviousId,
				NextId:     r.NextId,
				IsCategory: r.IsCategory,
			},
		},
	})
	return &corev1.CreateChannelResponse{
		ChannelId: channel.ChannelID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetGuild")
}

func (v1 *V1) GetGuild(c context.Context, r *corev1.GetGuildRequest) (*corev1.GetGuildResponse, error) {
	guild, err := v1.DB.GetGuildByID(r.Location.GuildId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, v1.Logger.ErrorResponse(codes.NotFound, err, responses.GuildNotFound)
		}
		return nil, err
	}
	return &corev1.GetGuildResponse{
		GuildName:    guild.GuildName,
		GuildOwner:   guild.OwnerID,
		GuildPicture: guild.PictureUrl,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyInvites,
	}, "/protocol.core.v1.CoreService/GetGuildInvites")
}

func (v1 *V1) GetGuildInvites(c context.Context, r *corev1.GetGuildInvitesRequest) (*corev1.GetGuildInvitesResponse, error) {
	invites, err := v1.DB.GetInvites(r.Location.GuildId)
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

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetGuildMembers")
}

func (v1 *V1) GetGuildMembers(c context.Context, r *corev1.GetGuildMembersRequest) (*corev1.GetGuildMembersResponse, error) {
	members, err := v1.DB.MembersInGuild(r.Location.GuildId)
	if err != nil {
		return nil, err
	}
	return &corev1.GetGuildMembersResponse{
		Members: members,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetGuildChannels")
}

func (v1 *V1) GetGuildChannels(c context.Context, r *corev1.GetGuildChannelsRequest) (*corev1.GetGuildChannelsResponse, error) {
	chans, err := v1.DB.ChannelsForGuild(r.Location.GuildId)
	if err != nil {
		return nil, err
	}
	ret := []*corev1.GetGuildChannelsResponse_Channel{}
	for _, channel := range chans {
		ret = append(ret, &corev1.GetGuildChannelsResponse_Channel{
			ChannelId:   channel.ChannelID,
			ChannelName: channel.ChannelName,
			IsCategory:  channel.Category,
			IsVoice:     channel.Isvoice,
		})
	}
	return &corev1.GetGuildChannelsResponse{
		Channels: ret,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetMessage")
}

func (v1 *V1) GetMessage(c context.Context, r *corev1.GetMessageRequest) (*corev1.GetMessageResponse, error) {
	message, err := v1.DB.GetMessage(r.Location.MessageId)
	if err != nil {
		return nil, err
	}
	createdAt, _ := ptypes.TimestampProto(message.CreatedAt.UTC())
	var editedAt *timestamppb.Timestamp
	if message.EditedAt.Valid {
		editedAt, _ = ptypes.TimestampProto(editedAt.AsTime().UTC())
	}
	var embeds []*corev1.Embed
	var actions []*corev1.Action
	var override *corev1.Override
	if err := json.Unmarshal(message.Embeds, &embeds); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(message.Actions, &actions); err != nil {
		return nil, err
	}
	if len(message.Overrides) > 0 {
		override = new(corev1.Override)
		if err := proto.Unmarshal(message.Overrides, override); err != nil {
			return nil, err
		}
	}
	return &corev1.GetMessageResponse{
		Message: &corev1.Message{
			Location: &corev1.Location{
				MessageId: message.MessageID,
				GuildId:   message.GuildID,
				ChannelId: message.ChannelID,
			},
			AuthorId:  message.UserID,
			CreatedAt: createdAt,
			EditedAt:  editedAt,
			Content:   message.Content,
			Embeds:    embeds,
			Actions:   actions,
			Overrides: override,
			InReplyTo: uint64(message.ReplyToID.Int64),
		},
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetChannelMessages")
}

func (v1 *V1) GetChannelMessages(c context.Context, r *corev1.GetChannelMessagesRequest) (*corev1.GetChannelMessagesResponse, error) {
	var err error
	var messages []queries.Message
	if r.BeforeMessage != 0 {
		time, err := v1.DB.GetMessageDate(r.BeforeMessage)
		if err != nil {
			return nil, err
		}
		messages, err = v1.DB.GetMessagesBefore(r.Location.GuildId, r.Location.ChannelId, time)
	} else {
		messages, err = v1.DB.GetMessages(r.Location.GuildId, r.Location.ChannelId)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &corev1.GetChannelMessagesResponse{
		Messages: func() (ret []*corev1.Message) {
			for _, message := range messages {
				createdAt, _ := ptypes.TimestampProto(message.CreatedAt.UTC())
				var editedAt *timestamppb.Timestamp
				if message.EditedAt.Valid {
					editedAt, _ = ptypes.TimestampProto(editedAt.AsTime().UTC())
				}
				var embeds []*corev1.Embed
				var actions []*corev1.Action
				var overrides *corev1.Override
				if err := json.Unmarshal(message.Embeds, &embeds); err != nil {
					continue
				}
				if err := json.Unmarshal(message.Actions, &actions); err != nil {
					continue
				}
				if len(message.Overrides) > 0 {
					overrides = new(corev1.Override)
					if err := proto.Unmarshal(message.Overrides, overrides); err != nil {
						continue
					}
				}
				ret = append(ret, &corev1.Message{
					Location: &corev1.Location{
						MessageId: message.MessageID,
						GuildId:   message.GuildID,
						ChannelId: message.ChannelID,
					},
					AuthorId:  message.UserID,
					CreatedAt: createdAt,
					EditedAt:  editedAt,
					Content:   message.Content,
					Embeds:    embeds,
					Actions:   actions,
					Overrides: overrides,
					InReplyTo: uint64(message.ReplyToID.Int64),
				})
			}
			return
		}(),
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyGuild,
	}, "/protocol.core.v1.CoreService/UpdateGuildName")
}

func (v1 *V1) UpdateGuildName(c context.Context, r *corev1.UpdateGuildNameRequest) (*empty.Empty, error) {
	if err := v1.DB.UpdateGuildName(r.Location.GuildId, r.NewGuildName); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_EditedGuild{
			EditedGuild: &corev1.GuildEvent_GuildUpdated{
				Name:       r.NewGuildName,
				UpdateName: true,
			},
		},
	})
	return &empty.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyChannels,
	}, "/protocol.core.v1.CoreService/UpdateChannelName")
}

func (v1 *V1) UpdateChannelName(c context.Context, r *corev1.UpdateChannelNameRequest) (*empty.Empty, error) {
	if err := v1.DB.SetChannelName(r.Location.GuildId, r.Location.ChannelId, r.NewChannelName); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_EditedChannel{
			EditedChannel: &corev1.GuildEvent_ChannelUpdated{
				Location:   r.Location,
				Name:       r.NewChannelName,
				UpdateName: true,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
		Permission: middleware.ModifyChannels,
	}, "/protocol.core.v1.CoreService/UpdateChannelOrder")
}

func (v1 *V1) UpdateChannelOrder(c context.Context, r *corev1.UpdateChannelOrderRequest) (*empty.Empty, error) {
	if err := v1.DB.MoveChannel(r.Location.GuildId, r.Location.ChannelId, r.PreviousId, r.NextId); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_EditedChannel{
			EditedChannel: &corev1.GuildEvent_ChannelUpdated{
				Location:    r.Location,
				PreviousId:  r.PreviousId,
				NextId:      r.NextId,
				UpdateOrder: true,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation | middleware.AuthorLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/UpdateMessage")
}

func (v1 *V1) UpdateMessage(c context.Context, r *corev1.UpdateMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if !r.UpdateActions && !r.UpdateEmbeds && !r.UpdateContent && !r.UpdateOverrides {
		return nil, status.Error(codes.InvalidArgument, responses.InvalidRequest)
	}

	owner, err := v1.DB.GetMessageOwner(r.Location.MessageId)
	if err != nil {
		return nil, err
	}
	if owner != ctx.UserID {
		return nil, NoPermissionsError
	}

	var actions *[]byte
	var embeds *[]byte
	var overrides *[]byte
	if r.UpdateActions {
		val := v1.ProtoToActions(r.Actions)
		actions = &val
	}
	if r.UpdateEmbeds {
		val := v1.ProtoToEmbeds(r.Embeds)
		embeds = &val
	}
	if r.UpdateOverrides {
		val := v1.ProtoToOverrides(r.Overrides)
		overrides = &val
	}
	tiempo, err := v1.DB.UpdateMessage(r.Location.MessageId, &r.Content, embeds, actions, overrides)
	if err != nil {
		return nil, err
	}
	editedAt, _ := ptypes.TimestampProto(tiempo.UTC())
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_EditedMessage{
			EditedMessage: &corev1.GuildEvent_MessageUpdated{
				Location:      r.Location,
				Content:       r.Content,
				UpdateContent: r.UpdateContent,
				Embeds:        r.Embeds,
				UpdateEmbeds:  r.UpdateEmbeds,
				Actions:       r.Actions,
				UpdateActions: r.UpdateActions,
				Overrides:     r.Overrides,
				EditedAt:      editedAt,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 15 * time.Second,
			Burst:    1,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: middleware.Owner,
	}, "/protocol.core.v1.CoreService/DeleteGuild")
}

func (v1 *V1) DeleteGuild(c context.Context, r *corev1.DeleteGuildRequest) (*empty.Empty, error) {
	err := v1.DB.DeleteGuild(r.Location.GuildId)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_DeletedGuild{
			DeletedGuild: &corev1.GuildEvent_GuildDeleted{},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: middleware.ModifyInvites,
	}, "/protocol.core.v1.CoreService/DeleteInvite")
}

func (v1 *V1) DeleteInvite(c context.Context, r *corev1.DeleteInviteRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteInvite(r.InviteId); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation,
		Permission: middleware.ModifyChannels,
	}, "/protocol.core.v1.CoreService/DeleteChannel")
}

func (v1 *V1) DeleteChannel(c context.Context, r *corev1.DeleteChannelRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteChannelFromGuild(r.Location.GuildId, r.Location.ChannelId); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_DeletedChannel{
			DeletedChannel: &corev1.GuildEvent_ChannelDeleted{
				Location: r.Location,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation | middleware.AuthorLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/DeleteMessage")
}

func (v1 *V1) DeleteMessage(c context.Context, r *corev1.DeleteMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	owner, err := v1.DB.GetMessageOwner(r.Location.MessageId)
	if err != nil {
		return nil, err
	}
	if ctx.UserID != owner {
		return nil, NoPermissionsError
	}
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_DeletedMessage{
			DeletedMessage: &corev1.GuildEvent_MessageDeleted{
				Location: &corev1.Location{
					MessageId: r.Location.MessageId,
					ChannelId: r.Location.ChannelId,
					GuildId:   r.Location.GuildId,
				},
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.NoLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/JoinGuild")
}

func (v1 *V1) JoinGuild(c context.Context, r *corev1.JoinGuildRequest) (*corev1.JoinGuildResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	guildID, err := v1.DB.ResolveGuildID(r.InviteId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, v1.Logger.ErrorResponse(codes.NotFound, err, responses.GuildNotFound)
		}
		return nil, err
	}
	if err := v1.DB.AddMemberToGuild(ctx.UserID, guildID); err != nil {
		return nil, err
	}
	if err := v1.DB.IncrementInvite(r.InviteId); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(guildID, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_JoinedMember{
			JoinedMember: &corev1.GuildEvent_MemberJoined{
				MemberId: ctx.UserID,
			},
		},
	})
	return &corev1.JoinGuildResponse{
		Location: &corev1.Location{
			GuildId: guildID,
		},
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/LeaveGuild")
}

func (v1 *V1) LeaveGuild(c context.Context, r *corev1.LeaveGuildRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if isOwner, err := v1.DB.IsOwner(r.Location.GuildId, ctx.UserID); err != nil {
		return nil, err
	} else if isOwner {
		return nil, status.Error(codes.FailedPrecondition, responses.InvalidRequest)
	}
	v1.PubSub.Guild.UnsubscribeUserFromGuild(r.Location.GuildId, ctx.UserID)
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_LeftMember{
			LeftMember: &corev1.GuildEvent_MemberLeft{
				MemberId: ctx.UserID,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/StreamGuildEvents")
}

func (v1 *V1) StreamGuildEvents(r *corev1.StreamGuildEventsRequest, s corev1.CoreService_StreamGuildEventsServer) error {
	userID, err := middleware.AuthHandler(v1.DB, s.Context())
	if err != nil {
		return err
	}
	if err := middleware.LocationHandler(v1.DB, r, "/protocol.core.v1.CoreService/StreamGuildEvents", userID); err != nil {
		return err
	}
	ok, err := v1.DB.UserInGuild(userID, r.Location.GuildId)
	if err != nil {
		return err
	}
	if !ok {
		return NotInGuild
	}
	<-v1.PubSub.Guild.Subscribe(r.Location.GuildId, userID, s)
	return nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:       true,
		Location:   middleware.NoLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/StreamActionEvents")
}

func (v1 *V1) StreamActionEvents(r *corev1.StreamActionEventsRequest, s corev1.CoreService_StreamActionEventsServer) error {
	wrappedStream := s.(middleware.IHarmonyWrappedServerStream)
	<-v1.PubSub.Actions.Subscribe(wrappedStream.GetWrappedContext().UserID, s)
	return nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    20,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/TriggerAction")
}

func (v1 *V1) TriggerAction(c context.Context, r *corev1.TriggerActionRequest) (*emptypb.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	msg, err := v1.DB.GetMessage(r.Location.MessageId)
	if err != nil {
		return nil, err
	}
	if msg.ChannelID != r.Location.ChannelId || msg.GuildID != r.Location.GuildId {
		return nil, status.Error(codes.InvalidArgument, responses.InvalidRequest)
	}
	for _, action := range v1.ActionsToProto(msg.Actions) {
		if action.Id == r.ActionId {
			v1.PubSub.Actions.Broadcast(ctx.UserID, &corev1.ActionEvent{
				Event: &corev1.ActionEvent_Action_{
					Action: &corev1.ActionEvent_Action{
						Location:   r.Location,
						ActionData: r.ActionData,
						ActionId:   r.ActionId,
					},
				},
			})
			return &emptypb.Empty{}, nil
		}
	}
	return nil, status.Error(codes.InvalidArgument, responses.InvalidRequest)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    1500,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/SendMessage")
}

func (v1 *V1) SendMessage(c context.Context, r *corev1.SendMessageRequest) (*corev1.SendMessageResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	messageID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, v1.Logger.ErrorResponse(codes.Unknown, err, responses.UnknownError)
	}
	msg, err := v1.DB.AddMessage(
		r.Location.ChannelId,
		r.Location.GuildId,
		ctx.UserID,
		messageID,
		r.Content,
		r.Attachments,
		v1.ProtoToEmbeds(r.Embeds),
		v1.ProtoToActions(r.Actions),
		v1.ProtoToOverrides(r.Overrides),
		sql.NullInt64{
			Int64: int64(r.InReplyTo),
			Valid: r.InReplyTo != 0,
		},
	)
	if err != nil {
		return nil, err
	}
	message := corev1.Message{
		Location:    r.Location,
		AuthorId:    ctx.UserID,
		Content:     r.Content,
		Attachments: r.Attachments,
		Embeds:      r.Embeds,
		Actions:     r.Actions,
		Overrides:   r.Overrides,
		InReplyTo:   r.InReplyTo,
	}
	createdAt, _ := ptypes.TimestampProto(msg.CreatedAt.UTC())
	message.CreatedAt = createdAt
	message.Location.MessageId = msg.MessageID
	message.AuthorId = ctx.UserID
	v1.PubSub.Guild.Broadcast(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_SentMessage{
			SentMessage: &corev1.GuildEvent_MessageSent{
				Message: &message,
			},
		},
	})
	return &corev1.SendMessageResponse{
		MessageId: messageID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.NoLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/GetGuildList")
}

func (v1 *V1) GetGuildList(c context.Context, r *corev1.GetGuildListRequest) (*corev1.GetGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	data, err := v1.DB.GetGuildList(ctx.UserID)
	if err != nil {
		return nil, err
	}
	var out []*corev1.GetGuildListResponse_GuildListEntry
	for _, guildEntry := range data {
		out = append(out, &corev1.GetGuildListResponse_GuildListEntry{
			GuildId: guildEntry.GuildID,
			Host:    guildEntry.HomeServer,
		})
	}
	return &corev1.GetGuildListResponse{
		Guilds: out,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Local:      true,
		Location:   middleware.NoLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/AddGuildToGuildList")
}

func (v1 *V1) AddGuildToGuildList(c context.Context, r *corev1.AddGuildToGuildListRequest) (*corev1.AddGuildToGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.AddGuildToList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Homeserver.Broadcast(ctx.UserID, &corev1.HomeserverEvent{
		Event: &corev1.HomeserverEvent_GuildAddedToList_{
			GuildAddedToList: &corev1.HomeserverEvent_GuildAddedToList{
				GuildId:    r.GuildId,
				Homeserver: r.Homeserver,
			},
		},
	})
	return &corev1.AddGuildToGuildListResponse{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Local:      true,
		Location:   middleware.NoLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/RemoveGuildFromGuildList")
}

func (v1 *V1) RemoveGuildFromGuildList(c context.Context, r *corev1.RemoveGuildFromGuildListRequest) (*corev1.RemoveGuildFromGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.RemoveGuildFromList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Homeserver.Broadcast(ctx.UserID, &corev1.HomeserverEvent{
		Event: &corev1.HomeserverEvent_GuildRemovedFromList_{
			GuildRemovedFromList: &corev1.HomeserverEvent_GuildRemovedFromList{
				GuildId:    r.GuildId,
				Homeserver: r.Homeserver,
			},
		},
	})
	return &corev1.RemoveGuildFromGuildListResponse{}, nil
}

func (v1 *V1) StreamHomeserverEvents(r *corev1.StreamHomeserverEventsRequest, s corev1.CoreService_StreamHomeserverEventsServer) error {
	userID, err := middleware.AuthHandler(v1.DB, s.Context())
	if err != nil {
		return err
	}
	err = v1.DB.UserIsLocal(userID)
	if err != nil {
		return err
	}
	<-v1.PubSub.Homeserver.Subscribe(userID, s)
	return nil
}
