// LINT: LEGACY_NAMES

// Copyright 2021 The TensorFlow Authors. All Rights Reserved.
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

// Protocol messages for describing parameters of convolution-related
// operations.

// For Google-internal use only.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/util/autotune_maps/conv_parameters.proto

package autotune_maps

import (
	_ "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/stream_executor"
	framework "github.com/airenas/go-tf-serving-protogen/tensorflow/core/framework"
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

// This is the underlying data structure of class ConvParameters, which are used
// as the keys in cuDNN autotuning maps for retrieving corresponding cuDNN
// algorithms. This is used as a serialization format for saving/loading
// autotuning databases.
type ConvParametersProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch     int64   `protobuf:"varint,1,opt,name=batch,proto3" json:"batch,omitempty"`
	InDepths  int64   `protobuf:"varint,2,opt,name=in_depths,json=inDepths,proto3" json:"in_depths,omitempty"`
	OutDepths int64   `protobuf:"varint,3,opt,name=out_depths,json=outDepths,proto3" json:"out_depths,omitempty"`
	In        []int64 `protobuf:"varint,4,rep,packed,name=in,proto3" json:"in,omitempty"`
	// data_format corresponds to type TensorFormat in
	// third_party/tensorflow/core/util/tensor_format.h.
	DataFormat int32              `protobuf:"varint,5,opt,name=data_format,json=dataFormat,proto3" json:"data_format,omitempty"`
	Filter     []int64            `protobuf:"varint,6,rep,packed,name=filter,proto3" json:"filter,omitempty"`
	Dilation   []int64            `protobuf:"varint,7,rep,packed,name=dilation,proto3" json:"dilation,omitempty"`
	Stride     []int64            `protobuf:"varint,8,rep,packed,name=stride,proto3" json:"stride,omitempty"`
	Padding    []int64            `protobuf:"varint,9,rep,packed,name=padding,proto3" json:"padding,omitempty"`
	Dtype      framework.DataType `protobuf:"varint,10,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	GroupCount int32              `protobuf:"varint,11,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	// A string uniquely identifying a particular GPU model, e.g. V100 vs RTX
	// 2080.
	DeviceIdentifier string                      `protobuf:"bytes,12,opt,name=device_identifier,json=deviceIdentifier,proto3" json:"device_identifier,omitempty"`
	Fusion           *ConvParametersProto_Fusion `protobuf:"bytes,13,opt,name=fusion,proto3" json:"fusion,omitempty"`
	// The version number of ConvParameters class. Offline autotune results
	// whose version number is different from the runtime's version number
	// (defined in ConvParameters::kVersion) will be rejected and ignored by
	// LoadSerializedAutotuneMaps. This ensures that we will not load out-of-date
	// autotune results.
	Version int32 `protobuf:"varint,14,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ConvParametersProto) Reset() {
	*x = ConvParametersProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvParametersProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvParametersProto) ProtoMessage() {}

func (x *ConvParametersProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvParametersProto.ProtoReflect.Descriptor instead.
func (*ConvParametersProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescGZIP(), []int{0}
}

func (x *ConvParametersProto) GetBatch() int64 {
	if x != nil {
		return x.Batch
	}
	return 0
}

func (x *ConvParametersProto) GetInDepths() int64 {
	if x != nil {
		return x.InDepths
	}
	return 0
}

func (x *ConvParametersProto) GetOutDepths() int64 {
	if x != nil {
		return x.OutDepths
	}
	return 0
}

func (x *ConvParametersProto) GetIn() []int64 {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *ConvParametersProto) GetDataFormat() int32 {
	if x != nil {
		return x.DataFormat
	}
	return 0
}

func (x *ConvParametersProto) GetFilter() []int64 {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ConvParametersProto) GetDilation() []int64 {
	if x != nil {
		return x.Dilation
	}
	return nil
}

func (x *ConvParametersProto) GetStride() []int64 {
	if x != nil {
		return x.Stride
	}
	return nil
}

func (x *ConvParametersProto) GetPadding() []int64 {
	if x != nil {
		return x.Padding
	}
	return nil
}

func (x *ConvParametersProto) GetDtype() framework.DataType {
	if x != nil {
		return x.Dtype
	}
	return framework.DataType(0)
}

func (x *ConvParametersProto) GetGroupCount() int32 {
	if x != nil {
		return x.GroupCount
	}
	return 0
}

func (x *ConvParametersProto) GetDeviceIdentifier() string {
	if x != nil {
		return x.DeviceIdentifier
	}
	return ""
}

func (x *ConvParametersProto) GetFusion() *ConvParametersProto_Fusion {
	if x != nil {
		return x.Fusion
	}
	return nil
}

func (x *ConvParametersProto) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type MatmulParametersProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbDtype          framework.DataType      `protobuf:"varint,1,opt,name=ab_dtype,json=abDtype,proto3,enum=tensorflow.DataType" json:"ab_dtype,omitempty"`
	CDtype           framework.DataType      `protobuf:"varint,2,opt,name=c_dtype,json=cDtype,proto3,enum=tensorflow.DataType" json:"c_dtype,omitempty"`
	TransA           bool                    `protobuf:"varint,3,opt,name=trans_a,json=transA,proto3" json:"trans_a,omitempty"`
	TransB           bool                    `protobuf:"varint,4,opt,name=trans_b,json=transB,proto3" json:"trans_b,omitempty"`
	M                uint64                  `protobuf:"varint,5,opt,name=m,proto3" json:"m,omitempty"`
	N                uint64                  `protobuf:"varint,6,opt,name=n,proto3" json:"n,omitempty"`
	K                uint64                  `protobuf:"varint,7,opt,name=k,proto3" json:"k,omitempty"`
	Lda              int64                   `protobuf:"varint,8,opt,name=lda,proto3" json:"lda,omitempty"`
	Ldb              int64                   `protobuf:"varint,9,opt,name=ldb,proto3" json:"ldb,omitempty"`
	Ldc              int64                   `protobuf:"varint,10,opt,name=ldc,proto3" json:"ldc,omitempty"`
	ActivationMode   protobuf.ActivationMode `protobuf:"varint,11,opt,name=activation_mode,json=activationMode,proto3,enum=stream_executor.dnn.ActivationMode" json:"activation_mode,omitempty"`
	DeviceIdentifier string                  `protobuf:"bytes,12,opt,name=device_identifier,json=deviceIdentifier,proto3" json:"device_identifier,omitempty"`
	Version          int32                   `protobuf:"varint,14,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *MatmulParametersProto) Reset() {
	*x = MatmulParametersProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatmulParametersProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatmulParametersProto) ProtoMessage() {}

func (x *MatmulParametersProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatmulParametersProto.ProtoReflect.Descriptor instead.
func (*MatmulParametersProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescGZIP(), []int{1}
}

func (x *MatmulParametersProto) GetAbDtype() framework.DataType {
	if x != nil {
		return x.AbDtype
	}
	return framework.DataType(0)
}

func (x *MatmulParametersProto) GetCDtype() framework.DataType {
	if x != nil {
		return x.CDtype
	}
	return framework.DataType(0)
}

func (x *MatmulParametersProto) GetTransA() bool {
	if x != nil {
		return x.TransA
	}
	return false
}

func (x *MatmulParametersProto) GetTransB() bool {
	if x != nil {
		return x.TransB
	}
	return false
}

func (x *MatmulParametersProto) GetM() uint64 {
	if x != nil {
		return x.M
	}
	return 0
}

func (x *MatmulParametersProto) GetN() uint64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *MatmulParametersProto) GetK() uint64 {
	if x != nil {
		return x.K
	}
	return 0
}

func (x *MatmulParametersProto) GetLda() int64 {
	if x != nil {
		return x.Lda
	}
	return 0
}

func (x *MatmulParametersProto) GetLdb() int64 {
	if x != nil {
		return x.Ldb
	}
	return 0
}

func (x *MatmulParametersProto) GetLdc() int64 {
	if x != nil {
		return x.Ldc
	}
	return 0
}

func (x *MatmulParametersProto) GetActivationMode() protobuf.ActivationMode {
	if x != nil {
		return x.ActivationMode
	}
	return protobuf.ActivationMode(0)
}

func (x *MatmulParametersProto) GetDeviceIdentifier() string {
	if x != nil {
		return x.DeviceIdentifier
	}
	return ""
}

func (x *MatmulParametersProto) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// This stores the information for fused convolution operations where an
// activation and a side input might follow the convolution.
type ConvParametersProto_Fusion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, this proto corresponds to a FusedConvBiasActivation operation
	// implemented in the contrib library, otherwise this proto corresponds to
	// the FusedConv operation implemented in the core library.
	// Compared with FusedConv, FusedConvBiasActivation supports more types of
	// activation function (including no activation) as well as the side_input.
	// For now they have same type of keys in autotune maps, but the semantics
	// of some fields (like padding) are different. So we add this field to
	// distinguish them.
	// TODO(b/177365158) Remove this field once these two operations are merged.
	IsContrib      bool                    `protobuf:"varint,1,opt,name=is_contrib,json=isContrib,proto3" json:"is_contrib,omitempty"`
	ActivationMode protobuf.ActivationMode `protobuf:"varint,2,opt,name=activation_mode,json=activationMode,proto3,enum=stream_executor.dnn.ActivationMode" json:"activation_mode,omitempty"`
	ConvScale      float64                 `protobuf:"fixed64,3,opt,name=conv_scale,json=convScale,proto3" json:"conv_scale,omitempty"`
	SideInputScale float64                 `protobuf:"fixed64,4,opt,name=side_input_scale,json=sideInputScale,proto3" json:"side_input_scale,omitempty"`
}

func (x *ConvParametersProto_Fusion) Reset() {
	*x = ConvParametersProto_Fusion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvParametersProto_Fusion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvParametersProto_Fusion) ProtoMessage() {}

func (x *ConvParametersProto_Fusion) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvParametersProto_Fusion.ProtoReflect.Descriptor instead.
func (*ConvParametersProto_Fusion) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConvParametersProto_Fusion) GetIsContrib() bool {
	if x != nil {
		return x.IsContrib
	}
	return false
}

func (x *ConvParametersProto_Fusion) GetActivationMode() protobuf.ActivationMode {
	if x != nil {
		return x.ActivationMode
	}
	return protobuf.ActivationMode(0)
}

func (x *ConvParametersProto_Fusion) GetConvScale() float64 {
	if x != nil {
		return x.ConvScale
	}
	return 0
}

func (x *ConvParametersProto_Fusion) GetSideInputScale() float64 {
	if x != nil {
		return x.SideInputScale
	}
	return 0
}

var File_tensorflow_core_util_autotune_maps_conv_parameters_proto protoreflect.FileDescriptor

var file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x31, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f,
	0x64, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x93, 0x05, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x76, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x74, 0x68, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x75, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6f, 0x75, 0x74, 0x44, 0x65, 0x70, 0x74, 0x68, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x09, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b,
	0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x06, 0x66,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x75, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xbe, 0x01, 0x0a, 0x06, 0x46, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x12,
	0x4c, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x76, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x9e, 0x03, 0x0a, 0x15, 0x4d, 0x61, 0x74, 0x6d, 0x75,
	0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2f, 0x0a, 0x08, 0x61, 0x62, 0x5f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x61, 0x62, 0x44, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x5f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x63, 0x44, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x5f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x42, 0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x6d,
	0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x6e, 0x12, 0x0c,
	0x0a, 0x01, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x64, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x64, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x64, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x64, 0x62,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x64, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c,
	0x64, 0x63, 0x12, 0x4c, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e,
	0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xbb, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x13, 0x43, 0x6f, 0x6e, 0x76,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69,
	0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x73, 0xa2,
	0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2,
	0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescOnce sync.Once
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescData = file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDesc
)

func file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescGZIP() []byte {
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescData)
	})
	return file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDescData
}

var file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_util_autotune_maps_conv_parameters_proto_goTypes = []interface{}{
	(*ConvParametersProto)(nil),        // 0: tensorflow.ConvParametersProto
	(*MatmulParametersProto)(nil),      // 1: tensorflow.MatmulParametersProto
	(*ConvParametersProto_Fusion)(nil), // 2: tensorflow.ConvParametersProto.Fusion
	(framework.DataType)(0),            // 3: tensorflow.DataType
	(protobuf.ActivationMode)(0),       // 4: stream_executor.dnn.ActivationMode
}
var file_tensorflow_core_util_autotune_maps_conv_parameters_proto_depIdxs = []int32{
	3, // 0: tensorflow.ConvParametersProto.dtype:type_name -> tensorflow.DataType
	2, // 1: tensorflow.ConvParametersProto.fusion:type_name -> tensorflow.ConvParametersProto.Fusion
	3, // 2: tensorflow.MatmulParametersProto.ab_dtype:type_name -> tensorflow.DataType
	3, // 3: tensorflow.MatmulParametersProto.c_dtype:type_name -> tensorflow.DataType
	4, // 4: tensorflow.MatmulParametersProto.activation_mode:type_name -> stream_executor.dnn.ActivationMode
	4, // 5: tensorflow.ConvParametersProto.Fusion.activation_mode:type_name -> stream_executor.dnn.ActivationMode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_core_util_autotune_maps_conv_parameters_proto_init() }
func file_tensorflow_core_util_autotune_maps_conv_parameters_proto_init() {
	if File_tensorflow_core_util_autotune_maps_conv_parameters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvParametersProto); i {
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
		file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatmulParametersProto); i {
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
		file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvParametersProto_Fusion); i {
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
			RawDescriptor: file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_util_autotune_maps_conv_parameters_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_util_autotune_maps_conv_parameters_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_util_autotune_maps_conv_parameters_proto_msgTypes,
	}.Build()
	File_tensorflow_core_util_autotune_maps_conv_parameters_proto = out.File
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_rawDesc = nil
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_goTypes = nil
	file_tensorflow_core_util_autotune_maps_conv_parameters_proto_depIdxs = nil
}
