// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: wxxcx/v1/music.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMusicListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Page      uint64 `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  uint64 `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetMusicListRequest) Reset() {
	*x = GetMusicListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_music_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMusicListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMusicListRequest) ProtoMessage() {}

func (x *GetMusicListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_music_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMusicListRequest.ProtoReflect.Descriptor instead.
func (*GetMusicListRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_music_proto_rawDescGZIP(), []int{0}
}

func (x *GetMusicListRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetMusicListRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetMusicListRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetMusicListRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetMusicListRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetMusicListRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetMusicListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetMusicListReply_Info `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetMusicListReply) Reset() {
	*x = GetMusicListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_music_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMusicListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMusicListReply) ProtoMessage() {}

func (x *GetMusicListReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_music_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMusicListReply.ProtoReflect.Descriptor instead.
func (*GetMusicListReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_music_proto_rawDescGZIP(), []int{1}
}

func (x *GetMusicListReply) GetResults() []*GetMusicListReply_Info {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetMusicListReply_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Duration  uint64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Singer    string `protobuf:"bytes,4,opt,name=singer,proto3" json:"singer,omitempty"`
	Url       string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	ImagePath string `protobuf:"bytes,6,opt,name=image_path,proto3" json:"image_path,omitempty"`
	Share     *Share `protobuf:"bytes,7,opt,name=share,proto3" json:"share,omitempty"`
}

func (x *GetMusicListReply_Info) Reset() {
	*x = GetMusicListReply_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_music_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMusicListReply_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMusicListReply_Info) ProtoMessage() {}

func (x *GetMusicListReply_Info) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_music_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMusicListReply_Info.ProtoReflect.Descriptor instead.
func (*GetMusicListReply_Info) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_music_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetMusicListReply_Info) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMusicListReply_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetMusicListReply_Info) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *GetMusicListReply_Info) GetSinger() string {
	if x != nil {
		return x.Singer
	}
	return ""
}

func (x *GetMusicListReply_Info) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetMusicListReply_Info) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *GetMusicListReply_Info) GetShare() *Share {
	if x != nil {
		return x.Share
	}
	return nil
}

var File_wxxcx_v1_music_proto protoreflect.FileDescriptor

var file_wxxcx_v1_music_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x96, 0x02,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xbe, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52,
	0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x05, 0x4d, 0x75, 0x73, 0x69, 0x63,
	0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x23, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x75, 0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f,
	0x2d, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_music_proto_rawDescOnce sync.Once
	file_wxxcx_v1_music_proto_rawDescData = file_wxxcx_v1_music_proto_rawDesc
)

func file_wxxcx_v1_music_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_music_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_music_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_music_proto_rawDescData)
	})
	return file_wxxcx_v1_music_proto_rawDescData
}

var file_wxxcx_v1_music_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wxxcx_v1_music_proto_goTypes = []interface{}{
	(*GetMusicListRequest)(nil),    // 0: wxxcx.v1.music.GetMusicListRequest
	(*GetMusicListReply)(nil),      // 1: wxxcx.v1.music.GetMusicListReply
	(*GetMusicListReply_Info)(nil), // 2: wxxcx.v1.music.GetMusicListReply.Info
	(*Share)(nil),                  // 3: wxxcx.v1.common.Share
	(*emptypb.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_wxxcx_v1_music_proto_depIdxs = []int32{
	2, // 0: wxxcx.v1.music.GetMusicListReply.results:type_name -> wxxcx.v1.music.GetMusicListReply.Info
	3, // 1: wxxcx.v1.music.GetMusicListReply.Info.share:type_name -> wxxcx.v1.common.Share
	0, // 2: wxxcx.v1.music.Music.GetMusicList:input_type -> wxxcx.v1.music.GetMusicListRequest
	4, // 3: wxxcx.v1.music.Music.Ping:input_type -> google.protobuf.Empty
	1, // 4: wxxcx.v1.music.Music.GetMusicList:output_type -> wxxcx.v1.music.GetMusicListReply
	4, // 5: wxxcx.v1.music.Music.Ping:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_music_proto_init() }
func file_wxxcx_v1_music_proto_init() {
	if File_wxxcx_v1_music_proto != nil {
		return
	}
	file_wxxcx_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_music_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMusicListRequest); i {
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
		file_wxxcx_v1_music_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMusicListReply); i {
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
		file_wxxcx_v1_music_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMusicListReply_Info); i {
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
			RawDescriptor: file_wxxcx_v1_music_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxxcx_v1_music_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_music_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_music_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_music_proto = out.File
	file_wxxcx_v1_music_proto_rawDesc = nil
	file_wxxcx_v1_music_proto_goTypes = nil
	file_wxxcx_v1_music_proto_depIdxs = nil
}
