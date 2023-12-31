// Test messages used in compare_test.py.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/python/util/protobuf/compare_test.proto

package protobuf

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

type Enum int32

const (
	Enum_A Enum = 0
	Enum_B Enum = 1
	Enum_C Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "A",
		1: "B",
		2: "C",
	}
	Enum_value = map[string]int32{
		"A": 0,
		"B": 1,
		"C": 2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_python_util_protobuf_compare_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_tensorflow_python_util_protobuf_compare_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum(num)
	return nil
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{0}
}

type Small struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings []string `protobuf:"bytes,1,rep,name=strings" json:"strings,omitempty"`
}

func (x *Small) Reset() {
	*x = Small{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Small) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Small) ProtoMessage() {}

func (x *Small) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Small.ProtoReflect.Descriptor instead.
func (*Small) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{0}
}

func (x *Small) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type Medium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32S []int32          `protobuf:"varint,1,rep,name=int32s" json:"int32s,omitempty"`
	Smalls []*Small         `protobuf:"bytes,2,rep,name=smalls" json:"smalls,omitempty"`
	Groupa []*Medium_GroupA `protobuf:"group,3,rep,name=GroupA,json=groupa" json:"groupa,omitempty"`
	Floats []float32        `protobuf:"fixed32,6,rep,name=floats" json:"floats,omitempty"`
}

func (x *Medium) Reset() {
	*x = Medium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medium) ProtoMessage() {}

func (x *Medium) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medium.ProtoReflect.Descriptor instead.
func (*Medium) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{1}
}

func (x *Medium) GetInt32S() []int32 {
	if x != nil {
		return x.Int32S
	}
	return nil
}

func (x *Medium) GetSmalls() []*Small {
	if x != nil {
		return x.Smalls
	}
	return nil
}

func (x *Medium) GetGroupa() []*Medium_GroupA {
	if x != nil {
		return x.Groupa
	}
	return nil
}

func (x *Medium) GetFloats() []float32 {
	if x != nil {
		return x.Floats
	}
	return nil
}

type Large struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	String_ *string  `protobuf:"bytes,1,opt,name=string_,json=string" json:"string_,omitempty"`
	Int64_  *int64   `protobuf:"varint,2,opt,name=int64_,json=int64" json:"int64_,omitempty"`
	Float_  *float32 `protobuf:"fixed32,3,opt,name=float_,json=float" json:"float_,omitempty"`
	Bool_   *bool    `protobuf:"varint,4,opt,name=bool_,json=bool" json:"bool_,omitempty"`
	Enum_   *Enum    `protobuf:"varint,5,opt,name=enum_,json=enum,enum=compare_test.Enum" json:"enum_,omitempty"`
	Int64S  []int64  `protobuf:"varint,6,rep,name=int64s" json:"int64s,omitempty"`
	Medium  *Medium  `protobuf:"bytes,7,opt,name=medium" json:"medium,omitempty"`
	Small   *Small   `protobuf:"bytes,8,opt,name=small" json:"small,omitempty"`
	Double_ *float64 `protobuf:"fixed64,9,opt,name=double_,json=double" json:"double_,omitempty"`
	WithMap *WithMap `protobuf:"bytes,10,opt,name=with_map,json=withMap" json:"with_map,omitempty"`
}

func (x *Large) Reset() {
	*x = Large{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Large) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Large) ProtoMessage() {}

func (x *Large) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Large.ProtoReflect.Descriptor instead.
func (*Large) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{2}
}

func (x *Large) GetString_() string {
	if x != nil && x.String_ != nil {
		return *x.String_
	}
	return ""
}

func (x *Large) GetInt64_() int64 {
	if x != nil && x.Int64_ != nil {
		return *x.Int64_
	}
	return 0
}

func (x *Large) GetFloat_() float32 {
	if x != nil && x.Float_ != nil {
		return *x.Float_
	}
	return 0
}

func (x *Large) GetBool_() bool {
	if x != nil && x.Bool_ != nil {
		return *x.Bool_
	}
	return false
}

func (x *Large) GetEnum_() Enum {
	if x != nil && x.Enum_ != nil {
		return *x.Enum_
	}
	return Enum_A
}

func (x *Large) GetInt64S() []int64 {
	if x != nil {
		return x.Int64S
	}
	return nil
}

func (x *Large) GetMedium() *Medium {
	if x != nil {
		return x.Medium
	}
	return nil
}

func (x *Large) GetSmall() *Small {
	if x != nil {
		return x.Small
	}
	return nil
}

func (x *Large) GetDouble_() float64 {
	if x != nil && x.Double_ != nil {
		return *x.Double_
	}
	return 0
}

func (x *Large) GetWithMap() *WithMap {
	if x != nil {
		return x.WithMap
	}
	return nil
}

type Labeled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required *int32 `protobuf:"varint,1,req,name=required" json:"required,omitempty"`
	Optional *int32 `protobuf:"varint,2,opt,name=optional" json:"optional,omitempty"`
}

func (x *Labeled) Reset() {
	*x = Labeled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labeled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labeled) ProtoMessage() {}

func (x *Labeled) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labeled.ProtoReflect.Descriptor instead.
func (*Labeled) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{3}
}

func (x *Labeled) GetRequired() int32 {
	if x != nil && x.Required != nil {
		return *x.Required
	}
	return 0
}

func (x *Labeled) GetOptional() int32 {
	if x != nil && x.Optional != nil {
		return *x.Optional
	}
	return 0
}

type WithMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueMessage map[int32]*Small  `protobuf:"bytes,1,rep,name=value_message,json=valueMessage" json:"value_message,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ValueString  map[string]string `protobuf:"bytes,2,rep,name=value_string,json=valueString" json:"value_string,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *WithMap) Reset() {
	*x = WithMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithMap) ProtoMessage() {}

func (x *WithMap) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithMap.ProtoReflect.Descriptor instead.
func (*WithMap) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{4}
}

func (x *WithMap) GetValueMessage() map[int32]*Small {
	if x != nil {
		return x.ValueMessage
	}
	return nil
}

func (x *WithMap) GetValueString() map[string]string {
	if x != nil {
		return x.ValueString
	}
	return nil
}

type Floats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Float_  *float32 `protobuf:"fixed32,1,opt,name=float_,json=float" json:"float_,omitempty"`
	Double_ *float64 `protobuf:"fixed64,2,opt,name=double_,json=double" json:"double_,omitempty"`
}

func (x *Floats) Reset() {
	*x = Floats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Floats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Floats) ProtoMessage() {}

func (x *Floats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Floats.ProtoReflect.Descriptor instead.
func (*Floats) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{5}
}

func (x *Floats) GetFloat_() float32 {
	if x != nil && x.Float_ != nil {
		return *x.Float_
	}
	return 0
}

func (x *Floats) GetDouble_() float64 {
	if x != nil && x.Double_ != nil {
		return *x.Double_
	}
	return 0
}

type RepeatedFloats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Float_  []float32 `protobuf:"fixed32,1,rep,packed,name=float_,json=float" json:"float_,omitempty"`
	Double_ []float64 `protobuf:"fixed64,2,rep,packed,name=double_,json=double" json:"double_,omitempty"`
}

func (x *RepeatedFloats) Reset() {
	*x = RepeatedFloats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepeatedFloats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepeatedFloats) ProtoMessage() {}

func (x *RepeatedFloats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepeatedFloats.ProtoReflect.Descriptor instead.
func (*RepeatedFloats) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{6}
}

func (x *RepeatedFloats) GetFloat_() []float32 {
	if x != nil {
		return x.Float_
	}
	return nil
}

func (x *RepeatedFloats) GetDouble_() []float64 {
	if x != nil {
		return x.Double_
	}
	return nil
}

type NestedFloats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Floats *Floats `protobuf:"bytes,1,opt,name=floats" json:"floats,omitempty"`
}

func (x *NestedFloats) Reset() {
	*x = NestedFloats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedFloats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedFloats) ProtoMessage() {}

func (x *NestedFloats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedFloats.ProtoReflect.Descriptor instead.
func (*NestedFloats) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{7}
}

func (x *NestedFloats) GetFloats() *Floats {
	if x != nil {
		return x.Floats
	}
	return nil
}

type MapFloats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntToFloats map[int64]*Floats `protobuf:"bytes,1,rep,name=int_to_floats,json=intToFloats" json:"int_to_floats,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *MapFloats) Reset() {
	*x = MapFloats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapFloats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapFloats) ProtoMessage() {}

func (x *MapFloats) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapFloats.ProtoReflect.Descriptor instead.
func (*MapFloats) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{8}
}

func (x *MapFloats) GetIntToFloats() map[int64]*Floats {
	if x != nil {
		return x.IntToFloats
	}
	return nil
}

type Medium_GroupA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groupb []*Medium_GroupA_GroupB `protobuf:"group,4,rep,name=GroupB,json=groupb" json:"groupb,omitempty"`
}

func (x *Medium_GroupA) Reset() {
	*x = Medium_GroupA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medium_GroupA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medium_GroupA) ProtoMessage() {}

func (x *Medium_GroupA) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medium_GroupA.ProtoReflect.Descriptor instead.
func (*Medium_GroupA) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Medium_GroupA) GetGroupb() []*Medium_GroupA_GroupB {
	if x != nil {
		return x.Groupb
	}
	return nil
}

type Medium_GroupA_GroupB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings *string `protobuf:"bytes,5,req,name=strings" json:"strings,omitempty"`
}

func (x *Medium_GroupA_GroupB) Reset() {
	*x = Medium_GroupA_GroupB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medium_GroupA_GroupB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medium_GroupA_GroupB) ProtoMessage() {}

func (x *Medium_GroupA_GroupB) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medium_GroupA_GroupB.ProtoReflect.Descriptor instead.
func (*Medium_GroupA_GroupB) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Medium_GroupA_GroupB) GetStrings() string {
	if x != nil && x.Strings != nil {
		return *x.Strings
	}
	return ""
}

var File_tensorflow_python_util_protobuf_compare_test_proto protoreflect.FileDescriptor

var file_tensorflow_python_util_protobuf_compare_test_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74,
	0x68, 0x6f, 0x6e, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x21, 0x0a, 0x05, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x6d, 0x61, 0x6c,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x73,
	0x6d, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x73, 0x1a, 0x68, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x12, 0x3a, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x64, 0x69,
	0x75, 0x6d, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x1a, 0x22, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x42, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xc8, 0x02, 0x0a,
	0x05, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x15, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x13, 0x0a,
	0x05, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6d, 0x61,
	0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x52, 0x07,
	0x77, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x22, 0x41, 0x0a, 0x07, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0xb8, 0x02, 0x0a, 0x07, 0x57,
	0x69, 0x74, 0x68, 0x4d, 0x61, 0x70, 0x12, 0x4c, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x61, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61,
	0x70, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x54, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x06, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22,
	0x48, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x73, 0x12, 0x19, 0x0a, 0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x02, 0x42, 0x02, 0x10, 0x01, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x07,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x52,
	0x06, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x54, 0x6f, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x54, 0x6f, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x73, 0x1a, 0x54, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x54, 0x6f, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x1b, 0x0a, 0x04, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x01, 0x12,
	0x05, 0x0a, 0x01, 0x43, 0x10, 0x02, 0x42, 0xbe, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x42, 0x10, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65,
	0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x43, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x65, 0x73,
	0x74, 0xca, 0x02, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x65, 0x73, 0x74, 0xe2,
	0x02, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x65, 0x54, 0x65, 0x73, 0x74,
}

var (
	file_tensorflow_python_util_protobuf_compare_test_proto_rawDescOnce sync.Once
	file_tensorflow_python_util_protobuf_compare_test_proto_rawDescData = file_tensorflow_python_util_protobuf_compare_test_proto_rawDesc
)

func file_tensorflow_python_util_protobuf_compare_test_proto_rawDescGZIP() []byte {
	file_tensorflow_python_util_protobuf_compare_test_proto_rawDescOnce.Do(func() {
		file_tensorflow_python_util_protobuf_compare_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_python_util_protobuf_compare_test_proto_rawDescData)
	})
	return file_tensorflow_python_util_protobuf_compare_test_proto_rawDescData
}

var file_tensorflow_python_util_protobuf_compare_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_tensorflow_python_util_protobuf_compare_test_proto_goTypes = []interface{}{
	(Enum)(0),                    // 0: compare_test.Enum
	(*Small)(nil),                // 1: compare_test.Small
	(*Medium)(nil),               // 2: compare_test.Medium
	(*Large)(nil),                // 3: compare_test.Large
	(*Labeled)(nil),              // 4: compare_test.Labeled
	(*WithMap)(nil),              // 5: compare_test.WithMap
	(*Floats)(nil),               // 6: compare_test.Floats
	(*RepeatedFloats)(nil),       // 7: compare_test.RepeatedFloats
	(*NestedFloats)(nil),         // 8: compare_test.NestedFloats
	(*MapFloats)(nil),            // 9: compare_test.MapFloats
	(*Medium_GroupA)(nil),        // 10: compare_test.Medium.GroupA
	(*Medium_GroupA_GroupB)(nil), // 11: compare_test.Medium.GroupA.GroupB
	nil,                          // 12: compare_test.WithMap.ValueMessageEntry
	nil,                          // 13: compare_test.WithMap.ValueStringEntry
	nil,                          // 14: compare_test.MapFloats.IntToFloatsEntry
}
var file_tensorflow_python_util_protobuf_compare_test_proto_depIdxs = []int32{
	1,  // 0: compare_test.Medium.smalls:type_name -> compare_test.Small
	10, // 1: compare_test.Medium.groupa:type_name -> compare_test.Medium.GroupA
	0,  // 2: compare_test.Large.enum_:type_name -> compare_test.Enum
	2,  // 3: compare_test.Large.medium:type_name -> compare_test.Medium
	1,  // 4: compare_test.Large.small:type_name -> compare_test.Small
	5,  // 5: compare_test.Large.with_map:type_name -> compare_test.WithMap
	12, // 6: compare_test.WithMap.value_message:type_name -> compare_test.WithMap.ValueMessageEntry
	13, // 7: compare_test.WithMap.value_string:type_name -> compare_test.WithMap.ValueStringEntry
	6,  // 8: compare_test.NestedFloats.floats:type_name -> compare_test.Floats
	14, // 9: compare_test.MapFloats.int_to_floats:type_name -> compare_test.MapFloats.IntToFloatsEntry
	11, // 10: compare_test.Medium.GroupA.groupb:type_name -> compare_test.Medium.GroupA.GroupB
	1,  // 11: compare_test.WithMap.ValueMessageEntry.value:type_name -> compare_test.Small
	6,  // 12: compare_test.MapFloats.IntToFloatsEntry.value:type_name -> compare_test.Floats
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_tensorflow_python_util_protobuf_compare_test_proto_init() }
func file_tensorflow_python_util_protobuf_compare_test_proto_init() {
	if File_tensorflow_python_util_protobuf_compare_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Small); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medium); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Large); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labeled); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithMap); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Floats); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepeatedFloats); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedFloats); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapFloats); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medium_GroupA); i {
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
		file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medium_GroupA_GroupB); i {
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
			RawDescriptor: file_tensorflow_python_util_protobuf_compare_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_python_util_protobuf_compare_test_proto_goTypes,
		DependencyIndexes: file_tensorflow_python_util_protobuf_compare_test_proto_depIdxs,
		EnumInfos:         file_tensorflow_python_util_protobuf_compare_test_proto_enumTypes,
		MessageInfos:      file_tensorflow_python_util_protobuf_compare_test_proto_msgTypes,
	}.Build()
	File_tensorflow_python_util_protobuf_compare_test_proto = out.File
	file_tensorflow_python_util_protobuf_compare_test_proto_rawDesc = nil
	file_tensorflow_python_util_protobuf_compare_test_proto_goTypes = nil
	file_tensorflow_python_util_protobuf_compare_test_proto_depIdxs = nil
}
