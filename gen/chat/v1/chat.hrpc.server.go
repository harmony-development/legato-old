package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"

import "github.com/golang/protobuf/ptypes/empty"

type ChatServiceServer interface {
	CreateGuild(ctx context.Context, r *CreateGuildRequest, headers http.Header) (resp *CreateGuildResponse, err error)

	CreateInvite(ctx context.Context, r *CreateInviteRequest, headers http.Header) (resp *CreateInviteResponse, err error)

	CreateChannel(ctx context.Context, r *CreateChannelRequest, headers http.Header) (resp *CreateChannelResponse, err error)

	CreateEmotePack(ctx context.Context, r *CreateEmotePackRequest, headers http.Header) (resp *CreateEmotePackResponse, err error)

	GetGuildList(ctx context.Context, r *GetGuildListRequest, headers http.Header) (resp *GetGuildListResponse, err error)

	AddGuildToGuildList(ctx context.Context, r *AddGuildToGuildListRequest, headers http.Header) (resp *AddGuildToGuildListResponse, err error)

	RemoveGuildFromGuildList(ctx context.Context, r *RemoveGuildFromGuildListRequest, headers http.Header) (resp *RemoveGuildFromGuildListResponse, err error)

	GetGuild(ctx context.Context, r *GetGuildRequest, headers http.Header) (resp *GetGuildResponse, err error)

	GetGuildInvites(ctx context.Context, r *GetGuildInvitesRequest, headers http.Header) (resp *GetGuildInvitesResponse, err error)

	GetGuildMembers(ctx context.Context, r *GetGuildMembersRequest, headers http.Header) (resp *GetGuildMembersResponse, err error)

	GetGuildChannels(ctx context.Context, r *GetGuildChannelsRequest, headers http.Header) (resp *GetGuildChannelsResponse, err error)

	GetChannelMessages(ctx context.Context, r *GetChannelMessagesRequest, headers http.Header) (resp *GetChannelMessagesResponse, err error)

	GetMessage(ctx context.Context, r *GetMessageRequest, headers http.Header) (resp *GetMessageResponse, err error)

	GetEmotePacks(ctx context.Context, r *GetEmotePacksRequest, headers http.Header) (resp *GetEmotePacksResponse, err error)

	GetEmotePackEmotes(ctx context.Context, r *GetEmotePackEmotesRequest, headers http.Header) (resp *GetEmotePackEmotesResponse, err error)

	UpdateGuildInformation(ctx context.Context, r *UpdateGuildInformationRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateChannelInformation(ctx context.Context, r *UpdateChannelInformationRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateChannelOrder(ctx context.Context, r *UpdateChannelOrderRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateMessage(ctx context.Context, r *UpdateMessageRequest, headers http.Header) (resp *empty.Empty, err error)

	AddEmoteToPack(ctx context.Context, r *AddEmoteToPackRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteGuild(ctx context.Context, r *DeleteGuildRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteInvite(ctx context.Context, r *DeleteInviteRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteChannel(ctx context.Context, r *DeleteChannelRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteMessage(ctx context.Context, r *DeleteMessageRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteEmoteFromPack(ctx context.Context, r *DeleteEmoteFromPackRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteEmotePack(ctx context.Context, r *DeleteEmotePackRequest, headers http.Header) (resp *empty.Empty, err error)

	DequipEmotePack(ctx context.Context, r *DequipEmotePackRequest, headers http.Header) (resp *empty.Empty, err error)

	JoinGuild(ctx context.Context, r *JoinGuildRequest, headers http.Header) (resp *JoinGuildResponse, err error)

	LeaveGuild(ctx context.Context, r *LeaveGuildRequest, headers http.Header) (resp *empty.Empty, err error)

	TriggerAction(ctx context.Context, r *TriggerActionRequest, headers http.Header) (resp *empty.Empty, err error)

	SendMessage(ctx context.Context, r *SendMessageRequest, headers http.Header) (resp *SendMessageResponse, err error)

	QueryHasPermission(ctx context.Context, r *QueryPermissionsRequest, headers http.Header) (resp *QueryPermissionsResponse, err error)

	SetPermissions(ctx context.Context, r *SetPermissionsRequest, headers http.Header) (resp *empty.Empty, err error)

	GetPermissions(ctx context.Context, r *GetPermissionsRequest, headers http.Header) (resp *GetPermissionsResponse, err error)

	MoveRole(ctx context.Context, r *MoveRoleRequest, headers http.Header) (resp *MoveRoleResponse, err error)

	GetGuildRoles(ctx context.Context, r *GetGuildRolesRequest, headers http.Header) (resp *GetGuildRolesResponse, err error)

	AddGuildRole(ctx context.Context, r *AddGuildRoleRequest, headers http.Header) (resp *AddGuildRoleResponse, err error)

	ModifyGuildRole(ctx context.Context, r *ModifyGuildRoleRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteGuildRole(ctx context.Context, r *DeleteGuildRoleRequest, headers http.Header) (resp *empty.Empty, err error)

	ManageUserRoles(ctx context.Context, r *ManageUserRolesRequest, headers http.Header) (resp *empty.Empty, err error)

	GetUserRoles(ctx context.Context, r *GetUserRolesRequest, headers http.Header) (resp *GetUserRolesResponse, err error)

	StreamEvents(ctx context.Context, in chan *StreamEventsRequest, out chan *Event, headers http.Header)

	Sync(ctx context.Context, r *SyncRequest, out chan *SyncEvent, headers http.Header)

	GetUser(ctx context.Context, r *GetUserRequest, headers http.Header) (resp *GetUserResponse, err error)

	GetUserMetadata(ctx context.Context, r *GetUserMetadataRequest, headers http.Header) (resp *GetUserMetadataResponse, err error)

	ProfileUpdate(ctx context.Context, r *ProfileUpdateRequest, headers http.Header) (resp *empty.Empty, err error)

	Typing(ctx context.Context, r *TypingRequest, headers http.Header) (resp *empty.Empty, err error)

	PreviewGuild(ctx context.Context, r *PreviewGuildRequest, headers http.Header) (resp *PreviewGuildResponse, err error)
}

type ChatServiceHandler struct {
	Server       ChatServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	upgrader     websocket.Upgrader
}

func NewChatServiceHandler(s ChatServiceServer, errHandler func(err error, w http.ResponseWriter)) *ChatServiceHandler {
	return &ChatServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *ChatServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.chat.v1.ChatService/CreateGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateInvite":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateInviteRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateInvite(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateChannel":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateChannelRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateChannel(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddGuildToGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddGuildToGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddGuildToGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/RemoveGuildFromGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(RemoveGuildFromGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.RemoveGuildFromGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildInvites":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildInvitesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildInvites(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildMembers":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildMembersRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildMembers(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildChannels":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildChannelsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildChannels(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetChannelMessages":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetChannelMessagesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetChannelMessages(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetEmotePacks":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetEmotePacksRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetEmotePacks(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetEmotePackEmotes":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetEmotePackEmotesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetEmotePackEmotes(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateGuildInformation":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateGuildInformationRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateGuildInformation(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelInformation":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateChannelInformationRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateChannelInformation(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelOrder":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateChannelOrderRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateChannelOrder(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddEmoteToPack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddEmoteToPackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddEmoteToPack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteInvite":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteInviteRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteInvite(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteChannel":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteChannelRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteChannel(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteEmoteFromPack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteEmoteFromPackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteEmoteFromPack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DequipEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DequipEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DequipEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/JoinGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(JoinGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.JoinGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/LeaveGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(LeaveGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.LeaveGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/TriggerAction":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(TriggerActionRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.TriggerAction(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/SendMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(SendMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.SendMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/QueryHasPermission":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(QueryPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.QueryHasPermission(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/SetPermissions":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(SetPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.SetPermissions(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetPermissions":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetPermissions(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/MoveRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(MoveRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.MoveRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ModifyGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ModifyGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ModifyGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ManageUserRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ManageUserRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ManageUserRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetUserRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUserRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/StreamEvents":
		{
			var err error

			in := make(chan *StreamEventsRequest)
			err = nil

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *Event)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				msgs := make(chan []byte)

				go func() {
					for {
						_, message, err := ws.ReadMessage()
						if err != nil {
							close(msgs)
							break
						}
						msgs <- message
					}
				}()

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case data, ok := <-msgs:
						if !ok {
							return
						}

						item := new(StreamEventsRequest)
						err = proto.Unmarshal(data, item)
						if err != nil {
							close(in)
							close(out)
							return
						}

						in <- item

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(in)

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(in)

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(in)

							close(out)
							return
						}
					}
				}
			}()

			h.Server.StreamEvents(req.Context(), in, out, req.Header)
		}

	case "/protocol.chat.v1.ChatService/Sync":
		{
			var err error

			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			in := new(SyncRequest)
			err = proto.Unmarshal(body, in)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *SyncEvent)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(out)
							return
						}
					}
				}
			}()

			h.Server.Sync(req.Context(), in, out, req.Header)
		}

	case "/protocol.chat.v1.ChatService/GetUser":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUser(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetUserMetadata":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserMetadataRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUserMetadata(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ProfileUpdate":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ProfileUpdateRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ProfileUpdate(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/Typing":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(TypingRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.Typing(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/PreviewGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(PreviewGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.PreviewGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	}
}
