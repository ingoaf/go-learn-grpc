// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: max.proto

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

type MaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *MaxRequest) Reset() {
	*x = MaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_max_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxRequest) ProtoMessage() {}

func (x *MaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_max_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxRequest.ProtoReflect.Descriptor instead.
func (*MaxRequest) Descriptor() ([]byte, []int) {
	return file_max_proto_rawDescGZIP(), []int{0}
}

func (x *MaxRequest) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum float64 `protobuf:"fixed64,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *MaxResponse) Reset() {
	*x = MaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_max_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxResponse) ProtoMessage() {}

func (x *MaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_max_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxResponse.ProtoReflect.Descriptor instead.
func (*MaxResponse) Descriptor() ([]byte, []int) {
	return file_max_proto_rawDescGZIP(), []int{1}
}

func (x *MaxResponse) GetMaximum() float64 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_max_proto protoreflect.FileDescriptor

var file_max_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x0a, 0x4d, 0x61, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x27, 0x0a,
	0x0b, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x67, 0x6f, 0x61, 0x66, 0x2f, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x61, 0x6e,
	0x64, 0x73, 0x2d, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_max_proto_rawDescOnce sync.Once
	file_max_proto_rawDescData = file_max_proto_rawDesc
)

func file_max_proto_rawDescGZIP() []byte {
	file_max_proto_rawDescOnce.Do(func() {
		file_max_proto_rawDescData = protoimpl.X.CompressGZIP(file_max_proto_rawDescData)
	})
	return file_max_proto_rawDescData
}

var file_max_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_max_proto_goTypes = []interface{}{
	(*MaxRequest)(nil),  // 0: calculator.MaxRequest
	(*MaxResponse)(nil), // 1: calculator.MaxResponse
}
var file_max_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_max_proto_init() }
func file_max_proto_init() {
	if File_max_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_max_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxRequest); i {
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
		file_max_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxResponse); i {
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
			RawDescriptor: file_max_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_max_proto_goTypes,
		DependencyIndexes: file_max_proto_depIdxs,
		MessageInfos:      file_max_proto_msgTypes,
	}.Build()
	File_max_proto = out.File
	file_max_proto_rawDesc = nil
	file_max_proto_goTypes = nil
	file_max_proto_depIdxs = nil
}
