// Protocol messages for describing the results of benchmarks and unit tests.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/util/test_log.proto

package util

import (
	protobuf "github.com/airenas/go-tf-serving-protogen/tensorflow/tsl/protobuf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/tsl/protobuf/test_log.proto.

type TestResults_BenchmarkType = protobuf.TestResults_BenchmarkType

const TestResults_UNKNOWN = protobuf.TestResults_UNKNOWN
const TestResults_CPP_MICROBENCHMARK = protobuf.TestResults_CPP_MICROBENCHMARK
const TestResults_PYTHON_BENCHMARK = protobuf.TestResults_PYTHON_BENCHMARK
const TestResults_ANDROID_BENCHMARK = protobuf.TestResults_ANDROID_BENCHMARK
const TestResults_EDGE_BENCHMARK = protobuf.TestResults_EDGE_BENCHMARK
const TestResults_IOS_BENCHMARK = protobuf.TestResults_IOS_BENCHMARK

var TestResults_BenchmarkType_name = protobuf.TestResults_BenchmarkType_name
var TestResults_BenchmarkType_value = protobuf.TestResults_BenchmarkType_value

type EntryValue = protobuf.EntryValue
type EntryValue_DoubleValue = protobuf.EntryValue_DoubleValue
type EntryValue_StringValue = protobuf.EntryValue_StringValue
type MetricEntry = protobuf.MetricEntry
type BenchmarkEntry = protobuf.BenchmarkEntry
type BenchmarkEntries = protobuf.BenchmarkEntries
type BuildConfiguration = protobuf.BuildConfiguration
type CommitId = protobuf.CommitId
type CommitId_Changelist = protobuf.CommitId_Changelist
type CommitId_Hash = protobuf.CommitId_Hash
type CPUInfo = protobuf.CPUInfo
type MemoryInfo = protobuf.MemoryInfo
type GPUInfo = protobuf.GPUInfo
type PlatformInfo = protobuf.PlatformInfo
type AvailableDeviceInfo = protobuf.AvailableDeviceInfo
type MachineConfiguration = protobuf.MachineConfiguration
type RunConfiguration = protobuf.RunConfiguration
type TestResults = protobuf.TestResults

var File_tensorflow_core_util_test_log_proto protoreflect.FileDescriptor

var file_tensorflow_core_util_test_log_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0xc5, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x42, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d,
	0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0xa2, 0x02, 0x03, 0x54, 0x44, 0x58, 0xaa, 0x02,
	0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0xca, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0xe2, 0x02, 0x1c, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x3a, 0x3a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_tensorflow_core_util_test_log_proto_goTypes = []interface{}{}
var file_tensorflow_core_util_test_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_util_test_log_proto_init() }
func file_tensorflow_core_util_test_log_proto_init() {
	if File_tensorflow_core_util_test_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_util_test_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_util_test_log_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_util_test_log_proto_depIdxs,
	}.Build()
	File_tensorflow_core_util_test_log_proto = out.File
	file_tensorflow_core_util_test_log_proto_rawDesc = nil
	file_tensorflow_core_util_test_log_proto_goTypes = nil
	file_tensorflow_core_util_test_log_proto_depIdxs = nil
}
