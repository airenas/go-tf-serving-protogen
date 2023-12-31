// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/pjrt/execute_options.proto

package pjrt

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

type ExecutionModeProto int32

const (
	ExecutionModeProto_EXECUTION_MODE_UNSPECIFIED  ExecutionModeProto = 0
	ExecutionModeProto_EXECUTION_MODE_DEFAULT      ExecutionModeProto = 1
	ExecutionModeProto_EXECUTION_MODE_SYNCHRONOUS  ExecutionModeProto = 2
	ExecutionModeProto_EXECUTION_MODE_ASYNCHRONOUS ExecutionModeProto = 3
)

// Enum value maps for ExecutionModeProto.
var (
	ExecutionModeProto_name = map[int32]string{
		0: "EXECUTION_MODE_UNSPECIFIED",
		1: "EXECUTION_MODE_DEFAULT",
		2: "EXECUTION_MODE_SYNCHRONOUS",
		3: "EXECUTION_MODE_ASYNCHRONOUS",
	}
	ExecutionModeProto_value = map[string]int32{
		"EXECUTION_MODE_UNSPECIFIED":  0,
		"EXECUTION_MODE_DEFAULT":      1,
		"EXECUTION_MODE_SYNCHRONOUS":  2,
		"EXECUTION_MODE_ASYNCHRONOUS": 3,
	}
)

func (x ExecutionModeProto) Enum() *ExecutionModeProto {
	p := new(ExecutionModeProto)
	*p = x
	return p
}

func (x ExecutionModeProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecutionModeProto) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_xla_pjrt_execute_options_proto_enumTypes[0].Descriptor()
}

func (ExecutionModeProto) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_xla_pjrt_execute_options_proto_enumTypes[0]
}

func (x ExecutionModeProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecutionModeProto.Descriptor instead.
func (ExecutionModeProto) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescGZIP(), []int{0}
}

// Mirrors `xla::ExecuteOptions`.
type ExecuteOptionsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArgumentsAreTupled       bool               `protobuf:"varint,1,opt,name=arguments_are_tupled,json=argumentsAreTupled,proto3" json:"arguments_are_tupled,omitempty"`
	UntupleResult            bool               `protobuf:"varint,2,opt,name=untuple_result,json=untupleResult,proto3" json:"untuple_result,omitempty"`
	LaunchId                 int32              `protobuf:"varint,3,opt,name=launch_id,json=launchId,proto3" json:"launch_id,omitempty"`
	StrictShapeChecking      bool               `protobuf:"varint,4,opt,name=strict_shape_checking,json=strictShapeChecking,proto3" json:"strict_shape_checking,omitempty"`
	ExecutionMode            ExecutionModeProto `protobuf:"varint,6,opt,name=execution_mode,json=executionMode,proto3,enum=xla.ExecutionModeProto" json:"execution_mode,omitempty"`
	NonDonatableInputIndices []int32            `protobuf:"varint,7,rep,packed,name=non_donatable_input_indices,json=nonDonatableInputIndices,proto3" json:"non_donatable_input_indices,omitempty"`
}

func (x *ExecuteOptionsProto) Reset() {
	*x = ExecuteOptionsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_pjrt_execute_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteOptionsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteOptionsProto) ProtoMessage() {}

func (x *ExecuteOptionsProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_pjrt_execute_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteOptionsProto.ProtoReflect.Descriptor instead.
func (*ExecuteOptionsProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteOptionsProto) GetArgumentsAreTupled() bool {
	if x != nil {
		return x.ArgumentsAreTupled
	}
	return false
}

func (x *ExecuteOptionsProto) GetUntupleResult() bool {
	if x != nil {
		return x.UntupleResult
	}
	return false
}

func (x *ExecuteOptionsProto) GetLaunchId() int32 {
	if x != nil {
		return x.LaunchId
	}
	return 0
}

func (x *ExecuteOptionsProto) GetStrictShapeChecking() bool {
	if x != nil {
		return x.StrictShapeChecking
	}
	return false
}

func (x *ExecuteOptionsProto) GetExecutionMode() ExecutionModeProto {
	if x != nil {
		return x.ExecutionMode
	}
	return ExecutionModeProto_EXECUTION_MODE_UNSPECIFIED
}

func (x *ExecuteOptionsProto) GetNonDonatableInputIndices() []int32 {
	if x != nil {
		return x.NonDonatableInputIndices
	}
	return nil
}

var File_tensorflow_compiler_xla_pjrt_execute_options_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x6a, 0x72, 0x74, 0x2f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x78, 0x6c, 0x61, 0x22, 0xbe, 0x02, 0x0a, 0x13, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61,
	0x72, 0x65, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x72, 0x65, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x75, 0x6e, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x6e, 0x74,
	0x75, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x75, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x0e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6e,
	0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x6e, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x18, 0x6e, 0x6f, 0x6e, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x2a, 0x91, 0x01, 0x0a, 0x12, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x1e, 0x0a,
	0x1a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x02, 0x12, 0x1f, 0x0a,
	0x1b, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x41, 0x53, 0x59, 0x4e, 0x43, 0x48, 0x52, 0x4f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x03, 0x42, 0x92,
	0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x42, 0x13, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69,
	0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72,
	0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x6a, 0x72, 0x74, 0xa2, 0x02, 0x03, 0x58, 0x58, 0x58, 0xaa,
	0x02, 0x03, 0x58, 0x6c, 0x61, 0xca, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xe2, 0x02, 0x0f, 0x58, 0x6c,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03,
	0x58, 0x6c, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescData = file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDesc
)

func file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDescData
}

var file_tensorflow_compiler_xla_pjrt_execute_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_xla_pjrt_execute_options_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_compiler_xla_pjrt_execute_options_proto_goTypes = []interface{}{
	(ExecutionModeProto)(0),     // 0: xla.ExecutionModeProto
	(*ExecuteOptionsProto)(nil), // 1: xla.ExecuteOptionsProto
}
var file_tensorflow_compiler_xla_pjrt_execute_options_proto_depIdxs = []int32{
	0, // 0: xla.ExecuteOptionsProto.execution_mode:type_name -> xla.ExecutionModeProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_pjrt_execute_options_proto_init() }
func file_tensorflow_compiler_xla_pjrt_execute_options_proto_init() {
	if File_tensorflow_compiler_xla_pjrt_execute_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_pjrt_execute_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteOptionsProto); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_pjrt_execute_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_pjrt_execute_options_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_xla_pjrt_execute_options_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_xla_pjrt_execute_options_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_pjrt_execute_options_proto = out.File
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_rawDesc = nil
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_goTypes = nil
	file_tensorflow_compiler_xla_pjrt_execute_options_proto_depIdxs = nil
}
