// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: wxxcx/v1/wechatmp.proto

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

type LoginWechatMpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoginWechatMpRequest) Reset() {
	*x = LoginWechatMpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginWechatMpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWechatMpRequest) ProtoMessage() {}

func (x *LoginWechatMpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWechatMpRequest.ProtoReflect.Descriptor instead.
func (*LoginWechatMpRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_wechatmp_proto_rawDescGZIP(), []int{0}
}

func (x *LoginWechatMpRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *LoginWechatMpRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *LoginWechatMpRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *LoginWechatMpRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginWechatMpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Openid     string `protobuf:"bytes,1,opt,name=openid,proto3" json:"openid,omitempty"`
	SessionKey string `protobuf:"bytes,2,opt,name=session_key,proto3" json:"session_key,omitempty"`
	Unionid    string `protobuf:"bytes,3,opt,name=unionid,proto3" json:"unionid,omitempty"`
}

func (x *LoginWechatMpReply) Reset() {
	*x = LoginWechatMpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginWechatMpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginWechatMpReply) ProtoMessage() {}

func (x *LoginWechatMpReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginWechatMpReply.ProtoReflect.Descriptor instead.
func (*LoginWechatMpReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_wechatmp_proto_rawDescGZIP(), []int{1}
}

func (x *LoginWechatMpReply) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *LoginWechatMpReply) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *LoginWechatMpReply) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

type SecurityCheckMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Openid    string `protobuf:"bytes,5,opt,name=openid,proto3" json:"openid,omitempty"`
}

func (x *SecurityCheckMsgRequest) Reset() {
	*x = SecurityCheckMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityCheckMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityCheckMsgRequest) ProtoMessage() {}

func (x *SecurityCheckMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityCheckMsgRequest.ProtoReflect.Descriptor instead.
func (*SecurityCheckMsgRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_wechatmp_proto_rawDescGZIP(), []int{2}
}

func (x *SecurityCheckMsgRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *SecurityCheckMsgRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *SecurityCheckMsgRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *SecurityCheckMsgRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SecurityCheckMsgRequest) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

type SecurityCheckMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggest string `protobuf:"bytes,1,opt,name=suggest,proto3" json:"suggest,omitempty"` // ????????????risky???pass???review?????????
	Label   uint32 `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`    // ????????????????????????100 ?????????10001 ?????????20001 ?????????20002 ?????????20003 ?????????20006 ???????????????20008 ?????????20012 ?????????20013 ?????????21000 ??????
}

func (x *SecurityCheckMsgReply) Reset() {
	*x = SecurityCheckMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityCheckMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityCheckMsgReply) ProtoMessage() {}

func (x *SecurityCheckMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_wechatmp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityCheckMsgReply.ProtoReflect.Descriptor instead.
func (*SecurityCheckMsgReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_wechatmp_proto_rawDescGZIP(), []int{3}
}

func (x *SecurityCheckMsgReply) GetSuggest() string {
	if x != nil {
		return x.Suggest
	}
	return ""
}

func (x *SecurityCheckMsgReply) GetLabel() uint32 {
	if x != nil {
		return x.Label
	}
	return 0
}

var File_wxxcx_v1_wechatmp_proto protoreflect.FileDescriptor

var file_wxxcx_v1_wechatmp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x70, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x68,
	0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x22, 0x47,
	0x0a, 0x15, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0xa5, 0x02, 0x0a, 0x08, 0x57, 0x65, 0x63, 0x68,
	0x61, 0x74, 0x4d, 0x70, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65,
	0x63, 0x68, 0x61, 0x74, 0x4d, 0x70, 0x12, 0x27, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x6d, 0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x4d,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16,
	0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x70,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x93, 0x01, 0x0a, 0x10, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x2a,
	0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x6d, 0x70, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x78, 0x78,
	0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x6d, 0x70, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x77,
	0x78, 0x78, 0x63, 0x78, 0x2f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x70, 0x2f, 0x73,
	0x65, 0x63, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x6d, 0x73, 0x67, 0x3a, 0x01, 0x2a, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x75,
	0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_wechatmp_proto_rawDescOnce sync.Once
	file_wxxcx_v1_wechatmp_proto_rawDescData = file_wxxcx_v1_wechatmp_proto_rawDesc
)

func file_wxxcx_v1_wechatmp_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_wechatmp_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_wechatmp_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_wechatmp_proto_rawDescData)
	})
	return file_wxxcx_v1_wechatmp_proto_rawDescData
}

var file_wxxcx_v1_wechatmp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wxxcx_v1_wechatmp_proto_goTypes = []interface{}{
	(*LoginWechatMpRequest)(nil),    // 0: wxxcx.v1.wechatmp.LoginWechatMpRequest
	(*LoginWechatMpReply)(nil),      // 1: wxxcx.v1.wechatmp.LoginWechatMpReply
	(*SecurityCheckMsgRequest)(nil), // 2: wxxcx.v1.wechatmp.SecurityCheckMsgRequest
	(*SecurityCheckMsgReply)(nil),   // 3: wxxcx.v1.wechatmp.SecurityCheckMsgReply
}
var file_wxxcx_v1_wechatmp_proto_depIdxs = []int32{
	0, // 0: wxxcx.v1.wechatmp.WechatMp.LoginWechatMp:input_type -> wxxcx.v1.wechatmp.LoginWechatMpRequest
	2, // 1: wxxcx.v1.wechatmp.WechatMp.SecurityCheckMsg:input_type -> wxxcx.v1.wechatmp.SecurityCheckMsgRequest
	1, // 2: wxxcx.v1.wechatmp.WechatMp.LoginWechatMp:output_type -> wxxcx.v1.wechatmp.LoginWechatMpReply
	3, // 3: wxxcx.v1.wechatmp.WechatMp.SecurityCheckMsg:output_type -> wxxcx.v1.wechatmp.SecurityCheckMsgReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_wechatmp_proto_init() }
func file_wxxcx_v1_wechatmp_proto_init() {
	if File_wxxcx_v1_wechatmp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_wechatmp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginWechatMpRequest); i {
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
		file_wxxcx_v1_wechatmp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginWechatMpReply); i {
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
		file_wxxcx_v1_wechatmp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityCheckMsgRequest); i {
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
		file_wxxcx_v1_wechatmp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityCheckMsgReply); i {
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
			RawDescriptor: file_wxxcx_v1_wechatmp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxxcx_v1_wechatmp_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_wechatmp_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_wechatmp_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_wechatmp_proto = out.File
	file_wxxcx_v1_wechatmp_proto_rawDesc = nil
	file_wxxcx_v1_wechatmp_proto_goTypes = nil
	file_wxxcx_v1_wechatmp_proto_depIdxs = nil
}
