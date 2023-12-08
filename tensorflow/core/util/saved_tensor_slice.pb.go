// Protocol buffers for saved tensor slices. It's used for the brain tensor
// ops checkpoints and the V3 checkpoints in dist_belief.

// A checkpoint file is an sstable. The value for each record is a serialized
// SavedTensorSlices message (defined below).
//
// Each checkpoint file has a record with the empty key (""), which corresponds
// to a SavedTensorSlices message that contains a "meta", that serves as a
// table of contents on all the tensor slices saved in this file. Since the key
// is "", it's always the first record in each file.
//
// Each of the rest of the records in a checkpoint stores the raw data of a
// particular tensor slice, in SavedSlice format. The corresponding key is an
// ordered code that encodes the name of the tensor and the slice
// information. The name is also stored in the SaveSlice message for ease of
// debugging and manual examination.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/util/saved_tensor_slice.proto

package util

import (
	framework "github.com/airenas/go-tf-serving-protogen/tensorflow/core/framework"
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

// Metadata describing the set of slices of the same tensor saved in a
// checkpoint file.
type SavedSliceMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the tensor.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Shape of the tensor
	Shape *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	// Type of the tensor
	Type framework.DataType `protobuf:"varint,3,opt,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	// Explicit list of slices saved in the checkpoint file.
	Slice []*framework.TensorSliceProto `protobuf:"bytes,4,rep,name=slice,proto3" json:"slice,omitempty"`
}

func (x *SavedSliceMeta) Reset() {
	*x = SavedSliceMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedSliceMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedSliceMeta) ProtoMessage() {}

func (x *SavedSliceMeta) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedSliceMeta.ProtoReflect.Descriptor instead.
func (*SavedSliceMeta) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_saved_tensor_slice_proto_rawDescGZIP(), []int{0}
}

func (x *SavedSliceMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SavedSliceMeta) GetShape() *framework.TensorShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *SavedSliceMeta) GetType() framework.DataType {
	if x != nil {
		return x.Type
	}
	return framework.DataType(0)
}

func (x *SavedSliceMeta) GetSlice() []*framework.TensorSliceProto {
	if x != nil {
		return x.Slice
	}
	return nil
}

// Metadata describing the set of tensor slices saved in a checkpoint file.
// It is always stored at the beginning of each checkpoint file.
type SavedTensorSliceMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each SavedSliceMeta describes the slices for one tensor.
	Tensor []*SavedSliceMeta `protobuf:"bytes,1,rep,name=tensor,proto3" json:"tensor,omitempty"`
	// Compatibility version of this checkpoint.  See core/public/version.h
	// for version history.
	Versions *framework.VersionDef `protobuf:"bytes,2,opt,name=versions,proto3" json:"versions,omitempty"`
}

func (x *SavedTensorSliceMeta) Reset() {
	*x = SavedTensorSliceMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedTensorSliceMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedTensorSliceMeta) ProtoMessage() {}

func (x *SavedTensorSliceMeta) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedTensorSliceMeta.ProtoReflect.Descriptor instead.
func (*SavedTensorSliceMeta) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_saved_tensor_slice_proto_rawDescGZIP(), []int{1}
}

func (x *SavedTensorSliceMeta) GetTensor() []*SavedSliceMeta {
	if x != nil {
		return x.Tensor
	}
	return nil
}

func (x *SavedTensorSliceMeta) GetVersions() *framework.VersionDef {
	if x != nil {
		return x.Versions
	}
	return nil
}

// Saved tensor slice: it stores the name of the tensors, the slice, and the
// raw data.
type SavedSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the tensor that this slice belongs to. This must be identical to
	// the name used to encode the key for this record.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Extent of the slice.  Must have one entry for each of the dimension of the
	// tensor that this slice belongs to.
	Slice *framework.TensorSliceProto `protobuf:"bytes,2,opt,name=slice,proto3" json:"slice,omitempty"`
	// The raw data of the slice is stored as a TensorProto. Only raw data are
	// stored (we don't fill in fields such as dtype or tensor_shape).
	Data *framework.TensorProto `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SavedSlice) Reset() {
	*x = SavedSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedSlice) ProtoMessage() {}

func (x *SavedSlice) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedSlice.ProtoReflect.Descriptor instead.
func (*SavedSlice) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_saved_tensor_slice_proto_rawDescGZIP(), []int{2}
}

func (x *SavedSlice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SavedSlice) GetSlice() *framework.TensorSliceProto {
	if x != nil {
		return x.Slice
	}
	return nil
}

func (x *SavedSlice) GetData() *framework.TensorProto {
	if x != nil {
		return x.Data
	}
	return nil
}

// Each record in a v3 checkpoint file is a serialized SavedTensorSlices
// message.
type SavedTensorSlices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is only present at the first item of each checkpoint file and serves
	// as a table of contents, listing all the tensor slices saved in this file.
	Meta *SavedTensorSliceMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// This exists in all but the first item of each checkpoint file.
	Data *SavedSlice `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SavedTensorSlices) Reset() {
	*x = SavedTensorSlices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedTensorSlices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedTensorSlices) ProtoMessage() {}

func (x *SavedTensorSlices) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedTensorSlices.ProtoReflect.Descriptor instead.
func (*SavedTensorSlices) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_util_saved_tensor_slice_proto_rawDescGZIP(), []int{3}
}

func (x *SavedTensorSlices) GetMeta() *SavedTensorSliceMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SavedTensorSlices) GetData() *SavedSlice {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tensorflow_core_util_saved_tensor_slice_proto protoreflect.FileDescriptor

var file_tensorflow_core_util_saved_tensor_slice_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2c, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x7e, 0x0a, 0x14, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x32, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x52,
	0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0a, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x6c, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x75, 0x0a,
	0x11, 0x53, 0x61, 0x76, 0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0xb2, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x15, 0x53, 0x61, 0x76, 0x65, 0x64, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72,
	0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_core_util_saved_tensor_slice_proto_rawDescOnce sync.Once
	file_tensorflow_core_util_saved_tensor_slice_proto_rawDescData = file_tensorflow_core_util_saved_tensor_slice_proto_rawDesc
)

func file_tensorflow_core_util_saved_tensor_slice_proto_rawDescGZIP() []byte {
	file_tensorflow_core_util_saved_tensor_slice_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_util_saved_tensor_slice_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_util_saved_tensor_slice_proto_rawDescData)
	})
	return file_tensorflow_core_util_saved_tensor_slice_proto_rawDescData
}

var file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_util_saved_tensor_slice_proto_goTypes = []interface{}{
	(*SavedSliceMeta)(nil),             // 0: tensorflow.SavedSliceMeta
	(*SavedTensorSliceMeta)(nil),       // 1: tensorflow.SavedTensorSliceMeta
	(*SavedSlice)(nil),                 // 2: tensorflow.SavedSlice
	(*SavedTensorSlices)(nil),          // 3: tensorflow.SavedTensorSlices
	(*framework.TensorShapeProto)(nil), // 4: tensorflow.TensorShapeProto
	(framework.DataType)(0),            // 5: tensorflow.DataType
	(*framework.TensorSliceProto)(nil), // 6: tensorflow.TensorSliceProto
	(*framework.VersionDef)(nil),       // 7: tensorflow.VersionDef
	(*framework.TensorProto)(nil),      // 8: tensorflow.TensorProto
}
var file_tensorflow_core_util_saved_tensor_slice_proto_depIdxs = []int32{
	4, // 0: tensorflow.SavedSliceMeta.shape:type_name -> tensorflow.TensorShapeProto
	5, // 1: tensorflow.SavedSliceMeta.type:type_name -> tensorflow.DataType
	6, // 2: tensorflow.SavedSliceMeta.slice:type_name -> tensorflow.TensorSliceProto
	0, // 3: tensorflow.SavedTensorSliceMeta.tensor:type_name -> tensorflow.SavedSliceMeta
	7, // 4: tensorflow.SavedTensorSliceMeta.versions:type_name -> tensorflow.VersionDef
	6, // 5: tensorflow.SavedSlice.slice:type_name -> tensorflow.TensorSliceProto
	8, // 6: tensorflow.SavedSlice.data:type_name -> tensorflow.TensorProto
	1, // 7: tensorflow.SavedTensorSlices.meta:type_name -> tensorflow.SavedTensorSliceMeta
	2, // 8: tensorflow.SavedTensorSlices.data:type_name -> tensorflow.SavedSlice
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_tensorflow_core_util_saved_tensor_slice_proto_init() }
func file_tensorflow_core_util_saved_tensor_slice_proto_init() {
	if File_tensorflow_core_util_saved_tensor_slice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedSliceMeta); i {
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
		file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedTensorSliceMeta); i {
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
		file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedSlice); i {
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
		file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedTensorSlices); i {
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
			RawDescriptor: file_tensorflow_core_util_saved_tensor_slice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_util_saved_tensor_slice_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_util_saved_tensor_slice_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_util_saved_tensor_slice_proto_msgTypes,
	}.Build()
	File_tensorflow_core_util_saved_tensor_slice_proto = out.File
	file_tensorflow_core_util_saved_tensor_slice_proto_rawDesc = nil
	file_tensorflow_core_util_saved_tensor_slice_proto_goTypes = nil
	file_tensorflow_core_util_saved_tensor_slice_proto_depIdxs = nil
}
