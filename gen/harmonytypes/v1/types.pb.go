// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: harmonytypes/v1/types.proto

package harmonytypesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Metadata for methods. These are set in individual RPC endpoints and are
// typically used by servers.
type HarmonyMethodMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// whether the method requires authentication.
	RequiresAuthentication bool `protobuf:"varint,1,opt,name=requires_authentication,json=requiresAuthentication,proto3" json:"requires_authentication,omitempty"`
	// whether the method allows federation or not.
	RequiresLocal bool `protobuf:"varint,2,opt,name=requires_local,json=requiresLocal,proto3" json:"requires_local,omitempty"`
	// the permission nodes required to invoke the method.
	RequiresPermissionNode string `protobuf:"bytes,3,opt,name=requires_permission_node,json=requiresPermissionNode,proto3" json:"requires_permission_node,omitempty"`
}

func (x *HarmonyMethodMetadata) Reset() {
	*x = HarmonyMethodMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HarmonyMethodMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HarmonyMethodMetadata) ProtoMessage() {}

func (x *HarmonyMethodMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HarmonyMethodMetadata.ProtoReflect.Descriptor instead.
func (*HarmonyMethodMetadata) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *HarmonyMethodMetadata) GetRequiresAuthentication() bool {
	if x != nil {
		return x.RequiresAuthentication
	}
	return false
}

func (x *HarmonyMethodMetadata) GetRequiresLocal() bool {
	if x != nil {
		return x.RequiresLocal
	}
	return false
}

func (x *HarmonyMethodMetadata) GetRequiresPermissionNode() string {
	if x != nil {
		return x.RequiresPermissionNode
	}
	return ""
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind      string                `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Extension map[string]*anypb.Any `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Metadata) GetExtension() map[string]*anypb.Any {
	if x != nil {
		return x.Extension
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier   string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	HumanMessage string `protobuf:"bytes,2,opt,name=human_message,json=humanMessage,proto3" json:"human_message,omitempty"`
	MoreDetails  []byte `protobuf:"bytes,3,opt,name=more_details,json=moreDetails,proto3" json:"more_details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Error) GetHumanMessage() string {
	if x != nil {
		return x.HumanMessage
	}
	return ""
}

func (x *Error) GetMoreDetails() []byte {
	if x != nil {
		return x.MoreDetails
	}
	return nil
}

// Token that will be used for authentication.
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ed25519 signature of the following serialized protobuf data, signed
	// with a private key. Which private key used to sign will be described
	// in the documentation.
	//
	// Has to be 64 bytes long, otherwise it will be rejected.
	Sig []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	// Serialized protobuf data.
	// The protobuf type of this serialized data is dependent on the API endpoint
	// used.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *Token) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// An object representing an item position between two other items.
type ItemPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Position:
	//	*ItemPosition_Top_
	//	*ItemPosition_Between_
	//	*ItemPosition_Bottom_
	Position isItemPosition_Position `protobuf_oneof:"position"`
}

func (x *ItemPosition) Reset() {
	*x = ItemPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPosition) ProtoMessage() {}

func (x *ItemPosition) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPosition.ProtoReflect.Descriptor instead.
func (*ItemPosition) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{4}
}

func (m *ItemPosition) GetPosition() isItemPosition_Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (x *ItemPosition) GetTop() *ItemPosition_Top {
	if x, ok := x.GetPosition().(*ItemPosition_Top_); ok {
		return x.Top
	}
	return nil
}

func (x *ItemPosition) GetBetween() *ItemPosition_Between {
	if x, ok := x.GetPosition().(*ItemPosition_Between_); ok {
		return x.Between
	}
	return nil
}

func (x *ItemPosition) GetBottom() *ItemPosition_Bottom {
	if x, ok := x.GetPosition().(*ItemPosition_Bottom_); ok {
		return x.Bottom
	}
	return nil
}

type isItemPosition_Position interface {
	isItemPosition_Position()
}

type ItemPosition_Top_ struct {
	Top *ItemPosition_Top `protobuf:"bytes,1,opt,name=top,proto3,oneof"`
}

type ItemPosition_Between_ struct {
	Between *ItemPosition_Between `protobuf:"bytes,2,opt,name=between,proto3,oneof"`
}

type ItemPosition_Bottom_ struct {
	Bottom *ItemPosition_Bottom `protobuf:"bytes,3,opt,name=bottom,proto3,oneof"`
}

func (*ItemPosition_Top_) isItemPosition_Position() {}

func (*ItemPosition_Between_) isItemPosition_Position() {}

func (*ItemPosition_Bottom_) isItemPosition_Position() {}

// An object that represents the top of an ordered list.
type ItemPosition_Top struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextId uint64 `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
}

func (x *ItemPosition_Top) Reset() {
	*x = ItemPosition_Top{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPosition_Top) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPosition_Top) ProtoMessage() {}

func (x *ItemPosition_Top) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPosition_Top.ProtoReflect.Descriptor instead.
func (*ItemPosition_Top) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ItemPosition_Top) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

// An object that represents a place between two items in an ordered list.
type ItemPosition_Between struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousId uint64 `protobuf:"varint,1,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	NextId     uint64 `protobuf:"varint,2,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
}

func (x *ItemPosition_Between) Reset() {
	*x = ItemPosition_Between{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPosition_Between) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPosition_Between) ProtoMessage() {}

func (x *ItemPosition_Between) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPosition_Between.ProtoReflect.Descriptor instead.
func (*ItemPosition_Between) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{4, 1}
}

func (x *ItemPosition_Between) GetPreviousId() uint64 {
	if x != nil {
		return x.PreviousId
	}
	return 0
}

func (x *ItemPosition_Between) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

// An object that represents the bottom of an ordered list.
type ItemPosition_Bottom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousId uint64 `protobuf:"varint,1,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
}

func (x *ItemPosition_Bottom) Reset() {
	*x = ItemPosition_Bottom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_harmonytypes_v1_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemPosition_Bottom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemPosition_Bottom) ProtoMessage() {}

func (x *ItemPosition_Bottom) ProtoReflect() protoreflect.Message {
	mi := &file_harmonytypes_v1_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemPosition_Bottom.ProtoReflect.Descriptor instead.
func (*ItemPosition_Bottom) Descriptor() ([]byte, []int) {
	return file_harmonytypes_v1_types_proto_rawDescGZIP(), []int{4, 2}
}

func (x *ItemPosition_Bottom) GetPreviousId() uint64 {
	if x != nil {
		return x.PreviousId
	}
	return 0
}

var file_harmonytypes_v1_types_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*HarmonyMethodMetadata)(nil),
		Field:         1091,
		Name:          "protocol.harmonytypes.v1.metadata",
		Tag:           "bytes,1091,opt,name=metadata",
		Filename:      "harmonytypes/v1/types.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional protocol.harmonytypes.v1.HarmonyMethodMetadata metadata = 1091;
	E_Metadata = &file_harmonytypes_v1_types_proto_extTypes[0]
)

var File_harmonytypes_v1_types_proto protoreflect.FileDescriptor

var file_harmonytypes_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x48, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37,
	0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x38,
	0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x4f, 0x0a, 0x09, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x52, 0x0a, 0x0e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6f,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x68, 0x75, 0x6d, 0x61, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x6f, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x2d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xff,
	0x02, 0x0a, 0x0c, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x48, 0x00, 0x52, 0x03, 0x74, 0x6f, 0x70, 0x12,
	0x4a, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d,
	0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e,
	0x48, 0x00, 0x52, 0x07, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x62,
	0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x48, 0x00, 0x52, 0x06, 0x62, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x1a, 0x1e, 0x0a, 0x03, 0x54, 0x6f, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x65,
	0x78, 0x74, 0x49, 0x64, 0x1a, 0x43, 0x0a, 0x07, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x1a, 0x29, 0x0a, 0x06, 0x42, 0x6f, 0x74,
	0x74, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x6c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x68,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xf6,
	0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e,
	0x79, 0x2d, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x48, 0x58, 0xaa, 0x02, 0x18,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5c, 0x48, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x48,
	0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x48, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_harmonytypes_v1_types_proto_rawDescOnce sync.Once
	file_harmonytypes_v1_types_proto_rawDescData = file_harmonytypes_v1_types_proto_rawDesc
)

func file_harmonytypes_v1_types_proto_rawDescGZIP() []byte {
	file_harmonytypes_v1_types_proto_rawDescOnce.Do(func() {
		file_harmonytypes_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_harmonytypes_v1_types_proto_rawDescData)
	})
	return file_harmonytypes_v1_types_proto_rawDescData
}

var file_harmonytypes_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_harmonytypes_v1_types_proto_goTypes = []interface{}{
	(*HarmonyMethodMetadata)(nil),      // 0: protocol.harmonytypes.v1.HarmonyMethodMetadata
	(*Metadata)(nil),                   // 1: protocol.harmonytypes.v1.Metadata
	(*Error)(nil),                      // 2: protocol.harmonytypes.v1.Error
	(*Token)(nil),                      // 3: protocol.harmonytypes.v1.Token
	(*ItemPosition)(nil),               // 4: protocol.harmonytypes.v1.ItemPosition
	nil,                                // 5: protocol.harmonytypes.v1.Metadata.ExtensionEntry
	(*ItemPosition_Top)(nil),           // 6: protocol.harmonytypes.v1.ItemPosition.Top
	(*ItemPosition_Between)(nil),       // 7: protocol.harmonytypes.v1.ItemPosition.Between
	(*ItemPosition_Bottom)(nil),        // 8: protocol.harmonytypes.v1.ItemPosition.Bottom
	(*anypb.Any)(nil),                  // 9: google.protobuf.Any
	(*descriptorpb.MethodOptions)(nil), // 10: google.protobuf.MethodOptions
}
var file_harmonytypes_v1_types_proto_depIdxs = []int32{
	5,  // 0: protocol.harmonytypes.v1.Metadata.extension:type_name -> protocol.harmonytypes.v1.Metadata.ExtensionEntry
	6,  // 1: protocol.harmonytypes.v1.ItemPosition.top:type_name -> protocol.harmonytypes.v1.ItemPosition.Top
	7,  // 2: protocol.harmonytypes.v1.ItemPosition.between:type_name -> protocol.harmonytypes.v1.ItemPosition.Between
	8,  // 3: protocol.harmonytypes.v1.ItemPosition.bottom:type_name -> protocol.harmonytypes.v1.ItemPosition.Bottom
	9,  // 4: protocol.harmonytypes.v1.Metadata.ExtensionEntry.value:type_name -> google.protobuf.Any
	10, // 5: protocol.harmonytypes.v1.metadata:extendee -> google.protobuf.MethodOptions
	0,  // 6: protocol.harmonytypes.v1.metadata:type_name -> protocol.harmonytypes.v1.HarmonyMethodMetadata
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	6,  // [6:7] is the sub-list for extension type_name
	5,  // [5:6] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_harmonytypes_v1_types_proto_init() }
func file_harmonytypes_v1_types_proto_init() {
	if File_harmonytypes_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_harmonytypes_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HarmonyMethodMetadata); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPosition); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPosition_Top); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPosition_Between); i {
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
		file_harmonytypes_v1_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemPosition_Bottom); i {
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
	file_harmonytypes_v1_types_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ItemPosition_Top_)(nil),
		(*ItemPosition_Between_)(nil),
		(*ItemPosition_Bottom_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_harmonytypes_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_harmonytypes_v1_types_proto_goTypes,
		DependencyIndexes: file_harmonytypes_v1_types_proto_depIdxs,
		MessageInfos:      file_harmonytypes_v1_types_proto_msgTypes,
		ExtensionInfos:    file_harmonytypes_v1_types_proto_extTypes,
	}.Build()
	File_harmonytypes_v1_types_proto = out.File
	file_harmonytypes_v1_types_proto_rawDesc = nil
	file_harmonytypes_v1_types_proto_goTypes = nil
	file_harmonytypes_v1_types_proto_depIdxs = nil
}
