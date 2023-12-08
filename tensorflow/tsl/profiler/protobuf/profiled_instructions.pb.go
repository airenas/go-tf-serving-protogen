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
// source: tensorflow/tsl/profiler/protobuf/profiled_instructions.proto

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

// Next ID: 3
type ProfiledInstructionsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Costs     []*ProfiledInstructionsProto_InstructionCost `protobuf:"bytes,1,rep,name=costs,proto3" json:"costs,omitempty"`
	Latencies []*ProfiledInstructionsProto_Latency         `protobuf:"bytes,2,rep,name=latencies,proto3" json:"latencies,omitempty"`
}

func (x *ProfiledInstructionsProto) Reset() {
	*x = ProfiledInstructionsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfiledInstructionsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfiledInstructionsProto) ProtoMessage() {}

func (x *ProfiledInstructionsProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfiledInstructionsProto.ProtoReflect.Descriptor instead.
func (*ProfiledInstructionsProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescGZIP(), []int{0}
}

func (x *ProfiledInstructionsProto) GetCosts() []*ProfiledInstructionsProto_InstructionCost {
	if x != nil {
		return x.Costs
	}
	return nil
}

func (x *ProfiledInstructionsProto) GetLatencies() []*ProfiledInstructionsProto_Latency {
	if x != nil {
		return x.Latencies
	}
	return nil
}

type ProfiledInstructionsProto_InstructionCost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CostUs float64 `protobuf:"fixed64,2,opt,name=cost_us,json=costUs,proto3" json:"cost_us,omitempty"`
}

func (x *ProfiledInstructionsProto_InstructionCost) Reset() {
	*x = ProfiledInstructionsProto_InstructionCost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfiledInstructionsProto_InstructionCost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfiledInstructionsProto_InstructionCost) ProtoMessage() {}

func (x *ProfiledInstructionsProto_InstructionCost) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfiledInstructionsProto_InstructionCost.ProtoReflect.Descriptor instead.
func (*ProfiledInstructionsProto_InstructionCost) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProfiledInstructionsProto_InstructionCost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfiledInstructionsProto_InstructionCost) GetCostUs() float64 {
	if x != nil {
		return x.CostUs
	}
	return 0
}

type ProfiledInstructionsProto_Latency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source    string  `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target    string  `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	LatencyUs float64 `protobuf:"fixed64,3,opt,name=latency_us,json=latencyUs,proto3" json:"latency_us,omitempty"`
}

func (x *ProfiledInstructionsProto_Latency) Reset() {
	*x = ProfiledInstructionsProto_Latency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfiledInstructionsProto_Latency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfiledInstructionsProto_Latency) ProtoMessage() {}

func (x *ProfiledInstructionsProto_Latency) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfiledInstructionsProto_Latency.ProtoReflect.Descriptor instead.
func (*ProfiledInstructionsProto_Latency) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ProfiledInstructionsProto_Latency) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ProfiledInstructionsProto_Latency) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ProfiledInstructionsProto_Latency) GetLatencyUs() float64 {
	if x != nil {
		return x.LatencyUs
	}
	return 0
}

var File_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x22, 0xe1, 0x02, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x64,
	0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x54, 0x0a, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x63, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x09, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x3e, 0x0a,
	0x0f, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x73, 0x1a, 0x58, 0x0a,
	0x07, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x74, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x61,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x73, 0x42, 0xed, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x42, 0x19, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72,
	0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54,
	0x50, 0x58, 0xaa, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xca, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xe2, 0x02,
	0x1f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_goTypes = []interface{}{
	(*ProfiledInstructionsProto)(nil),                 // 0: tensorflow.profiler.ProfiledInstructionsProto
	(*ProfiledInstructionsProto_InstructionCost)(nil), // 1: tensorflow.profiler.ProfiledInstructionsProto.InstructionCost
	(*ProfiledInstructionsProto_Latency)(nil),         // 2: tensorflow.profiler.ProfiledInstructionsProto.Latency
}
var file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_depIdxs = []int32{
	1, // 0: tensorflow.profiler.ProfiledInstructionsProto.costs:type_name -> tensorflow.profiler.ProfiledInstructionsProto.InstructionCost
	2, // 1: tensorflow.profiler.ProfiledInstructionsProto.latencies:type_name -> tensorflow.profiler.ProfiledInstructionsProto.Latency
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfiledInstructionsProto); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfiledInstructionsProto_InstructionCost); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfiledInstructionsProto_Latency); i {
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
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_profiled_instructions_proto_depIdxs = nil
}