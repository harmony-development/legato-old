// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: voice/v1/voice.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type ClientSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*ClientSignal_Answer_
	//	*ClientSignal_Candidate_
	Event isClientSignal_Event `protobuf_oneof:"event"`
}

func (x *ClientSignal) Reset() {
	*x = ClientSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSignal) ProtoMessage() {}

func (x *ClientSignal) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSignal.ProtoReflect.Descriptor instead.
func (*ClientSignal) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{0}
}

func (m *ClientSignal) GetEvent() isClientSignal_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ClientSignal) GetAnswer() *ClientSignal_Answer {
	if x, ok := x.GetEvent().(*ClientSignal_Answer_); ok {
		return x.Answer
	}
	return nil
}

func (x *ClientSignal) GetCandidate() *ClientSignal_Candidate {
	if x, ok := x.GetEvent().(*ClientSignal_Candidate_); ok {
		return x.Candidate
	}
	return nil
}

type isClientSignal_Event interface {
	isClientSignal_Event()
}

type ClientSignal_Answer_ struct {
	Answer *ClientSignal_Answer `protobuf:"bytes,1,opt,name=answer,proto3,oneof"`
}

type ClientSignal_Candidate_ struct {
	Candidate *ClientSignal_Candidate `protobuf:"bytes,2,opt,name=candidate,proto3,oneof"`
}

func (*ClientSignal_Answer_) isClientSignal_Event() {}

func (*ClientSignal_Candidate_) isClientSignal_Event() {}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*Signal_Candidate
	//	*Signal_Offer_
	Event isSignal_Event `protobuf_oneof:"event"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{1}
}

func (m *Signal) GetEvent() isSignal_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *Signal) GetCandidate() *Signal_ICECandidate {
	if x, ok := x.GetEvent().(*Signal_Candidate); ok {
		return x.Candidate
	}
	return nil
}

func (x *Signal) GetOffer() *Signal_Offer {
	if x, ok := x.GetEvent().(*Signal_Offer_); ok {
		return x.Offer
	}
	return nil
}

type isSignal_Event interface {
	isSignal_Event()
}

type Signal_Candidate struct {
	Candidate *Signal_ICECandidate `protobuf:"bytes,1,opt,name=candidate,proto3,oneof"`
}

type Signal_Offer_ struct {
	Offer *Signal_Offer `protobuf:"bytes,2,opt,name=offer,proto3,oneof"`
}

func (*Signal_Candidate) isSignal_Event() {}

func (*Signal_Offer_) isSignal_Event() {}

type ClientSignal_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *ClientSignal_Answer) Reset() {
	*x = ClientSignal_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSignal_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSignal_Answer) ProtoMessage() {}

func (x *ClientSignal_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSignal_Answer.ProtoReflect.Descriptor instead.
func (*ClientSignal_Answer) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientSignal_Answer) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type ClientSignal_Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate string `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *ClientSignal_Candidate) Reset() {
	*x = ClientSignal_Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientSignal_Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientSignal_Candidate) ProtoMessage() {}

func (x *ClientSignal_Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientSignal_Candidate.ProtoReflect.Descriptor instead.
func (*ClientSignal_Candidate) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ClientSignal_Candidate) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

type Signal_ICECandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate string `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *Signal_ICECandidate) Reset() {
	*x = Signal_ICECandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal_ICECandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal_ICECandidate) ProtoMessage() {}

func (x *Signal_ICECandidate) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal_ICECandidate.ProtoReflect.Descriptor instead.
func (*Signal_ICECandidate) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Signal_ICECandidate) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

type Signal_Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offer string `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *Signal_Offer) Reset() {
	*x = Signal_Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voice_v1_voice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal_Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal_Offer) ProtoMessage() {}

func (x *Signal_Offer) ProtoReflect() protoreflect.Message {
	mi := &file_voice_v1_voice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal_Offer.ProtoReflect.Descriptor instead.
func (*Signal_Offer) Descriptor() ([]byte, []int) {
	return file_voice_v1_voice_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Signal_Offer) GetOffer() string {
	if x != nil {
		return x.Offer
	}
	return ""
}

var File_voice_v1_voice_proto protoreflect.FileDescriptor

var file_voice_v1_voice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x09, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x1a, 0x20, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x1a, 0x29, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xdf, 0x01, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x2e, 0x49, 0x43, 0x45, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x00, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a,
	0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x2c, 0x0a, 0x0c, 0x49, 0x43, 0x45, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x1a, 0x1d, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x5b, 0x0a, 0x0c,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voice_v1_voice_proto_rawDescOnce sync.Once
	file_voice_v1_voice_proto_rawDescData = file_voice_v1_voice_proto_rawDesc
)

func file_voice_v1_voice_proto_rawDescGZIP() []byte {
	file_voice_v1_voice_proto_rawDescOnce.Do(func() {
		file_voice_v1_voice_proto_rawDescData = protoimpl.X.CompressGZIP(file_voice_v1_voice_proto_rawDescData)
	})
	return file_voice_v1_voice_proto_rawDescData
}

var file_voice_v1_voice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_voice_v1_voice_proto_goTypes = []interface{}{
	(*ClientSignal)(nil),           // 0: protocol.voice.v1.ClientSignal
	(*Signal)(nil),                 // 1: protocol.voice.v1.Signal
	(*ClientSignal_Answer)(nil),    // 2: protocol.voice.v1.ClientSignal.Answer
	(*ClientSignal_Candidate)(nil), // 3: protocol.voice.v1.ClientSignal.Candidate
	(*Signal_ICECandidate)(nil),    // 4: protocol.voice.v1.Signal.ICECandidate
	(*Signal_Offer)(nil),           // 5: protocol.voice.v1.Signal.Offer
}
var file_voice_v1_voice_proto_depIdxs = []int32{
	2, // 0: protocol.voice.v1.ClientSignal.answer:type_name -> protocol.voice.v1.ClientSignal.Answer
	3, // 1: protocol.voice.v1.ClientSignal.candidate:type_name -> protocol.voice.v1.ClientSignal.Candidate
	4, // 2: protocol.voice.v1.Signal.candidate:type_name -> protocol.voice.v1.Signal.ICECandidate
	5, // 3: protocol.voice.v1.Signal.offer:type_name -> protocol.voice.v1.Signal.Offer
	0, // 4: protocol.voice.v1.VoiceService.Connect:input_type -> protocol.voice.v1.ClientSignal
	1, // 5: protocol.voice.v1.VoiceService.Connect:output_type -> protocol.voice.v1.Signal
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_voice_v1_voice_proto_init() }
func file_voice_v1_voice_proto_init() {
	if File_voice_v1_voice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voice_v1_voice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSignal); i {
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
		file_voice_v1_voice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
		file_voice_v1_voice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSignal_Answer); i {
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
		file_voice_v1_voice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientSignal_Candidate); i {
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
		file_voice_v1_voice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal_ICECandidate); i {
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
		file_voice_v1_voice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal_Offer); i {
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
	file_voice_v1_voice_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientSignal_Answer_)(nil),
		(*ClientSignal_Candidate_)(nil),
	}
	file_voice_v1_voice_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Signal_Candidate)(nil),
		(*Signal_Offer_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voice_v1_voice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voice_v1_voice_proto_goTypes,
		DependencyIndexes: file_voice_v1_voice_proto_depIdxs,
		MessageInfos:      file_voice_v1_voice_proto_msgTypes,
	}.Build()
	File_voice_v1_voice_proto = out.File
	file_voice_v1_voice_proto_rawDesc = nil
	file_voice_v1_voice_proto_goTypes = nil
	file_voice_v1_voice_proto_depIdxs = nil
}
