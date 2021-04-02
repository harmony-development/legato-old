package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/responses"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB             types.IHarmonyDB
	Logger         logger.ILogger
	Sonyflake      *sonyflake.Sonyflake
	Streams        StreamManager
	Perms          *permissions.Manager
	Config         *config.Config
	StorageBackend backend.AttachmentBackend
	Middleware     *middleware.Middlewares
}

// V1 contains the gRPC handler for v1
type V1 struct {
	Dependencies
}

// ActionsToProto is a utility function
func (v1 *V1) ActionsToProto(msgs json.RawMessage) (ret []*harmonytypesv1.Action) {
	if err := json.Unmarshal([]byte(msgs), &ret); err != nil {
		panic(err)
	}
	return
}

// ProtoToActions is a utility function
func (v1 *V1) ProtoToActions(msgs []*harmonytypesv1.Action) (ret []byte) {
	ret, _ = json.Marshal(msgs)
	return
}

// EmbedsToProto is a utility function
func (v1 *V1) EmbedsToProto(embeds json.RawMessage) (ret []*harmonytypesv1.Embed) {
	if err := json.Unmarshal([]byte(embeds), &ret); err != nil {
		panic(err)
	}
	return
}

// ProtoToEmbeds is a utility function
func (v1 *V1) ProtoToEmbeds(embeds []*harmonytypesv1.Embed) (ret []byte) {
	ret, _ = json.Marshal(embeds)
	return
}

// OverridesToProto is a utility function
func (v1 *V1) OverridesToProto(overrides []byte) (ret *harmonytypesv1.Override) {
	if len(overrides) == 0 {
		return
	}
	ret = new(harmonytypesv1.Override)
	err := proto.Unmarshal(overrides, ret)
	if err != nil {
		panic(err)
	}
	return
}

// ProtoToOverrides is a utility function
func (v1 *V1) ProtoToOverrides(overrides *harmonytypesv1.Override) (ret []byte) {
	ret, _ = proto.Marshal(overrides)
	return
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    1,
		},
	}, "/protocol.chat.v1.ChatService/CreateGuild")
}

// CreateGuild implements the CreateGuild RPC
func (v1 *V1) CreateGuild(c echo.Context, r *chatv1.CreateGuildRequest) (*chatv1.CreateGuildResponse, error) {
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
	return &chatv1.CreateGuildResponse{
		GuildId: guild.GuildID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/CreateInvite")
}

// CreateInvite implements the CreateInvite RPC
func (v1 *V1) CreateInvite(c echo.Context, r *chatv1.CreateInviteRequest) (*chatv1.CreateInviteResponse, error) {
	inv := int32(-1)
	if r.PossibleUses != 0 {
		inv = r.PossibleUses
	}
	invite, err := v1.DB.CreateInvite(r.GuildId, inv, r.Name)
	if err != nil {
		return nil, err
	}
	return &chatv1.CreateInviteResponse{
		Name: invite.InviteID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/CreateChannel")
}

// CreateChannel implements the CreateChannel RPC
func (v1 *V1) CreateChannel(c echo.Context, r *chatv1.CreateChannelRequest) (*chatv1.CreateChannelResponse, error) {
	channel, err := v1.DB.AddChannelToGuild(r.GuildId, r.ChannelName, r.PreviousId, r.NextId, r.IsCategory, r.Metadata)
	if err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_CreatedChannel{
			CreatedChannel: &chatv1.Event_ChannelCreated{
				GuildId:    r.GuildId,
				ChannelId:  channel.ChannelID,
				Name:       r.ChannelName,
				PreviousId: r.PreviousId,
				NextId:     r.NextId,
				IsCategory: r.IsCategory,
			},
		},
	})
	return &chatv1.CreateChannelResponse{
		ChannelId: channel.ChannelID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/GetGuild")
}

// GetGuild implements the GetGuild RPC
func (v1 *V1) GetGuild(c echo.Context, r *chatv1.GetGuildRequest) (*chatv1.GetGuildResponse, error) {
	guild, err := v1.DB.GetGuildByID(r.GuildId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, responses.NewError(responses.BadGuildID)
		}
		return nil, err
	}
	return &chatv1.GetGuildResponse{
		GuildName:    guild.GuildName,
		GuildOwner:   guild.GuildOwner,
		GuildPicture: guild.GuildPicture,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/GetGuildInvites")
}

// GetGuildInvites implements the GetGuildInvites RPC
func (v1 *V1) GetGuildInvites(c echo.Context, r *chatv1.GetGuildInvitesRequest) (*chatv1.GetGuildInvitesResponse, error) {
	invites, err := v1.DB.GetInvites(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &chatv1.GetGuildInvitesResponse{
		Invites: func() (ret []*chatv1.GetGuildInvitesResponse_Invite) {
			for _, inv := range invites {
				ret = append(ret, &chatv1.GetGuildInvitesResponse_Invite{
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

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/GetGuildMembers")
}

// GetGuildMembers implements the GetGuildMembers RPC
func (v1 *V1) GetGuildMembers(c echo.Context, r *chatv1.GetGuildMembersRequest) (*chatv1.GetGuildMembersResponse, error) {
	members, err := v1.DB.MembersInGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	return &chatv1.GetGuildMembersResponse{
		Members: members,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    15,
		},

		Location:   middleware.GuildLocation | middleware.JoinedLocation,
		WantsRoles: true,
	}, "/protocol.chat.v1.ChatService/GetGuildChannels")
}

// GetGuildChannels implements the GetGuildChannels RPC
func (v1 *V1) GetGuildChannels(c echo.Context, r *chatv1.GetGuildChannelsRequest) (*chatv1.GetGuildChannelsResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	chans, err := v1.DB.ChannelsForGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	ret := []*chatv1.GetGuildChannelsResponse_Channel{}
	roles := ctx.UserRoles

	for _, channel := range chans {
		if ctx.IsOwner || v1.Perms.Check("messages.view", roles, r.GuildId, channel.ChannelId) {
			ret = append(ret, channel)
		}
	}
	return &chatv1.GetGuildChannelsResponse{
		Channels: ret,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/GetMessage")
}

// GetMessage implements the GetMessage RPC
func (v1 *V1) GetMessage(c echo.Context, r *chatv1.GetMessageRequest) (*chatv1.GetMessageResponse, error) {
	data, err := v1.DB.GetMessage(r.MessageId)
	if err != nil {
		return nil, err
	}
	return &chatv1.GetMessageResponse{
		Message: data,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},
	}, "/protocol.chat.v1.ChatService/PreviewGuild")
}

func (v1 *V1) PreviewGuild(c echo.Context, r *chatv1.PreviewGuildRequest) (*chatv1.PreviewGuildResponse, error) {
	data, err := v1.DB.ResolveGuildID(r.InviteId)
	if err != nil {
		return nil, err
	}

	ret := &chatv1.PreviewGuildResponse{}
	guildData, err := v1.DB.GetGuildByID(data)
	if err != nil {
		return nil, err
	}

	ret.Avatar = guildData.GuildPicture
	ret.Name = guildData.GuildName
	count := int64(0)
	count, err = v1.DB.CountMembersInGuild(data)
	if err != nil {
		return nil, err
	}

	ret.MemberCount = uint64(count)
	return ret, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/GetChannelMessages")
}

// GetChannelMessages implements the GetChannelMessages RPC
func (v1 *V1) GetChannelMessages(c echo.Context, r *chatv1.GetChannelMessagesRequest) (*chatv1.GetChannelMessagesResponse, error) {
	var err error
	var messages []*harmonytypesv1.Message
	if r.BeforeMessage != 0 {
		time, err := v1.DB.GetMessageDate(r.BeforeMessage)
		if err != nil {
			return nil, err
		}
		messages, err = v1.DB.GetMessagesBefore(r.GuildId, r.ChannelId, time)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		messages, err = v1.DB.GetMessages(r.GuildId, r.ChannelId)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
	}
	return &chatv1.GetChannelMessagesResponse{
		ReachedTop: len(messages) < v1.Config.Server.Policies.APIs.Messages.MaximumGetAmount,
		Messages:   messages,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    2,
		},

		Location: middleware.GuildLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/UpdateGuildInformation")
}

// UpdateGuildInformation implements the UpdateGuildInformation RPC
func (v1 *V1) UpdateGuildInformation(c echo.Context, r *chatv1.UpdateGuildInformationRequest) (*empty.Empty, error) {
	err := v1.DB.UpdateGuildInformation(r.GuildId, r.NewGuildName, r.NewGuildPicture, r.Metadata, r.UpdateGuildName, r.UpdateGuildPicture, r.UpdateMetadata)
	if err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_EditedGuild{
			EditedGuild: &chatv1.Event_GuildUpdated{
				GuildId:        r.GuildId,
				Name:           r.NewGuildName,
				UpdateName:     r.UpdateGuildName,
				Picture:        r.NewGuildPicture,
				UpdatePicture:  r.UpdateGuildPicture,
				Metadata:       r.Metadata,
				UpdateMetadata: r.UpdateMetadata,
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

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/UpdateChannelInformation")
}

// UpdateChannelInformation implements the UpdateChannelInformation RPC
func (v1 *V1) UpdateChannelInformation(c echo.Context, r *chatv1.UpdateChannelInformationRequest) (*empty.Empty, error) {
	if err := v1.DB.UpdateChannelInformation(r.GuildId, r.ChannelId, r.Name, r.UpdateName, r.Metadata, r.UpdateMetadata); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_EditedChannel{
			EditedChannel: &chatv1.Event_ChannelUpdated{
				GuildId:        r.GuildId,
				ChannelId:      r.ChannelId,
				Name:           r.Name,
				UpdateName:     r.UpdateName,
				Metadata:       r.Metadata,
				UpdateMetadata: r.UpdateMetadata,
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

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.JoinedLocation,
	}, "/protocol.chat.v1.ChatService/UpdateChannelOrder")
}

// UpdateChannelOrder implements the UpdateChannelOrder RPC
func (v1 *V1) UpdateChannelOrder(c echo.Context, r *chatv1.UpdateChannelOrderRequest) (*empty.Empty, error) {
	if err := v1.DB.MoveChannel(r.GuildId, r.ChannelId, r.PreviousId, r.NextId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_EditedChannel{
			EditedChannel: &chatv1.Event_ChannelUpdated{
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

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation | middleware.AuthorLocation,
	}, "/protocol.chat.v1.ChatService/UpdateMessage")
}

// UpdateMessage implements the UpdateMessage RPC
func (v1 *V1) UpdateMessageText(c echo.Context, r *chatv1.UpdateMessageTextRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	owner, err := v1.DB.GetMessageOwner(r.MessageId)
	if err != nil {
		return nil, err
	}
	if owner != ctx.UserID {
		return nil, responses.NewError(responses.NotOwner)
	}

	tenpo, err := v1.DB.UpdateTextMessage(r.MessageId, r.NewContent)
	editedAt, _ := ptypes.TimestampProto(tenpo.UTC())

	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_EditedMessage{
			EditedMessage: &chatv1.Event_MessageUpdated{
				GuildId:   r.GuildId,
				ChannelId: r.ChannelId,
				MessageId: r.MessageId,
				Content:   r.NewContent,
				EditedAt:  editedAt,
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

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/DeleteGuild")
}

// DeleteGuild implements the DeleteGuild RPC
func (v1 *V1) DeleteGuild(c echo.Context, r *chatv1.DeleteGuildRequest) (*empty.Empty, error) {
	err := v1.DB.DeleteGuild(r.GuildId)
	if err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_DeletedGuild{
			DeletedGuild: &chatv1.Event_GuildDeleted{
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

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/DeleteInvite")
}

// DeleteInvite implements the DeleteInvite RPC
func (v1 *V1) DeleteInvite(c echo.Context, r *chatv1.DeleteInviteRequest) (*empty.Empty, error) {
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

		Location: middleware.GuildLocation | middleware.ChannelLocation,
	}, "/protocol.chat.v1.ChatService/DeleteChannel")
}

// DeleteChannel implements the DeleteChannel RPC
func (v1 *V1) DeleteChannel(c echo.Context, r *chatv1.DeleteChannelRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteChannelFromGuild(r.GuildId, r.ChannelId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_DeletedChannel{
			DeletedChannel: &chatv1.Event_ChannelDeleted{
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

		WantsRoles: true,
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation,
	}, "/protocol.chat.v1.ChatService/DeleteMessage")
}

// DeleteMessage implements the DeleteMessage RPC
func (v1 *V1) DeleteMessage(c echo.Context, r *chatv1.DeleteMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	owner, err := v1.DB.GetMessageOwner(r.MessageId)
	if err != nil {
		return nil, err
	}
	if ctx.UserID != owner && !(ctx.IsOwner || v1.Perms.Check("messages.manage.delete", ctx.UserRoles, r.GuildId, r.ChannelId)) {
		return nil, responses.NewError(responses.NotEnoughPermissions)
	}
	if err := v1.DB.DeleteMessage(r.MessageId, r.ChannelId, r.GuildId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_DeletedMessage{
			DeletedMessage: &chatv1.Event_MessageDeleted{
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

		Location: middleware.NoLocation,
	}, "/protocol.chat.v1.ChatService/GetUserBulk")
}

func (v1 *V1) GetUserBulk(c echo.Context, r *chatv1.GetUserBulkRequest) (*chatv1.GetUserBulkResponse, error) {
	panic("unimplemented")
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},

		Location: middleware.NoLocation,
	}, "/protocol.chat.v1.ChatService/JoinGuild")
}

// JoinGuild implements the JoinGuild RPC
func (v1 *V1) JoinGuild(c echo.Context, r *chatv1.JoinGuildRequest) (*chatv1.JoinGuildResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	guildID, err := v1.DB.ResolveGuildID(r.InviteId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, responses.NewError(responses.BadGuildID)
		}
		return nil, err
	}
	if banned, err := v1.DB.IsBanned(guildID, ctx.UserID); banned || err != nil {
		if err != nil {
			return nil, responses.NewError(responses.InternalServerError)
		}
		return nil, responses.NewError(responses.BannedFromGuild)
	}
	if err := v1.DB.AddMemberToGuild(ctx.UserID, guildID); err != nil {
		return nil, err
	}
	if err := v1.DB.IncrementInvite(r.InviteId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(guildID, &chatv1.Event{
		Event: &chatv1.Event_JoinedMember{
			JoinedMember: &chatv1.Event_MemberJoined{
				GuildId:  guildID,
				MemberId: ctx.UserID,
			},
		},
	})
	return &chatv1.JoinGuildResponse{
		GuildId: guildID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/LeaveGuild")
}

// LeaveGuild implements the LeaveGuild RPC
func (v1 *V1) LeaveGuild(c echo.Context, r *chatv1.LeaveGuildRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if isOwner, err := v1.DB.IsOwner(r.GuildId, ctx.UserID); err != nil {
		return nil, err
	} else if isOwner {
		return nil, responses.NewError(responses.IsOwner)
	}

	if err := v1.DB.DeleteMember(r.GuildId, ctx.UserID); err != nil {
		return nil, err
	}

	if err := v1.DB.RemoveGuildFromList(ctx.UserID, r.GuildId, ""); err != nil {
		return nil, err
	}

	v1.Streams.RemoveGuildSubscription(r.GuildId, ctx.UserID)
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_LeftMember{
			LeftMember: &chatv1.Event_MemberLeft{
				MemberId: ctx.UserID,
				GuildId:  r.GuildId,
			},
		},
	})

	v1.Streams.BroadcastHomeserver(ctx.UserID, &chatv1.Event{
		Event: &chatv1.Event_GuildRemovedFromList_{
			GuildRemovedFromList: &chatv1.Event_GuildRemovedFromList{
				GuildId:    r.GuildId,
				Homeserver: "",
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

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/StreamEvents")
}

// StreamEvents implements the StreamEvents RPC
func (v1 *V1) StreamEvents(c echo.Context, in chan *chatv1.StreamEventsRequest, out chan *chatv1.Event) {
	userID, err := v1.Middleware.AuthHandler(c)
	if err != nil {
		close(in)
		v1.Logger.Exception(err)
		return
	}
	done := make(chan struct{})
	v1.Streams.RegisterClient(userID, out, done)
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for {
			dat, ok := <-in
			if !ok {
				return
			}

			switch x := dat.Request.(type) {
			case *chatv1.StreamEventsRequest_SubscribeToGuild_:
				ok, err := v1.DB.UserInGuild(userID, x.SubscribeToGuild.GuildId)
				if err != nil {
					fmt.Println(err)
					break
				}
				if !ok {
					fmt.Println("user not in guild")
					break
				}
				v1.Streams.AddGuildSubscription(out, x.SubscribeToGuild.GuildId)
			case *chatv1.StreamEventsRequest_SubscribeToActions_:
				v1.Streams.AddActionSubscription(out)
			case *chatv1.StreamEventsRequest_SubscribeToHomeserverEvents_:
				err = v1.DB.UserIsLocal(userID)
				if err != nil {
					continue
				}
				v1.Streams.AddHomeserverSubscription(out)
			}
		}
	}()
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    20,
		},

		Location: middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation,
	}, "/protocol.chat.v1.ChatService/TriggerAction")
}

func (v1 *V1) TriggerAction(c echo.Context, r *chatv1.TriggerActionRequest) (*emptypb.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	msg, err := v1.DB.GetMessage(r.MessageId)
	if err != nil {
		return nil, err
	}
	if msg.ChannelId != r.ChannelId || msg.GuildId != r.GuildId {
		return nil, responses.NewError(responses.BadAction)
	}
	for _, action := range msg.Content.Actions {
		if action.Id == r.ActionId {
			v1.Streams.BroadcastAction(ctx.UserID, &chatv1.Event{
				Event: &chatv1.Event_ActionPerformed_{
					ActionPerformed: &chatv1.Event_ActionPerformed{
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
	return nil, responses.NewError(responses.BadAction)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 10 * time.Second,
			Burst:    50,
		},

		Location: middleware.GuildLocation | middleware.ChannelLocation,
	}, "/protocol.chat.v1.ChatService/SendMessage")
}

// SendMessage implements the SendMessage RPC
func (v1 *V1) SendMessage(c echo.Context, r *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	messageID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	replyTo := sql.NullInt64{
		Int64: int64(r.InReplyTo),
		Valid: r.InReplyTo != 0,
	}

	var (
		t time.Time
	)

	switch e := r.Content.Content.(type) {
	case *harmonytypesv1.Content_EmbedMessage:
		t, err = v1.DB.AddEmbedMessage(r.GuildId, r.ChannelId, messageID, ctx.UserID, r.Content.Actions, r.Overrides, replyTo, r.Metadata, e.EmbedMessage.Embeds)
	case *harmonytypesv1.Content_FilesMessage:
		t, err = v1.DB.AddFilesMessage(r.GuildId, r.ChannelId, messageID, ctx.UserID, r.Content.Actions, r.Overrides, replyTo, r.Metadata, e.FilesMessage.Attachments)
	case *harmonytypesv1.Content_TextMessage:
		t, err = v1.DB.AddTextMessage(r.GuildId, r.ChannelId, messageID, ctx.UserID, r.Content.Actions, r.Overrides, replyTo, r.Metadata, e.TextMessage.Content)
	}

	if err != nil {
		return nil, err
	}

	createdAt, _ := ptypes.TimestampProto(t.UTC())
	message := harmonytypesv1.Message{
		GuildId:   r.GuildId,
		ChannelId: r.ChannelId,
		MessageId: messageID,
		AuthorId:  ctx.UserID,
		Content:   r.Content,
		Overrides: r.Overrides,
		InReplyTo: r.InReplyTo,
		Metadata:  r.Metadata,
		CreatedAt: createdAt,
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_SentMessage{
			SentMessage: &chatv1.Event_MessageSent{
				EchoId:  r.EchoId,
				Message: &message,
			},
		},
	})
	return &chatv1.SendMessageResponse{
		MessageId: messageID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.NoLocation,
	}, "/protocol.chat.v1.ChatService/GetGuildList")
}

// GetGuildList implements the GetGuildList RPC
func (v1 *V1) GetGuildList(c echo.Context, r *chatv1.GetGuildListRequest) (*chatv1.GetGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	data, err := v1.DB.GetGuildList(ctx.UserID)
	if err != nil {
		return nil, err
	}
	var out []*chatv1.GetGuildListResponse_GuildListEntry
	for _, guildEntry := range data {
		out = append(out, &chatv1.GetGuildListResponse_GuildListEntry{
			GuildId: guildEntry.GuildID,
			Host:    guildEntry.HomeServer,
		})
	}
	return &chatv1.GetGuildListResponse{
		Guilds: out,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.NoLocation,
	}, "/protocol.chat.v1.ChatService/AddGuildToGuildList")
}

// AddGuildToGuildList implements the AddGuildToGuildList RPC
func (v1 *V1) AddGuildToGuildList(c echo.Context, r *chatv1.AddGuildToGuildListRequest) (*chatv1.AddGuildToGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.AddGuildToList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.Streams.BroadcastHomeserver(ctx.UserID, &chatv1.Event{
		Event: &chatv1.Event_GuildAddedToList_{
			GuildAddedToList: &chatv1.Event_GuildAddedToList{
				GuildId:    r.GuildId,
				Homeserver: r.Homeserver,
			},
		},
	})
	return &chatv1.AddGuildToGuildListResponse{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.NoLocation,
	}, "/protocol.chat.v1.ChatService/RemoveGuildFromGuildList")
}

// RemoveGuildFromGuildList implements the RemoveGuildFromGuildList RPC
func (v1 *V1) RemoveGuildFromGuildList(c echo.Context, r *chatv1.RemoveGuildFromGuildListRequest) (*chatv1.RemoveGuildFromGuildListResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	err := v1.DB.RemoveGuildFromList(ctx.UserID, r.GuildId, r.Homeserver)
	if err != nil {
		return nil, err
	}
	v1.Streams.BroadcastHomeserver(ctx.UserID, &chatv1.Event{
		Event: &chatv1.Event_GuildRemovedFromList_{
			GuildRemovedFromList: &chatv1.Event_GuildRemovedFromList{
				GuildId:    r.GuildId,
				Homeserver: r.Homeserver,
			},
		},
	})
	return &chatv1.RemoveGuildFromGuildListResponse{}, nil
}

// CreateEmotePack implements the CreateEmotePack RPC
func (v1 *V1) CreateEmotePack(c echo.Context, r *chatv1.CreateEmotePackRequest) (*chatv1.CreateEmotePackResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	packID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	if err := v1.DB.CreateEmotePack(ctx.UserID, packID, r.PackName); err != nil {
		return nil, err
	}

	return &chatv1.CreateEmotePackResponse{
		PackId: packID,
	}, nil
}

// AddEmoteToPack implements the AddEmoteToPack RPC
func (v1 *V1) AddEmoteToPack(c echo.Context, r *chatv1.AddEmoteToPackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, responses.NewError(responses.NotOwner)
	}
	if err := v1.DB.AddEmoteToPack(r.PackId, r.ImageId, r.Name); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DeleteEmoteFromPack implements the DeleteEmoteFromPack RPC
func (v1 *V1) DeleteEmoteFromPack(c echo.Context, r *chatv1.DeleteEmoteFromPackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, responses.NewError(responses.NotOwner)
	}
	if err := v1.DB.DeleteEmoteFromPack(r.PackId, r.ImageId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DeleteEmotePack implements the DeleteEmotePack RPC
func (v1 *V1) DeleteEmotePack(c echo.Context, r *chatv1.DeleteEmotePackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if isOwner, err := v1.DB.IsPackOwner(ctx.UserID, r.PackId); err != nil {
		return nil, err
	} else if !isOwner {
		return nil, responses.NewError(responses.NotOwner)
	}
	if err := v1.DB.DeleteEmotePack(r.PackId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// DequipEmotePack implements the DequipEmotePack RPC
func (v1 *V1) DequipEmotePack(c echo.Context, r *chatv1.DequipEmotePackRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)

	if err := v1.DB.DequipEmotePack(ctx.UserID, r.PackId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// GetEmotePacks implements the GetEmotePacks RPC
func (v1 *V1) GetEmotePacks(c echo.Context, r *chatv1.GetEmotePacksRequest) (*chatv1.GetEmotePacksResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	packs, err := v1.DB.GetEmotePacks(ctx.UserID)
	if err != nil {
		return nil, err
	}
	outPacks := []*chatv1.GetEmotePacksResponse_EmotePack{}
	for _, pack := range packs {
		outPacks = append(outPacks, &chatv1.GetEmotePacksResponse_EmotePack{
			PackId:    pack.ID,
			PackOwner: pack.Edges.Owner.ID,
			PackName:  pack.Name,
		})
	}
	return &chatv1.GetEmotePacksResponse{
		Packs: outPacks,
	}, nil
}

// GetEmotePackEmotes implements the GetEmotePackEmotes RPC
func (v1 *V1) GetEmotePackEmotes(c echo.Context, r *chatv1.GetEmotePackEmotesRequest) (*chatv1.GetEmotePackEmotesResponse, error) {
	emotes, err := v1.DB.GetEmotePackEmotes(r.PackId)
	if err != nil {
		return nil, err
	}
	outEmotes := []*chatv1.GetEmotePackEmotesResponse_Emote{}
	for _, emote := range emotes {
		outEmotes = append(outEmotes, &chatv1.GetEmotePackEmotesResponse_Emote{
			ImageId: emote.ID,
			Name:    emote.Name,
		})
	}
	return &chatv1.GetEmotePackEmotesResponse{
		Emotes: outEmotes,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/AddGuildRole")
}

// AddGuildRole implements the AddGuildRole RPC
func (v1 *V1) AddGuildRole(c echo.Context, r *chatv1.AddGuildRoleRequest) (*chatv1.AddGuildRoleResponse, error) {
	roleID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	r.Role.RoleId = roleID
	err = v1.DB.AddRoleToGuild(r.GuildId, r.Role)
	if err != nil {
		return nil, err
	}

	return &chatv1.AddGuildRoleResponse{
		RoleId: roleID,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/AddGuildRole")
}

// DeleteGuildRole implements the DeleteGuildRole RPC
func (v1 *V1) DeleteGuildRole(c echo.Context, r *chatv1.DeleteGuildRoleRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.RemoveRoleFromGuild(r.GuildId, r.RoleId)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/MoveRole")
}

// MoveRole implements the MoveRole RPC
func (v1 *V1) MoveRole(c echo.Context, r *chatv1.MoveRoleRequest) (*chatv1.MoveRoleResponse, error) {
	err := v1.DB.MoveRole(r.GuildId, r.RoleId, r.BeforeId, r.AfterId)
	return &chatv1.MoveRoleResponse{}, err
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/GetGuildRoles")
}

// GetGuildRoles implements the GetGuildRoles RPC
func (v1 *V1) GetGuildRoles(c echo.Context, r *chatv1.GetGuildRolesRequest) (*chatv1.GetGuildRolesResponse, error) {
	roles, err := v1.DB.GetGuildRoles(r.GuildId)
	return &chatv1.GetGuildRolesResponse{
		Roles: roles,
	}, err
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/SetPermissions")
}

// SetPermissions implements the SetPermissions RPC
func (v1 *V1) SetPermissions(c echo.Context, r *chatv1.SetPermissionsRequest) (*empty.Empty, error) {
	return &emptypb.Empty{}, v1.Perms.SetPermissions(r.Perms.Permissions, r.GuildId, r.ChannelId, r.RoleId)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/GetPermissions")
}

// GetPermissions implements the GetPermissions RPC
func (v1 *V1) GetPermissions(c echo.Context, r *chatv1.GetPermissionsRequest) (*chatv1.GetPermissionsResponse, error) {
	return &chatv1.GetPermissionsResponse{Perms: &chatv1.PermissionList{Permissions: v1.Perms.GetPermissions(r.GuildId, r.ChannelId, r.RoleId)}}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		WantsRoles: true,
		Location:   middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/QueryHasPermission")
}

// QueryHasPermission implements the QueryHasPermission RPC
func (v1 *V1) QueryHasPermission(c echo.Context, r *chatv1.QueryPermissionsRequest) (*chatv1.QueryPermissionsResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	if r.As == 0 {
		r.As = c.(middleware.HarmonyContext).UserID
	} else if !(ctx.IsOwner || v1.Perms.Check("permissions.query", ctx.UserRoles, r.GuildId, r.ChannelId)) {
		return nil, responses.NewError(responses.NotEnoughPermissions)
	}

	owner, err := v1.DB.GetOwner(r.GuildId)
	if err != nil {
		return nil, err
	}

	roles, err := v1.DB.RolesForUser(r.GuildId, r.As)
	if err != nil {
		return nil, err
	}
	return &chatv1.QueryPermissionsResponse{
		Ok: owner == r.As || v1.Perms.Check(r.CheckFor, roles, r.GuildId, r.ChannelId),
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/ManageUserRoles")
}

func (v1 *V1) ManageUserRoles(c echo.Context, r *chatv1.ManageUserRolesRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.ManageRoles(r.GuildId, r.UserId, r.GiveRoleIds, r.TakeRoleIds)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/ModifyGuildRole")
}

func (v1 *V1) ModifyGuildRole(c echo.Context, r *chatv1.ModifyGuildRoleRequest) (*empty.Empty, error) {
	return &empty.Empty{}, v1.DB.ModifyRole(r.GuildId, r.Role.RoleId, r.Role.Name, r.Role.Color, r.Role.Hoist, r.Role.Pingable, r.ModifyName, r.ModifyColor, r.ModifyHoist, r.ModifyPingable)
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    10,
		},

		WantsRoles: true,
		Location:   middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/GetUserRoles")
}

func (v1 *V1) GetUserRoles(c echo.Context, r *chatv1.GetUserRolesRequest) (*chatv1.GetUserRolesResponse, error) {
	ctx := c.(middleware.HarmonyContext)

	if r.UserId == 0 {
		return &chatv1.GetUserRolesResponse{
			Roles: ctx.UserRoles,
		}, nil
	}

	if !(ctx.IsOwner || v1.Perms.Check("roles.users.get", ctx.UserRoles, r.GuildId, 0)) {
		return nil, responses.NewError(responses.NotEnoughPermissions)
	}

	roles, err := v1.DB.RolesForUser(r.GuildId, r.UserId)
	if err != nil {
		return nil, err
	}

	return &chatv1.GetUserRolesResponse{
		Roles: roles,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 10 * time.Second,
			Burst:    64,
		},
	}, "/protocol.chat.v1.ChatService/GetUser")
}

// GetUser handles the protocol's GetUser request
func (v1 *V1) GetUser(c echo.Context, r *chatv1.GetUserRequest) (*chatv1.GetUserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	res, err := v1.DB.GetUserByID(r.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, responses.NewError(responses.BadUserID)
		}
		v1.Logger.Exception(err)
		return nil, err
	}
	return &chatv1.GetUserResponse{
		UserName:   res.Username,
		UserAvatar: res.Avatar.String,
		UserStatus: harmonytypesv1.UserStatus(res.Status),
		IsBot:      res.IsBot,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    4,
		},
	}, "/protocol.chat.v1.ChatService/GetUserMetadata")
}

// GetUserMetadata handles the protocol's GetUserMetadata request
func (v1 *V1) GetUserMetadata(c echo.Context, r *chatv1.GetUserMetadataRequest) (*chatv1.GetUserMetadataResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	meta, err := v1.DB.GetUserMetadata(0, r.AppId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, responses.NewError(responses.NoMetadata)
		}
		v1.Logger.Exception(err)
		return nil, err
	}
	return &chatv1.GetUserMetadataResponse{
		Metadata: meta,
	}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},
	}, "/protocol.chat.v1.ChatService/ProfileUpdate")
}

// ProfileUpdate handles the protocol's ProfileUpdate request
func (v1 *V1) ProfileUpdate(c echo.Context, r *chatv1.ProfileUpdateRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := r.Validate(); err != nil {
		return nil, err
	}
	if r.UpdateStatus {
		if err := v1.DB.SetStatus(ctx.UserID, r.NewStatus); err != nil {
			v1.Logger.Exception(err)
			return nil, err
		}
	}
	if r.UpdateUsername {
		if err := v1.DB.SetUsername(ctx.UserID, r.NewUsername); err != nil {
			v1.Logger.Exception(err)
			return nil, err
		}
	}
	if r.UpdateAvatar {
		if err := v1.DB.SetAvatar(ctx.UserID, r.NewAvatar); err != nil {
			v1.Logger.Exception(err)
			return nil, err
		}
	}

	if r.UpdateIsBot {
		if err := v1.DB.SetIsBot(ctx.UserID, r.IsBot); err != nil {
			v1.Logger.Exception(err)
			return nil, err
		}
	}

	guilds, err := v1.DB.GetLocalGuilds(ctx.UserID)
	if err != nil {
		return nil, err
	}

	for _, g := range guilds {
		v1.Streams.BroadcastGuild(g, &chatv1.Event{
			Event: &chatv1.Event_ProfileUpdated_{
				ProfileUpdated: &chatv1.Event_ProfileUpdated{
					UserId:         ctx.UserID,
					NewUsername:    r.NewUsername,
					UpdateUsername: r.UpdateUsername,
					NewAvatar:      r.NewAvatar,
					UpdateAvatar:   r.UpdateAvatar,
					NewStatus:      r.NewStatus,
					UpdateStatus:   r.UpdateStatus,
					IsBot:          r.IsBot,
					UpdateIsBot:    r.UpdateIsBot,
				},
			},
		})
	}

	return &emptypb.Empty{}, nil
}

func (v1 *V1) Sync(c echo.Context, r *chatv1.SyncRequest, out chan *chatv1.SyncEvent) {

}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},

		Location: middleware.GuildLocation | middleware.ChannelLocation,
	}, "/protocol.chat.v1.ChatService/Typing")
}

// Typing handles the protocol's Typing request
func (v1 *V1) Typing(c echo.Context, r *chatv1.TypingRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := r.Validate(); err != nil {
		return nil, err
	}

	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_Typing_{
			Typing: &chatv1.Event_Typing{
				UserId:    ctx.UserID,
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
			Burst:    4,
		},
		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/BanUser")
}

func (v1 *V1) BanUser(c echo.Context, r *chatv1.BanUserRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteMember(r.GuildId, r.UserId); err != nil {
		return nil, err
	}
	if err := v1.DB.BanUser(r.GuildId, r.UserId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_LeftMember{
			LeftMember: &chatv1.Event_MemberLeft{
				MemberId:    r.UserId,
				GuildId:     r.GuildId,
				LeaveReason: chatv1.Event_banned,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},
		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/KickUser")
}

func (v1 *V1) KickUser(c echo.Context, r *chatv1.KickUserRequest) (*empty.Empty, error) {
	if err := v1.DB.DeleteMember(r.GuildId, r.UserId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_LeftMember{
			LeftMember: &chatv1.Event_MemberLeft{
				MemberId:    r.UserId,
				GuildId:     r.GuildId,
				LeaveReason: chatv1.Event_kicked,
			},
		},
	})
	return &emptypb.Empty{}, nil
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},
		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.ChatService/UnbanUser")
}

func (v1 *V1) UnbanUser(c echo.Context, r *chatv1.UnbanUserRequest) (*empty.Empty, error) {
	if err := v1.DB.UnbanUser(r.GuildId, r.UserId); err != nil {
		return nil, err
	}
	v1.Streams.BroadcastGuild(r.GuildId, &chatv1.Event{
		Event: &chatv1.Event_LeftMember{
			LeftMember: &chatv1.Event_MemberLeft{
				MemberId:    r.UserId,
				GuildId:     r.GuildId,
				LeaveReason: chatv1.Event_kicked,
			},
		},
	})
	return &emptypb.Empty{}, nil
}
