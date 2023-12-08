// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/mlir/lite/quantization/quantization_info.proto

package quantization

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

// Represents the quantization parameters for a list of named tensors.
type QuantizationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of quantization parameters for tensors.
	Entries []*QuantizationInfo_QuantParams `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *QuantizationInfo) Reset() {
	*x = QuantizationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo) ProtoMessage() {}

func (x *QuantizationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo.ProtoReflect.Descriptor instead.
func (*QuantizationInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0}
}

func (x *QuantizationInfo) GetEntries() []*QuantizationInfo_QuantParams {
	if x != nil {
		return x.Entries
	}
	return nil
}

// min/max of the per axis value range. To quantize the value, the metadata
// of the target properties should be specified or read from the ops
// quantization specification.
type QuantizationInfo_MinMax struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min float32 `protobuf:"fixed32,1,opt,name=min,proto3" json:"min,omitempty"`
	Max float32 `protobuf:"fixed32,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *QuantizationInfo_MinMax) Reset() {
	*x = QuantizationInfo_MinMax{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo_MinMax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo_MinMax) ProtoMessage() {}

func (x *QuantizationInfo_MinMax) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo_MinMax.ProtoReflect.Descriptor instead.
func (*QuantizationInfo_MinMax) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0, 0}
}

func (x *QuantizationInfo_MinMax) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *QuantizationInfo_MinMax) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

// Affine parameters to quantize the per axis value. The metadata of the
// target properties should be specified as well.
type QuantizationInfo_AffineParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scale     float32 `protobuf:"fixed32,1,opt,name=scale,proto3" json:"scale,omitempty"`
	ZeroPoint int32   `protobuf:"varint,2,opt,name=zero_point,json=zeroPoint,proto3" json:"zero_point,omitempty"`
}

func (x *QuantizationInfo_AffineParams) Reset() {
	*x = QuantizationInfo_AffineParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo_AffineParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo_AffineParams) ProtoMessage() {}

func (x *QuantizationInfo_AffineParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo_AffineParams.ProtoReflect.Descriptor instead.
func (*QuantizationInfo_AffineParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0, 1}
}

func (x *QuantizationInfo_AffineParams) GetScale() float32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *QuantizationInfo_AffineParams) GetZeroPoint() int32 {
	if x != nil {
		return x.ZeroPoint
	}
	return 0
}

// Params to quantize the axis. Only one of the field can be used.
type QuantizationInfo_PerAxisParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ParamsOneof:
	//
	//	*QuantizationInfo_PerAxisParams_MinMax
	//	*QuantizationInfo_PerAxisParams_AffineParams
	ParamsOneof isQuantizationInfo_PerAxisParams_ParamsOneof `protobuf_oneof:"params_oneof"`
}

func (x *QuantizationInfo_PerAxisParams) Reset() {
	*x = QuantizationInfo_PerAxisParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo_PerAxisParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo_PerAxisParams) ProtoMessage() {}

func (x *QuantizationInfo_PerAxisParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo_PerAxisParams.ProtoReflect.Descriptor instead.
func (*QuantizationInfo_PerAxisParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0, 2}
}

func (m *QuantizationInfo_PerAxisParams) GetParamsOneof() isQuantizationInfo_PerAxisParams_ParamsOneof {
	if m != nil {
		return m.ParamsOneof
	}
	return nil
}

func (x *QuantizationInfo_PerAxisParams) GetMinMax() *QuantizationInfo_MinMax {
	if x, ok := x.GetParamsOneof().(*QuantizationInfo_PerAxisParams_MinMax); ok {
		return x.MinMax
	}
	return nil
}

func (x *QuantizationInfo_PerAxisParams) GetAffineParams() *QuantizationInfo_AffineParams {
	if x, ok := x.GetParamsOneof().(*QuantizationInfo_PerAxisParams_AffineParams); ok {
		return x.AffineParams
	}
	return nil
}

type isQuantizationInfo_PerAxisParams_ParamsOneof interface {
	isQuantizationInfo_PerAxisParams_ParamsOneof()
}

type QuantizationInfo_PerAxisParams_MinMax struct {
	// min/max of the ranges.
	MinMax *QuantizationInfo_MinMax `protobuf:"bytes,1,opt,name=min_max,json=minMax,proto3,oneof"`
}

type QuantizationInfo_PerAxisParams_AffineParams struct {
	// affine parameters to quantize the per axis value.
	AffineParams *QuantizationInfo_AffineParams `protobuf:"bytes,2,opt,name=affine_params,json=affineParams,proto3,oneof"`
}

func (*QuantizationInfo_PerAxisParams_MinMax) isQuantizationInfo_PerAxisParams_ParamsOneof() {}

func (*QuantizationInfo_PerAxisParams_AffineParams) isQuantizationInfo_PerAxisParams_ParamsOneof() {}

// The metadata defines the target properties.
type QuantizationInfo_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bit number of fixed-point data the target kernel supports.
	NumBits int32 `protobuf:"varint,1,opt,name=num_bits,json=numBits,proto3" json:"num_bits,omitempty"`
	// The quantized axis index if it is per-axis quantization.
	QuantizeAxis int32 `protobuf:"varint,2,opt,name=quantize_axis,json=quantizeAxis,proto3" json:"quantize_axis,omitempty"`
	// The minimum allowed value of the fixed-point data range.
	// This can also be used to derive the sign of storage type.
	RangeMin int32 `protobuf:"varint,3,opt,name=range_min,json=rangeMin,proto3" json:"range_min,omitempty"`
	// The minimum allowed value of the fixed-point data range.
	RangeMax int32 `protobuf:"varint,4,opt,name=range_max,json=rangeMax,proto3" json:"range_max,omitempty"`
}

func (x *QuantizationInfo_Metadata) Reset() {
	*x = QuantizationInfo_Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo_Metadata) ProtoMessage() {}

func (x *QuantizationInfo_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo_Metadata.ProtoReflect.Descriptor instead.
func (*QuantizationInfo_Metadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0, 3}
}

func (x *QuantizationInfo_Metadata) GetNumBits() int32 {
	if x != nil {
		return x.NumBits
	}
	return 0
}

func (x *QuantizationInfo_Metadata) GetQuantizeAxis() int32 {
	if x != nil {
		return x.QuantizeAxis
	}
	return 0
}

func (x *QuantizationInfo_Metadata) GetRangeMin() int32 {
	if x != nil {
		return x.RangeMin
	}
	return 0
}

func (x *QuantizationInfo_Metadata) GetRangeMax() int32 {
	if x != nil {
		return x.RangeMax
	}
	return 0
}

// The quantization parameters for a named tensor.
type QuantizationInfo_QuantParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tensor name has the following convention:
	//
	//	tensor_name := op_name | op_name  ’:’  port_number.
	//
	// If the op has only one port, op_name can be used.
	// If the op has internal states, such as fused LSTM, the port_number should
	// follow a predefined convention.
	//
	// Types that are assignable to NameOneof:
	//
	//	*QuantizationInfo_QuantParams_Name
	//	*QuantizationInfo_QuantParams_NameRegex
	NameOneof isQuantizationInfo_QuantParams_NameOneof `protobuf_oneof:"name_oneof"`
	// The quantization parameters for the tensor. If it is for per-axis, the
	// parameters should be defined for each axis, otherwise, if it is for
	// per-tensor, this repeated field should only contain a single element.
	Params []*QuantizationInfo_PerAxisParams `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	// Metadata about the quantization parameters.
	Meta *QuantizationInfo_Metadata `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *QuantizationInfo_QuantParams) Reset() {
	*x = QuantizationInfo_QuantParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuantizationInfo_QuantParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuantizationInfo_QuantParams) ProtoMessage() {}

func (x *QuantizationInfo_QuantParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuantizationInfo_QuantParams.ProtoReflect.Descriptor instead.
func (*QuantizationInfo_QuantParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP(), []int{0, 4}
}

func (m *QuantizationInfo_QuantParams) GetNameOneof() isQuantizationInfo_QuantParams_NameOneof {
	if m != nil {
		return m.NameOneof
	}
	return nil
}

func (x *QuantizationInfo_QuantParams) GetName() string {
	if x, ok := x.GetNameOneof().(*QuantizationInfo_QuantParams_Name); ok {
		return x.Name
	}
	return ""
}

func (x *QuantizationInfo_QuantParams) GetNameRegex() string {
	if x, ok := x.GetNameOneof().(*QuantizationInfo_QuantParams_NameRegex); ok {
		return x.NameRegex
	}
	return ""
}

func (x *QuantizationInfo_QuantParams) GetParams() []*QuantizationInfo_PerAxisParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *QuantizationInfo_QuantParams) GetMeta() *QuantizationInfo_Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type isQuantizationInfo_QuantParams_NameOneof interface {
	isQuantizationInfo_QuantParams_NameOneof()
}

type QuantizationInfo_QuantParams_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type QuantizationInfo_QuantParams_NameRegex struct {
	// An regex can be used to match multiple tensors.
	NameRegex string `protobuf:"bytes,2,opt,name=name_regex,json=nameRegex,proto3,oneof"`
}

func (*QuantizationInfo_QuantParams_Name) isQuantizationInfo_QuantParams_NameOneof() {}

func (*QuantizationInfo_QuantParams_NameRegex) isQuantizationInfo_QuantParams_NameOneof() {}

var File_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDesc = []byte{
	0x0a, 0x42, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x22, 0xd8, 0x05, 0x0a, 0x10, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x2c, 0x0a, 0x06, 0x4d, 0x69, 0x6e,
	0x4d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x1a, 0x43, 0x0a, 0x0c, 0x41, 0x66, 0x66, 0x69, 0x6e,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x7a, 0x65, 0x72, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0xb1, 0x01, 0x0a,
	0x0d, 0x50, 0x65, 0x72, 0x41, 0x78, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3e,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x69,
	0x6e, 0x4d, 0x61, 0x78, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x50,
	0x0a, 0x0d, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x2e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x66, 0x66, 0x69, 0x6e, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x48, 0x00, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x69, 0x6e, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x0e, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x1a, 0x84, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6e, 0x75, 0x6d, 0x42, 0x69, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x7a, 0x65, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x65, 0x41, 0x78, 0x69, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x1a, 0xd1, 0x01, 0x0a, 0x0b, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x42,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2e, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x65, 0x72,
	0x41, 0x78, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x2e, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x0c, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0xc9, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x42, 0x15,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74,
	0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x6c, 0x69, 0x74, 0x65,
	0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x4d, 0x51, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x6c, 0x69, 0x72, 0x2e, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0xca, 0x02, 0x0a, 0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x51, 0x75, 0x61, 0x6e, 0x74,
	0xe2, 0x02, 0x16, 0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4d, 0x6c, 0x69, 0x72,
	0x3a, 0x3a, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescData = file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDesc
)

func file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescData)
	})
	return file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDescData
}

var file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_goTypes = []interface{}{
	(*QuantizationInfo)(nil),               // 0: mlir.quant.QuantizationInfo
	(*QuantizationInfo_MinMax)(nil),        // 1: mlir.quant.QuantizationInfo.MinMax
	(*QuantizationInfo_AffineParams)(nil),  // 2: mlir.quant.QuantizationInfo.AffineParams
	(*QuantizationInfo_PerAxisParams)(nil), // 3: mlir.quant.QuantizationInfo.PerAxisParams
	(*QuantizationInfo_Metadata)(nil),      // 4: mlir.quant.QuantizationInfo.Metadata
	(*QuantizationInfo_QuantParams)(nil),   // 5: mlir.quant.QuantizationInfo.QuantParams
}
var file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_depIdxs = []int32{
	5, // 0: mlir.quant.QuantizationInfo.entries:type_name -> mlir.quant.QuantizationInfo.QuantParams
	1, // 1: mlir.quant.QuantizationInfo.PerAxisParams.min_max:type_name -> mlir.quant.QuantizationInfo.MinMax
	2, // 2: mlir.quant.QuantizationInfo.PerAxisParams.affine_params:type_name -> mlir.quant.QuantizationInfo.AffineParams
	3, // 3: mlir.quant.QuantizationInfo.QuantParams.params:type_name -> mlir.quant.QuantizationInfo.PerAxisParams
	4, // 4: mlir.quant.QuantizationInfo.QuantParams.meta:type_name -> mlir.quant.QuantizationInfo.Metadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_init() }
func file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_init() {
	if File_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo); i {
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
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo_MinMax); i {
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
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo_AffineParams); i {
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
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo_PerAxisParams); i {
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
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo_Metadata); i {
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
		file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuantizationInfo_QuantParams); i {
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
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*QuantizationInfo_PerAxisParams_MinMax)(nil),
		(*QuantizationInfo_PerAxisParams_AffineParams)(nil),
	}
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*QuantizationInfo_QuantParams_Name)(nil),
		(*QuantizationInfo_QuantParams_NameRegex)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto = out.File
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_rawDesc = nil
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_goTypes = nil
	file_tensorflow_compiler_mlir_lite_quantization_quantization_info_proto_depIdxs = nil
}
