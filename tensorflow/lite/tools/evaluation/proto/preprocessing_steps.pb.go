// Copyright 2019 The TensorFlow Authors. All Rights Reserved.
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
// source: tensorflow/lite/tools/evaluation/proto/preprocessing_steps.proto

package proto

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

// Defines the preprocesing steps available.
//
// Next ID: 5
type ImagePreprocessingStepParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Params:
	//
	//	*ImagePreprocessingStepParams_CroppingParams
	//	*ImagePreprocessingStepParams_ResizingParams
	//	*ImagePreprocessingStepParams_PaddingParams
	//	*ImagePreprocessingStepParams_NormalizationParams
	Params isImagePreprocessingStepParams_Params `protobuf_oneof:"params"`
}

func (x *ImagePreprocessingStepParams) Reset() {
	*x = ImagePreprocessingStepParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImagePreprocessingStepParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImagePreprocessingStepParams) ProtoMessage() {}

func (x *ImagePreprocessingStepParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImagePreprocessingStepParams.ProtoReflect.Descriptor instead.
func (*ImagePreprocessingStepParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{0}
}

func (m *ImagePreprocessingStepParams) GetParams() isImagePreprocessingStepParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *ImagePreprocessingStepParams) GetCroppingParams() *CroppingParams {
	if x, ok := x.GetParams().(*ImagePreprocessingStepParams_CroppingParams); ok {
		return x.CroppingParams
	}
	return nil
}

func (x *ImagePreprocessingStepParams) GetResizingParams() *ResizingParams {
	if x, ok := x.GetParams().(*ImagePreprocessingStepParams_ResizingParams); ok {
		return x.ResizingParams
	}
	return nil
}

func (x *ImagePreprocessingStepParams) GetPaddingParams() *PaddingParams {
	if x, ok := x.GetParams().(*ImagePreprocessingStepParams_PaddingParams); ok {
		return x.PaddingParams
	}
	return nil
}

func (x *ImagePreprocessingStepParams) GetNormalizationParams() *NormalizationParams {
	if x, ok := x.GetParams().(*ImagePreprocessingStepParams_NormalizationParams); ok {
		return x.NormalizationParams
	}
	return nil
}

type isImagePreprocessingStepParams_Params interface {
	isImagePreprocessingStepParams_Params()
}

type ImagePreprocessingStepParams_CroppingParams struct {
	CroppingParams *CroppingParams `protobuf:"bytes,1,opt,name=cropping_params,json=croppingParams,oneof"`
}

type ImagePreprocessingStepParams_ResizingParams struct {
	ResizingParams *ResizingParams `protobuf:"bytes,2,opt,name=resizing_params,json=resizingParams,oneof"`
}

type ImagePreprocessingStepParams_PaddingParams struct {
	PaddingParams *PaddingParams `protobuf:"bytes,3,opt,name=padding_params,json=paddingParams,oneof"`
}

type ImagePreprocessingStepParams_NormalizationParams struct {
	NormalizationParams *NormalizationParams `protobuf:"bytes,4,opt,name=normalization_params,json=normalizationParams,oneof"`
}

func (*ImagePreprocessingStepParams_CroppingParams) isImagePreprocessingStepParams_Params() {}

func (*ImagePreprocessingStepParams_ResizingParams) isImagePreprocessingStepParams_Params() {}

func (*ImagePreprocessingStepParams_PaddingParams) isImagePreprocessingStepParams_Params() {}

func (*ImagePreprocessingStepParams_NormalizationParams) isImagePreprocessingStepParams_Params() {}

// Defines the size of an image.
//
// Next ID: 3
type ImageSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Width of the image.
	Width *uint32 `protobuf:"varint,1,req,name=width" json:"width,omitempty"`
	// Height of the image.
	Height *uint32 `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
}

func (x *ImageSize) Reset() {
	*x = ImageSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageSize) ProtoMessage() {}

func (x *ImageSize) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageSize.ProtoReflect.Descriptor instead.
func (*ImageSize) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{1}
}

func (x *ImageSize) GetWidth() uint32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

func (x *ImageSize) GetHeight() uint32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

// Defines parameters for central-cropping.
//
// Next ID: 4
type CroppingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Params:
	//
	//	*CroppingParams_CroppingFraction
	//	*CroppingParams_TargetSize
	Params isCroppingParams_Params `protobuf_oneof:"params"`
	// Crops to a square image.
	SquareCropping *bool `protobuf:"varint,3,opt,name=square_cropping,json=squareCropping" json:"square_cropping,omitempty"`
}

// Default values for CroppingParams fields.
const (
	Default_CroppingParams_CroppingFraction = float32(0.875)
)

func (x *CroppingParams) Reset() {
	*x = CroppingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CroppingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CroppingParams) ProtoMessage() {}

func (x *CroppingParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CroppingParams.ProtoReflect.Descriptor instead.
func (*CroppingParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{2}
}

func (m *CroppingParams) GetParams() isCroppingParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *CroppingParams) GetCroppingFraction() float32 {
	if x, ok := x.GetParams().(*CroppingParams_CroppingFraction); ok {
		return x.CroppingFraction
	}
	return Default_CroppingParams_CroppingFraction
}

func (x *CroppingParams) GetTargetSize() *ImageSize {
	if x, ok := x.GetParams().(*CroppingParams_TargetSize); ok {
		return x.TargetSize
	}
	return nil
}

func (x *CroppingParams) GetSquareCropping() bool {
	if x != nil && x.SquareCropping != nil {
		return *x.SquareCropping
	}
	return false
}

type isCroppingParams_Params interface {
	isCroppingParams_Params()
}

type CroppingParams_CroppingFraction struct {
	// Fraction for central-cropping.
	// A central cropping-fraction of 0.875 is considered best for Inception
	// models, hence the default value. See:
	// https://github.com/tensorflow/tpu/blob/master/models/experimental/inception/inception_preprocessing.py#L296
	// Set to 0 to disable cropping.
	CroppingFraction float32 `protobuf:"fixed32,1,opt,name=cropping_fraction,json=croppingFraction,oneof,def=0.875"`
}

type CroppingParams_TargetSize struct {
	// The target size after cropping.
	TargetSize *ImageSize `protobuf:"bytes,2,opt,name=target_size,json=targetSize,oneof"`
}

func (*CroppingParams_CroppingFraction) isCroppingParams_Params() {}

func (*CroppingParams_TargetSize) isCroppingParams_Params() {}

// Defines parameters for bilinear central-resizing.
//
// Next ID: 3
type ResizingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Size of the image after resizing.
	TargetSize *ImageSize `protobuf:"bytes,1,req,name=target_size,json=targetSize" json:"target_size,omitempty"`
	// If this flag is true, the resize function will preserve the image's aspect
	// ratio. Note that in this case, the size of output image may not equal to
	// the target size defined above.
	AspectPreserving *bool `protobuf:"varint,2,req,name=aspect_preserving,json=aspectPreserving" json:"aspect_preserving,omitempty"`
}

func (x *ResizingParams) Reset() {
	*x = ResizingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizingParams) ProtoMessage() {}

func (x *ResizingParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizingParams.ProtoReflect.Descriptor instead.
func (*ResizingParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{3}
}

func (x *ResizingParams) GetTargetSize() *ImageSize {
	if x != nil {
		return x.TargetSize
	}
	return nil
}

func (x *ResizingParams) GetAspectPreserving() bool {
	if x != nil && x.AspectPreserving != nil {
		return *x.AspectPreserving
	}
	return false
}

// Defines parameters for central-padding.
//
// Next ID: 4
type PaddingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Params:
	//
	//	*PaddingParams_TargetSize
	//	*PaddingParams_SquarePadding
	Params isPaddingParams_Params `protobuf_oneof:"params"`
	// Padding value.
	PaddingValue *int32 `protobuf:"varint,3,req,name=padding_value,json=paddingValue" json:"padding_value,omitempty"`
}

func (x *PaddingParams) Reset() {
	*x = PaddingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaddingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaddingParams) ProtoMessage() {}

func (x *PaddingParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaddingParams.ProtoReflect.Descriptor instead.
func (*PaddingParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{4}
}

func (m *PaddingParams) GetParams() isPaddingParams_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (x *PaddingParams) GetTargetSize() *ImageSize {
	if x, ok := x.GetParams().(*PaddingParams_TargetSize); ok {
		return x.TargetSize
	}
	return nil
}

func (x *PaddingParams) GetSquarePadding() bool {
	if x, ok := x.GetParams().(*PaddingParams_SquarePadding); ok {
		return x.SquarePadding
	}
	return false
}

func (x *PaddingParams) GetPaddingValue() int32 {
	if x != nil && x.PaddingValue != nil {
		return *x.PaddingValue
	}
	return 0
}

type isPaddingParams_Params interface {
	isPaddingParams_Params()
}

type PaddingParams_TargetSize struct {
	// Size of the image after padding.
	TargetSize *ImageSize `protobuf:"bytes,1,opt,name=target_size,json=targetSize,oneof"`
}

type PaddingParams_SquarePadding struct {
	// Pads to a square image.
	SquarePadding bool `protobuf:"varint,2,opt,name=square_padding,json=squarePadding,oneof"`
}

func (*PaddingParams_TargetSize) isPaddingParams_Params() {}

func (*PaddingParams_SquarePadding) isPaddingParams_Params() {}

// Defines parameters for normalization.
// The normalization formula is: output = (input - mean) * scale.
//
// Next ID: 4
type NormalizationParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Mean:
	//
	//	*NormalizationParams_ChannelwiseMean
	//	*NormalizationParams_Means
	Mean isNormalizationParams_Mean `protobuf_oneof:"mean"`
	// Scale value in the normalization.
	Scale *float32 `protobuf:"fixed32,3,req,name=scale,def=1" json:"scale,omitempty"`
}

// Default values for NormalizationParams fields.
const (
	Default_NormalizationParams_Scale = float32(1)
)

func (x *NormalizationParams) Reset() {
	*x = NormalizationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalizationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizationParams) ProtoMessage() {}

func (x *NormalizationParams) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizationParams.ProtoReflect.Descriptor instead.
func (*NormalizationParams) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{5}
}

func (m *NormalizationParams) GetMean() isNormalizationParams_Mean {
	if m != nil {
		return m.Mean
	}
	return nil
}

func (x *NormalizationParams) GetChannelwiseMean() float32 {
	if x, ok := x.GetMean().(*NormalizationParams_ChannelwiseMean); ok {
		return x.ChannelwiseMean
	}
	return 0
}

func (x *NormalizationParams) GetMeans() *NormalizationParams_PerChannelMeanValues {
	if x, ok := x.GetMean().(*NormalizationParams_Means); ok {
		return x.Means
	}
	return nil
}

func (x *NormalizationParams) GetScale() float32 {
	if x != nil && x.Scale != nil {
		return *x.Scale
	}
	return Default_NormalizationParams_Scale
}

type isNormalizationParams_Mean interface {
	isNormalizationParams_Mean()
}

type NormalizationParams_ChannelwiseMean struct {
	// Channelwise mean value.
	ChannelwiseMean float32 `protobuf:"fixed32,1,opt,name=channelwise_mean,json=channelwiseMean,oneof"`
}

type NormalizationParams_Means struct {
	// Per-Channel mean values.
	Means *NormalizationParams_PerChannelMeanValues `protobuf:"bytes,2,opt,name=means,oneof"`
}

func (*NormalizationParams_ChannelwiseMean) isNormalizationParams_Mean() {}

func (*NormalizationParams_Means) isNormalizationParams_Mean() {}

type NormalizationParams_PerChannelMeanValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mean values of r channel.
	RMean *float32 `protobuf:"fixed32,1,req,name=r_mean,json=rMean" json:"r_mean,omitempty"`
	// The mean values of g channel.
	GMean *float32 `protobuf:"fixed32,2,req,name=g_mean,json=gMean" json:"g_mean,omitempty"`
	// The mean values of b channel.
	BMean *float32 `protobuf:"fixed32,3,req,name=b_mean,json=bMean" json:"b_mean,omitempty"`
}

func (x *NormalizationParams_PerChannelMeanValues) Reset() {
	*x = NormalizationParams_PerChannelMeanValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalizationParams_PerChannelMeanValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalizationParams_PerChannelMeanValues) ProtoMessage() {}

func (x *NormalizationParams_PerChannelMeanValues) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalizationParams_PerChannelMeanValues.ProtoReflect.Descriptor instead.
func (*NormalizationParams_PerChannelMeanValues) Descriptor() ([]byte, []int) {
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP(), []int{5, 0}
}

func (x *NormalizationParams_PerChannelMeanValues) GetRMean() float32 {
	if x != nil && x.RMean != nil {
		return *x.RMean
	}
	return 0
}

func (x *NormalizationParams_PerChannelMeanValues) GetGMean() float32 {
	if x != nil && x.GMean != nil {
		return *x.GMean
	}
	return 0
}

func (x *NormalizationParams_PerChannelMeanValues) GetBMean() float32 {
	if x != nil && x.BMean != nil {
		return *x.BMean
	}
	return 0
}

var File_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto protoreflect.FileDescriptor

var file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDesc = []byte{
	0x0a, 0x40, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xec, 0x02, 0x0a, 0x1c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x4c, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x49, 0x0a, 0x0e, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x66, 0x6c,
	0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x0d,
	0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x5b, 0x0a,
	0x14, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x66,
	0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x48, 0x00, 0x52, 0x13, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x39, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0xba, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x34, 0x0a, 0x11, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x05, 0x30,
	0x2e, 0x38, 0x37, 0x35, 0x48, 0x00, 0x52, 0x10, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x43, 0x72, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x7c, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3d,
	0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x10, 0x61, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x50,
	0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3f, 0x0a, 0x0b,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x48,
	0x00, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a,
	0x0e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x50,
	0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x70,
	0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a,
	0x10, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x77, 0x69, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x61,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x77, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x61, 0x6e, 0x12, 0x53, 0x0a, 0x05, 0x6d, 0x65,
	0x61, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x74, 0x66, 0x6c, 0x69,
	0x74, 0x65, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x61, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x65, 0x61, 0x6e, 0x73, 0x12,
	0x17, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x02, 0x3a, 0x01,
	0x31, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x1a, 0x5b, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x02,
	0x52, 0x05, 0x72, 0x4d, 0x65, 0x61, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x5f, 0x6d, 0x65, 0x61,
	0x6e, 0x18, 0x02, 0x20, 0x02, 0x28, 0x02, 0x52, 0x05, 0x67, 0x4d, 0x65, 0x61, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x62, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x02, 0x52, 0x05,
	0x62, 0x4d, 0x65, 0x61, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6e, 0x42, 0xea, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x17, 0x50, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x65, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x11,
	0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0xca, 0x02, 0x11, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x1d, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x3a, 0x3a,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
}

var (
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescOnce sync.Once
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescData = file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDesc
)

func file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescGZIP() []byte {
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescOnce.Do(func() {
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescData)
	})
	return file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDescData
}

var file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_goTypes = []interface{}{
	(*ImagePreprocessingStepParams)(nil),             // 0: tflite.evaluation.ImagePreprocessingStepParams
	(*ImageSize)(nil),                                // 1: tflite.evaluation.ImageSize
	(*CroppingParams)(nil),                           // 2: tflite.evaluation.CroppingParams
	(*ResizingParams)(nil),                           // 3: tflite.evaluation.ResizingParams
	(*PaddingParams)(nil),                            // 4: tflite.evaluation.PaddingParams
	(*NormalizationParams)(nil),                      // 5: tflite.evaluation.NormalizationParams
	(*NormalizationParams_PerChannelMeanValues)(nil), // 6: tflite.evaluation.NormalizationParams.PerChannelMeanValues
}
var file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_depIdxs = []int32{
	2, // 0: tflite.evaluation.ImagePreprocessingStepParams.cropping_params:type_name -> tflite.evaluation.CroppingParams
	3, // 1: tflite.evaluation.ImagePreprocessingStepParams.resizing_params:type_name -> tflite.evaluation.ResizingParams
	4, // 2: tflite.evaluation.ImagePreprocessingStepParams.padding_params:type_name -> tflite.evaluation.PaddingParams
	5, // 3: tflite.evaluation.ImagePreprocessingStepParams.normalization_params:type_name -> tflite.evaluation.NormalizationParams
	1, // 4: tflite.evaluation.CroppingParams.target_size:type_name -> tflite.evaluation.ImageSize
	1, // 5: tflite.evaluation.ResizingParams.target_size:type_name -> tflite.evaluation.ImageSize
	1, // 6: tflite.evaluation.PaddingParams.target_size:type_name -> tflite.evaluation.ImageSize
	6, // 7: tflite.evaluation.NormalizationParams.means:type_name -> tflite.evaluation.NormalizationParams.PerChannelMeanValues
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_init() }
func file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_init() {
	if File_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImagePreprocessingStepParams); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageSize); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CroppingParams); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizingParams); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaddingParams); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalizationParams); i {
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
		file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalizationParams_PerChannelMeanValues); i {
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
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImagePreprocessingStepParams_CroppingParams)(nil),
		(*ImagePreprocessingStepParams_ResizingParams)(nil),
		(*ImagePreprocessingStepParams_PaddingParams)(nil),
		(*ImagePreprocessingStepParams_NormalizationParams)(nil),
	}
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CroppingParams_CroppingFraction)(nil),
		(*CroppingParams_TargetSize)(nil),
	}
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*PaddingParams_TargetSize)(nil),
		(*PaddingParams_SquarePadding)(nil),
	}
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*NormalizationParams_ChannelwiseMean)(nil),
		(*NormalizationParams_Means)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_goTypes,
		DependencyIndexes: file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_depIdxs,
		MessageInfos:      file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_msgTypes,
	}.Build()
	File_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto = out.File
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_rawDesc = nil
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_goTypes = nil
	file_tensorflow_lite_tools_evaluation_proto_preprocessing_steps_proto_depIdxs = nil
}
