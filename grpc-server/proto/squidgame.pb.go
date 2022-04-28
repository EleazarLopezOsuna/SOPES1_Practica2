// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/squidgame.proto

package proto

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

type MatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId        int32 `protobuf:"varint,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
	NumberPlayers int32 `protobuf:"varint,2,opt,name=numberPlayers,proto3" json:"numberPlayers,omitempty"`
}

func (x *MatchRequest) Reset() {
	*x = MatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchRequest) ProtoMessage() {}

func (x *MatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchRequest.ProtoReflect.Descriptor instead.
func (*MatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{0}
}

func (x *MatchRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *MatchRequest) GetNumberPlayers() int32 {
	if x != nil {
		return x.NumberPlayers
	}
	return 0
}

type MatchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MatchReply) Reset() {
	*x = MatchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_squidgame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchReply) ProtoMessage() {}

func (x *MatchReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_squidgame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchReply.ProtoReflect.Descriptor instead.
func (*MatchReply) Descriptor() ([]byte, []int) {
	return file_proto_squidgame_proto_rawDescGZIP(), []int{1}
}

func (x *MatchReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_squidgame_proto protoreflect.FileDescriptor

var file_proto_squidgame_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x22, 0x26, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x17, 0x2e, 0x73, 0x71, 0x75, 0x69, 0x64, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x71, 0x75, 0x69, 0x64,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x6c, 0x6c, 0x56, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x31, 0x5f, 0x70, 0x72, 0x6f,
	0x79, 0x65, 0x63, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x73, 0x65, 0x32, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x32, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_squidgame_proto_rawDescOnce sync.Once
	file_proto_squidgame_proto_rawDescData = file_proto_squidgame_proto_rawDesc
)

func file_proto_squidgame_proto_rawDescGZIP() []byte {
	file_proto_squidgame_proto_rawDescOnce.Do(func() {
		file_proto_squidgame_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_squidgame_proto_rawDescData)
	})
	return file_proto_squidgame_proto_rawDescData
}

var file_proto_squidgame_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_squidgame_proto_goTypes = []interface{}{
	(*MatchRequest)(nil), // 0: squidgame.MatchRequest
	(*MatchReply)(nil),   // 1: squidgame.MatchReply
}
var file_proto_squidgame_proto_depIdxs = []int32{
	0, // 0: squidgame.Matches.AddMatch:input_type -> squidgame.MatchRequest
	1, // 1: squidgame.Matches.AddMatch:output_type -> squidgame.MatchReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_squidgame_proto_init() }
func file_proto_squidgame_proto_init() {
	if File_proto_squidgame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_squidgame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchRequest); i {
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
		file_proto_squidgame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchReply); i {
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
			RawDescriptor: file_proto_squidgame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_squidgame_proto_goTypes,
		DependencyIndexes: file_proto_squidgame_proto_depIdxs,
		MessageInfos:      file_proto_squidgame_proto_msgTypes,
	}.Build()
	File_proto_squidgame_proto = out.File
	file_proto_squidgame_proto_rawDesc = nil
	file_proto_squidgame_proto_goTypes = nil
	file_proto_squidgame_proto_depIdxs = nil
}