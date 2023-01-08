// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: wxxcx/v1/voice.proto

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

type GetVoiceDefaultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *GetVoiceDefaultRequest) Reset() {
	*x = GetVoiceDefaultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceDefaultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceDefaultRequest) ProtoMessage() {}

func (x *GetVoiceDefaultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceDefaultRequest.ProtoReflect.Descriptor instead.
func (*GetVoiceDefaultRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{0}
}

func (x *GetVoiceDefaultRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetVoiceDefaultRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetVoiceDefaultRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type GetVoiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type    uint32 `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	MpAppid string `protobuf:"bytes,5,opt,name=mp_appid,proto3" json:"mp_appid,omitempty"`
	Works   string `protobuf:"bytes,6,opt,name=works,proto3" json:"works,omitempty"`
	Share   *Share `protobuf:"bytes,7,opt,name=share,proto3" json:"share,omitempty"`
	MpUrl   string `protobuf:"bytes,8,opt,name=mp_url,proto3" json:"mp_url,omitempty"`
}

func (x *GetVoiceReply) Reset() {
	*x = GetVoiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceReply) ProtoMessage() {}

func (x *GetVoiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceReply.ProtoReflect.Descriptor instead.
func (*GetVoiceReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{1}
}

func (x *GetVoiceReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetVoiceReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetVoiceReply) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetVoiceReply) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetVoiceReply) GetMpAppid() string {
	if x != nil {
		return x.MpAppid
	}
	return ""
}

func (x *GetVoiceReply) GetWorks() string {
	if x != nil {
		return x.Works
	}
	return ""
}

func (x *GetVoiceReply) GetShare() *Share {
	if x != nil {
		return x.Share
	}
	return nil
}

func (x *GetVoiceReply) GetMpUrl() string {
	if x != nil {
		return x.MpUrl
	}
	return ""
}

type GetVoiceByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	Id        uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVoiceByIdRequest) Reset() {
	*x = GetVoiceByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceByIdRequest) ProtoMessage() {}

func (x *GetVoiceByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceByIdRequest.ProtoReflect.Descriptor instead.
func (*GetVoiceByIdRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{2}
}

func (x *GetVoiceByIdRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetVoiceByIdRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetVoiceByIdRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *GetVoiceByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetVoiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Sign      string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (x *GetVoiceListRequest) Reset() {
	*x = GetVoiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceListRequest) ProtoMessage() {}

func (x *GetVoiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceListRequest.ProtoReflect.Descriptor instead.
func (*GetVoiceListRequest) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{3}
}

func (x *GetVoiceListRequest) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *GetVoiceListRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *GetVoiceListRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

type GetVoiceListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*GetVoiceListReply_Info `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetVoiceListReply) Reset() {
	*x = GetVoiceListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceListReply) ProtoMessage() {}

func (x *GetVoiceListReply) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceListReply.ProtoReflect.Descriptor instead.
func (*GetVoiceListReply) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{4}
}

func (x *GetVoiceListReply) GetResults() []*GetVoiceListReply_Info {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetVoiceListReply_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          uint32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code          string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	MpAppid       string `protobuf:"bytes,5,opt,name=mp_appid,proto3" json:"mp_appid,omitempty"`
	ShareImageUrl string `protobuf:"bytes,6,opt,name=share_image_url,proto3" json:"share_image_url,omitempty"`
	MpUrl         string `protobuf:"bytes,7,opt,name=mp_url,proto3" json:"mp_url,omitempty"`
}

func (x *GetVoiceListReply_Info) Reset() {
	*x = GetVoiceListReply_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wxxcx_v1_voice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoiceListReply_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoiceListReply_Info) ProtoMessage() {}

func (x *GetVoiceListReply_Info) ProtoReflect() protoreflect.Message {
	mi := &file_wxxcx_v1_voice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoiceListReply_Info.ProtoReflect.Descriptor instead.
func (*GetVoiceListReply_Info) Descriptor() ([]byte, []int) {
	return file_wxxcx_v1_voice_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetVoiceListReply_Info) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetVoiceListReply_Info) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetVoiceListReply_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetVoiceListReply_Info) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetVoiceListReply_Info) GetMpAppid() string {
	if x != nil {
		return x.MpAppid
	}
	return ""
}

func (x *GetVoiceListReply_Info) GetShareImageUrl() string {
	if x != nil {
		return x.ShareImageUrl
	}
	return ""
}

func (x *GetVoiceListReply_Info) GetMpUrl() string {
	if x != nil {
		return x.MpUrl
	}
	return ""
}

var File_wxxcx_v1_voice_proto protoreflect.FileDescriptor

var file_wxxcx_v1_voice_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0xd3, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x70, 0x5f, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x70, 0x5f, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x78, 0x78, 0x63,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x70, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c,
	0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x88,
	0x02, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0xb0, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x70, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x70, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x32, 0xb5, 0x03, 0x0a, 0x05, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x26, 0x2e, 0x77, 0x78, 0x78, 0x63,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78,
	0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x6e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23,
	0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x77, 0x78, 0x78,
	0x63, 0x78, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x51,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11,
	0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x69, 0x6e,
	0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x75, 0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x78, 0x78, 0x63,
	0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wxxcx_v1_voice_proto_rawDescOnce sync.Once
	file_wxxcx_v1_voice_proto_rawDescData = file_wxxcx_v1_voice_proto_rawDesc
)

func file_wxxcx_v1_voice_proto_rawDescGZIP() []byte {
	file_wxxcx_v1_voice_proto_rawDescOnce.Do(func() {
		file_wxxcx_v1_voice_proto_rawDescData = protoimpl.X.CompressGZIP(file_wxxcx_v1_voice_proto_rawDescData)
	})
	return file_wxxcx_v1_voice_proto_rawDescData
}

var file_wxxcx_v1_voice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wxxcx_v1_voice_proto_goTypes = []interface{}{
	(*GetVoiceDefaultRequest)(nil), // 0: wxxcx.v1.voice.GetVoiceDefaultRequest
	(*GetVoiceReply)(nil),          // 1: wxxcx.v1.voice.GetVoiceReply
	(*GetVoiceByIdRequest)(nil),    // 2: wxxcx.v1.voice.GetVoiceByIdRequest
	(*GetVoiceListRequest)(nil),    // 3: wxxcx.v1.voice.GetVoiceListRequest
	(*GetVoiceListReply)(nil),      // 4: wxxcx.v1.voice.GetVoiceListReply
	(*GetVoiceListReply_Info)(nil), // 5: wxxcx.v1.voice.GetVoiceListReply.Info
	(*Share)(nil),                  // 6: wxxcx.v1.common.Share
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_wxxcx_v1_voice_proto_depIdxs = []int32{
	6, // 0: wxxcx.v1.voice.GetVoiceReply.share:type_name -> wxxcx.v1.common.Share
	5, // 1: wxxcx.v1.voice.GetVoiceListReply.results:type_name -> wxxcx.v1.voice.GetVoiceListReply.Info
	3, // 2: wxxcx.v1.voice.Voice.GetVoiceList:input_type -> wxxcx.v1.voice.GetVoiceListRequest
	0, // 3: wxxcx.v1.voice.Voice.GetVoiceDefault:input_type -> wxxcx.v1.voice.GetVoiceDefaultRequest
	2, // 4: wxxcx.v1.voice.Voice.GetVoiceById:input_type -> wxxcx.v1.voice.GetVoiceByIdRequest
	7, // 5: wxxcx.v1.voice.Voice.Ping:input_type -> google.protobuf.Empty
	4, // 6: wxxcx.v1.voice.Voice.GetVoiceList:output_type -> wxxcx.v1.voice.GetVoiceListReply
	1, // 7: wxxcx.v1.voice.Voice.GetVoiceDefault:output_type -> wxxcx.v1.voice.GetVoiceReply
	1, // 8: wxxcx.v1.voice.Voice.GetVoiceById:output_type -> wxxcx.v1.voice.GetVoiceReply
	7, // 9: wxxcx.v1.voice.Voice.Ping:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wxxcx_v1_voice_proto_init() }
func file_wxxcx_v1_voice_proto_init() {
	if File_wxxcx_v1_voice_proto != nil {
		return
	}
	file_wxxcx_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wxxcx_v1_voice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceDefaultRequest); i {
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
		file_wxxcx_v1_voice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceReply); i {
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
		file_wxxcx_v1_voice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceByIdRequest); i {
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
		file_wxxcx_v1_voice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceListRequest); i {
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
		file_wxxcx_v1_voice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceListReply); i {
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
		file_wxxcx_v1_voice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoiceListReply_Info); i {
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
			RawDescriptor: file_wxxcx_v1_voice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wxxcx_v1_voice_proto_goTypes,
		DependencyIndexes: file_wxxcx_v1_voice_proto_depIdxs,
		MessageInfos:      file_wxxcx_v1_voice_proto_msgTypes,
	}.Build()
	File_wxxcx_v1_voice_proto = out.File
	file_wxxcx_v1_voice_proto_rawDesc = nil
	file_wxxcx_v1_voice_proto_goTypes = nil
	file_wxxcx_v1_voice_proto_depIdxs = nil
}
