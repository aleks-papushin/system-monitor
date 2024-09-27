// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: api/stat.proto

package gen

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

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N int32 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	M int32 `protobuf:"varint,2,opt,name=m,proto3" json:"m,omitempty"`
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_api_stat_proto_rawDescGZIP(), []int{0}
}

func (x *StatsRequest) GetN() int32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *StatsRequest) GetM() int32 {
	if x != nil {
		return x.M
	}
	return 0
}

type StatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoadAverage float32 `protobuf:"fixed32,1,opt,name=loadAverage,proto3" json:"loadAverage,omitempty"`
	UserUsage   float32 `protobuf:"fixed32,2,opt,name=userUsage,proto3" json:"userUsage,omitempty"`
	SysUsage    float32 `protobuf:"fixed32,3,opt,name=sysUsage,proto3" json:"sysUsage,omitempty"`
	Idle        float32 `protobuf:"fixed32,4,opt,name=idle,proto3" json:"idle,omitempty"`
	Timestamp   string  `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StatsResponse) Reset() {
	*x = StatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsResponse) ProtoMessage() {}

func (x *StatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsResponse.ProtoReflect.Descriptor instead.
func (*StatsResponse) Descriptor() ([]byte, []int) {
	return file_api_stat_proto_rawDescGZIP(), []int{1}
}

func (x *StatsResponse) GetLoadAverage() float32 {
	if x != nil {
		return x.LoadAverage
	}
	return 0
}

func (x *StatsResponse) GetUserUsage() float32 {
	if x != nil {
		return x.UserUsage
	}
	return 0
}

func (x *StatsResponse) GetSysUsage() float32 {
	if x != nil {
		return x.SysUsage
	}
	return 0
}

func (x *StatsResponse) GetIdle() float32 {
	if x != nil {
		return x.Idle
	}
	return 0
}

func (x *StatsResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_api_stat_proto protoreflect.FileDescriptor

var file_api_stat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x73, 0x74, 0x61, 0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x6d, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x73, 0x79, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04,
	0x69, 0x64, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x32, 0x44, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x12, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_stat_proto_rawDescOnce sync.Once
	file_api_stat_proto_rawDescData = file_api_stat_proto_rawDesc
)

func file_api_stat_proto_rawDescGZIP() []byte {
	file_api_stat_proto_rawDescOnce.Do(func() {
		file_api_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_stat_proto_rawDescData)
	})
	return file_api_stat_proto_rawDescData
}

var file_api_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_stat_proto_goTypes = []any{
	(*StatsRequest)(nil),  // 0: stat.StatsRequest
	(*StatsResponse)(nil), // 1: stat.StatsResponse
}
var file_api_stat_proto_depIdxs = []int32{
	0, // 0: stat.StatService.GetStats:input_type -> stat.StatsRequest
	1, // 1: stat.StatService.GetStats:output_type -> stat.StatsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_stat_proto_init() }
func file_api_stat_proto_init() {
	if File_api_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_stat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatsRequest); i {
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
		file_api_stat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatsResponse); i {
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
			RawDescriptor: file_api_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_stat_proto_goTypes,
		DependencyIndexes: file_api_stat_proto_depIdxs,
		MessageInfos:      file_api_stat_proto_msgTypes,
	}.Build()
	File_api_stat_proto = out.File
	file_api_stat_proto_rawDesc = nil
	file_api_stat_proto_goTypes = nil
	file_api_stat_proto_depIdxs = nil
}
