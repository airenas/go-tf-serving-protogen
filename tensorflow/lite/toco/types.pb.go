// Copyright 2017 The TensorFlow Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/lite/toco/types.proto

package toco

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

// IODataType describes the numeric data types of input and output arrays
// of a model.
type IODataType int32

const (
	IODataType_IO_DATA_TYPE_UNKNOWN IODataType = 0
	// Float32, not quantized
	IODataType_FLOAT IODataType = 1
	// Uint8, quantized
	IODataType_QUANTIZED_UINT8 IODataType = 2
	// Int32, not quantized
	IODataType_INT32 IODataType = 3
	// Int64, not quantized
	IODataType_INT64 IODataType = 4
	// String, not quantized
	IODataType_STRING IODataType = 5
	// Int16, quantized
	IODataType_QUANTIZED_INT16 IODataType = 6
	// Boolean
	IODataType_BOOL IODataType = 7
	// Complex64, not quantized
	IODataType_COMPLEX64 IODataType = 8
	// Int8, quantized based on QuantizationParameters in schema.
	IODataType_QUANTIZED_INT8 IODataType = 9
	// Half precision float, not quantized.
	IODataType_FLOAT16 IODataType = 10
	// Double precision float, not quantized.
	IODataType_FLOAT64 IODataType = 11
	// Complex128, not quantized
	IODataType_COMPLEX128 IODataType = 12
	// Uint64, not quantized
	IODataType_UINT64 IODataType = 13
	// Resource type
	IODataType_RESOURCE IODataType = 14
	// Variant type
	IODataType_VARIANT IODataType = 15
	// Uint32
	IODataType_UINT32 IODataType = 16
	// Uint8, not quantized
	IODataType_UINT8 IODataType = 17
	// Int8, not quantized
	IODataType_INT8 IODataType = 18
	// Int16, not quantized
	IODataType_INT16 IODataType = 19
	// Uint16, not quantized
	IODataType_UINT16 IODataType = 20
)

// Enum value maps for IODataType.
var (
	IODataType_name = map[int32]string{
		0:  "IO_DATA_TYPE_UNKNOWN",
		1:  "FLOAT",
		2:  "QUANTIZED_UINT8",
		3:  "INT32",
		4:  "INT64",
		5:  "STRING",
		6:  "QUANTIZED_INT16",
		7:  "BOOL",
		8:  "COMPLEX64",
		9:  "QUANTIZED_INT8",
		10: "FLOAT16",
		11: "FLOAT64",
		12: "COMPLEX128",
		13: "UINT64",
		14: "RESOURCE",
		15: "VARIANT",
		16: "UINT32",
		17: "UINT8",
		18: "INT8",
		19: "INT16",
		20: "UINT16",
	}
	IODataType_value = map[string]int32{
		"IO_DATA_TYPE_UNKNOWN": 0,
		"FLOAT":                1,
		"QUANTIZED_UINT8":      2,
		"INT32":                3,
		"INT64":                4,
		"STRING":               5,
		"QUANTIZED_INT16":      6,
		"BOOL":                 7,
		"COMPLEX64":            8,
		"QUANTIZED_INT8":       9,
		"FLOAT16":              10,
		"FLOAT64":              11,
		"COMPLEX128":           12,
		"UINT64":               13,
		"RESOURCE":             14,
		"VARIANT":              15,
		"UINT32":               16,
		"UINT8":                17,
		"INT8":                 18,
		"INT16":                19,
		"UINT16":               20,
	}
)

func (x IODataType) Enum() *IODataType {
	p := new(IODataType)
	*p = x
	return p
}

func (x IODataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IODataType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_lite_toco_types_proto_enumTypes[0].Descriptor()
}

func (IODataType) Type() protoreflect.EnumType {
	return &file_tensorflow_lite_toco_types_proto_enumTypes[0]
}

func (x IODataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *IODataType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = IODataType(num)
	return nil
}

// Deprecated: Use IODataType.Descriptor instead.
func (IODataType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_lite_toco_types_proto_rawDescGZIP(), []int{0}
}

var File_tensorflow_lite_toco_types_proto protoreflect.FileDescriptor

var file_tensorflow_lite_toco_types_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x74, 0x6f, 0x63, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x63, 0x6f, 0x2a, 0xb3, 0x02, 0x0a, 0x0a, 0x49, 0x4f, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4f, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x5a, 0x45, 0x44,
	0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c,
	0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x36, 0x34, 0x10,
	0x08, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55, 0x41, 0x4e, 0x54, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x54, 0x38, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x31, 0x36,
	0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x0b, 0x12,
	0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32, 0x38, 0x10, 0x0c, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x41, 0x52,
	0x49, 0x41, 0x4e, 0x54, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32,
	0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x11, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x31, 0x36,
	0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x14, 0x42, 0x86,
	0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x63, 0x6f, 0x42, 0x0a, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f,
	0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x6c, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x63, 0x6f, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa,
	0x02, 0x04, 0x54, 0x6f, 0x63, 0x6f, 0xca, 0x02, 0x04, 0x54, 0x6f, 0x63, 0x6f, 0xe2, 0x02, 0x10,
	0x54, 0x6f, 0x63, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x04, 0x54, 0x6f, 0x63, 0x6f,
}

var (
	file_tensorflow_lite_toco_types_proto_rawDescOnce sync.Once
	file_tensorflow_lite_toco_types_proto_rawDescData = file_tensorflow_lite_toco_types_proto_rawDesc
)

func file_tensorflow_lite_toco_types_proto_rawDescGZIP() []byte {
	file_tensorflow_lite_toco_types_proto_rawDescOnce.Do(func() {
		file_tensorflow_lite_toco_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_lite_toco_types_proto_rawDescData)
	})
	return file_tensorflow_lite_toco_types_proto_rawDescData
}

var file_tensorflow_lite_toco_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_lite_toco_types_proto_goTypes = []interface{}{
	(IODataType)(0), // 0: toco.IODataType
}
var file_tensorflow_lite_toco_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_lite_toco_types_proto_init() }
func file_tensorflow_lite_toco_types_proto_init() {
	if File_tensorflow_lite_toco_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_lite_toco_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_lite_toco_types_proto_goTypes,
		DependencyIndexes: file_tensorflow_lite_toco_types_proto_depIdxs,
		EnumInfos:         file_tensorflow_lite_toco_types_proto_enumTypes,
	}.Build()
	File_tensorflow_lite_toco_types_proto = out.File
	file_tensorflow_lite_toco_types_proto_rawDesc = nil
	file_tensorflow_lite_toco_types_proto_goTypes = nil
	file_tensorflow_lite_toco_types_proto_depIdxs = nil
}
