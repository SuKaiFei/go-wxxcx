// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: wxxcx/v1/navigation.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetNavigationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetNavigationsRequest) Reset() {
	*x = GetNavigationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_navigation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNavigationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNavigationsRequest) ProtoMessage() {}

func (x *GetNavigationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_navigation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNavigationsRequest.ProtoReflect.Descriptor instead.
func (*GetNavigationsRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_navigation_proto_rawDescGZIP(), []int{0}
}

func (x *GetNavigationsRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetNavigationsRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetNavigationsRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetNavigationsRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetNavigationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ImagePath string `protobuf:"bytes,2,opt,name=image_path,proto3" json:"image_path,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Describe  string `protobuf:"bytes,4,opt,name=describe,proto3" json:"describe,omitempty"`
	Sort      int64  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	MpAppid   string `protobuf:"bytes,6,opt,name=mp_appid,proto3" json:"mp_appid,omitempty"`
	Url       string `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetNavigationReply) Reset() {
	*x = GetNavigationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_navigation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNavigationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNavigationReply) ProtoMessage() {}

func (x *GetNavigationReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_navigation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNavigationReply.ProtoReflect.Descriptor instead.
func (*GetNavigationReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_navigation_proto_rawDescGZIP(), []int{1}
}

func (x *GetNavigationReply) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetNavigationReply) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *GetNavigationReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetNavigationReply) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *GetNavigationReply) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GetNavigationReply) GetMpAppid() string {
	if x != nil {
		return x.MpAppid
	}
	return ""
}

func (x *GetNavigationReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetNavigationsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetNavigationReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetNavigationsReply) Reset() {
	*x = GetNavigationsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_navigation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNavigationsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNavigationsReply) ProtoMessage() {}

func (x *GetNavigationsReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_navigation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNavigationsReply.ProtoReflect.Descriptor instead.
func (*GetNavigationsReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_navigation_proto_rawDescGZIP(), []int{2}
}

func (x *GetNavigationsReply) GetResults() []*GetNavigationReply {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_wxxcx_v1_navigation_proto protoreflect.FileDescriptor

var file_wxxcx_v1_navigation_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xbc, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x51, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x8f,
	0x01, 0x0a, 0x0a, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x80, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x78, 0x78, 0x63,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x75, 0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_navigation_proto_rawDescOnce sync.Once
	file_wxxcx_v1_navigation_proto_rawDescData = file_wxxcx_v1_navigation_proto_rawDesc
)

func file_wxxcx_v1_navigation_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_navigation_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_navigation_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_navigation_proto_rawDescData)
	})
	return file_wxxcx_v1_navigation_proto_rawDescData
}

var file_wxxcx_v1_navigation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wxxcx_v1_navigation_proto_goTypes = []interface{}{
	(*GetNavigationsRequest)(nil), // 0: api.wxxcx.v1.GetNavigationsRequest
	(*GetNavigationReply)(nil),    // 1: api.wxxcx.v1.GetNavigationReply
	(*GetNavigationsReply)(nil),   // 2: api.wxxcx.v1.GetNavigationsReply
}
var file_wxxcx_v1_navigation_proto_depIdxs = []int32{
	1, // 0: api.wxxcx.v1.GetNavigationsReply.results:type_name -> api.wxxcx.v1.GetNavigationReply
	0, // 1: api.wxxcx.v1.Navigation.GetNavigations:input_type -> api.wxxcx.v1.GetNavigationsRequest
	2, // 2: api.wxxcx.v1.Navigation.GetNavigations:output_type -> api.wxxcx.v1.GetNavigationsReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_navigation_proto_init() }
func file_wxxcx_v1_navigation_proto_init() {
	if File_wxxcx_v1_navigation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_navigation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNavigationsRequest); i {
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
		file_wxxcx_v1_navigation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNavigationReply); i {
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
		file_wxxcx_v1_navigation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNavigationsReply); i {
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
			RawDescriptor: file_wxxcx_v1_navigation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxxcx_v1_navigation_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_navigation_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_navigation_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_navigation_proto = out.File
	file_wxxcx_v1_navigation_proto_rawDesc = nil
	file_wxxcx_v1_navigation_proto_goTypes = nil
	file_wxxcx_v1_navigation_proto_depIdxs = nil
}