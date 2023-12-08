// Copyright 2022 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/autotune_results.proto

package xla

import (
	protobuf "github.com/airenas/go-tf-serving-protogen/tensorflow/tsl/protobuf"
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

// A collection of algorithms for particular dot/convs.  Usually this is "the
// best" algorithm for the particular dot/conv, although that's not strictly
// required.
//
// Users don't interact with this proto directly.  It's used internally to
// facilitate ahead-of-time autotuning -- The string used by
// xla::{Serialize,Load}AutotuneResults is, internally, a serialization of this
// proto.
//
// LINT.IfChange
type AutotuneResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32                    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Dots    []*AutotuneResults_Entry `protobuf:"bytes,2,rep,name=dots,proto3" json:"dots,omitempty"`
	Convs   []*AutotuneResults_Entry `protobuf:"bytes,3,rep,name=convs,proto3" json:"convs,omitempty"`
}

func (x *AutotuneResults) Reset() {
	*x = AutotuneResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResults) ProtoMessage() {}

func (x *AutotuneResults) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResults.ProtoReflect.Descriptor instead.
func (*AutotuneResults) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_autotune_results_proto_rawDescGZIP(), []int{0}
}

func (x *AutotuneResults) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AutotuneResults) GetDots() []*AutotuneResults_Entry {
	if x != nil {
		return x.Dots
	}
	return nil
}

func (x *AutotuneResults) GetConvs() []*AutotuneResults_Entry {
	if x != nil {
		return x.Convs
	}
	return nil
}

type AutotuneResults_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Hlo    string `protobuf:"bytes,2,opt,name=hlo,proto3" json:"hlo,omitempty"`
	// nb: These results are always tied to a particular version of
	// cublas/cudnn, but this is *especially* true for cublasLt results.  For
	// cublasLt gemms, the result is an index into the list of candidate
	// algorithms returned by cublasLt.  Different version of cublasLt ->
	// different list of algos -> different interpretation of results!
	Result *protobuf.AutotuneResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AutotuneResults_Entry) Reset() {
	*x = AutotuneResults_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutotuneResults_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutotuneResults_Entry) ProtoMessage() {}

func (x *AutotuneResults_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutotuneResults_Entry.ProtoReflect.Descriptor instead.
func (*AutotuneResults_Entry) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_autotune_results_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AutotuneResults_Entry) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *AutotuneResults_Entry) GetHlo() string {
	if x != nil {
		return x.Hlo
	}
	return ""
}

func (x *AutotuneResults_Entry) GetResult() *protobuf.AutotuneResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_tensorflow_compiler_xla_autotune_results_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_autotune_results_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75,
	0x6e, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x78, 0x6c, 0x61, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf4, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x6c,
	0x61, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x6f, 0x74, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x63, 0x6f, 0x6e, 0x76, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78,
	0x6c, 0x61, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x76, 0x73, 0x1a,
	0x65, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x68, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68,
	0x6c, 0x6f, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x8e, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x78,
	0x6c, 0x61, 0x42, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67,
	0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0xa2, 0x02, 0x03,
	0x58, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xca, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xe2,
	0x02, 0x0f, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x03, 0x58, 0x6c, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_autotune_results_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_autotune_results_proto_rawDescData = file_tensorflow_compiler_xla_autotune_results_proto_rawDesc
)

func file_tensorflow_compiler_xla_autotune_results_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_autotune_results_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_autotune_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_autotune_results_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_autotune_results_proto_rawDescData
}

var file_tensorflow_compiler_xla_autotune_results_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_compiler_xla_autotune_results_proto_goTypes = []interface{}{
	(*AutotuneResults)(nil),         // 0: xla.AutotuneResults
	(*AutotuneResults_Entry)(nil),   // 1: xla.AutotuneResults.Entry
	(*protobuf.AutotuneResult)(nil), // 2: tensorflow.AutotuneResult
}
var file_tensorflow_compiler_xla_autotune_results_proto_depIdxs = []int32{
	1, // 0: xla.AutotuneResults.dots:type_name -> xla.AutotuneResults.Entry
	1, // 1: xla.AutotuneResults.convs:type_name -> xla.AutotuneResults.Entry
	2, // 2: xla.AutotuneResults.Entry.result:type_name -> tensorflow.AutotuneResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_autotune_results_proto_init() }
func file_tensorflow_compiler_xla_autotune_results_proto_init() {
	if File_tensorflow_compiler_xla_autotune_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResults); i {
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
		file_tensorflow_compiler_xla_autotune_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutotuneResults_Entry); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_autotune_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_autotune_results_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_autotune_results_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_autotune_results_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_autotune_results_proto = out.File
	file_tensorflow_compiler_xla_autotune_results_proto_rawDesc = nil
	file_tensorflow_compiler_xla_autotune_results_proto_goTypes = nil
	file_tensorflow_compiler_xla_autotune_results_proto_depIdxs = nil
}
