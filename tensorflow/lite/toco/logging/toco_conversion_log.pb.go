// Copyright 2019 The TensorFlow Authors. All Rights Reserved.
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
// source: tensorflow/lite/toco/logging/toco_conversion_log.proto

package logging

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

// TocoConversionLog contains the analytics to be gathered when user converts
// a model to TF Lite using TOCO.
// Next ID to USE: 14.
type TocoConversionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total ops listed by name.
	OpList []string `protobuf:"bytes,1,rep,name=op_list,json=opList" json:"op_list,omitempty"`
	// Counts of built-in ops.
	// Key is op name and value is the count.
	BuiltInOps map[string]int32 `protobuf:"bytes,2,rep,name=built_in_ops,json=builtInOps" json:"built_in_ops,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Counts of custom ops.
	CustomOps map[string]int32 `protobuf:"bytes,3,rep,name=custom_ops,json=customOps" json:"custom_ops,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Counts of select ops.
	SelectOps map[string]int32 `protobuf:"bytes,4,rep,name=select_ops,json=selectOps" json:"select_ops,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// The signature of operators. Including ops input/output types and shapes,
	// op name and version.
	OpSignatures []string `protobuf:"bytes,5,rep,name=op_signatures,json=opSignatures" json:"op_signatures,omitempty"`
	// Input tensor types.
	InputTensorTypes []string `protobuf:"bytes,6,rep,name=input_tensor_types,json=inputTensorTypes" json:"input_tensor_types,omitempty"`
	// Output tensor types.
	OutputTensorTypes []string `protobuf:"bytes,7,rep,name=output_tensor_types,json=outputTensorTypes" json:"output_tensor_types,omitempty"`
	// Log generation time in micro-seconds.
	LogGenerationTs *int64 `protobuf:"varint,8,opt,name=log_generation_ts,json=logGenerationTs" json:"log_generation_ts,omitempty"`
	// Total number of ops in the model.
	ModelSize *int32 `protobuf:"varint,9,opt,name=model_size,json=modelSize" json:"model_size,omitempty"`
	// Tensorflow Lite runtime version.
	TfLiteVersion *string `protobuf:"bytes,10,opt,name=tf_lite_version,json=tfLiteVersion" json:"tf_lite_version,omitempty"`
	// Operating System info.
	OsVersion *string `protobuf:"bytes,11,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	// Model hash string.
	ModelHash *string `protobuf:"bytes,12,opt,name=model_hash,json=modelHash" json:"model_hash,omitempty"`
	// Error messages emitted by TOCO during conversion.
	TocoErrLogs *string `protobuf:"bytes,13,opt,name=toco_err_logs,json=tocoErrLogs" json:"toco_err_logs,omitempty"`
}

func (x *TocoConversionLog) Reset() {
	*x = TocoConversionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_toco_logging_toco_conversion_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TocoConversionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TocoConversionLog) ProtoMessage() {}

func (x *TocoConversionLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_toco_logging_toco_conversion_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TocoConversionLog.ProtoReflect.Descriptor instead.
func (*TocoConversionLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescGZIP(), []int{0}
}

func (x *TocoConversionLog) GetOpList() []string {
	if x != nil {
		return x.OpList
	}
	return nil
}

func (x *TocoConversionLog) GetBuiltInOps() map[string]int32 {
	if x != nil {
		return x.BuiltInOps
	}
	return nil
}

func (x *TocoConversionLog) GetCustomOps() map[string]int32 {
	if x != nil {
		return x.CustomOps
	}
	return nil
}

func (x *TocoConversionLog) GetSelectOps() map[string]int32 {
	if x != nil {
		return x.SelectOps
	}
	return nil
}

func (x *TocoConversionLog) GetOpSignatures() []string {
	if x != nil {
		return x.OpSignatures
	}
	return nil
}

func (x *TocoConversionLog) GetInputTensorTypes() []string {
	if x != nil {
		return x.InputTensorTypes
	}
	return nil
}

func (x *TocoConversionLog) GetOutputTensorTypes() []string {
	if x != nil {
		return x.OutputTensorTypes
	}
	return nil
}

func (x *TocoConversionLog) GetLogGenerationTs() int64 {
	if x != nil && x.LogGenerationTs != nil {
		return *x.LogGenerationTs
	}
	return 0
}

func (x *TocoConversionLog) GetModelSize() int32 {
	if x != nil && x.ModelSize != nil {
		return *x.ModelSize
	}
	return 0
}

func (x *TocoConversionLog) GetTfLiteVersion() string {
	if x != nil && x.TfLiteVersion != nil {
		return *x.TfLiteVersion
	}
	return ""
}

func (x *TocoConversionLog) GetOsVersion() string {
	if x != nil && x.OsVersion != nil {
		return *x.OsVersion
	}
	return ""
}

func (x *TocoConversionLog) GetModelHash() string {
	if x != nil && x.ModelHash != nil {
		return *x.ModelHash
	}
	return ""
}

func (x *TocoConversionLog) GetTocoErrLogs() string {
	if x != nil && x.TocoErrLogs != nil {
		return *x.TocoErrLogs
	}
	return ""
}

var File_tensorflow_lite_toco_logging_toco_conversion_log_proto protoreflect.FileDescriptor

var file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x74, 0x6f, 0x63, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x74,
	0x6f, 0x63, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x63, 0x6f, 0x22, 0x98,
	0x06, 0x0a, 0x11, 0x54, 0x6f, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x49, 0x0a,
	0x0c, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x6f, 0x63, 0x6f, 0x2e, 0x54, 0x6f, 0x63, 0x6f, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x74, 0x49, 0x6e, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x75,
	0x69, 0x6c, 0x74, 0x49, 0x6e, 0x4f, 0x70, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74,
	0x6f, 0x63, 0x6f, 0x2e, 0x54, 0x6f, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x73, 0x12,
	0x45, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x6f, 0x63, 0x6f, 0x2e, 0x54, 0x6f, 0x63, 0x6f, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x4f, 0x70, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x5f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x66, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74,
	0x66, 0x4c, 0x69, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f,
	0x63, 0x6f, 0x5f, 0x65, 0x72, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x6f, 0x63, 0x6f, 0x45, 0x72, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x3d,
	0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x49, 0x6e, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a,
	0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x9a, 0x01, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x6f, 0x63, 0x6f, 0x42, 0x16, 0x54, 0x6f, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72,
	0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x63, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x04, 0x54, 0x6f, 0x63, 0x6f, 0xca, 0x02, 0x04, 0x54, 0x6f, 0x63, 0x6f, 0xe2, 0x02, 0x10, 0x54,
	0x6f, 0x63, 0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x04, 0x54, 0x6f, 0x63, 0x6f,
}

var (
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescOnce sync.Once
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescData = file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDesc
)

func file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescGZIP() []byte {
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescOnce.Do(func() {
		file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescData)
	})
	return file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDescData
}

var file_tensorflow_lite_toco_logging_toco_conversion_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_lite_toco_logging_toco_conversion_log_proto_goTypes = []interface{}{
	(*TocoConversionLog)(nil), // 0: toco.TocoConversionLog
	nil,                       // 1: toco.TocoConversionLog.BuiltInOpsEntry
	nil,                       // 2: toco.TocoConversionLog.CustomOpsEntry
	nil,                       // 3: toco.TocoConversionLog.SelectOpsEntry
}
var file_tensorflow_lite_toco_logging_toco_conversion_log_proto_depIdxs = []int32{
	1, // 0: toco.TocoConversionLog.built_in_ops:type_name -> toco.TocoConversionLog.BuiltInOpsEntry
	2, // 1: toco.TocoConversionLog.custom_ops:type_name -> toco.TocoConversionLog.CustomOpsEntry
	3, // 2: toco.TocoConversionLog.select_ops:type_name -> toco.TocoConversionLog.SelectOpsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_lite_toco_logging_toco_conversion_log_proto_init() }
func file_tensorflow_lite_toco_logging_toco_conversion_log_proto_init() {
	if File_tensorflow_lite_toco_logging_toco_conversion_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_lite_toco_logging_toco_conversion_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TocoConversionLog); i {
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
			RawDescriptor: file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_lite_toco_logging_toco_conversion_log_proto_goTypes,
		DependencyIndexes: file_tensorflow_lite_toco_logging_toco_conversion_log_proto_depIdxs,
		MessageInfos:      file_tensorflow_lite_toco_logging_toco_conversion_log_proto_msgTypes,
	}.Build()
	File_tensorflow_lite_toco_logging_toco_conversion_log_proto = out.File
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_rawDesc = nil
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_goTypes = nil
	file_tensorflow_lite_toco_logging_toco_conversion_log_proto_depIdxs = nil
}
