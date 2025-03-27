// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: protobuf/entity/errors.proto

package entity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotFoundError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *string                `protobuf:"bytes,1,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotFoundError) Reset() {
	*x = NotFoundError{}
	mi := &file_protobuf_entity_errors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotFoundError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotFoundError) ProtoMessage() {}

func (x *NotFoundError) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_entity_errors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotFoundError.ProtoReflect.Descriptor instead.
func (*NotFoundError) Descriptor() ([]byte, []int) {
	return file_protobuf_entity_errors_proto_rawDescGZIP(), []int{0}
}

func (x *NotFoundError) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type BadRequestError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *string                `protobuf:"bytes,1,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BadRequestError) Reset() {
	*x = BadRequestError{}
	mi := &file_protobuf_entity_errors_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BadRequestError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadRequestError) ProtoMessage() {}

func (x *BadRequestError) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_entity_errors_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadRequestError.ProtoReflect.Descriptor instead.
func (*BadRequestError) Descriptor() ([]byte, []int) {
	return file_protobuf_entity_errors_proto_rawDescGZIP(), []int{1}
}

func (x *BadRequestError) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type InternalError struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *string                `protobuf:"bytes,1,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InternalError) Reset() {
	*x = InternalError{}
	mi := &file_protobuf_entity_errors_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InternalError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalError) ProtoMessage() {}

func (x *InternalError) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_entity_errors_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalError.ProtoReflect.Descriptor instead.
func (*InternalError) Descriptor() ([]byte, []int) {
	return file_protobuf_entity_errors_proto_rawDescGZIP(), []int{2}
}

func (x *InternalError) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_protobuf_entity_errors_proto protoreflect.FileDescriptor

const file_protobuf_entity_errors_proto_rawDesc = "" +
	"\n" +
	"\x1cprotobuf/entity/errors.proto\":\n" +
	"\rNotFoundError\x12\x1d\n" +
	"\amessage\x18\x01 \x01(\tH\x00R\amessage\x88\x01\x01B\n" +
	"\n" +
	"\b_message\"<\n" +
	"\x0fBadRequestError\x12\x1d\n" +
	"\amessage\x18\x01 \x01(\tH\x00R\amessage\x88\x01\x01B\n" +
	"\n" +
	"\b_message\":\n" +
	"\rInternalError\x12\x1d\n" +
	"\amessage\x18\x01 \x01(\tH\x00R\amessage\x88\x01\x01B\n" +
	"\n" +
	"\b_messageB?Z=github.com/Sujin1135/coupon-service-interface/protobuf/entityb\x06proto3"

var (
	file_protobuf_entity_errors_proto_rawDescOnce sync.Once
	file_protobuf_entity_errors_proto_rawDescData []byte
)

func file_protobuf_entity_errors_proto_rawDescGZIP() []byte {
	file_protobuf_entity_errors_proto_rawDescOnce.Do(func() {
		file_protobuf_entity_errors_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobuf_entity_errors_proto_rawDesc), len(file_protobuf_entity_errors_proto_rawDesc)))
	})
	return file_protobuf_entity_errors_proto_rawDescData
}

var file_protobuf_entity_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_entity_errors_proto_goTypes = []any{
	(*NotFoundError)(nil),   // 0: NotFoundError
	(*BadRequestError)(nil), // 1: BadRequestError
	(*InternalError)(nil),   // 2: InternalError
}
var file_protobuf_entity_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_entity_errors_proto_init() }
func file_protobuf_entity_errors_proto_init() {
	if File_protobuf_entity_errors_proto != nil {
		return
	}
	file_protobuf_entity_errors_proto_msgTypes[0].OneofWrappers = []any{}
	file_protobuf_entity_errors_proto_msgTypes[1].OneofWrappers = []any{}
	file_protobuf_entity_errors_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobuf_entity_errors_proto_rawDesc), len(file_protobuf_entity_errors_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_entity_errors_proto_goTypes,
		DependencyIndexes: file_protobuf_entity_errors_proto_depIdxs,
		MessageInfos:      file_protobuf_entity_errors_proto_msgTypes,
	}.Build()
	File_protobuf_entity_errors_proto = out.File
	file_protobuf_entity_errors_proto_goTypes = nil
	file_protobuf_entity_errors_proto_depIdxs = nil
}
