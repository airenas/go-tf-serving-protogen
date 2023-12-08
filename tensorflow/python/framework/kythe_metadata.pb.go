//
// Copyright 2019 The Kythe Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// The content of this file is copied from
// https://github.com/kythe/kythe/blob/master/kythe/proto/metadata.proto and
// https://github.com/kythe/kythe/blob/master/kythe/proto/storage.proto
// and shouldn't be modified unless the protos is outdated.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/python/framework/kythe_metadata.proto

package framework

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

type GeneratedCodeInfo_Type int32

const (
	GeneratedCodeInfo_NONE   GeneratedCodeInfo_Type = 0
	GeneratedCodeInfo_KYTHE0 GeneratedCodeInfo_Type = 1 // Initial metadata document type.
)

// Enum value maps for GeneratedCodeInfo_Type.
var (
	GeneratedCodeInfo_Type_name = map[int32]string{
		0: "NONE",
		1: "KYTHE0",
	}
	GeneratedCodeInfo_Type_value = map[string]int32{
		"NONE":   0,
		"KYTHE0": 1,
	}
)

func (x GeneratedCodeInfo_Type) Enum() *GeneratedCodeInfo_Type {
	p := new(GeneratedCodeInfo_Type)
	*p = x
	return p
}

func (x GeneratedCodeInfo_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeneratedCodeInfo_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[0].Descriptor()
}

func (GeneratedCodeInfo_Type) Type() protoreflect.EnumType {
	return &file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[0]
}

func (x GeneratedCodeInfo_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeneratedCodeInfo_Type.Descriptor instead.
func (GeneratedCodeInfo_Type) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type MappingRule_Type int32

const (
	MappingRule_NONE           MappingRule_Type = 0
	MappingRule_NOP            MappingRule_Type = 1 // Dummy rule that contains no relevant information.
	MappingRule_ANCHOR_DEFINES MappingRule_Type = 2 // Rule describing a generates edge between target
	// range and source definition.
	MappingRule_ANCHOR_ANCHOR MappingRule_Type = 3 // Rule describing an imputes edge between target range
)

// Enum value maps for MappingRule_Type.
var (
	MappingRule_Type_name = map[int32]string{
		0: "NONE",
		1: "NOP",
		2: "ANCHOR_DEFINES",
		3: "ANCHOR_ANCHOR",
	}
	MappingRule_Type_value = map[string]int32{
		"NONE":           0,
		"NOP":            1,
		"ANCHOR_DEFINES": 2,
		"ANCHOR_ANCHOR":  3,
	}
)

func (x MappingRule_Type) Enum() *MappingRule_Type {
	p := new(MappingRule_Type)
	*p = x
	return p
}

func (x MappingRule_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MappingRule_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[1].Descriptor()
}

func (MappingRule_Type) Type() protoreflect.EnumType {
	return &file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[1]
}

func (x MappingRule_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MappingRule_Type.Descriptor instead.
func (MappingRule_Type) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{1, 0}
}

type MappingRule_Semantic int32

const (
	MappingRule_SEMA_NONE       MappingRule_Semantic = 0
	MappingRule_SEMA_WRITE      MappingRule_Semantic = 1
	MappingRule_SEMA_READ_WRITE MappingRule_Semantic = 2
)

// Enum value maps for MappingRule_Semantic.
var (
	MappingRule_Semantic_name = map[int32]string{
		0: "SEMA_NONE",
		1: "SEMA_WRITE",
		2: "SEMA_READ_WRITE",
	}
	MappingRule_Semantic_value = map[string]int32{
		"SEMA_NONE":       0,
		"SEMA_WRITE":      1,
		"SEMA_READ_WRITE": 2,
	}
)

func (x MappingRule_Semantic) Enum() *MappingRule_Semantic {
	p := new(MappingRule_Semantic)
	*p = x
	return p
}

func (x MappingRule_Semantic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MappingRule_Semantic) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[2].Descriptor()
}

func (MappingRule_Semantic) Type() protoreflect.EnumType {
	return &file_tensorflow_python_framework_kythe_metadata_proto_enumTypes[2]
}

func (x MappingRule_Semantic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MappingRule_Semantic.Descriptor instead.
func (MappingRule_Semantic) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{1, 1}
}

// Schema for the JSON-encoded Kythe metadata describing the relationship
// between source and target code for generated code.
type GeneratedCodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type GeneratedCodeInfo_Type `protobuf:"varint,1,opt,name=type,proto3,enum=tensorflow.GeneratedCodeInfo_Type" json:"type,omitempty"`
	Meta []*MappingRule         `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty"` // Only relevant if type == kythe0.
}

func (x *GeneratedCodeInfo) Reset() {
	*x = GeneratedCodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratedCodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratedCodeInfo) ProtoMessage() {}

func (x *GeneratedCodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratedCodeInfo.ProtoReflect.Descriptor instead.
func (*GeneratedCodeInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *GeneratedCodeInfo) GetType() GeneratedCodeInfo_Type {
	if x != nil {
		return x.Type
	}
	return GeneratedCodeInfo_NONE
}

func (x *GeneratedCodeInfo) GetMeta() []*MappingRule {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Metadata for a single mapping between a generated source range and a node
// in the source language or file.
type MappingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MappingRule_Type `protobuf:"varint,1,opt,name=type,proto3,enum=tensorflow.MappingRule_Type" json:"type,omitempty"`
	// If type == anchor_defines, this should generally be a reverse generates
	// edge, %/kythe/edge/generates, indicating that the specified vname generated
	// the source range.
	// If type == anchor_anchor, this should generally be a forward imputes edge,
	// /kythe/edge/imputes, indicating that the range in the source file produced
	// the text in the target file.
	// If semantic is not NONE, this field is ignored and the identified
	// declaration at the indicated text range is given the associated semantic.
	Edge string `protobuf:"bytes,2,opt,name=edge,proto3" json:"edge,omitempty"`
	// Fields only relevant if type == anchor_defines.
	Vname *VName `protobuf:"bytes,3,opt,name=vname,proto3" json:"vname,omitempty"` // The semantic node in the source language
	// which generated the text range.
	Begin    uint32               `protobuf:"varint,4,opt,name=begin,proto3" json:"begin,omitempty"` // Beginning of the range to match in the generated text.
	End      uint32               `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`     // End of the range to match in the generated text.
	Semantic MappingRule_Semantic `protobuf:"varint,11,opt,name=semantic,proto3,enum=tensorflow.MappingRule_Semantic" json:"semantic,omitempty"`
	// Fields only relevant if type == anchor_anchor.
	SourceVname *VName `protobuf:"bytes,6,opt,name=source_vname,json=sourceVname,proto3" json:"source_vname,omitempty"` // Anchor node in the generating source file.
	// Note: the signature in this vname, if present,
	// will typically be replaced by the target indexer
	// using its own anchor-construction rules based on
	// source_begin and source_end.
	SourceBegin uint32 `protobuf:"varint,7,opt,name=source_begin,json=sourceBegin,proto3" json:"source_begin,omitempty"` // loc/start of the anchor node in the source file.
	SourceEnd   uint32 `protobuf:"varint,8,opt,name=source_end,json=sourceEnd,proto3" json:"source_end,omitempty"`       // loc/end of the anchor node in the source file.
	TargetBegin uint32 `protobuf:"varint,9,opt,name=target_begin,json=targetBegin,proto3" json:"target_begin,omitempty"` // Start of the range in the generated text.
	TargetEnd   uint32 `protobuf:"varint,10,opt,name=target_end,json=targetEnd,proto3" json:"target_end,omitempty"`      // End of the range in the generated text.
}

func (x *MappingRule) Reset() {
	*x = MappingRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MappingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MappingRule) ProtoMessage() {}

func (x *MappingRule) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MappingRule.ProtoReflect.Descriptor instead.
func (*MappingRule) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *MappingRule) GetType() MappingRule_Type {
	if x != nil {
		return x.Type
	}
	return MappingRule_NONE
}

func (x *MappingRule) GetEdge() string {
	if x != nil {
		return x.Edge
	}
	return ""
}

func (x *MappingRule) GetVname() *VName {
	if x != nil {
		return x.Vname
	}
	return nil
}

func (x *MappingRule) GetBegin() uint32 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *MappingRule) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *MappingRule) GetSemantic() MappingRule_Semantic {
	if x != nil {
		return x.Semantic
	}
	return MappingRule_SEMA_NONE
}

func (x *MappingRule) GetSourceVname() *VName {
	if x != nil {
		return x.SourceVname
	}
	return nil
}

func (x *MappingRule) GetSourceBegin() uint32 {
	if x != nil {
		return x.SourceBegin
	}
	return 0
}

func (x *MappingRule) GetSourceEnd() uint32 {
	if x != nil {
		return x.SourceEnd
	}
	return 0
}

func (x *MappingRule) GetTargetBegin() uint32 {
	if x != nil {
		return x.TargetBegin
	}
	return 0
}

func (x *MappingRule) GetTargetEnd() uint32 {
	if x != nil {
		return x.TargetEnd
	}
	return 0
}

type VName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A language-specific signature assigned by the analyzer.
	// e.g., "com.google.common.collect.Lists.newLinkedList<#1>()"
	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The corpus this name belongs to.
	// e.g., "kythe", "chromium", "github.com/creachadair/imath", "aosp"
	// The corpus label "kythe" is reserved for internal use.
	Corpus string `protobuf:"bytes,2,opt,name=corpus,proto3" json:"corpus,omitempty"`
	// A corpus-specific root label, designating a subordinate collection within
	// the corpus.  If a corpus stores files in unrelated directory structures,
	// for example, the root can be used to distinguish them.  Or, if a corpus
	// incorporates subprojects, the root can be a project ID that it governs.
	// This may also be used to distinguish virtual subgroups of a corpus such as
	// generated files.
	Root string `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// A path-structured label describing the location of this object relative to
	// the corpus and the root.  For code, this will generally be the relative
	// path to the file containing the code, e.g., "storage/service.go" in kythe.
	// The individual elements of the path are separated by "/".
	// The path must not start with "/".
	// The path must be normalized to a canonical form (with no path
	// elements "", ".", or "..").
	//
	// However, this need not be a true file path; virtual objects like figments
	// can assign an ad-hoc abstract ID, or omit it entirely.
	//
	// Examples:
	//
	//	"devools/kythe/platform/go/datastore.go" (a file)
	//	"type/cpp/void.cc" (a type figment)
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// The language this name belongs to.
	// e.g., "c++", "python", "elisp", "haskell", "java"
	//
	// The schema will define specific labels for each supported language, so we
	// don't wind up with a confusion of names like "cxx", "cpp", "C++", etc.
	// Prototype: Official language name converted to lowercase.  If a version
	// number is necessary, include it, e.g., "python3".
	Language string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *VName) Reset() {
	*x = VName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VName) ProtoMessage() {}

func (x *VName) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VName.ProtoReflect.Descriptor instead.
func (*VName) Descriptor() ([]byte, []int) {
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *VName) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *VName) GetCorpus() string {
	if x != nil {
		return x.Corpus
	}
	return ""
}

func (x *VName) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *VName) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *VName) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_tensorflow_python_framework_kythe_metadata_proto protoreflect.FileDescriptor

var file_tensorflow_python_framework_kythe_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74,
	0x68, 0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x96,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4b,
	0x59, 0x54, 0x48, 0x45, 0x30, 0x10, 0x01, 0x22, 0x9e, 0x04, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x64, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x64, 0x67, 0x65, 0x12, 0x27, 0x0a,
	0x05, 0x76, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x05, 0x76, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x3c,
	0x0a, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x45, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x50, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x5f,
	0x41, 0x4e, 0x43, 0x48, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x3e, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x4d, 0x41, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x4d, 0x41, 0x5f, 0x57, 0x52, 0x49, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x4d, 0x41, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x56, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x72, 0x70, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0xb3, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42,
	0x12, 0x4b, 0x79, 0x74, 0x68, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x79, 0x74, 0x68,
	0x6f, 0x6e, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0xa2, 0x02, 0x03, 0x54,
	0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca,
	0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_python_framework_kythe_metadata_proto_rawDescOnce sync.Once
	file_tensorflow_python_framework_kythe_metadata_proto_rawDescData = file_tensorflow_python_framework_kythe_metadata_proto_rawDesc
)

func file_tensorflow_python_framework_kythe_metadata_proto_rawDescGZIP() []byte {
	file_tensorflow_python_framework_kythe_metadata_proto_rawDescOnce.Do(func() {
		file_tensorflow_python_framework_kythe_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_python_framework_kythe_metadata_proto_rawDescData)
	})
	return file_tensorflow_python_framework_kythe_metadata_proto_rawDescData
}

var file_tensorflow_python_framework_kythe_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tensorflow_python_framework_kythe_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_python_framework_kythe_metadata_proto_goTypes = []interface{}{
	(GeneratedCodeInfo_Type)(0), // 0: tensorflow.GeneratedCodeInfo.Type
	(MappingRule_Type)(0),       // 1: tensorflow.MappingRule.Type
	(MappingRule_Semantic)(0),   // 2: tensorflow.MappingRule.Semantic
	(*GeneratedCodeInfo)(nil),   // 3: tensorflow.GeneratedCodeInfo
	(*MappingRule)(nil),         // 4: tensorflow.MappingRule
	(*VName)(nil),               // 5: tensorflow.VName
}
var file_tensorflow_python_framework_kythe_metadata_proto_depIdxs = []int32{
	0, // 0: tensorflow.GeneratedCodeInfo.type:type_name -> tensorflow.GeneratedCodeInfo.Type
	4, // 1: tensorflow.GeneratedCodeInfo.meta:type_name -> tensorflow.MappingRule
	1, // 2: tensorflow.MappingRule.type:type_name -> tensorflow.MappingRule.Type
	5, // 3: tensorflow.MappingRule.vname:type_name -> tensorflow.VName
	2, // 4: tensorflow.MappingRule.semantic:type_name -> tensorflow.MappingRule.Semantic
	5, // 5: tensorflow.MappingRule.source_vname:type_name -> tensorflow.VName
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_python_framework_kythe_metadata_proto_init() }
func file_tensorflow_python_framework_kythe_metadata_proto_init() {
	if File_tensorflow_python_framework_kythe_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratedCodeInfo); i {
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
		file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MappingRule); i {
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
		file_tensorflow_python_framework_kythe_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VName); i {
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
			RawDescriptor: file_tensorflow_python_framework_kythe_metadata_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_python_framework_kythe_metadata_proto_goTypes,
		DependencyIndexes: file_tensorflow_python_framework_kythe_metadata_proto_depIdxs,
		EnumInfos:         file_tensorflow_python_framework_kythe_metadata_proto_enumTypes,
		MessageInfos:      file_tensorflow_python_framework_kythe_metadata_proto_msgTypes,
	}.Build()
	File_tensorflow_python_framework_kythe_metadata_proto = out.File
	file_tensorflow_python_framework_kythe_metadata_proto_rawDesc = nil
	file_tensorflow_python_framework_kythe_metadata_proto_goTypes = nil
	file_tensorflow_python_framework_kythe_metadata_proto_depIdxs = nil
}
