// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: examples/internal/proto/examplepb/response_body_service.proto

package examplepb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RepeatedResponseBodyOut_Response_ResponseType int32

const (
	// UNKNOWN
	RepeatedResponseBodyOut_Response_UNKNOWN RepeatedResponseBodyOut_Response_ResponseType = 0
	// A is 1
	RepeatedResponseBodyOut_Response_A RepeatedResponseBodyOut_Response_ResponseType = 1
	// B is 2
	RepeatedResponseBodyOut_Response_B RepeatedResponseBodyOut_Response_ResponseType = 2
)

// Enum value maps for RepeatedResponseBodyOut_Response_ResponseType.
var (
	RepeatedResponseBodyOut_Response_ResponseType_name = map[int32]string{
		0: "UNKNOWN",
		1: "A",
		2: "B",
	}
	RepeatedResponseBodyOut_Response_ResponseType_value = map[string]int32{
		"UNKNOWN": 0,
		"A":       1,
		"B":       2,
	}
)

func (x RepeatedResponseBodyOut_Response_ResponseType) Enum() *RepeatedResponseBodyOut_Response_ResponseType {
	p := new(RepeatedResponseBodyOut_Response_ResponseType)
	*p = x
	return p
}

func (x RepeatedResponseBodyOut_Response_ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RepeatedResponseBodyOut_Response_ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_examples_internal_proto_examplepb_response_body_service_proto_enumTypes[0].Descriptor()
}

func (RepeatedResponseBodyOut_Response_ResponseType) Type() protoreflect.EnumType {
	return &file_examples_internal_proto_examplepb_response_body_service_proto_enumTypes[0]
}

func (x RepeatedResponseBodyOut_Response_ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RepeatedResponseBodyOut_Response_ResponseType.Descriptor instead.
func (RepeatedResponseBodyOut_Response_ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{2, 0, 0}
}

type ResponseBodyIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseBodyIn) Reset() {
	*x = ResponseBodyIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBodyIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBodyIn) ProtoMessage() {}

func (x *ResponseBodyIn) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBodyIn.ProtoReflect.Descriptor instead.
func (*ResponseBodyIn) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseBodyIn) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type ResponseBodyOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *ResponseBodyOut_Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ResponseBodyOut) Reset() {
	*x = ResponseBodyOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBodyOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBodyOut) ProtoMessage() {}

func (x *ResponseBodyOut) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBodyOut.ProtoReflect.Descriptor instead.
func (*ResponseBodyOut) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseBodyOut) GetResponse() *ResponseBodyOut_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type RepeatedResponseBodyOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*RepeatedResponseBodyOut_Response `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *RepeatedResponseBodyOut) Reset() {
	*x = RepeatedResponseBodyOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedResponseBodyOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedResponseBodyOut) ProtoMessage() {}

func (x *RepeatedResponseBodyOut) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedResponseBodyOut.ProtoReflect.Descriptor instead.
func (*RepeatedResponseBodyOut) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{2}
}

func (x *RepeatedResponseBodyOut) GetResponse() []*RepeatedResponseBodyOut_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type RepeatedResponseStrings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *RepeatedResponseStrings) Reset() {
	*x = RepeatedResponseStrings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedResponseStrings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedResponseStrings) ProtoMessage() {}

func (x *RepeatedResponseStrings) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedResponseStrings.ProtoReflect.Descriptor instead.
func (*RepeatedResponseStrings) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{3}
}

func (x *RepeatedResponseStrings) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type ResponseBodyOut_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseBodyOut_Response) Reset() {
	*x = ResponseBodyOut_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBodyOut_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBodyOut_Response) ProtoMessage() {}

func (x *ResponseBodyOut_Response) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBodyOut_Response.ProtoReflect.Descriptor instead.
func (*ResponseBodyOut_Response) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResponseBodyOut_Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type RepeatedResponseBodyOut_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string                                        `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Type RepeatedResponseBodyOut_Response_ResponseType `protobuf:"varint,3,opt,name=type,proto3,enum=grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut_Response_ResponseType" json:"type,omitempty"`
}

func (x *RepeatedResponseBodyOut_Response) Reset() {
	*x = RepeatedResponseBodyOut_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedResponseBodyOut_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedResponseBodyOut_Response) ProtoMessage() {}

func (x *RepeatedResponseBodyOut_Response) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedResponseBodyOut_Response.ProtoReflect.Descriptor instead.
func (*RepeatedResponseBodyOut_Response) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RepeatedResponseBodyOut_Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *RepeatedResponseBodyOut_Response) GetType() RepeatedResponseBodyOut_Response_ResponseType {
	if x != nil {
		return x.Type
	}
	return RepeatedResponseBodyOut_Response_UNKNOWN
}

var File_examples_internal_proto_examplepb_response_body_service_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_response_body_service_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x12, 0x64, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1e, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc6, 0x02,
	0x0a, 0x17, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x12, 0x6c, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x4f, 0x75, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0xbc, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x71, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x12,
	0x05, 0x0a, 0x01, 0x42, 0x10, 0x02, 0x22, 0x31, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xb2, 0x06, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xba, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x49, 0x6e, 0x1a, 0x3f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x14,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x6f, 0x64, 0x79, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x61, 0x7d, 0x62, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xc7,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x6f, 0x64, 0x79, 0x49, 0x6e, 0x1a, 0x47, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x22, 0x28,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x16, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x6f, 0x64, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x7d, 0x62, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xc7, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e,
	0x1a, 0x47, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x12, 0x17, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x7d, 0x62, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0xc9, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3e, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x49, 0x6e, 0x1a, 0x3f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x4f, 0x75, 0x74, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x1b, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x62, 0x6f, 0x64, 0x79, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x7b, 0x64, 0x61, 0x74,
	0x61, 0x7d, 0x62, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4d,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_response_body_service_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_response_body_service_proto_rawDescData = file_examples_internal_proto_examplepb_response_body_service_proto_rawDesc
)

func file_examples_internal_proto_examplepb_response_body_service_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_response_body_service_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_response_body_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_response_body_service_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_response_body_service_proto_rawDescData
}

var file_examples_internal_proto_examplepb_response_body_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_examples_internal_proto_examplepb_response_body_service_proto_goTypes = []interface{}{
	(RepeatedResponseBodyOut_Response_ResponseType)(0), // 0: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.Response.ResponseType
	(*ResponseBodyIn)(nil),                             // 1: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyIn
	(*ResponseBodyOut)(nil),                            // 2: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut
	(*RepeatedResponseBodyOut)(nil),                    // 3: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut
	(*RepeatedResponseStrings)(nil),                    // 4: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseStrings
	(*ResponseBodyOut_Response)(nil),                   // 5: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut.Response
	(*RepeatedResponseBodyOut_Response)(nil),           // 6: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.Response
}
var file_examples_internal_proto_examplepb_response_body_service_proto_depIdxs = []int32{
	5, // 0: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut.response:type_name -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut.Response
	6, // 1: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.response:type_name -> grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.Response
	0, // 2: grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.Response.type:type_name -> grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut.Response.ResponseType
	1, // 3: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.GetResponseBody:input_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyIn
	1, // 4: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.ListResponseBodies:input_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyIn
	1, // 5: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.ListResponseStrings:input_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyIn
	1, // 6: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.GetResponseBodyStream:input_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyIn
	2, // 7: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.GetResponseBody:output_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut
	3, // 8: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.ListResponseBodies:output_type -> grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseBodyOut
	4, // 9: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.ListResponseStrings:output_type -> grpc.gateway.examples.internal.proto.examplepb.RepeatedResponseStrings
	2, // 10: grpc.gateway.examples.internal.proto.examplepb.ResponseBodyService.GetResponseBodyStream:output_type -> grpc.gateway.examples.internal.proto.examplepb.ResponseBodyOut
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_response_body_service_proto_init() }
func file_examples_internal_proto_examplepb_response_body_service_proto_init() {
	if File_examples_internal_proto_examplepb_response_body_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBodyIn); i {
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
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBodyOut); i {
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
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedResponseBodyOut); i {
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
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedResponseStrings); i {
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
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBodyOut_Response); i {
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
		file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedResponseBodyOut_Response); i {
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
			RawDescriptor: file_examples_internal_proto_examplepb_response_body_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_response_body_service_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_response_body_service_proto_depIdxs,
		EnumInfos:         file_examples_internal_proto_examplepb_response_body_service_proto_enumTypes,
		MessageInfos:      file_examples_internal_proto_examplepb_response_body_service_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_response_body_service_proto = out.File
	file_examples_internal_proto_examplepb_response_body_service_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_response_body_service_proto_goTypes = nil
	file_examples_internal_proto_examplepb_response_body_service_proto_depIdxs = nil
}
