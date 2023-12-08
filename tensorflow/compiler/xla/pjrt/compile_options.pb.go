// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/pjrt/compile_options.proto

package pjrt

import (
	xla "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla"
	data "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/data"
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

// A serialization of xla::ExecutableBuildOptions.
// Next id: 14.
type ExecutableBuildOptionsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, this is the device to build the computation for. Valid
	// device_ordinal values are: 0 to # of devices - 1. These values are
	// identical to the device ordinal values used by StreamExecutor. The built
	// executable will be executable on any device equivalent to the specified
	// device as determined by Backend::devices_equivalent(). A value of -1
	// indicates this option has not been set.
	DeviceOrdinal int64 `protobuf:"varint,1,opt,name=device_ordinal,json=deviceOrdinal,proto3" json:"device_ordinal,omitempty"`
	// If set, this specifies the layout of the result of the computation. If not
	// set, the service will chose the layout of the result. A Shape is used to
	// store the layout to accommodate tuple result shapes. A value of nullptr
	// indicates the option has not been set.
	ResultLayout *data.ShapeProto `protobuf:"bytes,2,opt,name=result_layout,json=resultLayout,proto3" json:"result_layout,omitempty"`
	// Expose access to the XLA compilation environments, which will be passed to
	// the compilation process.
	CompEnvs *xla.CompilationEnvironmentsProto `protobuf:"bytes,13,opt,name=comp_envs,json=compEnvs,proto3" json:"comp_envs,omitempty"`
	// Expose access to the XLA debug options which will be passed to the
	// compilation process.
	DebugOptions *xla.DebugOptions `protobuf:"bytes,3,opt,name=debug_options,json=debugOptions,proto3" json:"debug_options,omitempty"`
	// The number of replicas of this computation that are to be executed.
	// Defaults to 1.
	NumReplicas int64 `protobuf:"varint,4,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
	// The number of partitions in this computation. Defaults to 1.
	NumPartitions int64 `protobuf:"varint,5,opt,name=num_partitions,json=numPartitions,proto3" json:"num_partitions,omitempty"`
	// Indicates whether to use SPMD (true) or MPMD (false) partitioning when
	// num_partitions > 1 and XLA is requested to partition the input program.
	UseSpmdPartitioning bool `protobuf:"varint,6,opt,name=use_spmd_partitioning,json=useSpmdPartitioning,proto3" json:"use_spmd_partitioning,omitempty"`
	// Whether to automatically generate XLA shardings for SPMD partitioner.
	UseAutoSpmdPartitioning bool `protobuf:"varint,7,opt,name=use_auto_spmd_partitioning,json=useAutoSpmdPartitioning,proto3" json:"use_auto_spmd_partitioning,omitempty"`
	// Whether HLOs should be deduplicated.
	DeduplicateHlo bool `protobuf:"varint,8,opt,name=deduplicate_hlo,json=deduplicateHlo,proto3" json:"deduplicate_hlo,omitempty"`
	// If set, this specifies a static device assignment for the computation.
	// Otherwise, the computation will be compiled generically and can be run with
	// any device assignment compatible with the computation's replica and
	// partition counts.
	DeviceAssignment *data.DeviceAssignmentProto `protobuf:"bytes,9,opt,name=device_assignment,json=deviceAssignment,proto3" json:"device_assignment,omitempty"`
	// Whether input and output buffers are aliased if the associated parameter is
	// passed-through XLA modules without being changed.
	AliasPassthroughParams bool `protobuf:"varint,10,opt,name=alias_passthrough_params,json=aliasPassthroughParams,proto3" json:"alias_passthrough_params,omitempty"`
	// By default, XLA builds an executable by invoking standard compilation, i.e.
	// running Compiler::Compile, or both Compiler::RunHloPasses and
	// Compiler::RunBackend. When run_backend_only is set to true, XLA builds an
	// executable by invoking only RunBackend and skip invoking RunHloPasses,
	// which can be used to compile post-optimizations HLO modules.
	RunBackendOnly bool `protobuf:"varint,11,opt,name=run_backend_only,json=runBackendOnly,proto3" json:"run_backend_only,omitempty"`
	// Allows sharding propagation to propagate to the outputs. This changes the
	// output shape of the computation (which is undesirable), but it can be used
	// to allow to run partial compilation to determine what would be the output
	// sharding of a computation if XLA would be allowed to propagate the sharding
	// which can be used by higher level framework as a way to query intermediate
	// sharding of operations when multiple computation would be chained and
	// merged together.
	AllowSpmdShardingPropagationToOutput bool `protobuf:"varint,12,opt,name=allow_spmd_sharding_propagation_to_output,json=allowSpmdShardingPropagationToOutput,proto3" json:"allow_spmd_sharding_propagation_to_output,omitempty"`
}

func (x *ExecutableBuildOptionsProto) Reset() {
	*x = ExecutableBuildOptionsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutableBuildOptionsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutableBuildOptionsProto) ProtoMessage() {}

func (x *ExecutableBuildOptionsProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutableBuildOptionsProto.ProtoReflect.Descriptor instead.
func (*ExecutableBuildOptionsProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutableBuildOptionsProto) GetDeviceOrdinal() int64 {
	if x != nil {
		return x.DeviceOrdinal
	}
	return 0
}

func (x *ExecutableBuildOptionsProto) GetResultLayout() *data.ShapeProto {
	if x != nil {
		return x.ResultLayout
	}
	return nil
}

func (x *ExecutableBuildOptionsProto) GetCompEnvs() *xla.CompilationEnvironmentsProto {
	if x != nil {
		return x.CompEnvs
	}
	return nil
}

func (x *ExecutableBuildOptionsProto) GetDebugOptions() *xla.DebugOptions {
	if x != nil {
		return x.DebugOptions
	}
	return nil
}

func (x *ExecutableBuildOptionsProto) GetNumReplicas() int64 {
	if x != nil {
		return x.NumReplicas
	}
	return 0
}

func (x *ExecutableBuildOptionsProto) GetNumPartitions() int64 {
	if x != nil {
		return x.NumPartitions
	}
	return 0
}

func (x *ExecutableBuildOptionsProto) GetUseSpmdPartitioning() bool {
	if x != nil {
		return x.UseSpmdPartitioning
	}
	return false
}

func (x *ExecutableBuildOptionsProto) GetUseAutoSpmdPartitioning() bool {
	if x != nil {
		return x.UseAutoSpmdPartitioning
	}
	return false
}

func (x *ExecutableBuildOptionsProto) GetDeduplicateHlo() bool {
	if x != nil {
		return x.DeduplicateHlo
	}
	return false
}

func (x *ExecutableBuildOptionsProto) GetDeviceAssignment() *data.DeviceAssignmentProto {
	if x != nil {
		return x.DeviceAssignment
	}
	return nil
}

func (x *ExecutableBuildOptionsProto) GetAliasPassthroughParams() bool {
	if x != nil {
		return x.AliasPassthroughParams
	}
	return false
}

func (x *ExecutableBuildOptionsProto) GetRunBackendOnly() bool {
	if x != nil {
		return x.RunBackendOnly
	}
	return false
}

func (x *ExecutableBuildOptionsProto) GetAllowSpmdShardingPropagationToOutput() bool {
	if x != nil {
		return x.AllowSpmdShardingPropagationToOutput
	}
	return false
}

type CompileOptionsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Refer CompileOptions for documentation of fields.
	ArgumentLayouts            []*data.ShapeProto           `protobuf:"bytes,1,rep,name=argument_layouts,json=argumentLayouts,proto3" json:"argument_layouts,omitempty"`
	ParameterIsTupledArguments bool                         `protobuf:"varint,2,opt,name=parameter_is_tupled_arguments,json=parameterIsTupledArguments,proto3" json:"parameter_is_tupled_arguments,omitempty"`
	ExecutableBuildOptions     *ExecutableBuildOptionsProto `protobuf:"bytes,3,opt,name=executable_build_options,json=executableBuildOptions,proto3" json:"executable_build_options,omitempty"`
	CompilePortableExecutable  bool                         `protobuf:"varint,4,opt,name=compile_portable_executable,json=compilePortableExecutable,proto3" json:"compile_portable_executable,omitempty"`
	ProfileVersion             int64                        `protobuf:"varint,5,opt,name=profile_version,json=profileVersion,proto3" json:"profile_version,omitempty"`
	SerializedMultiSliceConfig []byte                       `protobuf:"bytes,6,opt,name=serialized_multi_slice_config,json=serializedMultiSliceConfig,proto3" json:"serialized_multi_slice_config,omitempty"`
}

func (x *CompileOptionsProto) Reset() {
	*x = CompileOptionsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompileOptionsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompileOptionsProto) ProtoMessage() {}

func (x *CompileOptionsProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompileOptionsProto.ProtoReflect.Descriptor instead.
func (*CompileOptionsProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescGZIP(), []int{1}
}

func (x *CompileOptionsProto) GetArgumentLayouts() []*data.ShapeProto {
	if x != nil {
		return x.ArgumentLayouts
	}
	return nil
}

func (x *CompileOptionsProto) GetParameterIsTupledArguments() bool {
	if x != nil {
		return x.ParameterIsTupledArguments
	}
	return false
}

func (x *CompileOptionsProto) GetExecutableBuildOptions() *ExecutableBuildOptionsProto {
	if x != nil {
		return x.ExecutableBuildOptions
	}
	return nil
}

func (x *CompileOptionsProto) GetCompilePortableExecutable() bool {
	if x != nil {
		return x.CompilePortableExecutable
	}
	return false
}

func (x *CompileOptionsProto) GetProfileVersion() int64 {
	if x != nil {
		return x.ProfileVersion
	}
	return 0
}

func (x *CompileOptionsProto) GetSerializedMultiSliceConfig() []byte {
	if x != nil {
		return x.SerializedMultiSliceConfig
	}
	return nil
}

var File_tensorflow_compiler_xla_pjrt_compile_options_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x6a, 0x72, 0x74, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x78, 0x6c, 0x61, 0x1a, 0x21, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78,
	0x6c, 0x61, 0x2f, 0x78, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x78, 0x6c, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x05, 0x0a, 0x1b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x5f, 0x65, 0x6e, 0x76, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x45, 0x6e, 0x76,
	0x73, 0x12, 0x36, 0x0a, 0x0d, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x70, 0x6d, 0x64, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x75, 0x73, 0x65, 0x53, 0x70, 0x6d, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6d, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x75, 0x73, 0x65,
	0x41, 0x75, 0x74, 0x6f, 0x53, 0x70, 0x6d, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x68, 0x6c, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64,
	0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x6c, 0x6f, 0x12, 0x47, 0x0a,
	0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x50,
	0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x28, 0x0a, 0x10, 0x72, 0x75, 0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x29, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x70, 0x6d, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x24, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x70, 0x6d, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x9c, 0x03, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x0a, 0x10, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x64, 0x5f, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x49, 0x73, 0x54, 0x75, 0x70, 0x6c, 0x65,
	0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x5a, 0x0a, 0x18, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78,
	0x6c, 0x61, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x16,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x0a, 0x1d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x1a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x92, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x42, 0x13,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x6a, 0x72, 0x74, 0xa2, 0x02, 0x03,
	0x58, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xca, 0x02, 0x03, 0x58, 0x6c, 0x61, 0xe2,
	0x02, 0x0f, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x03, 0x58, 0x6c, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescData = file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDesc
)

func file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDescData
}

var file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_compiler_xla_pjrt_compile_options_proto_goTypes = []interface{}{
	(*ExecutableBuildOptionsProto)(nil),      // 0: xla.ExecutableBuildOptionsProto
	(*CompileOptionsProto)(nil),              // 1: xla.CompileOptionsProto
	(*data.ShapeProto)(nil),                  // 2: xla.ShapeProto
	(*xla.CompilationEnvironmentsProto)(nil), // 3: xla.CompilationEnvironmentsProto
	(*xla.DebugOptions)(nil),                 // 4: xla.DebugOptions
	(*data.DeviceAssignmentProto)(nil),       // 5: xla.DeviceAssignmentProto
}
var file_tensorflow_compiler_xla_pjrt_compile_options_proto_depIdxs = []int32{
	2, // 0: xla.ExecutableBuildOptionsProto.result_layout:type_name -> xla.ShapeProto
	3, // 1: xla.ExecutableBuildOptionsProto.comp_envs:type_name -> xla.CompilationEnvironmentsProto
	4, // 2: xla.ExecutableBuildOptionsProto.debug_options:type_name -> xla.DebugOptions
	5, // 3: xla.ExecutableBuildOptionsProto.device_assignment:type_name -> xla.DeviceAssignmentProto
	2, // 4: xla.CompileOptionsProto.argument_layouts:type_name -> xla.ShapeProto
	0, // 5: xla.CompileOptionsProto.executable_build_options:type_name -> xla.ExecutableBuildOptionsProto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_pjrt_compile_options_proto_init() }
func file_tensorflow_compiler_xla_pjrt_compile_options_proto_init() {
	if File_tensorflow_compiler_xla_pjrt_compile_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutableBuildOptionsProto); i {
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
		file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompileOptionsProto); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_pjrt_compile_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_pjrt_compile_options_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_pjrt_compile_options_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_pjrt_compile_options_proto = out.File
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_rawDesc = nil
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_goTypes = nil
	file_tensorflow_compiler_xla_pjrt_compile_options_proto_depIdxs = nil
}
