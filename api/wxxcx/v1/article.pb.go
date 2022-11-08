// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: wxxcx/v1/article.proto

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

type GetArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleRequest.ProtoReflect.Descriptor instead.
func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *GetArticleRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetArticleRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetArticleRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetArticleRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetArticleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImagePath string `protobuf:"bytes,1,opt,name=image_path,proto3" json:"image_path,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetArticleReply) Reset() {
	*x = GetArticleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleReply) ProtoMessage() {}

func (x *GetArticleReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleReply.ProtoReflect.Descriptor instead.
func (*GetArticleReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *GetArticleReply) GetImagePath() string {
	if x != nil {
		return x.ImagePath
	}
	return ""
}

func (x *GetArticleReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetArticleReply) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type GetArticlesRequest struct {
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

func (x *GetArticlesRequest) Reset() {
	*x = GetArticlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesRequest) ProtoMessage() {}

func (x *GetArticlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesRequest.ProtoReflect.Descriptor instead.
func (*GetArticlesRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_article_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticlesRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetArticlesRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetArticlesRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetArticlesRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetArticlesRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetArticlesRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetArticlesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetArticleReply `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetArticlesReply) Reset() {
	*x = GetArticlesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArticlesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticlesReply) ProtoMessage() {}

func (x *GetArticlesReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticlesReply.ProtoReflect.Descriptor instead.
func (*GetArticlesReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_article_proto_rawDescGZIP(), []int{3}
}

func (x *GetArticlesReply) GetResults() []*GetArticleReply {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_wxxcx_v1_article_proto protoreflect.FileDescriptor

var file_wxxcx_v1_article_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x61, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x32, 0xdd, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x64, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x77, 0x78,
	0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x77, 0x78,
	0x78, 0x63, 0x78, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x62, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x6c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x1c, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x75, 0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x78, 0x78, 0x63,
	0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_article_proto_rawDescOnce sync.Once
	file_wxxcx_v1_article_proto_rawDescData = file_wxxcx_v1_article_proto_rawDesc
)

func file_wxxcx_v1_article_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_article_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_article_proto_rawDescData)
	})
	return file_wxxcx_v1_article_proto_rawDescData
}

var file_wxxcx_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wxxcx_v1_article_proto_goTypes = []interface{}{
	(*GetArticleRequest)(nil),  // 0: wxxcx.v1.GetArticleRequest
	(*GetArticleReply)(nil),    // 1: wxxcx.v1.GetArticleReply
	(*GetArticlesRequest)(nil), // 2: wxxcx.v1.GetArticlesRequest
	(*GetArticlesReply)(nil),   // 3: wxxcx.v1.GetArticlesReply
}
var file_wxxcx_v1_article_proto_depIdxs = []int32{
	1, // 0: wxxcx.v1.GetArticlesReply.results:type_name -> wxxcx.v1.GetArticleReply
	0, // 1: wxxcx.v1.Article.GetArticle:input_type -> wxxcx.v1.GetArticleRequest
	2, // 2: wxxcx.v1.Article.GetArticles:input_type -> wxxcx.v1.GetArticlesRequest
	1, // 3: wxxcx.v1.Article.GetArticle:output_type -> wxxcx.v1.GetArticleReply
	3, // 4: wxxcx.v1.Article.GetArticles:output_type -> wxxcx.v1.GetArticlesReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_article_proto_init() }
func file_wxxcx_v1_article_proto_init() {
	if File_wxxcx_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleRequest); i {
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
		file_wxxcx_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticleReply); i {
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
		file_wxxcx_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesRequest); i {
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
		file_wxxcx_v1_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArticlesReply); i {
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
			RawDescriptor: file_wxxcx_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxxcx_v1_article_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_article_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_article_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_article_proto = out.File
	file_wxxcx_v1_article_proto_rawDesc = nil
	file_wxxcx_v1_article_proto_goTypes = nil
	file_wxxcx_v1_article_proto_depIdxs = nil
}