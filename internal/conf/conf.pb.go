// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server      *Server      `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data        *Data        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Application *Application `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetApplication() *Application {
	if x != nil {
		return x.Application
	}
	return nil
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mp         map[string]*Application_App `protobuf:"bytes,1,rep,name=mp,proto3" json:"mp,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StaticPath string                      `protobuf:"bytes,2,opt,name=staticPath,proto3" json:"staticPath,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Application) GetMp() map[string]*Application_App {
	if x != nil {
		return x.Mp
	}
	return nil
}

func (x *Application) GetStaticPath() string {
	if x != nil {
		return x.StaticPath
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Application_App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key            string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Token          string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	EncodingAESKey string `protobuf:"bytes,5,opt,name=encodingAESKey,proto3" json:"encodingAESKey,omitempty"`
	Secret         string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *Application_App) Reset() {
	*x = Application_App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application_App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application_App) ProtoMessage() {}

func (x *Application_App) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application_App.ProtoReflect.Descriptor instead.
func (*Application_App) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Application_App) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Application_App) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Application_App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application_App) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Application_App) GetEncodingAESKey() string {
	if x != nil {
		return x.EncodingAESKey
	}
	return ""
}

func (x *Application_App) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc6, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x02, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x91, 0x01, 0x0a, 0x03, 0x41, 0x70,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x45, 0x53, 0x4b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x41,
	0x45, 0x53, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x52, 0x0a,
	0x07, 0x4d, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48,
	0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x69, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x22, 0x79, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x75,
	0x4b, 0x61, 0x69, 0x46, 0x65, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x77, 0x78, 0x78, 0x63, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData = file_conf_conf_proto_rawDesc
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_conf_proto_rawDescData)
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Application)(nil),         // 1: kratos.api.Application
	(*Server)(nil),              // 2: kratos.api.Server
	(*Data)(nil),                // 3: kratos.api.Data
	(*Application_App)(nil),     // 4: kratos.api.Application.App
	nil,                         // 5: kratos.api.Application.MpEntry
	(*Server_HTTP)(nil),         // 6: kratos.api.Server.HTTP
	(*Data_Database)(nil),       // 7: kratos.api.Data.Database
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	2, // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	3, // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	1, // 2: kratos.api.Bootstrap.application:type_name -> kratos.api.Application
	5, // 3: kratos.api.Application.mp:type_name -> kratos.api.Application.MpEntry
	6, // 4: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	7, // 5: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	4, // 6: kratos.api.Application.MpEntry.value:type_name -> kratos.api.Application.App
	8, // 7: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application_App); i {
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
		file_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
			RawDescriptor: file_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_rawDesc = nil
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
