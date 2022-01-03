// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/rpc/webrtc/v1/grpc.proto

package v1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A PacketMessage is used to packetize large messages (> 64KiB) to be able to safely
// transmit over WebRTC data channels.
type PacketMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Eom  bool   `protobuf:"varint,2,opt,name=eom,proto3" json:"eom,omitempty"`
}

func (x *PacketMessage) Reset() {
	*x = PacketMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketMessage) ProtoMessage() {}

func (x *PacketMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketMessage.ProtoReflect.Descriptor instead.
func (*PacketMessage) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *PacketMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PacketMessage) GetEom() bool {
	if x != nil {
		return x.Eom
	}
	return false
}

// A Stream represents an instance of a gRPC stream between
// a client and a server.
type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Stream) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// A Request is a frame coming from a client. It is always
// associated with a stream where the client assigns the stream
// identifier. Servers will drop frames where the stream identifier
// has no association (if a non-header frames are sent).
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream *Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	// Types that are assignable to Type:
	//	*Request_Headers
	//	*Request_Message
	Type isRequest_Type `protobuf_oneof:"type"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (m *Request) GetType() isRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Request) GetHeaders() *RequestHeaders {
	if x, ok := x.GetType().(*Request_Headers); ok {
		return x.Headers
	}
	return nil
}

func (x *Request) GetMessage() *RequestMessage {
	if x, ok := x.GetType().(*Request_Message); ok {
		return x.Message
	}
	return nil
}

type isRequest_Type interface {
	isRequest_Type()
}

type Request_Headers struct {
	Headers *RequestHeaders `protobuf:"bytes,2,opt,name=headers,proto3,oneof"`
}

type Request_Message struct {
	Message *RequestMessage `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

func (*Request_Headers) isRequest_Type() {}

func (*Request_Message) isRequest_Type() {}

// RequestHeaders describe the unary or streaming call to make.
type RequestHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method   string               `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Metadata *Metadata            `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Timeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *RequestHeaders) Reset() {
	*x = RequestHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeaders) ProtoMessage() {}

func (x *RequestHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeaders.ProtoReflect.Descriptor instead.
func (*RequestHeaders) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *RequestHeaders) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RequestHeaders) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RequestHeaders) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

// A RequestMessage contains individual gRPC messages and a potential
// end-of-stream (EOS) marker.
type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasMessage    bool           `protobuf:"varint,1,opt,name=has_message,json=hasMessage,proto3" json:"has_message,omitempty"`
	PacketMessage *PacketMessage `protobuf:"bytes,2,opt,name=packet_message,json=packetMessage,proto3" json:"packet_message,omitempty"`
	Eos           bool           `protobuf:"varint,3,opt,name=eos,proto3" json:"eos,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *RequestMessage) GetHasMessage() bool {
	if x != nil {
		return x.HasMessage
	}
	return false
}

func (x *RequestMessage) GetPacketMessage() *PacketMessage {
	if x != nil {
		return x.PacketMessage
	}
	return nil
}

func (x *RequestMessage) GetEos() bool {
	if x != nil {
		return x.Eos
	}
	return false
}

// A Response is a frame coming from a server. It is always
// associated with a stream where the client assigns the stream
// identifier. Clients will drop frames where the stream identifier
// has no association.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream *Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	// Types that are assignable to Type:
	//	*Response_Headers
	//	*Response_Message
	//	*Response_Trailers
	Type isResponse_Type `protobuf_oneof:"type"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (m *Response) GetType() isResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Response) GetHeaders() *ResponseHeaders {
	if x, ok := x.GetType().(*Response_Headers); ok {
		return x.Headers
	}
	return nil
}

func (x *Response) GetMessage() *ResponseMessage {
	if x, ok := x.GetType().(*Response_Message); ok {
		return x.Message
	}
	return nil
}

func (x *Response) GetTrailers() *ResponseTrailers {
	if x, ok := x.GetType().(*Response_Trailers); ok {
		return x.Trailers
	}
	return nil
}

type isResponse_Type interface {
	isResponse_Type()
}

type Response_Headers struct {
	Headers *ResponseHeaders `protobuf:"bytes,2,opt,name=headers,proto3,oneof"`
}

type Response_Message struct {
	Message *ResponseMessage `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type Response_Trailers struct {
	Trailers *ResponseTrailers `protobuf:"bytes,4,opt,name=trailers,proto3,oneof"`
}

func (*Response_Headers) isResponse_Type() {}

func (*Response_Message) isResponse_Type() {}

func (*Response_Trailers) isResponse_Type() {}

// ResponseHeaders contain custom metadata that are sent to the client
// before any message or trailers (unless only trailers are sent).
type ResponseHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ResponseHeaders) Reset() {
	*x = ResponseHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseHeaders) ProtoMessage() {}

func (x *ResponseHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseHeaders.ProtoReflect.Descriptor instead.
func (*ResponseHeaders) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseHeaders) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ResponseMessage contains the data of a response to a call.
type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketMessage *PacketMessage `protobuf:"bytes,1,opt,name=packet_message,json=packetMessage,proto3" json:"packet_message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *ResponseMessage) GetPacketMessage() *PacketMessage {
	if x != nil {
		return x.PacketMessage
	}
	return nil
}

// ResponseTrailers contain the status of a response and any custom metadata.
type ResponseTrailers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Metadata *Metadata      `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ResponseTrailers) Reset() {
	*x = ResponseTrailers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseTrailers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTrailers) ProtoMessage() {}

func (x *ResponseTrailers) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTrailers.ProtoReflect.Descriptor instead.
func (*ResponseTrailers) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *ResponseTrailers) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ResponseTrailers) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Strings are a series of values.
type Strings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Strings) Reset() {
	*x = Strings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strings) ProtoMessage() {}

func (x *Strings) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strings.ProtoReflect.Descriptor instead.
func (*Strings) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *Strings) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// Metadata is for custom key values provided by a client or server
// during a stream.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md map[string]*Strings `protobuf:"bytes,1,rep,name=md,proto3" json:"md,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *Metadata) GetMd() map[string]*Strings {
	if x != nil {
		return x.Md
	}
	return nil
}

var File_proto_rpc_webrtc_v1_grpc_proto protoreflect.FileDescriptor

var file_proto_rpc_webrtc_v1_grpc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x65, 0x6f, 0x6d, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xc8, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x3f, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65,
	0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77,
	0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68,
	0x61, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65,
	0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x65, 0x6f, 0x73, 0x22, 0x90, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x40, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x48,
	0x00, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4c, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77,
	0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x21, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x35, 0x0a, 0x02, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x02, 0x6d, 0x64, 0x1a, 0x53, 0x0a, 0x07, 0x4d, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x65, 0x62, 0x72,
	0x74, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_webrtc_v1_grpc_proto_rawDescOnce sync.Once
	file_proto_rpc_webrtc_v1_grpc_proto_rawDescData = file_proto_rpc_webrtc_v1_grpc_proto_rawDesc
)

func file_proto_rpc_webrtc_v1_grpc_proto_rawDescGZIP() []byte {
	file_proto_rpc_webrtc_v1_grpc_proto_rawDescOnce.Do(func() {
		file_proto_rpc_webrtc_v1_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_webrtc_v1_grpc_proto_rawDescData)
	})
	return file_proto_rpc_webrtc_v1_grpc_proto_rawDescData
}

var file_proto_rpc_webrtc_v1_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_rpc_webrtc_v1_grpc_proto_goTypes = []interface{}{
	(*PacketMessage)(nil),       // 0: proto.rpc.webrtc.v1.PacketMessage
	(*Stream)(nil),              // 1: proto.rpc.webrtc.v1.Stream
	(*Request)(nil),             // 2: proto.rpc.webrtc.v1.Request
	(*RequestHeaders)(nil),      // 3: proto.rpc.webrtc.v1.RequestHeaders
	(*RequestMessage)(nil),      // 4: proto.rpc.webrtc.v1.RequestMessage
	(*Response)(nil),            // 5: proto.rpc.webrtc.v1.Response
	(*ResponseHeaders)(nil),     // 6: proto.rpc.webrtc.v1.ResponseHeaders
	(*ResponseMessage)(nil),     // 7: proto.rpc.webrtc.v1.ResponseMessage
	(*ResponseTrailers)(nil),    // 8: proto.rpc.webrtc.v1.ResponseTrailers
	(*Strings)(nil),             // 9: proto.rpc.webrtc.v1.Strings
	(*Metadata)(nil),            // 10: proto.rpc.webrtc.v1.Metadata
	nil,                         // 11: proto.rpc.webrtc.v1.Metadata.MdEntry
	(*durationpb.Duration)(nil), // 12: google.protobuf.Duration
	(*status.Status)(nil),       // 13: google.rpc.Status
}
var file_proto_rpc_webrtc_v1_grpc_proto_depIdxs = []int32{
	1,  // 0: proto.rpc.webrtc.v1.Request.stream:type_name -> proto.rpc.webrtc.v1.Stream
	3,  // 1: proto.rpc.webrtc.v1.Request.headers:type_name -> proto.rpc.webrtc.v1.RequestHeaders
	4,  // 2: proto.rpc.webrtc.v1.Request.message:type_name -> proto.rpc.webrtc.v1.RequestMessage
	10, // 3: proto.rpc.webrtc.v1.RequestHeaders.metadata:type_name -> proto.rpc.webrtc.v1.Metadata
	12, // 4: proto.rpc.webrtc.v1.RequestHeaders.timeout:type_name -> google.protobuf.Duration
	0,  // 5: proto.rpc.webrtc.v1.RequestMessage.packet_message:type_name -> proto.rpc.webrtc.v1.PacketMessage
	1,  // 6: proto.rpc.webrtc.v1.Response.stream:type_name -> proto.rpc.webrtc.v1.Stream
	6,  // 7: proto.rpc.webrtc.v1.Response.headers:type_name -> proto.rpc.webrtc.v1.ResponseHeaders
	7,  // 8: proto.rpc.webrtc.v1.Response.message:type_name -> proto.rpc.webrtc.v1.ResponseMessage
	8,  // 9: proto.rpc.webrtc.v1.Response.trailers:type_name -> proto.rpc.webrtc.v1.ResponseTrailers
	10, // 10: proto.rpc.webrtc.v1.ResponseHeaders.metadata:type_name -> proto.rpc.webrtc.v1.Metadata
	0,  // 11: proto.rpc.webrtc.v1.ResponseMessage.packet_message:type_name -> proto.rpc.webrtc.v1.PacketMessage
	13, // 12: proto.rpc.webrtc.v1.ResponseTrailers.status:type_name -> google.rpc.Status
	10, // 13: proto.rpc.webrtc.v1.ResponseTrailers.metadata:type_name -> proto.rpc.webrtc.v1.Metadata
	11, // 14: proto.rpc.webrtc.v1.Metadata.md:type_name -> proto.rpc.webrtc.v1.Metadata.MdEntry
	9,  // 15: proto.rpc.webrtc.v1.Metadata.MdEntry.value:type_name -> proto.rpc.webrtc.v1.Strings
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_proto_rpc_webrtc_v1_grpc_proto_init() }
func file_proto_rpc_webrtc_v1_grpc_proto_init() {
	if File_proto_rpc_webrtc_v1_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketMessage); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeaders); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseHeaders); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseTrailers); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strings); i {
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
		file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
	file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Request_Headers)(nil),
		(*Request_Message)(nil),
	}
	file_proto_rpc_webrtc_v1_grpc_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Response_Headers)(nil),
		(*Response_Message)(nil),
		(*Response_Trailers)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rpc_webrtc_v1_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_rpc_webrtc_v1_grpc_proto_goTypes,
		DependencyIndexes: file_proto_rpc_webrtc_v1_grpc_proto_depIdxs,
		MessageInfos:      file_proto_rpc_webrtc_v1_grpc_proto_msgTypes,
	}.Build()
	File_proto_rpc_webrtc_v1_grpc_proto = out.File
	file_proto_rpc_webrtc_v1_grpc_proto_rawDesc = nil
	file_proto_rpc_webrtc_v1_grpc_proto_goTypes = nil
	file_proto_rpc_webrtc_v1_grpc_proto_depIdxs = nil
}
