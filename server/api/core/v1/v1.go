package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

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
	guild, err := v1.DB.CreateGuild(ctx.UserID, guildID, r.GuildName, r.PictureUrl)
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
	channel, err := v1.DB.AddChannelToGuild(r.Location.GuildId, r.ChannelName)
	if err != nil {
		return nil, err
	}
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
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
		Location:   middleware.GuildLocation | middleware.ChannelLocation | middleware.MessageLocation | middleware.AuthorLocation,
		Permission: middleware.NoPermission,
	}, "/protocol.core.v1.CoreService/UpdateMessage")
}

func (v1 *V1) UpdateMessage(c context.Context, r *corev1.UpdateMessageRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if !r.UpdateActions && !r.UpdateEmbeds && !r.UpdateContent {
		return nil, errors.New("bad request; nothing is being edited")
	}

	owner, err := v1.DB.GetMessageOwner(r.Location.MessageId)
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
	tiempo, err := v1.DB.UpdateMessage(r.Location.MessageId, &r.Content, embeds, actions)
	if err != nil {
		return nil, err
	}
	editedAt, _ := ptypes.TimestampProto(tiempo.UTC())
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_EditedMessage{
			EditedMessage: &corev1.GuildEvent_MessageUpdated{
				Location:      r.Location,
				Content:       r.Content,
				UpdateContent: r.UpdateContent,
				Embeds:        r.Embeds,
				UpdateEmbeds:  r.UpdateEmbeds,
				Actions:       r.Actions,
				UpdateActions: r.UpdateActions,
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
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
		return nil, err
	}
	if err := v1.DB.AddMemberToGuild(ctx.UserID, guildID); err != nil {
		return nil, err
	}
	if err := v1.DB.IncrementInvite(r.InviteId); err != nil {
		return nil, err
	}
	streamState.BroadcastGuild(guildID, &corev1.GuildEvent{
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
		return nil, errors.New("You cannot leave a guild you own")
	}
	streamState.RemoveUserFromGuild(r.Location.GuildId, ctx.UserID)
	streamState.BroadcastGuild(r.Location.GuildId, &corev1.GuildEvent{
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
	userID, err := middleware.CheckAuth(v1.DB, s.Context())
	if err != nil {
		return err
	}
	ok, err := v1.DB.UserInGuild(userID, r.Location.GuildId)
	if err != nil {
		return err
	}
	if !ok {
		return NotInGuild
	}
	streamState.Add(r.Location.GuildId, userID, s)
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
	userID, err := middleware.CheckAuth(v1.DB, s.Context())
	if err != nil {
		return err
	}
	streamState.AddAction(userID, s)
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
		return nil, errors.New("Invalid location")
	}
	for _, action := range v1.ActionsToProto(msg.Actions) {
		if action.Id == r.ActionId {
			streamState.BroadcastAction(ctx.UserID, &corev1.ActionEvent{
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
	return nil, errors.New("Invalid action ID")
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

func (v1 *V1) SendMessage(c context.Context, r *corev1.SendMessageRequest) (*emptypb.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	msg, err := v1.DB.AddMessage(
		r.Message.Location.ChannelId,
		r.Message.Location.GuildId,
		ctx.UserID,
		r.Message.Content,
		r.Message.Attachments,
		v1.ProtoToEmbeds(r.Message.Embeds),
		v1.ProtoToActions(r.Message.Actions),
	)
	if err != nil {
		return nil, err
	}
	message := r.Message
	createdAt, _ := ptypes.TimestampProto(msg.CreatedAt.UTC())
	message.CreatedAt = createdAt
	message.Location.MessageId = msg.MessageID
	message.AuthorId = ctx.UserID
	streamState.BroadcastGuild(r.Message.Location.GuildId, &corev1.GuildEvent{
		Event: &corev1.GuildEvent_SentMessage{
			SentMessage: &corev1.GuildEvent_MessageSent{
				Message: message,
			},
		},
	})
	return &emptypb.Empty{}, nil
}
