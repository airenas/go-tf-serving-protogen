// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/tools/proto_splitter/testdata/test_message.proto

package testdata

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

type RepeatedString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings []string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
}

func (x *RepeatedString) Reset() {
	*x = RepeatedString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedString) ProtoMessage() {}

func (x *RepeatedString) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedString.ProtoReflect.Descriptor instead.
func (*RepeatedString) Descriptor() ([]byte, []int) {
	return file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescGZIP(), []int{0}
}

func (x *RepeatedString) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type RepeatedRepeatedString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FillerField int32             `protobuf:"varint,1,opt,name=filler_field,json=fillerField,proto3" json:"filler_field,omitempty"`
	Rs          []*RepeatedString `protobuf:"bytes,2,rep,name=rs,proto3" json:"rs,omitempty"`
}

func (x *RepeatedRepeatedString) Reset() {
	*x = RepeatedRepeatedString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedRepeatedString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedRepeatedString) ProtoMessage() {}

func (x *RepeatedRepeatedString) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedRepeatedString.ProtoReflect.Descriptor instead.
func (*RepeatedRepeatedString) Descriptor() ([]byte, []int) {
	return file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescGZIP(), []int{1}
}

func (x *RepeatedRepeatedString) GetFillerField() int32 {
	if x != nil {
		return x.FillerField
	}
	return 0
}

func (x *RepeatedRepeatedString) GetRs() []*RepeatedString {
	if x != nil {
		return x.Rs
	}
	return nil
}

type ManyFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldOne            *ManyFields          `protobuf:"bytes,1,opt,name=field_one,json=fieldOne,proto3" json:"field_one,omitempty"`
	RepeatedField       []*ManyFields        `protobuf:"bytes,2,rep,name=repeated_field,json=repeatedField,proto3" json:"repeated_field,omitempty"`
	StringField         string               `protobuf:"bytes,3,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	RepeatedStringField []string             `protobuf:"bytes,4,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
	MapFieldUint32      map[uint32]string    `protobuf:"bytes,5,rep,name=map_field_uint32,json=mapFieldUint32,proto3" json:"map_field_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapFieldInt64       map[int64]string     `protobuf:"bytes,6,rep,name=map_field_int64,json=mapFieldInt64,proto3" json:"map_field_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NestedMapBool       map[bool]*ManyFields `protobuf:"bytes,7,rep,name=nested_map_bool,json=nestedMapBool,proto3" json:"nested_map_bool,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ManyFields) Reset() {
	*x = ManyFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyFields) ProtoMessage() {}

func (x *ManyFields) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyFields.ProtoReflect.Descriptor instead.
func (*ManyFields) Descriptor() ([]byte, []int) {
	return file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescGZIP(), []int{2}
}

func (x *ManyFields) GetFieldOne() *ManyFields {
	if x != nil {
		return x.FieldOne
	}
	return nil
}

func (x *ManyFields) GetRepeatedField() []*ManyFields {
	if x != nil {
		return x.RepeatedField
	}
	return nil
}

func (x *ManyFields) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *ManyFields) GetRepeatedStringField() []string {
	if x != nil {
		return x.RepeatedStringField
	}
	return nil
}

func (x *ManyFields) GetMapFieldUint32() map[uint32]string {
	if x != nil {
		return x.MapFieldUint32
	}
	return nil
}

func (x *ManyFields) GetMapFieldInt64() map[int64]string {
	if x != nil {
		return x.MapFieldInt64
	}
	return nil
}

func (x *ManyFields) GetNestedMapBool() map[bool]*ManyFields {
	if x != nil {
		return x.NestedMapBool
	}
	return nil
}

type StringNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val        string        `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	ChildNodes []*StringNode `protobuf:"bytes,2,rep,name=child_nodes,json=childNodes,proto3" json:"child_nodes,omitempty"`
}

func (x *StringNode) Reset() {
	*x = StringNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringNode) ProtoMessage() {}

func (x *StringNode) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringNode.ProtoReflect.Descriptor instead.
func (*StringNode) Descriptor() ([]byte, []int) {
	return file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescGZIP(), []int{3}
}

func (x *StringNode) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *StringNode) GetChildNodes() []*StringNode {
	if x != nil {
		return x.ChildNodes
	}
	return nil
}

var File_tensorflow_tools_proto_splitter_testdata_test_message_proto protoreflect.FileDescriptor

var file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x74, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x37, 0x0a, 0x02, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x02, 0x72, 0x73, 0x22, 0x80, 0x06, 0x0a, 0x0a, 0x4d, 0x61, 0x6e,
	0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x6e, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x61, 0x0a, 0x10,
	0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73,
	0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0e, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x5e, 0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x5e, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x1a,
	0x41, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x65, 0x0a, 0x12, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x61,
	0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x64, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x0b, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x42, 0xf7, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x15, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x54,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0xca, 0x02, 0x15, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53,
	0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0xe2,
	0x02, 0x21, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x54,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescOnce sync.Once
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescData = file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDesc
)

func file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescGZIP() []byte {
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescOnce.Do(func() {
		file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescData)
	})
	return file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDescData
}

var file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_tools_proto_splitter_testdata_test_message_proto_goTypes = []interface{}{
	(*RepeatedString)(nil),         // 0: proto_splitter_testdata.RepeatedString
	(*RepeatedRepeatedString)(nil), // 1: proto_splitter_testdata.RepeatedRepeatedString
	(*ManyFields)(nil),             // 2: proto_splitter_testdata.ManyFields
	(*StringNode)(nil),             // 3: proto_splitter_testdata.StringNode
	nil,                            // 4: proto_splitter_testdata.ManyFields.MapFieldUint32Entry
	nil,                            // 5: proto_splitter_testdata.ManyFields.MapFieldInt64Entry
	nil,                            // 6: proto_splitter_testdata.ManyFields.NestedMapBoolEntry
}
var file_tensorflow_tools_proto_splitter_testdata_test_message_proto_depIdxs = []int32{
	0, // 0: proto_splitter_testdata.RepeatedRepeatedString.rs:type_name -> proto_splitter_testdata.RepeatedString
	2, // 1: proto_splitter_testdata.ManyFields.field_one:type_name -> proto_splitter_testdata.ManyFields
	2, // 2: proto_splitter_testdata.ManyFields.repeated_field:type_name -> proto_splitter_testdata.ManyFields
	4, // 3: proto_splitter_testdata.ManyFields.map_field_uint32:type_name -> proto_splitter_testdata.ManyFields.MapFieldUint32Entry
	5, // 4: proto_splitter_testdata.ManyFields.map_field_int64:type_name -> proto_splitter_testdata.ManyFields.MapFieldInt64Entry
	6, // 5: proto_splitter_testdata.ManyFields.nested_map_bool:type_name -> proto_splitter_testdata.ManyFields.NestedMapBoolEntry
	3, // 6: proto_splitter_testdata.StringNode.child_nodes:type_name -> proto_splitter_testdata.StringNode
	2, // 7: proto_splitter_testdata.ManyFields.NestedMapBoolEntry.value:type_name -> proto_splitter_testdata.ManyFields
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_tools_proto_splitter_testdata_test_message_proto_init() }
func file_tensorflow_tools_proto_splitter_testdata_test_message_proto_init() {
	if File_tensorflow_tools_proto_splitter_testdata_test_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedString); i {
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
		file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedRepeatedString); i {
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
		file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyFields); i {
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
		file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringNode); i {
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
			RawDescriptor: file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tools_proto_splitter_testdata_test_message_proto_goTypes,
		DependencyIndexes: file_tensorflow_tools_proto_splitter_testdata_test_message_proto_depIdxs,
		MessageInfos:      file_tensorflow_tools_proto_splitter_testdata_test_message_proto_msgTypes,
	}.Build()
	File_tensorflow_tools_proto_splitter_testdata_test_message_proto = out.File
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_rawDesc = nil
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_goTypes = nil
	file_tensorflow_tools_proto_splitter_testdata_test_message_proto_depIdxs = nil
}
