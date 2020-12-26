// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: mediaproxy/v1/mediaproxy.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SiteMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteTitle   string `protobuf:"bytes,1,opt,name=site_title,json=siteTitle,proto3" json:"site_title,omitempty"`
	PageTitle   string `protobuf:"bytes,2,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	Kind        string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *SiteMetadata) Reset() {
	*x = SiteMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteMetadata) ProtoMessage() {}

func (x *SiteMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteMetadata.ProtoReflect.Descriptor instead.
func (*SiteMetadata) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{0}
}

func (x *SiteMetadata) GetSiteTitle() string {
	if x != nil {
		return x.SiteTitle
	}
	return ""
}

func (x *SiteMetadata) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *SiteMetadata) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SiteMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SiteMetadata) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SiteMetadata) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type FetchLinkMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchLinkMetadataRequest) Reset() {
	*x = FetchLinkMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLinkMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLinkMetadataRequest) ProtoMessage() {}

func (x *FetchLinkMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLinkMetadataRequest.ProtoReflect.Descriptor instead.
func (*FetchLinkMetadataRequest) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{1}
}

func (x *FetchLinkMetadataRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type InstantViewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *InstantViewRequest) Reset() {
	*x = InstantViewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantViewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantViewRequest) ProtoMessage() {}

func (x *InstantViewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantViewRequest.ProtoReflect.Descriptor instead.
func (*InstantViewRequest) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{2}
}

func (x *InstantViewRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type InstantViewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *SiteMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Content  string        `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *InstantViewResponse) Reset() {
	*x = InstantViewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantViewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantViewResponse) ProtoMessage() {}

func (x *InstantViewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mediaproxy_v1_mediaproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantViewResponse.ProtoReflect.Descriptor instead.
func (*InstantViewResponse) Descriptor() ([]byte, []int) {
	return file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP(), []int{3}
}

func (x *InstantViewResponse) GetMetadata() *SiteMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *InstantViewResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_mediaproxy_v1_mediaproxy_proto protoreflect.FileDescriptor

var file_mediaproxy_v1_mediaproxy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xaa, 0x01, 0x0a, 0x0c, 0x53, 0x69, 0x74,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69,
	0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x26, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x71, 0x0a, 0x13, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xe8,
	0x01, 0x0a, 0x11, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e,
	0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x66, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mediaproxy_v1_mediaproxy_proto_rawDescOnce sync.Once
	file_mediaproxy_v1_mediaproxy_proto_rawDescData = file_mediaproxy_v1_mediaproxy_proto_rawDesc
)

func file_mediaproxy_v1_mediaproxy_proto_rawDescGZIP() []byte {
	file_mediaproxy_v1_mediaproxy_proto_rawDescOnce.Do(func() {
		file_mediaproxy_v1_mediaproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_mediaproxy_v1_mediaproxy_proto_rawDescData)
	})
	return file_mediaproxy_v1_mediaproxy_proto_rawDescData
}

var file_mediaproxy_v1_mediaproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mediaproxy_v1_mediaproxy_proto_goTypes = []interface{}{
	(*SiteMetadata)(nil),             // 0: protocol.mediaproxy.v1.SiteMetadata
	(*FetchLinkMetadataRequest)(nil), // 1: protocol.mediaproxy.v1.FetchLinkMetadataRequest
	(*InstantViewRequest)(nil),       // 2: protocol.mediaproxy.v1.InstantViewRequest
	(*InstantViewResponse)(nil),      // 3: protocol.mediaproxy.v1.InstantViewResponse
}
var file_mediaproxy_v1_mediaproxy_proto_depIdxs = []int32{
	0, // 0: protocol.mediaproxy.v1.InstantViewResponse.metadata:type_name -> protocol.mediaproxy.v1.SiteMetadata
	1, // 1: protocol.mediaproxy.v1.MediaProxyService.FetchLinkMetadata:input_type -> protocol.mediaproxy.v1.FetchLinkMetadataRequest
	2, // 2: protocol.mediaproxy.v1.MediaProxyService.InstantView:input_type -> protocol.mediaproxy.v1.InstantViewRequest
	0, // 3: protocol.mediaproxy.v1.MediaProxyService.FetchLinkMetadata:output_type -> protocol.mediaproxy.v1.SiteMetadata
	3, // 4: protocol.mediaproxy.v1.MediaProxyService.InstantView:output_type -> protocol.mediaproxy.v1.InstantViewResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mediaproxy_v1_mediaproxy_proto_init() }
func file_mediaproxy_v1_mediaproxy_proto_init() {
	if File_mediaproxy_v1_mediaproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteMetadata); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLinkMetadataRequest); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantViewRequest); i {
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
		file_mediaproxy_v1_mediaproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantViewResponse); i {
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
			RawDescriptor: file_mediaproxy_v1_mediaproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mediaproxy_v1_mediaproxy_proto_goTypes,
		DependencyIndexes: file_mediaproxy_v1_mediaproxy_proto_depIdxs,
		MessageInfos:      file_mediaproxy_v1_mediaproxy_proto_msgTypes,
	}.Build()
	File_mediaproxy_v1_mediaproxy_proto = out.File
	file_mediaproxy_v1_mediaproxy_proto_rawDesc = nil
	file_mediaproxy_v1_mediaproxy_proto_goTypes = nil
	file_mediaproxy_v1_mediaproxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MediaProxyServiceClient is the client API for MediaProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaProxyServiceClient interface {
	FetchLinkMetadata(ctx context.Context, in *FetchLinkMetadataRequest, opts ...grpc.CallOption) (*SiteMetadata, error)
	InstantView(ctx context.Context, in *InstantViewRequest, opts ...grpc.CallOption) (*InstantViewResponse, error)
}

type mediaProxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaProxyServiceClient(cc grpc.ClientConnInterface) MediaProxyServiceClient {
	return &mediaProxyServiceClient{cc}
}

func (c *mediaProxyServiceClient) FetchLinkMetadata(ctx context.Context, in *FetchLinkMetadataRequest, opts ...grpc.CallOption) (*SiteMetadata, error) {
	out := new(SiteMetadata)
	err := c.cc.Invoke(ctx, "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaProxyServiceClient) InstantView(ctx context.Context, in *InstantViewRequest, opts ...grpc.CallOption) (*InstantViewResponse, error) {
	out := new(InstantViewResponse)
	err := c.cc.Invoke(ctx, "/protocol.mediaproxy.v1.MediaProxyService/InstantView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaProxyServiceServer is the server API for MediaProxyService service.
type MediaProxyServiceServer interface {
	FetchLinkMetadata(context.Context, *FetchLinkMetadataRequest) (*SiteMetadata, error)
	InstantView(context.Context, *InstantViewRequest) (*InstantViewResponse, error)
}

// UnimplementedMediaProxyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMediaProxyServiceServer struct {
}

func (*UnimplementedMediaProxyServiceServer) FetchLinkMetadata(context.Context, *FetchLinkMetadataRequest) (*SiteMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLinkMetadata not implemented")
}
func (*UnimplementedMediaProxyServiceServer) InstantView(context.Context, *InstantViewRequest) (*InstantViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstantView not implemented")
}

func RegisterMediaProxyServiceServer(s *grpc.Server, srv MediaProxyServiceServer) {
	s.RegisterService(&_MediaProxyService_serviceDesc, srv)
}

func _MediaProxyService_FetchLinkMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLinkMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaProxyServiceServer).FetchLinkMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaProxyServiceServer).FetchLinkMetadata(ctx, req.(*FetchLinkMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaProxyService_InstantView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstantViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaProxyServiceServer).InstantView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.mediaproxy.v1.MediaProxyService/InstantView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaProxyServiceServer).InstantView(ctx, req.(*InstantViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.mediaproxy.v1.MediaProxyService",
	HandlerType: (*MediaProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchLinkMetadata",
			Handler:    _MediaProxyService_FetchLinkMetadata_Handler,
		},
		{
			MethodName: "InstantView",
			Handler:    _MediaProxyService_InstantView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mediaproxy/v1/mediaproxy.proto",
}
