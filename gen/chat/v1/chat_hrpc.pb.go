// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package chatv1

import (
	context "context"
	errors "errors"
	server "github.com/harmony-development/hrpc/server"
	proto "google.golang.org/protobuf/proto"
)

type ChatServiceServer interface {
	// Endpoint to create a guild.
	CreateGuild(context.Context, *CreateGuildRequest) (*CreateGuildResponse, error)
	// Endpoint to create an invite.
	CreateInvite(context.Context, *CreateInviteRequest) (*CreateInviteResponse, error)
	// Endpoint to create a channel.
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	// Endpoint to get your guild list.
	GetGuildList(context.Context, *GetGuildListRequest) (*GetGuildListResponse, error)
	// Endpoint to get information about a guild.
	GetGuild(context.Context, *GetGuildRequest) (*GetGuildResponse, error)
	// Endpoint to get the invites of a guild.
	//
	// This requires the "invites.view" permission.
	GetGuildInvites(context.Context, *GetGuildInvitesRequest) (*GetGuildInvitesResponse, error)
	// Endpoint to get the members of a guild.
	GetGuildMembers(context.Context, *GetGuildMembersRequest) (*GetGuildMembersResponse, error)
	// Endpoint to get the channels of a guild.
	//
	// You will only be informed of channels you have the "messages.view"
	// permission for.
	GetGuildChannels(context.Context, *GetGuildChannelsRequest) (*GetGuildChannelsResponse, error)
	// Endpoint to get the messages from a guild channel.
	GetChannelMessages(context.Context, *GetChannelMessagesRequest) (*GetChannelMessagesResponse, error)
	// Endpoint to get a single message from a guild channel.
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	// Endpoint to update a guild's information.
	UpdateGuildInformation(context.Context, *UpdateGuildInformationRequest) (*UpdateGuildInformationResponse, error)
	// Endpoint to update a channel's information.
	UpdateChannelInformation(context.Context, *UpdateChannelInformationRequest) (*UpdateChannelInformationResponse, error)
	// Endpoint to change the position of a channel in the channel list.
	UpdateChannelOrder(context.Context, *UpdateChannelOrderRequest) (*UpdateChannelOrderResponse, error)
	// Endpoint to change the position of all channels in the guild;
	// must specifcy all channels or fails
	UpdateAllChannelOrder(context.Context, *UpdateAllChannelOrderRequest) (*UpdateAllChannelOrderResponse, error)
	// Endpoint to change the text of a message.
	UpdateMessageText(context.Context, *UpdateMessageTextRequest) (*UpdateMessageTextResponse, error)
	// Endpoint to delete a guild.
	DeleteGuild(context.Context, *DeleteGuildRequest) (*DeleteGuildResponse, error)
	// Endpoint to delete an invite.
	DeleteInvite(context.Context, *DeleteInviteRequest) (*DeleteInviteResponse, error)
	// Endpoint to delete a channel.
	DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error)
	// Endpoint to delete a message.
	//
	// This requires the "messages.manage.delete" permission if you are not the
	// message author.
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	// Endpoint to join a guild.
	JoinGuild(context.Context, *JoinGuildRequest) (*JoinGuildResponse, error)
	// Endpoint to leave a guild.
	LeaveGuild(context.Context, *LeaveGuildRequest) (*LeaveGuildResponse, error)
	// Endpont to trigger an action.
	TriggerAction(context.Context, *TriggerActionRequest) (*TriggerActionResponse, error)
	// Endpoint to send a message.
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	// Endpoint to query if a user has a permission.
	QueryHasPermission(context.Context, *QueryHasPermissionRequest) (*QueryHasPermissionResponse, error)
	// Endpoint to set permissions for a role in a guild.
	SetPermissions(context.Context, *SetPermissionsRequest) (*SetPermissionsResponse, error)
	// Endpoint to get permissions for a role in a guild.
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	// Endpoint to change the position of a role in the role list in a guild.
	MoveRole(context.Context, *MoveRoleRequest) (*MoveRoleResponse, error)
	// Endpoint to get the roles from a guild.
	GetGuildRoles(context.Context, *GetGuildRolesRequest) (*GetGuildRolesResponse, error)
	// Endpoint to add a role to a guild.
	AddGuildRole(context.Context, *AddGuildRoleRequest) (*AddGuildRoleResponse, error)
	// Endpoint to a modify a role from a guild.
	ModifyGuildRole(context.Context, *ModifyGuildRoleRequest) (*ModifyGuildRoleResponse, error)
	// Endpoint to delete a role from a guild.
	DeleteGuildRole(context.Context, *DeleteGuildRoleRequest) (*DeleteGuildRoleResponse, error)
	// Endpoint to manage the roles of a guild member.
	ManageUserRoles(context.Context, *ManageUserRolesRequest) (*ManageUserRolesResponse, error)
	// Endpoint to get the roles a guild member has.
	GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error)
	// Endpoint to send a typing notification in a guild channel.
	Typing(context.Context, *TypingRequest) (*TypingResponse, error)
	// Endpoint to "preview" a guild, which returns some information about a
	// guild.
	PreviewGuild(context.Context, *PreviewGuildRequest) (*PreviewGuildResponse, error)
	// Endpoint to ban a user from a guild.
	BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error)
	// Endpoint to kick a user from a guild.
	KickUser(context.Context, *KickUserRequest) (*KickUserResponse, error)
	// Endpoint to unban a user from a guild.
	UnbanUser(context.Context, *UnbanUserRequest) (*UnbanUserResponse, error)
	// Endpoint to get all pinned messages in a guild channel.
	GetPinnedMessages(context.Context, *GetPinnedMessagesRequest) (*GetPinnedMessagesResponse, error)
	// Endpoint to pin a message in a guild channel.
	PinMessage(context.Context, *PinMessageRequest) (*PinMessageResponse, error)
	// Endpoint to unpin a message in a guild channel.
	UnpinMessage(context.Context, *UnpinMessageRequest) (*UnpinMessageResponse, error)
	// Endpoint to stream events from the homeserver.
	StreamEvents(context.Context, chan *StreamEventsRequest) (chan *StreamEventsResponse, error)
}

type DefaultChatService struct{}

func (DefaultChatService) CreateGuild(context.Context, *CreateGuildRequest) (*CreateGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) CreateInvite(context.Context, *CreateInviteRequest) (*CreateInviteResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuildList(context.Context, *GetGuildListRequest) (*GetGuildListResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuild(context.Context, *GetGuildRequest) (*GetGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuildInvites(context.Context, *GetGuildInvitesRequest) (*GetGuildInvitesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuildMembers(context.Context, *GetGuildMembersRequest) (*GetGuildMembersResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuildChannels(context.Context, *GetGuildChannelsRequest) (*GetGuildChannelsResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetChannelMessages(context.Context, *GetChannelMessagesRequest) (*GetChannelMessagesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UpdateGuildInformation(context.Context, *UpdateGuildInformationRequest) (*UpdateGuildInformationResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UpdateChannelInformation(context.Context, *UpdateChannelInformationRequest) (*UpdateChannelInformationResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UpdateChannelOrder(context.Context, *UpdateChannelOrderRequest) (*UpdateChannelOrderResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UpdateAllChannelOrder(context.Context, *UpdateAllChannelOrderRequest) (*UpdateAllChannelOrderResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UpdateMessageText(context.Context, *UpdateMessageTextRequest) (*UpdateMessageTextResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) DeleteGuild(context.Context, *DeleteGuildRequest) (*DeleteGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) DeleteInvite(context.Context, *DeleteInviteRequest) (*DeleteInviteResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) DeleteChannel(context.Context, *DeleteChannelRequest) (*DeleteChannelResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) JoinGuild(context.Context, *JoinGuildRequest) (*JoinGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) LeaveGuild(context.Context, *LeaveGuildRequest) (*LeaveGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) TriggerAction(context.Context, *TriggerActionRequest) (*TriggerActionResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) QueryHasPermission(context.Context, *QueryHasPermissionRequest) (*QueryHasPermissionResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) SetPermissions(context.Context, *SetPermissionsRequest) (*SetPermissionsResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) MoveRole(context.Context, *MoveRoleRequest) (*MoveRoleResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetGuildRoles(context.Context, *GetGuildRolesRequest) (*GetGuildRolesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) AddGuildRole(context.Context, *AddGuildRoleRequest) (*AddGuildRoleResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) ModifyGuildRole(context.Context, *ModifyGuildRoleRequest) (*ModifyGuildRoleResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) DeleteGuildRole(context.Context, *DeleteGuildRoleRequest) (*DeleteGuildRoleResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) ManageUserRoles(context.Context, *ManageUserRolesRequest) (*ManageUserRolesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetUserRoles(context.Context, *GetUserRolesRequest) (*GetUserRolesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) Typing(context.Context, *TypingRequest) (*TypingResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) PreviewGuild(context.Context, *PreviewGuildRequest) (*PreviewGuildResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) BanUser(context.Context, *BanUserRequest) (*BanUserResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) KickUser(context.Context, *KickUserRequest) (*KickUserResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UnbanUser(context.Context, *UnbanUserRequest) (*UnbanUserResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) GetPinnedMessages(context.Context, *GetPinnedMessagesRequest) (*GetPinnedMessagesResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) PinMessage(context.Context, *PinMessageRequest) (*PinMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) UnpinMessage(context.Context, *UnpinMessageRequest) (*UnpinMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
func (DefaultChatService) StreamEvents(context.Context, chan *StreamEventsRequest) (chan *StreamEventsResponse, error) {
	return nil, errors.New("unimplemented")
}

type ChatServiceHandler struct {
	Server ChatServiceServer
}

func NewChatServiceHandler(server ChatServiceServer) *ChatServiceHandler {
	return &ChatServiceHandler{Server: server}
}
func (h *ChatServiceHandler) Name() string {
	return "ChatService"
}
func (h *ChatServiceHandler) Routes() map[string]server.RawHandler {
	return map[string]server.RawHandler{
		"/protocol.chat.v1.ChatService.CreateGuild/": server.NewUnaryHandler(&CreateGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.CreateGuild(c, req.(*CreateGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.CreateInvite/": server.NewUnaryHandler(&CreateInviteRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.CreateInvite(c, req.(*CreateInviteRequest))
		}),
		"/protocol.chat.v1.ChatService.CreateChannel/": server.NewUnaryHandler(&CreateChannelRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.CreateChannel(c, req.(*CreateChannelRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuildList/": server.NewUnaryHandler(&GetGuildListRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuildList(c, req.(*GetGuildListRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuild/": server.NewUnaryHandler(&GetGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuild(c, req.(*GetGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuildInvites/": server.NewUnaryHandler(&GetGuildInvitesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuildInvites(c, req.(*GetGuildInvitesRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuildMembers/": server.NewUnaryHandler(&GetGuildMembersRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuildMembers(c, req.(*GetGuildMembersRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuildChannels/": server.NewUnaryHandler(&GetGuildChannelsRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuildChannels(c, req.(*GetGuildChannelsRequest))
		}),
		"/protocol.chat.v1.ChatService.GetChannelMessages/": server.NewUnaryHandler(&GetChannelMessagesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetChannelMessages(c, req.(*GetChannelMessagesRequest))
		}),
		"/protocol.chat.v1.ChatService.GetMessage/": server.NewUnaryHandler(&GetMessageRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetMessage(c, req.(*GetMessageRequest))
		}),
		"/protocol.chat.v1.ChatService.UpdateGuildInformation/": server.NewUnaryHandler(&UpdateGuildInformationRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateGuildInformation(c, req.(*UpdateGuildInformationRequest))
		}),
		"/protocol.chat.v1.ChatService.UpdateChannelInformation/": server.NewUnaryHandler(&UpdateChannelInformationRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateChannelInformation(c, req.(*UpdateChannelInformationRequest))
		}),
		"/protocol.chat.v1.ChatService.UpdateChannelOrder/": server.NewUnaryHandler(&UpdateChannelOrderRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateChannelOrder(c, req.(*UpdateChannelOrderRequest))
		}),
		"/protocol.chat.v1.ChatService.UpdateAllChannelOrder/": server.NewUnaryHandler(&UpdateAllChannelOrderRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateAllChannelOrder(c, req.(*UpdateAllChannelOrderRequest))
		}),
		"/protocol.chat.v1.ChatService.UpdateMessageText/": server.NewUnaryHandler(&UpdateMessageTextRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UpdateMessageText(c, req.(*UpdateMessageTextRequest))
		}),
		"/protocol.chat.v1.ChatService.DeleteGuild/": server.NewUnaryHandler(&DeleteGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.DeleteGuild(c, req.(*DeleteGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.DeleteInvite/": server.NewUnaryHandler(&DeleteInviteRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.DeleteInvite(c, req.(*DeleteInviteRequest))
		}),
		"/protocol.chat.v1.ChatService.DeleteChannel/": server.NewUnaryHandler(&DeleteChannelRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.DeleteChannel(c, req.(*DeleteChannelRequest))
		}),
		"/protocol.chat.v1.ChatService.DeleteMessage/": server.NewUnaryHandler(&DeleteMessageRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.DeleteMessage(c, req.(*DeleteMessageRequest))
		}),
		"/protocol.chat.v1.ChatService.JoinGuild/": server.NewUnaryHandler(&JoinGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.JoinGuild(c, req.(*JoinGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.LeaveGuild/": server.NewUnaryHandler(&LeaveGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.LeaveGuild(c, req.(*LeaveGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.TriggerAction/": server.NewUnaryHandler(&TriggerActionRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.TriggerAction(c, req.(*TriggerActionRequest))
		}),
		"/protocol.chat.v1.ChatService.SendMessage/": server.NewUnaryHandler(&SendMessageRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.SendMessage(c, req.(*SendMessageRequest))
		}),
		"/protocol.chat.v1.ChatService.QueryHasPermission/": server.NewUnaryHandler(&QueryHasPermissionRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.QueryHasPermission(c, req.(*QueryHasPermissionRequest))
		}),
		"/protocol.chat.v1.ChatService.SetPermissions/": server.NewUnaryHandler(&SetPermissionsRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.SetPermissions(c, req.(*SetPermissionsRequest))
		}),
		"/protocol.chat.v1.ChatService.GetPermissions/": server.NewUnaryHandler(&GetPermissionsRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetPermissions(c, req.(*GetPermissionsRequest))
		}),
		"/protocol.chat.v1.ChatService.MoveRole/": server.NewUnaryHandler(&MoveRoleRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.MoveRole(c, req.(*MoveRoleRequest))
		}),
		"/protocol.chat.v1.ChatService.GetGuildRoles/": server.NewUnaryHandler(&GetGuildRolesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetGuildRoles(c, req.(*GetGuildRolesRequest))
		}),
		"/protocol.chat.v1.ChatService.AddGuildRole/": server.NewUnaryHandler(&AddGuildRoleRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.AddGuildRole(c, req.(*AddGuildRoleRequest))
		}),
		"/protocol.chat.v1.ChatService.ModifyGuildRole/": server.NewUnaryHandler(&ModifyGuildRoleRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.ModifyGuildRole(c, req.(*ModifyGuildRoleRequest))
		}),
		"/protocol.chat.v1.ChatService.DeleteGuildRole/": server.NewUnaryHandler(&DeleteGuildRoleRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.DeleteGuildRole(c, req.(*DeleteGuildRoleRequest))
		}),
		"/protocol.chat.v1.ChatService.ManageUserRoles/": server.NewUnaryHandler(&ManageUserRolesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.ManageUserRoles(c, req.(*ManageUserRolesRequest))
		}),
		"/protocol.chat.v1.ChatService.GetUserRoles/": server.NewUnaryHandler(&GetUserRolesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetUserRoles(c, req.(*GetUserRolesRequest))
		}),
		"/protocol.chat.v1.ChatService.Typing/": server.NewUnaryHandler(&TypingRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.Typing(c, req.(*TypingRequest))
		}),
		"/protocol.chat.v1.ChatService.PreviewGuild/": server.NewUnaryHandler(&PreviewGuildRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.PreviewGuild(c, req.(*PreviewGuildRequest))
		}),
		"/protocol.chat.v1.ChatService.BanUser/": server.NewUnaryHandler(&BanUserRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.BanUser(c, req.(*BanUserRequest))
		}),
		"/protocol.chat.v1.ChatService.KickUser/": server.NewUnaryHandler(&KickUserRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.KickUser(c, req.(*KickUserRequest))
		}),
		"/protocol.chat.v1.ChatService.UnbanUser/": server.NewUnaryHandler(&UnbanUserRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UnbanUser(c, req.(*UnbanUserRequest))
		}),
		"/protocol.chat.v1.ChatService.GetPinnedMessages/": server.NewUnaryHandler(&GetPinnedMessagesRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.GetPinnedMessages(c, req.(*GetPinnedMessagesRequest))
		}),
		"/protocol.chat.v1.ChatService.PinMessage/": server.NewUnaryHandler(&PinMessageRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.PinMessage(c, req.(*PinMessageRequest))
		}),
		"/protocol.chat.v1.ChatService.UnpinMessage/": server.NewUnaryHandler(&UnpinMessageRequest{}, func(c context.Context, req proto.Message) (proto.Message, error) {
			return h.Server.UnpinMessage(c, req.(*UnpinMessageRequest))
		}),
	}
}
