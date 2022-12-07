// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/rpc/examples/echoresource/v1/echoresource.proto

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

type EchoResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceRequest) Reset() {
	*x = EchoResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceRequest) ProtoMessage() {}

func (x *EchoResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceRequest.ProtoReflect.Descriptor instead.
func (*EchoResourceRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{0}
}

func (x *EchoResourceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EchoResourceRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceResponse) Reset() {
	*x = EchoResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceResponse) ProtoMessage() {}

func (x *EchoResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceResponse.ProtoReflect.Descriptor instead.
func (*EchoResourceResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResourceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResourceMultipleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceMultipleRequest) Reset() {
	*x = EchoResourceMultipleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceMultipleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceMultipleRequest) ProtoMessage() {}

func (x *EchoResourceMultipleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceMultipleRequest.ProtoReflect.Descriptor instead.
func (*EchoResourceMultipleRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{2}
}

func (x *EchoResourceMultipleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EchoResourceMultipleRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResourceMultipleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceMultipleResponse) Reset() {
	*x = EchoResourceMultipleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceMultipleResponse) ProtoMessage() {}

func (x *EchoResourceMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceMultipleResponse.ProtoReflect.Descriptor instead.
func (*EchoResourceMultipleResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{3}
}

func (x *EchoResourceMultipleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResourceBiDiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceBiDiRequest) Reset() {
	*x = EchoResourceBiDiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceBiDiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceBiDiRequest) ProtoMessage() {}

func (x *EchoResourceBiDiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceBiDiRequest.ProtoReflect.Descriptor instead.
func (*EchoResourceBiDiRequest) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{4}
}

func (x *EchoResourceBiDiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EchoResourceBiDiRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResourceBiDiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResourceBiDiResponse) Reset() {
	*x = EchoResourceBiDiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResourceBiDiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResourceBiDiResponse) ProtoMessage() {}

func (x *EchoResourceBiDiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResourceBiDiResponse.ProtoReflect.Descriptor instead.
func (*EchoResourceBiDiResponse) Descriptor() ([]byte, []int) {
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP(), []int{5}
}

func (x *EchoResourceBiDiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_rpc_examples_echoresource_v1_echoresource_proto protoreflect.FileDescriptor

var file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x43, 0x0a, 0x13, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x30, 0x0a, 0x14, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4b, 0x0a, 0x1b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x38, 0x0a, 0x1c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x17, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x69, 0x44, 0x69, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x69, 0x44, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd1, 0x03, 0x0a, 0x13, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x83, 0x01, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9d, 0x01, 0x0a, 0x14, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x12,
	0x3f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x93, 0x01, 0x0a, 0x10, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x69, 0x44, 0x69, 0x12, 0x3b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x69, 0x44,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x69, 0x44, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescOnce sync.Once
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescData = file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDesc
)

func file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescGZIP() []byte {
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescOnce.Do(func() {
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescData)
	})
	return file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDescData
}

var file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_rpc_examples_echoresource_v1_echoresource_proto_goTypes = []interface{}{
	(*EchoResourceRequest)(nil),          // 0: proto.rpc.examples.echoresource.v1.EchoResourceRequest
	(*EchoResourceResponse)(nil),         // 1: proto.rpc.examples.echoresource.v1.EchoResourceResponse
	(*EchoResourceMultipleRequest)(nil),  // 2: proto.rpc.examples.echoresource.v1.EchoResourceMultipleRequest
	(*EchoResourceMultipleResponse)(nil), // 3: proto.rpc.examples.echoresource.v1.EchoResourceMultipleResponse
	(*EchoResourceBiDiRequest)(nil),      // 4: proto.rpc.examples.echoresource.v1.EchoResourceBiDiRequest
	(*EchoResourceBiDiResponse)(nil),     // 5: proto.rpc.examples.echoresource.v1.EchoResourceBiDiResponse
}
var file_proto_rpc_examples_echoresource_v1_echoresource_proto_depIdxs = []int32{
	0, // 0: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResource:input_type -> proto.rpc.examples.echoresource.v1.EchoResourceRequest
	2, // 1: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResourceMultiple:input_type -> proto.rpc.examples.echoresource.v1.EchoResourceMultipleRequest
	4, // 2: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResourceBiDi:input_type -> proto.rpc.examples.echoresource.v1.EchoResourceBiDiRequest
	1, // 3: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResource:output_type -> proto.rpc.examples.echoresource.v1.EchoResourceResponse
	3, // 4: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResourceMultiple:output_type -> proto.rpc.examples.echoresource.v1.EchoResourceMultipleResponse
	5, // 5: proto.rpc.examples.echoresource.v1.EchoResourceService.EchoResourceBiDi:output_type -> proto.rpc.examples.echoresource.v1.EchoResourceBiDiResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_examples_echoresource_v1_echoresource_proto_init() }
func file_proto_rpc_examples_echoresource_v1_echoresource_proto_init() {
	if File_proto_rpc_examples_echoresource_v1_echoresource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceRequest); i {
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
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceResponse); i {
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
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceMultipleRequest); i {
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
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceMultipleResponse); i {
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
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceBiDiRequest); i {
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
		file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResourceBiDiResponse); i {
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
			RawDescriptor: file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_examples_echoresource_v1_echoresource_proto_goTypes,
		DependencyIndexes: file_proto_rpc_examples_echoresource_v1_echoresource_proto_depIdxs,
		MessageInfos:      file_proto_rpc_examples_echoresource_v1_echoresource_proto_msgTypes,
	}.Build()
	File_proto_rpc_examples_echoresource_v1_echoresource_proto = out.File
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_rawDesc = nil
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_goTypes = nil
	file_proto_rpc_examples_echoresource_v1_echoresource_proto_depIdxs = nil
}
