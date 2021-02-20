// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: chat/v1/messages.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	v1 "github.com/harmony-development/legato/gen/harmonytypes/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetChannelMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId       uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId     uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	BeforeMessage uint64 `protobuf:"varint,3,opt,name=before_message,json=beforeMessage,proto3" json:"before_message,omitempty"`
}

func (x *GetChannelMessagesRequest) Reset() {
	*x = GetChannelMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelMessagesRequest) ProtoMessage() {}

func (x *GetChannelMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetChannelMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GetChannelMessagesRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *GetChannelMessagesRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetChannelMessagesRequest) GetBeforeMessage() uint64 {
	if x != nil {
		return x.BeforeMessage
	}
	return 0
}

type GetChannelMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReachedTop bool          `protobuf:"varint,1,opt,name=reached_top,json=reachedTop,proto3" json:"reached_top,omitempty"`
	Messages   []*v1.Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetChannelMessagesResponse) Reset() {
	*x = GetChannelMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChannelMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelMessagesResponse) ProtoMessage() {}

func (x *GetChannelMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetChannelMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetChannelMessagesResponse) GetReachedTop() bool {
	if x != nil {
		return x.ReachedTop
	}
	return false
}

func (x *GetChannelMessagesResponse) GetMessages() []*v1.Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId   uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MessageId uint64 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *GetMessageRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetMessageRequest) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *v1.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageResponse) GetMessage() *v1.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type UpdateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId           uint64       `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId         uint64       `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MessageId         uint64       `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Content           string       `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	UpdateContent     bool         `protobuf:"varint,5,opt,name=update_content,json=updateContent,proto3" json:"update_content,omitempty"`
	Embeds            []*v1.Embed  `protobuf:"bytes,6,rep,name=embeds,proto3" json:"embeds,omitempty"`
	UpdateEmbeds      bool         `protobuf:"varint,7,opt,name=update_embeds,json=updateEmbeds,proto3" json:"update_embeds,omitempty"`
	Actions           []*v1.Action `protobuf:"bytes,8,rep,name=actions,proto3" json:"actions,omitempty"`
	UpdateActions     bool         `protobuf:"varint,9,opt,name=update_actions,json=updateActions,proto3" json:"update_actions,omitempty"`
	Attachments       []string     `protobuf:"bytes,10,rep,name=attachments,proto3" json:"attachments,omitempty"`
	UpdateAttachments bool         `protobuf:"varint,11,opt,name=update_attachments,json=updateAttachments,proto3" json:"update_attachments,omitempty"`
	Overrides         *v1.Override `protobuf:"bytes,12,opt,name=overrides,proto3" json:"overrides,omitempty"`
	UpdateOverrides   bool         `protobuf:"varint,13,opt,name=update_overrides,json=updateOverrides,proto3" json:"update_overrides,omitempty"`
	Metadata          *v1.Metadata `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`
	UpdateMetadata    bool         `protobuf:"varint,15,opt,name=update_metadata,json=updateMetadata,proto3" json:"update_metadata,omitempty"`
}

func (x *UpdateMessageRequest) Reset() {
	*x = UpdateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequest) ProtoMessage() {}

func (x *UpdateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMessageRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *UpdateMessageRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *UpdateMessageRequest) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *UpdateMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateMessageRequest) GetUpdateContent() bool {
	if x != nil {
		return x.UpdateContent
	}
	return false
}

func (x *UpdateMessageRequest) GetEmbeds() []*v1.Embed {
	if x != nil {
		return x.Embeds
	}
	return nil
}

func (x *UpdateMessageRequest) GetUpdateEmbeds() bool {
	if x != nil {
		return x.UpdateEmbeds
	}
	return false
}

func (x *UpdateMessageRequest) GetActions() []*v1.Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *UpdateMessageRequest) GetUpdateActions() bool {
	if x != nil {
		return x.UpdateActions
	}
	return false
}

func (x *UpdateMessageRequest) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *UpdateMessageRequest) GetUpdateAttachments() bool {
	if x != nil {
		return x.UpdateAttachments
	}
	return false
}

func (x *UpdateMessageRequest) GetOverrides() *v1.Override {
	if x != nil {
		return x.Overrides
	}
	return nil
}

func (x *UpdateMessageRequest) GetUpdateOverrides() bool {
	if x != nil {
		return x.UpdateOverrides
	}
	return false
}

func (x *UpdateMessageRequest) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateMessageRequest) GetUpdateMetadata() bool {
	if x != nil {
		return x.UpdateMetadata
	}
	return false
}

type DeleteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId   uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MessageId uint64 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMessageRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *DeleteMessageRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *DeleteMessageRequest) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type TriggerActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId    uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId  uint64 `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	MessageId  uint64 `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	ActionId   string `protobuf:"bytes,4,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	ActionData string `protobuf:"bytes,5,opt,name=action_data,json=actionData,proto3" json:"action_data,omitempty"`
}

func (x *TriggerActionRequest) Reset() {
	*x = TriggerActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerActionRequest) ProtoMessage() {}

func (x *TriggerActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerActionRequest.ProtoReflect.Descriptor instead.
func (*TriggerActionRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *TriggerActionRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *TriggerActionRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *TriggerActionRequest) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *TriggerActionRequest) GetActionId() string {
	if x != nil {
		return x.ActionId
	}
	return ""
}

func (x *TriggerActionRequest) GetActionData() string {
	if x != nil {
		return x.ActionData
	}
	return ""
}

// SendMessage
type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId     uint64       `protobuf:"varint,1,opt,name=guild_id,json=guildId,proto3" json:"guild_id,omitempty"`
	ChannelId   uint64       `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Content     string       `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Actions     []*v1.Action `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Embeds      []*v1.Embed  `protobuf:"bytes,5,rep,name=embeds,proto3" json:"embeds,omitempty"`
	Attachments []string     `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
	InReplyTo   uint64       `protobuf:"varint,7,opt,name=in_reply_to,json=inReplyTo,proto3" json:"in_reply_to,omitempty"`
	Overrides   *v1.Override `protobuf:"bytes,8,opt,name=overrides,proto3" json:"overrides,omitempty"`
	EchoId      uint64       `protobuf:"varint,9,opt,name=echo_id,json=echoId,proto3" json:"echo_id,omitempty"`
	Metadata    *v1.Metadata `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *SendMessageRequest) GetGuildId() uint64 {
	if x != nil {
		return x.GuildId
	}
	return 0
}

func (x *SendMessageRequest) GetChannelId() uint64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendMessageRequest) GetActions() []*v1.Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *SendMessageRequest) GetEmbeds() []*v1.Embed {
	if x != nil {
		return x.Embeds
	}
	return nil
}

func (x *SendMessageRequest) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *SendMessageRequest) GetInReplyTo() uint64 {
	if x != nil {
		return x.InReplyTo
	}
	return 0
}

func (x *SendMessageRequest) GetOverrides() *v1.Override {
	if x != nil {
		return x.Overrides
	}
	return nil
}

func (x *SendMessageRequest) GetEchoId() uint64 {
	if x != nil {
		return x.EchoId
	}
	return 0
}

func (x *SendMessageRequest) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId uint64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *SendMessageResponse) GetMessageId() uint64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

var File_chat_v1_messages_proto protoreflect.FileDescriptor

var file_chat_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x68, 0x61, 0x72, 0x6d,
	0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0e, 0x62, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x0d, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x7c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x54, 0x6f,
	0x70, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x78, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72,
	0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa4, 0x05,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x06,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x06, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x40,
	0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72,
	0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x7b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0xb9, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75,
	0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc2, 0x03,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72,
	0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x37, 0x0a, 0x06,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x06, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x40, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x52, 0x09,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x63, 0x68,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x63, 0x68, 0x6f,
	0x49, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x38, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30,
	0x01, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f,
	0x6e, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v1_messages_proto_rawDescOnce sync.Once
	file_chat_v1_messages_proto_rawDescData = file_chat_v1_messages_proto_rawDesc
)

func file_chat_v1_messages_proto_rawDescGZIP() []byte {
	file_chat_v1_messages_proto_rawDescOnce.Do(func() {
		file_chat_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v1_messages_proto_rawDescData)
	})
	return file_chat_v1_messages_proto_rawDescData
}

var file_chat_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_chat_v1_messages_proto_goTypes = []interface{}{
	(*GetChannelMessagesRequest)(nil),  // 0: protocol.chat.v1.GetChannelMessagesRequest
	(*GetChannelMessagesResponse)(nil), // 1: protocol.chat.v1.GetChannelMessagesResponse
	(*GetMessageRequest)(nil),          // 2: protocol.chat.v1.GetMessageRequest
	(*GetMessageResponse)(nil),         // 3: protocol.chat.v1.GetMessageResponse
	(*UpdateMessageRequest)(nil),       // 4: protocol.chat.v1.UpdateMessageRequest
	(*DeleteMessageRequest)(nil),       // 5: protocol.chat.v1.DeleteMessageRequest
	(*TriggerActionRequest)(nil),       // 6: protocol.chat.v1.TriggerActionRequest
	(*SendMessageRequest)(nil),         // 7: protocol.chat.v1.SendMessageRequest
	(*SendMessageResponse)(nil),        // 8: protocol.chat.v1.SendMessageResponse
	(*v1.Message)(nil),                 // 9: protocol.harmonytypes.v1.Message
	(*v1.Embed)(nil),                   // 10: protocol.harmonytypes.v1.Embed
	(*v1.Action)(nil),                  // 11: protocol.harmonytypes.v1.Action
	(*v1.Override)(nil),                // 12: protocol.harmonytypes.v1.Override
	(*v1.Metadata)(nil),                // 13: protocol.harmonytypes.v1.Metadata
}
var file_chat_v1_messages_proto_depIdxs = []int32{
	9,  // 0: protocol.chat.v1.GetChannelMessagesResponse.messages:type_name -> protocol.harmonytypes.v1.Message
	9,  // 1: protocol.chat.v1.GetMessageResponse.message:type_name -> protocol.harmonytypes.v1.Message
	10, // 2: protocol.chat.v1.UpdateMessageRequest.embeds:type_name -> protocol.harmonytypes.v1.Embed
	11, // 3: protocol.chat.v1.UpdateMessageRequest.actions:type_name -> protocol.harmonytypes.v1.Action
	12, // 4: protocol.chat.v1.UpdateMessageRequest.overrides:type_name -> protocol.harmonytypes.v1.Override
	13, // 5: protocol.chat.v1.UpdateMessageRequest.metadata:type_name -> protocol.harmonytypes.v1.Metadata
	11, // 6: protocol.chat.v1.SendMessageRequest.actions:type_name -> protocol.harmonytypes.v1.Action
	10, // 7: protocol.chat.v1.SendMessageRequest.embeds:type_name -> protocol.harmonytypes.v1.Embed
	12, // 8: protocol.chat.v1.SendMessageRequest.overrides:type_name -> protocol.harmonytypes.v1.Override
	13, // 9: protocol.chat.v1.SendMessageRequest.metadata:type_name -> protocol.harmonytypes.v1.Metadata
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_chat_v1_messages_proto_init() }
func file_chat_v1_messages_proto_init() {
	if File_chat_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelMessagesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChannelMessagesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerActionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_v1_messages_proto_goTypes,
		DependencyIndexes: file_chat_v1_messages_proto_depIdxs,
		MessageInfos:      file_chat_v1_messages_proto_msgTypes,
	}.Build()
	File_chat_v1_messages_proto = out.File
	file_chat_v1_messages_proto_rawDesc = nil
	file_chat_v1_messages_proto_goTypes = nil
	file_chat_v1_messages_proto_depIdxs = nil
}
