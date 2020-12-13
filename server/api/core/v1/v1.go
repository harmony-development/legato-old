package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	corev1 "github.com/harmony-development/legato/gen/core"
	"github.com/harmony-development/legato/server/api/core/v1/permissions"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/queries"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/responses"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// ErrNoPermissions : you're not authenticated to do this
	ErrNoPermissions = errors.New("No permissions")
	// ErrNotInGuild : you're not in the guild
	ErrNotInGuild = errors.New("Not in guild")
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB             db.IHarmonyDB
	Logger         logger.ILogger
	Sonyflake      *sonyflake.Sonyflake
	PubSub         SubscriptionManager
	Perms          *permissions.Manager
	Config         *config.Config
	StorageBackend backend.AttachmentBackend
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies
}

// ActionsToProto is a utility function
func (v1 *V1) ActionsToProto(msgs json.RawMessage) (ret []*corev1.Action) {
	json.Unmarshal([]byte(msgs), &ret)
	return
}

// ProtoToActions is a utility function
func (v1 *V1) ProtoToActions(msgs []*corev1.Action) (ret []byte) {
	ret, _ = json.Marshal(msgs)
	return
}

// EmbedsToProto is a utility function
func (v1 *V1) EmbedsToProto(embeds json.RawMessage) (ret []*corev1.Embed) {
	json.Unmarshal([]byte(embeds), &ret)
	return
}

// ProtoToEmbeds is a utility function
func (v1 *V1) ProtoToEmbeds(embeds []*corev1.Embed) (ret []byte) {
	ret, _ = json.Marshal(embeds)
	return
}

// OverridesToProto is a utility function
func (v1 *V1) OverridesToProto(overrides []byte) (ret *corev1.Override) {
	if len(overrides) == 0 {
		return
	}
	ret = new(corev1.Override)
	proto.Unmarshal(overrides, ret)
	return
}

// ProtoToOverrides is a utility function
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
		Auth: true,
	}, "/protocol.core.v1.CoreService/CreateGuild")
}

// CreateGuild implements the CreateGuild RPC
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
		Permission: "invites.manage.create",
	}, "/protocol.core.v1.CoreService/CreateInvite")
}

// CreateInvite implements the CreateInvite RPC
func (v1 *V1) CreateInvite(c context.Context, r *corev1.CreateInviteRequest) (*corev1.CreateInviteResponse, error) {
	inv := int32(-1)
	if r.PossibleUses != 0 {
		inv = r.PossibleUses
	}
	invite, err := v1.DB.CreateInvite(r.GuildId, inv, r.Name)
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
		Permission: "channels.manage.create",
	}, "/protocol.core.v1.CoreService/CreateChannel")
}

// CreateChannel implements the CreateChannel RPC
func (v1 *V1) CreateChannel(c context.Context, r *corev1.CreateChannelRequest) (*corev1.CreateChannelResponse, error) {
	channel, err := v1.DB.AddChannelToGuild(r.GuildId, r.ChannelName, r.PreviousId, r.NextId, r.IsCategory)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_CreatedChannel{
			CreatedChannel: &corev1.Event_ChannelCreated{
				GuildId:    r.GuildId,
				ChannelId:  channel.ChannelID,
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
		Auth:     true,
		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.core.v1.CoreService/GetGuild")
}

// GetGuild implements the GetGuild RPC
func (v1 *V1) GetGuild(c context.Context, r *corev1.GetGuildRequest) (*corev1.GetGuildResponse, error) {
	guild, err := v1.DB.GetGuildByID(r.GuildId)
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
		Permission: "invites.view",
	}, "/protocol.core.v1.CoreService/GetGuildInvites")
}

// GetGuildInvites implements the GetGuildInvites RPC
func (v1 *V1) GetGuildInvites(c context.Context, r *corev1.GetGuildInvitesRequest) (*corev1.GetGuildInvitesResponse, error) {
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

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},
		Auth:     true,
		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.core.v1.CoreService/GetGuildMembers")
}

// GetGuildMembers implements the GetGuildMembers RPC
func (v1 *V1) GetGuildMembers(c context.Context, r *corev1.GetGuildMembersRequest) (*corev1.GetGuildMembersResponse, error) {
	members, err := v1.DB.MembersInGuild(r.GuildId)
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
		WantsRoles: true,
	}, "/protocol.core.v1.CoreService/GetGuildChannels")
}

// GetGuildChannels implements the GetGuildChannels RPC
func (v1 *V1) GetGuildChannels(c context.Context, r *corev1.GetGuildChannelsRequest) (*corev1.GetGuildChannelsResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	chans, err := v1.DB.ChannelsForGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	ret := []*corev1.GetGuildChannelsResponse_Channel{}
	roles := ctx.UserRoles

	for _, channel := range chans {
		if ctx.IsOwner || v1.Perms.Check("messages.view", roles, r.GuildId, channel.ChannelID) {
			ret = append(ret, &corev1.GetGuildChannelsResponse_Channel{
				ChannelId:   channel.ChannelID,
				ChannelName: channel.ChannelName,
				IsCategory:  channel.Category,
				IsVoice:     channel.Isvoice,
			})
		}
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
		Permission: "messages.view",
	}, "/protocol.core.v1.CoreService/GetMessage")
}

// GetMessage implements the GetMessage RPC
func (v1 *V1) GetMessage(c context.Context, r *corev1.GetMessageRequest) (*corev1.GetMessageResponse, error) {
	message, err := v1.DB.GetMessage(r.MessageId)
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
	attachments := []*corev1.Attachment{}
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
	for _, a := range message.Attachments {
		contentType, fileName, size, err := v1.StorageBackend.GetMetadata(a)
		if err == nil {
			attachments = append(attachments, &corev1.Attachment{
				Id:   a,
				Name: fileName,
				Type: contentType,
				Size: size,
			})
		}
	}
	return &corev1.GetMessageResponse{
		Message: &corev1.Message{
			GuildId:     message.GuildID,
			ChannelId:   message.ChannelID,
			MessageId:   message.MessageID,
			AuthorId:    message.UserID,
			CreatedAt:   createdAt,
			EditedAt:    editedAt,
			Content:     message.Content,
			Embeds:      embeds,
			Attachments: attachments,
			Actions:     actions,
			Overrides:   override,
			InReplyTo:   uint64(message.ReplyToID.Int64),
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
		Permission: "messages.view",
	}, "/protocol.core.v1.CoreService/GetChannelMessages")
}

// GetChannelMessages implements the GetChannelMessages RPC
func (v1 *V1) GetChannelMessages(c context.Context, r *corev1.GetChannelMessagesRequest) (*corev1.GetChannelMessagesResponse, error) {
	var err error
	var messages []queries.Message
	if r.BeforeMessage != 0 {
		time, err := v1.DB.GetMessageDate(r.BeforeMessage)
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
		ReachedTop: len(messages) < v1.Config.Server.GetMessageCount,
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
				attachments := []*corev1.Attachment{}
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
				for _, a := range message.Attachments {
					contentType, fileName, size, err := v1.StorageBackend.GetMetadata(a)
					if err == nil {
						attachments = append(attachments, &corev1.Attachment{
							Id:   a,
							Name: fileName,
							Type: contentType,
							Size: size,
						})
					}
				}
				ret = append(ret, &corev1.Message{
					GuildId:     message.GuildID,
					ChannelId:   message.ChannelID,
					MessageId:   message.MessageID,
					AuthorId:    message.UserID,
					CreatedAt:   createdAt,
					EditedAt:    editedAt,
					Content:     message.Content,
					Attachments: attachments,
					Embeds:      embeds,
					Actions:     actions,
					Overrides:   overrides,
					InReplyTo:   uint64(message.ReplyToID.Int64),
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
		Permission: "guild.manage.change-name",
	}, "/protocol.core.v1.CoreService/UpdateGuildName")
}

// UpdateGuildName implements the UpdateGuildName RPC
func (v1 *V1) UpdateGuildName(c context.Context, r *corev1.UpdateGuildNameRequest) (*empty.Empty, error) {
	if err := v1.DB.UpdateGuildName(r.GuildId, r.NewGuildName); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_EditedGuild{
			EditedGuild: &corev1.Event_GuildUpdated{
				GuildId:    r.GuildId,
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
		Permission: "channels.manage.change-name",
	}, "/protocol.core.v1.CoreService/UpdateChannelName")
}

// UpdateChannelName implements the UpdateChannelName RPC
func (v1 *V1) UpdateChannelName(c context.Context, r *corev1.UpdateChannelNameRequest) (*empty.Empty, error) {
	if err := v1.DB.SetChannelName(r.GuildId, r.ChannelId, r.NewChannelName); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_EditedChannel{
			EditedChannel: &corev1.Event_ChannelUpdated{
				GuildId:    r.GuildId,
				ChannelId:  r.ChannelId,
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
		Permission: "channels.manage.move",
	}, "/protocol.core.v1.CoreService/UpdateChannelOrder")
}

// UpdateChannelOrder implements the UpdateChannelOrder RPC
func (v1 *V1) UpdateChannelOrder(c context.Context, r *corev1.UpdateChannelOrderRequest) (*empty.Empty, error) {
	if err := v1.DB.MoveChannel(r.GuildId, r.ChannelId, r.PreviousId, r.NextId); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_EditedChannel{
			EditedChannel: &corev1.Event_ChannelUpdated{
				GuildId:     r.GuildId,
				ChannelId:   r.ChannelId,
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
		Permission: "messages.send",
	}, "/protocol.core.v1.CoreService/UpdateMessage")
}

// UpdateMessage implements the UpdateMessage RPC
func (v1 *V1) UpdateMessage(c context.Context, r *corev1.UpdateMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if !r.UpdateActions && !r.UpdateEmbeds && !r.UpdateContent && !r.UpdateOverrides {
		return nil, status.Error(codes.InvalidArgument, responses.InvalidRequest)
	}

	owner, err := v1.DB.GetMessageOwner(r.MessageId)
	if err != nil {
		return nil, err
	}
	if owner != ctx.UserID {
		return nil, ErrNoPermissions
	}

	var actions *[]byte
	var embeds *[]byte
	var overrides *[]byte
	var attachments *[]string
	attachmentsData := []*corev1.Attachment{}
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
	if r.UpdateAttachments {
		attachments = &r.Attachments

		for _, a := range r.Attachments {
			contentType, fileName, size, err := v1.StorageBackend.GetMetadata(a)
			if err == nil {
				attachmentsData = append(attachmentsData, &corev1.Attachment{
					Id:   a,
					Name: fileName,
					Type: contentType,
					Size: size,
				})
			}
		}
	}
	tiempo, err := v1.DB.UpdateMessage(r.MessageId, &r.Content, embeds, actions, overrides, attachments)
	if err != nil {
		return nil, err
	}
	editedAt, _ := ptypes.TimestampProto(tiempo.UTC())
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_EditedMessage{
			EditedMessage: &corev1.Event_MessageUpdated{
				GuildId:       r.GuildId,
				ChannelId:     r.ChannelId,
				MessageId:     r.MessageId,
				Content:       r.Content,
				UpdateContent: r.UpdateContent,
				Embeds:        r.Embeds,
				UpdateEmbeds:  r.UpdateEmbeds,
				Actions:       r.Actions,
				UpdateActions: r.UpdateActions,
				Overrides:     r.Overrides,
				EditedAt:      editedAt,
				Attachments:   attachmentsData,
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
		Permission: "guild.manage.delete",
	}, "/protocol.core.v1.CoreService/DeleteGuild")
}

// DeleteGuild implements the DeleteGuild RPC
func (v1 *V1) DeleteGuild(c context.Context, r *corev1.DeleteGuildRequest) (*empty.Empty, error) {
	err := v1.DB.DeleteGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_DeletedGuild{
			DeletedGuild: &corev1.Event_GuildDeleted{
				GuildId: r.GuildId,
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
		Permission: "invites.manage.delete",
	}, "/protocol.core.v1.CoreService/DeleteInvite")
}

// DeleteInvite implements the DeleteInvite RPC
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
		Permission: "channels.manage.delete",
	}, "/protocol.core.v1.CoreService/DeleteChannel")
}

// DeleteChannel implements the DeleteChannel RPC
func (v1 *V1) DeleteChannel(c context.Context, r *corev1.DeleteChannelRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteChannelFromGuild(r.GuildId, r.ChannelId); err != nil {
		return nil, err
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_DeletedChannel{
			DeletedChannel: &corev1.Event_ChannelDeleted{
				GuildId:   r.GuildId,
				ChannelId: r.ChannelId,
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
		WantsRoles: true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation,
	}, "/protocol.core.v1.CoreService/DeleteMessage")
}

// DeleteMessage implements the DeleteMessage RPC
func (v1 *V1) DeleteMessage(c context.Context, r *corev1.DeleteMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	owner, err := v1.DB.GetMessageOwner(r.MessageId)
	if err != nil {
		return nil, err
	}
	if ctx.UserID != owner && !(ctx.IsOwner || v1.Perms.Check("messages.manage.delete", ctx.UserRoles, r.GuildId, r.ChannelId)) {
		return nil, ErrNoPermissions
	}
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_DeletedMessage{
			DeletedMessage: &corev1.Event_MessageDeleted{
				GuildId:   r.GuildId,
				ChannelId: r.ChannelId,
				MessageId: r.MessageId,
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
		Auth:     true,
		Location: middleware.NoLocation,
	}, "/protocol.core.v1.CoreService/JoinGuild")
}

// JoinGuild implements the JoinGuild RPC
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
	v1.PubSub.Guild.Broadcast(guildID, &corev1.Event{
		Event: &corev1.Event_JoinedMember{
			JoinedMember: &corev1.Event_MemberJoined{
				GuildId:  guildID,
				MemberId: ctx.UserID,
			},
		},
	})
	return &corev1.JoinGuildResponse{
		GuildId: guildID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Auth:     true,
		Location: middleware.GuildLocation,
	}, "/protocol.core.v1.CoreService/LeaveGuild")
}

// LeaveGuild implements the LeaveGuild RPC
func (v1 *V1) LeaveGuild(c context.Context, r *corev1.LeaveGuildRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if isOwner, err := v1.DB.IsOwner(r.GuildId, ctx.UserID); err != nil {
		return nil, err
	} else if isOwner {
		return nil, status.Error(codes.FailedPrecondition, responses.InvalidRequest)
	}
	v1.PubSub.Guild.UnsubscribeUserFromGuild(r.GuildId, ctx.UserID)
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_LeftMember{
			LeftMember: &corev1.Event_MemberLeft{
				MemberId: ctx.UserID,
				GuildId:  r.GuildId,
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
		Auth:     true,
		Location: middleware.GuildLocation,
	}, "/protocol.core.v1.CoreService/StreamEvents")
}

// StreamEvents implements the StreamEvents RPC
func (v1 *V1) StreamEvents(s corev1.CoreService_StreamEventsServer) error {
	userID, err := middleware.AuthHandler(v1.DB, s.Context())
	if err != nil {
		return err
	}
	for {
		in, err := s.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch x := in.Request.(type) {
		case *corev1.StreamEventsRequest_SubscribeToGuild_:
			if err := middleware.LocationHandler(v1.DB, x.SubscribeToGuild, "/protocol.core.v1.CoreService/StreamGuildEvents", userID); err != nil {
				fmt.Println(err)
				break
			}
			ok, err := v1.DB.UserInGuild(userID, x.SubscribeToGuild.GuildId)
			if err != nil {
				fmt.Println(err)
				break
			}
			if !ok {
				fmt.Println("user not in guild")
				break
			}
			v1.PubSub.Guild.Subscribe(x.SubscribeToGuild.GuildId, userID, s)
		case *corev1.StreamEventsRequest_SubscribeToActions_:
			v1.PubSub.Actions.Subscribe(userID, s)
		case *corev1.StreamEventsRequest_SubscribeToHomeserverEvents_:
			err = v1.DB.UserIsLocal(userID)
			if err != nil {
				break
			}
			v1.PubSub.Homeserver.Subscribe(userID, s)
		}
	}
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    20,
		},
		Auth:       true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation,
		Permission: "actions.trigger",
	}, "/protocol.core.v1.CoreService/TriggerAction")
}

func (v1 *V1) TriggerAction(c context.Context, r *corev1.TriggerActionRequest) (*emptypb.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	msg, err := v1.DB.GetMessage(r.MessageId)
	if err != nil {
		return nil, err
	}
	if msg.ChannelID != r.ChannelId || msg.GuildID != r.GuildId {
		return nil, status.Error(codes.InvalidArgument, responses.InvalidRequest)
	}
	for _, action := range v1.ActionsToProto(msg.Actions) {
		if action.Id == r.ActionId {
			v1.PubSub.Actions.Broadcast(ctx.UserID, &corev1.Event{
				Event: &corev1.Event_ActionPerformed_{
					ActionPerformed: &corev1.Event_ActionPerformed{
						GuildId:    r.GuildId,
						ChannelId:  r.ChannelId,
						MessageId:  r.MessageId,
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
		Permission: "messages.send",
	}, "/protocol.core.v1.CoreService/SendMessage")
}

// SendMessage implements the SendMessage RPC
func (v1 *V1) SendMessage(c context.Context, r *corev1.SendMessageRequest) (*corev1.SendMessageResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	messageID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, v1.Logger.ErrorResponse(codes.Unknown, err, responses.UnknownError)
	}
	msg, err := v1.DB.AddMessage(
		r.ChannelId,
		r.GuildId,
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

	attachments := []*corev1.Attachment{}

	for _, a := range r.Attachments {
		contentType, fileName, size, err := v1.StorageBackend.GetMetadata(a)
		if err == nil {
			attachments = append(attachments, &corev1.Attachment{
				Id:   a,
				Name: fileName,
				Type: contentType,
				Size: size,
			})
		}
	}

	message := corev1.Message{
		GuildId:     r.GuildId,
		ChannelId:   r.ChannelId,
		MessageId:   messageID,
		AuthorId:    ctx.UserID,
		Content:     r.Content,
		Attachments: attachments,
		Embeds:      r.Embeds,
		Actions:     r.Actions,
		Overrides:   r.Overrides,
		InReplyTo:   r.InReplyTo,
	}
	createdAt, _ := ptypes.TimestampProto(msg.CreatedAt.UTC())
	message.CreatedAt = createdAt
	message.AuthorId = ctx.UserID
	v1.PubSub.Guild.Broadcast(r.GuildId, &corev1.Event{
		Event: &corev1.Event_SentMessage{
			SentMessage: &corev1.Event_MessageSent{
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
		Auth:     true,
		Location: middleware.NoLocation,
	}, "/protocol.core.v1.CoreService/GetGuildList")
}

// GetGuildList implements the GetGuildList RPC
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
		Auth:     true,
		Local:    true,
		Location: middleware.NoLocation,
	}, "/protocol.core.v1.CoreService/AddGuildToGuildList")
}

// AddGuildToGuildList implements the AddGuildToGuildList RPC
func (v1 *V1) AddGuildToGuildList(c context.Context, r *corev1.AddGuildToGuildListRequest) (*corev1.AddGuildToGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.AddGuildToList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Homeserver.Broadcast(ctx.UserID, &corev1.Event{
		Event: &corev1.Event_GuildAddedToList_{
			GuildAddedToList: &corev1.Event_GuildAddedToList{
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
		Auth:     true,
		Local:    true,
		Location: middleware.NoLocation,
	}, "/protocol.core.v1.CoreService/RemoveGuildFromGuildList")
}

// RemoveGuildFromGuildList implements the RemoveGuildFromGuildList RPC
func (v1 *V1) RemoveGuildFromGuildList(c context.Context, r *corev1.RemoveGuildFromGuildListRequest) (*corev1.RemoveGuildFromGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.RemoveGuildFromList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.PubSub.Homeserver.Broadcast(ctx.UserID, &corev1.Event{
		Event: &corev1.Event_GuildRemovedFromList_{
			GuildRemovedFromList: &corev1.Event_GuildRemovedFromList{
				GuildId:    r.GuildId,
				Homeserver: r.Homeserver,
			},
		},
	})
	return &corev1.RemoveGuildFromGuildListResponse{}, nil
}

// CreateEmotePack implements the CreateEmotePack RPC
func (v1 *V1) CreateEmotePack(c context.Context, r *corev1.CreateEmotePackRequest) (*corev1.CreateEmotePackResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	packID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	if err := v1.DB.CreateEmotePack(ctx.UserID, packID, r.PackName); err != nil {
		return nil, err
	}

	return &corev1.CreateEmotePackResponse{
		PackId: packID,
	}, nil
}

// AddEmoteToPack implements the AddEmoteToPack RPC
func (v1 *V1) AddEmoteToPack(c context.Context, r *corev1.AddEmoteToPackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
	}
	if err := v1.DB.AddEmoteToPack(r.PackId, r.ImageId, r.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DeleteEmoteFromPack implements the DeleteEmoteFromPack RPC
func (v1 *V1) DeleteEmoteFromPack(c context.Context, r *corev1.DeleteEmoteFromPackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
	}
	if err := v1.DB.DeleteEmoteFromPack(r.PackId, r.ImageId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DeleteEmotePack implements the DeleteEmotePack RPC
func (v1 *V1) DeleteEmotePack(c context.Context, r *corev1.DeleteEmotePackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, status.Error(codes.PermissionDenied, responses.InsufficientPrivileges)
	}
	if err := v1.DB.DeleteEmotePack(r.PackId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DequipEmotePack implements the DequipEmotePack RPC
func (v1 *V1) DequipEmotePack(c context.Context, r *corev1.DequipEmotePackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if err := v1.DB.DequipEmotePack(ctx.UserID, r.PackId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// GetEmotePacks implements the GetEmotePacks RPC
func (v1 *V1) GetEmotePacks(c context.Context, r *corev1.GetEmotePacksRequest) (*corev1.GetEmotePacksResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	packs, err := v1.DB.GetEmotePacks(ctx.UserID)
	if err != nil {
		return nil, err
	}
	outPacks := []*corev1.GetEmotePacksResponse_EmotePack{}
	for _, pack := range packs {
		outPacks = append(outPacks, &corev1.GetEmotePacksResponse_EmotePack{
			PackId:    pack.PackID,
			PackOwner: pack.UserID,
			PackName:  pack.PackName,
		})
	}
	return &corev1.GetEmotePacksResponse{
		Packs: outPacks,
	}, nil
}

// GetEmotePackEmotes implements the GetEmotePackEmotes RPC
func (v1 *V1) GetEmotePackEmotes(c context.Context, r *corev1.GetEmotePackEmotesRequest) (*corev1.GetEmotePackEmotesResponse, error) {
	emotes, err := v1.DB.GetEmotePackEmotes(r.PackId)
	if err != nil {
		return nil, err
	}
	outEmotes := []*corev1.GetEmotePackEmotesResponse_Emote{}
	for _, emote := range emotes {
		outEmotes = append(outEmotes, &corev1.GetEmotePackEmotesResponse_Emote{
			ImageId: emote.ImageID,
			Name:    emote.EmoteName,
		})
	}
	return &corev1.GetEmotePackEmotesResponse{
		Emotes: outEmotes,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "roles.manage",
	}, "/protocol.core.v1.CoreService/AddGuildRole")
}

// AddGuildRole implements the AddGuildRole RPC
func (v1 *V1) AddGuildRole(c context.Context, r *corev1.AddGuildRoleRequest) (*corev1.AddGuildRoleResponse, error) {
	roleID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	r.Role.RoleId = roleID
	err = v1.DB.AddRoleToGuild(r.GuildId, r.Role)
	if err != nil {
		return nil, err
	}

	return &corev1.AddGuildRoleResponse{
		RoleId: roleID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "roles.manage",
	}, "/protocol.core.v1.CoreService/AddGuildRole")
}

// DeleteGuildRole implements the DeleteGuildRole RPC
func (v1 *V1) DeleteGuildRole(c context.Context, r *corev1.DeleteGuildRoleRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.RemoveRoleFromGuild(r.GuildId, r.RoleId)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "roles.manage",
	}, "/protocol.core.v1.CoreService/MoveRole")
}

// MoveRole implements the MoveRole RPC
func (v1 *V1) MoveRole(c context.Context, r *corev1.MoveRoleRequest) (*corev1.MoveRoleResponse, error) {
	err := v1.DB.MoveRole(r.GuildId, r.RoleId, r.BeforeId, r.AfterId)
	return &corev1.MoveRoleResponse{}, err
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "roles.get",
	}, "/protocol.core.v1.CoreService/GetGuildRoles")
}

// GetGuildRoles implements the GetGuildRoles RPC
func (v1 *V1) GetGuildRoles(c context.Context, r *corev1.GetGuildRolesRequest) (*corev1.GetGuildRolesResponse, error) {
	roles, err := v1.DB.GetGuildRoles(r.GuildId)
	return &corev1.GetGuildRolesResponse{
		Roles: roles,
	}, err
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "permissions.manage.set",
	}, "/protocol.core.v1.CoreService/SetPermissions")
}

// SetPermissions implements the SetPermissions RPC
func (v1 *V1) SetPermissions(c context.Context, r *corev1.SetPermissionsRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, v1.Perms.SetPermissions(r.Perms.Permissions, r.GuildId, r.ChannelId, r.RoleId)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "permissions.manage.get",
	}, "/protocol.core.v1.CoreService/GetPermissions")
}

// GetPermissions implements the GetPermissions RPC
func (v1 *V1) GetPermissions(c context.Context, r *corev1.GetPermissionsRequest) (*corev1.GetPermissionsResponse, error) {
	return &corev1.GetPermissionsResponse{Perms: &corev1.PermissionList{Permissions: v1.Perms.GetPermissions(r.GuildId, r.ChannelId, r.RoleId)}}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		WantsRoles: true,
		Location:   middleware.GuildLocation,
	}, "/protocol.core.v1.CoreService/QueryHasPermission")
}

// QueryHasPermission implements the QueryHasPermission RPC
func (v1 *V1) QueryHasPermission(c context.Context, r *corev1.QueryPermissionsRequest) (*corev1.QueryPermissionsResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	if r.As == 0 {
		r.As = c.(middleware.HarmonyContext).UserID
	} else if !(ctx.IsOwner || v1.Perms.Check("permissions.query", ctx.UserRoles, r.GuildId, r.ChannelId)) {
		return nil, ErrNoPermissions
	}

	owner, err := v1.DB.GetOwner(r.GuildId)
	if err != nil {
		return nil, err
	}

	roles, err := v1.DB.RolesForUser(r.GuildId, r.As)
	if err != nil {
		return nil, err
	}
	return &corev1.QueryPermissionsResponse{
		Ok: owner == r.As || v1.Perms.Check(r.CheckFor, roles, r.GuildId, r.ChannelId),
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Location:   middleware.GuildLocation,
		Permission: "roles.users.manage",
	}, "/protocol.core.v1.CoreService/ManageUserRoles")
}

func (v1 *V1) ManageUserRoles(c context.Context, r *corev1.ManageUserRolesRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.ManageRoles(r.GuildId, r.UserId, r.GiveRoleIds, r.TakeRoleIds)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		Permission: "roles.manage",
		Location:   middleware.GuildLocation,
	}, "/protocol.core.v1.CoreService/ModifyGuildRole")
}

func (v1 *V1) ModifyGuildRole(c context.Context, r *corev1.ModifyGuildRoleRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.ModifyRole(r.GuildId, r.Role.RoleId, r.Role.Name, r.Role.Color, r.Role.Hoist, r.Role.Pingable, r.ModifyName, r.ModifyColor, r.ModifyHoist, r.ModifyPingable)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},
		Auth:       true,
		WantsRoles: true,
		Location:   middleware.GuildLocation,
	}, "/protocol.core.v1.CoreService/GetUserRoles")
}

func (v1 *V1) GetUserRoles(c context.Context, r *corev1.GetUserRolesRequest) (*corev1.GetUserRolesResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	if r.UserId == 0 {
		return &corev1.GetUserRolesResponse{
			Roles: ctx.UserRoles,
		}, nil
	}

	if !(ctx.IsOwner || v1.Perms.Check("roles.users.get", ctx.UserRoles, r.GuildId, 0)) {
		return nil, ErrNoPermissions
	}

	roles, err := v1.DB.RolesForUser(r.GuildId, r.UserId)
	if err != nil {
		return nil, err
	}

	return &corev1.GetUserRolesResponse{
		Roles: roles,
	}, nil
}
