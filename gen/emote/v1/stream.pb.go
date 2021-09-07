// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: emote/v1/stream.proto

package emotev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event sent when an emote pack's information is changed.
//
// Should only be sent to users who have the pack equipped.
type EmotePackUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the pack that was updated.
	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
	// New pack name of the pack.
	NewPackName *string `protobuf:"bytes,2,opt,name=new_pack_name,json=newPackName,proto3,oneof" json:"new_pack_name,omitempty"`
}

func (x *EmotePackUpdated) Reset() {
	*x = EmotePackUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emote_v1_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmotePackUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmotePackUpdated) ProtoMessage() {}

func (x *EmotePackUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_emote_v1_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmotePackUpdated.ProtoReflect.Descriptor instead.
func (*EmotePackUpdated) Descriptor() ([]byte, []int) {
	return file_emote_v1_stream_proto_rawDescGZIP(), []int{0}
}

func (x *EmotePackUpdated) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

func (x *EmotePackUpdated) GetNewPackName() string {
	if x != nil && x.NewPackName != nil {
		return *x.NewPackName
	}
	return ""
}

// Event sent when an emote pack is deleted.
//
// Should only be sent to users who have the pack equipped.
// Should also be sent if a user dequips an emote pack, only to the user that dequipped it.
type EmotePackDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the pack that was deleted.
	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
}

func (x *EmotePackDeleted) Reset() {
	*x = EmotePackDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emote_v1_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmotePackDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmotePackDeleted) ProtoMessage() {}

func (x *EmotePackDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_emote_v1_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmotePackDeleted.ProtoReflect.Descriptor instead.
func (*EmotePackDeleted) Descriptor() ([]byte, []int) {
	return file_emote_v1_stream_proto_rawDescGZIP(), []int{1}
}

func (x *EmotePackDeleted) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

// Event sent when an emote pack is added.
//
// Should only be sent to the user who equipped the pack.
type EmotePackAdded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Emote pack that was equipped by the user.
	Pack *EmotePack `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty"`
}

func (x *EmotePackAdded) Reset() {
	*x = EmotePackAdded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emote_v1_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmotePackAdded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmotePackAdded) ProtoMessage() {}

func (x *EmotePackAdded) ProtoReflect() protoreflect.Message {
	mi := &file_emote_v1_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmotePackAdded.ProtoReflect.Descriptor instead.
func (*EmotePackAdded) Descriptor() ([]byte, []int) {
	return file_emote_v1_stream_proto_rawDescGZIP(), []int{2}
}

func (x *EmotePackAdded) GetPack() *EmotePack {
	if x != nil {
		return x.Pack
	}
	return nil
}

// Event sent when an emote pack's emotes were changed.
//
// Should only be sent to users who have the pack equipped.
type EmotePackEmotesUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the pack to update the emotes of.
	PackId uint64 `protobuf:"varint,1,opt,name=pack_id,json=packId,proto3" json:"pack_id,omitempty"`
	// The added emotes.
	AddedEmotes []*Emote `protobuf:"bytes,2,rep,name=added_emotes,json=addedEmotes,proto3" json:"added_emotes,omitempty"`
	// The names of the deleted emotes.
	DeletedEmotes []string `protobuf:"bytes,3,rep,name=deleted_emotes,json=deletedEmotes,proto3" json:"deleted_emotes,omitempty"`
}

func (x *EmotePackEmotesUpdated) Reset() {
	*x = EmotePackEmotesUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emote_v1_stream_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmotePackEmotesUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmotePackEmotesUpdated) ProtoMessage() {}

func (x *EmotePackEmotesUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_emote_v1_stream_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmotePackEmotesUpdated.ProtoReflect.Descriptor instead.
func (*EmotePackEmotesUpdated) Descriptor() ([]byte, []int) {
	return file_emote_v1_stream_proto_rawDescGZIP(), []int{3}
}

func (x *EmotePackEmotesUpdated) GetPackId() uint64 {
	if x != nil {
		return x.PackId
	}
	return 0
}

func (x *EmotePackEmotesUpdated) GetAddedEmotes() []*Emote {
	if x != nil {
		return x.AddedEmotes
	}
	return nil
}

func (x *EmotePackEmotesUpdated) GetDeletedEmotes() []string {
	if x != nil {
		return x.DeletedEmotes
	}
	return nil
}

// Describes an emote service event.
type StreamEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event type.
	//
	// Types that are assignable to Event:
	//	*StreamEvent_EmotePackAdded
	//	*StreamEvent_EmotePackUpdated
	//	*StreamEvent_EmotePackDeleted
	//	*StreamEvent_EmotePackEmotesUpdated
	Event isStreamEvent_Event `protobuf_oneof:"event"`
}

func (x *StreamEvent) Reset() {
	*x = StreamEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emote_v1_stream_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEvent) ProtoMessage() {}

func (x *StreamEvent) ProtoReflect() protoreflect.Message {
	mi := &file_emote_v1_stream_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEvent.ProtoReflect.Descriptor instead.
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return file_emote_v1_stream_proto_rawDescGZIP(), []int{4}
}

func (m *StreamEvent) GetEvent() isStreamEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamEvent) GetEmotePackAdded() *EmotePackAdded {
	if x, ok := x.GetEvent().(*StreamEvent_EmotePackAdded); ok {
		return x.EmotePackAdded
	}
	return nil
}

func (x *StreamEvent) GetEmotePackUpdated() *EmotePackUpdated {
	if x, ok := x.GetEvent().(*StreamEvent_EmotePackUpdated); ok {
		return x.EmotePackUpdated
	}
	return nil
}

func (x *StreamEvent) GetEmotePackDeleted() *EmotePackDeleted {
	if x, ok := x.GetEvent().(*StreamEvent_EmotePackDeleted); ok {
		return x.EmotePackDeleted
	}
	return nil
}

func (x *StreamEvent) GetEmotePackEmotesUpdated() *EmotePackEmotesUpdated {
	if x, ok := x.GetEvent().(*StreamEvent_EmotePackEmotesUpdated); ok {
		return x.EmotePackEmotesUpdated
	}
	return nil
}

type isStreamEvent_Event interface {
	isStreamEvent_Event()
}

type StreamEvent_EmotePackAdded struct {
	// Send the emote pack added event.
	EmotePackAdded *EmotePackAdded `protobuf:"bytes,1,opt,name=emote_pack_added,json=emotePackAdded,proto3,oneof"`
}

type StreamEvent_EmotePackUpdated struct {
	// Send the emote pack updated event.
	EmotePackUpdated *EmotePackUpdated `protobuf:"bytes,2,opt,name=emote_pack_updated,json=emotePackUpdated,proto3,oneof"`
}

type StreamEvent_EmotePackDeleted struct {
	// Send the emote pack deleted event.
	EmotePackDeleted *EmotePackDeleted `protobuf:"bytes,3,opt,name=emote_pack_deleted,json=emotePackDeleted,proto3,oneof"`
}

type StreamEvent_EmotePackEmotesUpdated struct {
	// Send the emote pack emotes updated event.
	EmotePackEmotesUpdated *EmotePackEmotesUpdated `protobuf:"bytes,4,opt,name=emote_pack_emotes_updated,json=emotePackEmotesUpdated,proto3,oneof"`
}

func (*StreamEvent_EmotePackAdded) isStreamEvent_Event() {}

func (*StreamEvent_EmotePackUpdated) isStreamEvent_Event() {}

func (*StreamEvent_EmotePackDeleted) isStreamEvent_Event() {}

func (*StreamEvent_EmotePackEmotesUpdated) isStreamEvent_Event() {}

var File_emote_v1_stream_proto protoreflect.FileDescriptor

var file_emote_v1_stream_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x10, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x10, 0x45, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70,
	0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61,
	0x63, 0x6b, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x52, 0x04, 0x70, 0x61, 0x63, 0x6b, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x45, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x3b, 0x0a,
	0x0c, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x0b, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x6f, 0x74, 0x65,
	0x73, 0x22, 0xf7, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x4d, 0x0a, 0x10, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x5f,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x53, 0x0a, 0x12, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x12, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50,
	0x61, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x66, 0x0a, 0x19, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74, 0x65,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x16, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x50, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0xc6, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x45, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x45, 0x6d, 0x6f, 0x74, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x45, 0x6d, 0x6f, 0x74, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emote_v1_stream_proto_rawDescOnce sync.Once
	file_emote_v1_stream_proto_rawDescData = file_emote_v1_stream_proto_rawDesc
)

func file_emote_v1_stream_proto_rawDescGZIP() []byte {
	file_emote_v1_stream_proto_rawDescOnce.Do(func() {
		file_emote_v1_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_emote_v1_stream_proto_rawDescData)
	})
	return file_emote_v1_stream_proto_rawDescData
}

var file_emote_v1_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_emote_v1_stream_proto_goTypes = []interface{}{
	(*EmotePackUpdated)(nil),       // 0: protocol.emote.v1.EmotePackUpdated
	(*EmotePackDeleted)(nil),       // 1: protocol.emote.v1.EmotePackDeleted
	(*EmotePackAdded)(nil),         // 2: protocol.emote.v1.EmotePackAdded
	(*EmotePackEmotesUpdated)(nil), // 3: protocol.emote.v1.EmotePackEmotesUpdated
	(*StreamEvent)(nil),            // 4: protocol.emote.v1.StreamEvent
	(*EmotePack)(nil),              // 5: protocol.emote.v1.EmotePack
	(*Emote)(nil),                  // 6: protocol.emote.v1.Emote
}
var file_emote_v1_stream_proto_depIdxs = []int32{
	5, // 0: protocol.emote.v1.EmotePackAdded.pack:type_name -> protocol.emote.v1.EmotePack
	6, // 1: protocol.emote.v1.EmotePackEmotesUpdated.added_emotes:type_name -> protocol.emote.v1.Emote
	2, // 2: protocol.emote.v1.StreamEvent.emote_pack_added:type_name -> protocol.emote.v1.EmotePackAdded
	0, // 3: protocol.emote.v1.StreamEvent.emote_pack_updated:type_name -> protocol.emote.v1.EmotePackUpdated
	1, // 4: protocol.emote.v1.StreamEvent.emote_pack_deleted:type_name -> protocol.emote.v1.EmotePackDeleted
	3, // 5: protocol.emote.v1.StreamEvent.emote_pack_emotes_updated:type_name -> protocol.emote.v1.EmotePackEmotesUpdated
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_emote_v1_stream_proto_init() }
func file_emote_v1_stream_proto_init() {
	if File_emote_v1_stream_proto != nil {
		return
	}
	file_emote_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_emote_v1_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmotePackUpdated); i {
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
		file_emote_v1_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmotePackDeleted); i {
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
		file_emote_v1_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmotePackAdded); i {
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
		file_emote_v1_stream_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmotePackEmotesUpdated); i {
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
		file_emote_v1_stream_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEvent); i {
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
	file_emote_v1_stream_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_emote_v1_stream_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*StreamEvent_EmotePackAdded)(nil),
		(*StreamEvent_EmotePackUpdated)(nil),
		(*StreamEvent_EmotePackDeleted)(nil),
		(*StreamEvent_EmotePackEmotesUpdated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_emote_v1_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_emote_v1_stream_proto_goTypes,
		DependencyIndexes: file_emote_v1_stream_proto_depIdxs,
		MessageInfos:      file_emote_v1_stream_proto_msgTypes,
	}.Build()
	File_emote_v1_stream_proto = out.File
	file_emote_v1_stream_proto_rawDesc = nil
	file_emote_v1_stream_proto_goTypes = nil
	file_emote_v1_stream_proto_depIdxs = nil
}
