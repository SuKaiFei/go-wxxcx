// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: wxxcx/v1/common.proto

package v1

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

type Share struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Share) Reset() {
	*x = Share{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Share) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Share) ProtoMessage() {}

func (x *Share) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Share.ProtoReflect.Descriptor instead.
func (*Share) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Share) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Share) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_wxxcx_v1_common_proto protoreflect.FileDescriptor

var file_wxxcx_v1_common_proto_rawDesc = []byte{
	0x0a, 0x15, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x75, 0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77,
	0x78, 0x78, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_common_proto_rawDescOnce sync.Once
	file_wxxcx_v1_common_proto_rawDescData = file_wxxcx_v1_common_proto_rawDesc
)

func file_wxxcx_v1_common_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_common_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_common_proto_rawDescData)
	})
	return file_wxxcx_v1_common_proto_rawDescData
}

var file_wxxcx_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wxxcx_v1_common_proto_goTypes = []interface{}{
	(*Share)(nil), // 0: wxxcx.v1.common.Share
}
var file_wxxcx_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_common_proto_init() }
func file_wxxcx_v1_common_proto_init() {
	if File_wxxcx_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Share); i {
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
			RawDescriptor: file_wxxcx_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wxxcx_v1_common_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_common_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_common_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_common_proto = out.File
	file_wxxcx_v1_common_proto_rawDesc = nil
	file_wxxcx_v1_common_proto_goTypes = nil
	file_wxxcx_v1_common_proto_depIdxs = nil
}
