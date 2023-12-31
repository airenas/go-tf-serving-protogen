// Copyright 2023 The TensorFlow Authors. All Rights Reserved.
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
// source: tensorflow/compiler/xla/python/pjrt_ifrt/xla_sharding.proto

package pjrt_ifrt

import (
	data "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/data"
	ifrt "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/python/ifrt"
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

// Wire format for `HloSharding`.
type HloShardingProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices       *ifrt.DeviceListProto `protobuf:"bytes,1,opt,name=devices,proto3" json:"devices,omitempty"`
	MemoryKind    *string               `protobuf:"bytes,3,opt,name=memory_kind,json=memoryKind,proto3,oneof" json:"memory_kind,omitempty"`
	XlaOpSharding *data.OpSharding      `protobuf:"bytes,2,opt,name=xla_op_sharding,json=xlaOpSharding,proto3" json:"xla_op_sharding,omitempty"`
}

func (x *HloShardingProto) Reset() {
	*x = HloShardingProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HloShardingProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HloShardingProto) ProtoMessage() {}

func (x *HloShardingProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HloShardingProto.ProtoReflect.Descriptor instead.
func (*HloShardingProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescGZIP(), []int{0}
}

func (x *HloShardingProto) GetDevices() *ifrt.DeviceListProto {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *HloShardingProto) GetMemoryKind() string {
	if x != nil && x.MemoryKind != nil {
		return *x.MemoryKind
	}
	return ""
}

func (x *HloShardingProto) GetXlaOpSharding() *data.OpSharding {
	if x != nil {
		return x.XlaOpSharding
	}
	return nil
}

var File_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2f, 0x70, 0x6a, 0x72, 0x74, 0x5f, 0x69, 0x66, 0x72, 0x74, 0x2f, 0x78, 0x6c, 0x61, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x78,
	0x6c, 0x61, 0x2e, 0x69, 0x66, 0x72, 0x74, 0x1a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61,
	0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x69, 0x66, 0x72, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c,
	0x61, 0x2f, 0x78, 0x6c, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x01, 0x0a, 0x10, 0x48, 0x6c, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x69, 0x66, 0x72,
	0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x37, 0x0a, 0x0f, 0x78, 0x6c, 0x61, 0x5f, 0x6f, 0x70, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x78, 0x6c, 0x61, 0x2e,
	0x4f, 0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x78, 0x6c, 0x61, 0x4f,
	0x70, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0xb5, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x69, 0x66, 0x72, 0x74, 0x42, 0x10, 0x58, 0x6c, 0x61, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e,
	0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c,
	0x61, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x70, 0x6a, 0x72, 0x74, 0x5f, 0x69, 0x66,
	0x72, 0x74, 0xa2, 0x02, 0x03, 0x58, 0x49, 0x58, 0xaa, 0x02, 0x08, 0x58, 0x6c, 0x61, 0x2e, 0x49,
	0x66, 0x72, 0x74, 0xca, 0x02, 0x08, 0x58, 0x6c, 0x61, 0x5c, 0x49, 0x66, 0x72, 0x74, 0xe2, 0x02,
	0x14, 0x58, 0x6c, 0x61, 0x5c, 0x49, 0x66, 0x72, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x58, 0x6c, 0x61, 0x3a, 0x3a, 0x49, 0x66, 0x72,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescData = file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDesc
)

func file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDescData
}

var file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_goTypes = []interface{}{
	(*HloShardingProto)(nil),     // 0: xla.ifrt.HloShardingProto
	(*ifrt.DeviceListProto)(nil), // 1: xla.ifrt.DeviceListProto
	(*data.OpSharding)(nil),      // 2: xla.OpSharding
}
var file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_depIdxs = []int32{
	1, // 0: xla.ifrt.HloShardingProto.devices:type_name -> xla.ifrt.DeviceListProto
	2, // 1: xla.ifrt.HloShardingProto.xla_op_sharding:type_name -> xla.OpSharding
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_init() }
func file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_init() {
	if File_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HloShardingProto); i {
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
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto = out.File
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_rawDesc = nil
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_goTypes = nil
	file_tensorflow_compiler_xla_python_pjrt_ifrt_xla_sharding_proto_depIdxs = nil
}
