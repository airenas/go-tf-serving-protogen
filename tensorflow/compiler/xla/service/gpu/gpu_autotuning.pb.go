// This is used for convolution logging.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/service/gpu/gpu_autotuning.proto

package gpu

import (
	xla "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla"
	data "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/data"
	service "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla/service"
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

type ConvInstructionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instruction      *service.HloInstructionProto `protobuf:"bytes,1,opt,name=instruction,proto3" json:"instruction,omitempty"`
	OperandShapes    []*data.ShapeProto           `protobuf:"bytes,2,rep,name=operand_shapes,json=operandShapes,proto3" json:"operand_shapes,omitempty"`
	ResultAddress    uint64                       `protobuf:"varint,3,opt,name=result_address,json=resultAddress,proto3" json:"result_address,omitempty"`
	OperandAddresses []uint64                     `protobuf:"varint,4,rep,packed,name=operand_addresses,json=operandAddresses,proto3" json:"operand_addresses,omitempty"`
}

func (x *ConvInstructionLog) Reset() {
	*x = ConvInstructionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvInstructionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvInstructionLog) ProtoMessage() {}

func (x *ConvInstructionLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvInstructionLog.ProtoReflect.Descriptor instead.
func (*ConvInstructionLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescGZIP(), []int{0}
}

func (x *ConvInstructionLog) GetInstruction() *service.HloInstructionProto {
	if x != nil {
		return x.Instruction
	}
	return nil
}

func (x *ConvInstructionLog) GetOperandShapes() []*data.ShapeProto {
	if x != nil {
		return x.OperandShapes
	}
	return nil
}

func (x *ConvInstructionLog) GetResultAddress() uint64 {
	if x != nil {
		return x.ResultAddress
	}
	return 0
}

func (x *ConvInstructionLog) GetOperandAddresses() []uint64 {
	if x != nil {
		return x.OperandAddresses
	}
	return nil
}

type DenylistedAlgorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TensorOps bool  `protobuf:"varint,2,opt,name=tensor_ops,json=tensorOps,proto3" json:"tensor_ops,omitempty"`
}

func (x *DenylistedAlgorithm) Reset() {
	*x = DenylistedAlgorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenylistedAlgorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenylistedAlgorithm) ProtoMessage() {}

func (x *DenylistedAlgorithm) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenylistedAlgorithm.ProtoReflect.Descriptor instead.
func (*DenylistedAlgorithm) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescGZIP(), []int{1}
}

func (x *DenylistedAlgorithm) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DenylistedAlgorithm) GetTensorOps() bool {
	if x != nil {
		return x.TensorOps
	}
	return false
}

type AlgorithmDenylistEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hlo          string                 `protobuf:"bytes,1,opt,name=hlo,proto3" json:"hlo,omitempty"`
	Cc           *xla.ComputeCapability `protobuf:"bytes,2,opt,name=cc,proto3" json:"cc,omitempty"`
	CudnnVersion *xla.CudnnVersion      `protobuf:"bytes,3,opt,name=cudnn_version,json=cudnnVersion,proto3" json:"cudnn_version,omitempty"`
	BlasVersion  string                 `protobuf:"bytes,5,opt,name=blas_version,json=blasVersion,proto3" json:"blas_version,omitempty"`
	Algos        []*DenylistedAlgorithm `protobuf:"bytes,4,rep,name=algos,proto3" json:"algos,omitempty"`
}

func (x *AlgorithmDenylistEntry) Reset() {
	*x = AlgorithmDenylistEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmDenylistEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmDenylistEntry) ProtoMessage() {}

func (x *AlgorithmDenylistEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmDenylistEntry.ProtoReflect.Descriptor instead.
func (*AlgorithmDenylistEntry) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescGZIP(), []int{2}
}

func (x *AlgorithmDenylistEntry) GetHlo() string {
	if x != nil {
		return x.Hlo
	}
	return ""
}

func (x *AlgorithmDenylistEntry) GetCc() *xla.ComputeCapability {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *AlgorithmDenylistEntry) GetCudnnVersion() *xla.CudnnVersion {
	if x != nil {
		return x.CudnnVersion
	}
	return nil
}

func (x *AlgorithmDenylistEntry) GetBlasVersion() string {
	if x != nil {
		return x.BlasVersion
	}
	return ""
}

func (x *AlgorithmDenylistEntry) GetAlgos() []*DenylistedAlgorithm {
	if x != nil {
		return x.Algos
	}
	return nil
}

type AlgorithmDenylist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*AlgorithmDenylistEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *AlgorithmDenylist) Reset() {
	*x = AlgorithmDenylist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlgorithmDenylist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlgorithmDenylist) ProtoMessage() {}

func (x *AlgorithmDenylist) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlgorithmDenylist.ProtoReflect.Descriptor instead.
func (*AlgorithmDenylist) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescGZIP(), []int{3}
}

func (x *AlgorithmDenylist) GetEntries() []*AlgorithmDenylistEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDesc = []byte{
	0x0a, 0x38, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x67, 0x70, 0x75, 0x2f, 0x67, 0x70, 0x75, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x78, 0x6c, 0x61, 0x2e,
	0x67, 0x70, 0x75, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68,
	0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c,
	0x61, 0x2f, 0x78, 0x6c, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdc, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x76, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x3a, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x78,
	0x6c, 0x61, 0x2e, 0x48, 0x6c, 0x6f, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x78, 0x6c,
	0x61, 0x2e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x10, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22,
	0x44, 0x0a, 0x13, 0x44, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x4f, 0x70, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x16, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x44, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x68, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68,
	0x6c, 0x6f, 0x12, 0x26, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x02, 0x63, 0x63, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x75,
	0x64, 0x6e, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x43, 0x75, 0x64, 0x6e, 0x6e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x75, 0x64, 0x6e, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x61, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x61, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x6c, 0x67, 0x6f, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x44,
	0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x67, 0x6f, 0x73, 0x22, 0x4e, 0x0a, 0x11, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x44, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x67, 0x70, 0x75, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x44, 0x65, 0x6e, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0xad, 0x01, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x67, 0x70, 0x75, 0x42, 0x12, 0x47, 0x70, 0x75, 0x41, 0x75,
	0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65,
	0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78,
	0x6c, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x70, 0x75, 0xa2, 0x02,
	0x03, 0x58, 0x47, 0x58, 0xaa, 0x02, 0x07, 0x58, 0x6c, 0x61, 0x2e, 0x47, 0x70, 0x75, 0xca, 0x02,
	0x07, 0x58, 0x6c, 0x61, 0x5c, 0x47, 0x70, 0x75, 0xe2, 0x02, 0x13, 0x58, 0x6c, 0x61, 0x5c, 0x47,
	0x70, 0x75, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x58, 0x6c, 0x61, 0x3a, 0x3a, 0x47, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescData = file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDesc
)

func file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDescData
}

var file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_goTypes = []interface{}{
	(*ConvInstructionLog)(nil),          // 0: xla.gpu.ConvInstructionLog
	(*DenylistedAlgorithm)(nil),         // 1: xla.gpu.DenylistedAlgorithm
	(*AlgorithmDenylistEntry)(nil),      // 2: xla.gpu.AlgorithmDenylistEntry
	(*AlgorithmDenylist)(nil),           // 3: xla.gpu.AlgorithmDenylist
	(*service.HloInstructionProto)(nil), // 4: xla.HloInstructionProto
	(*data.ShapeProto)(nil),             // 5: xla.ShapeProto
	(*xla.ComputeCapability)(nil),       // 6: xla.ComputeCapability
	(*xla.CudnnVersion)(nil),            // 7: xla.CudnnVersion
}
var file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_depIdxs = []int32{
	4, // 0: xla.gpu.ConvInstructionLog.instruction:type_name -> xla.HloInstructionProto
	5, // 1: xla.gpu.ConvInstructionLog.operand_shapes:type_name -> xla.ShapeProto
	6, // 2: xla.gpu.AlgorithmDenylistEntry.cc:type_name -> xla.ComputeCapability
	7, // 3: xla.gpu.AlgorithmDenylistEntry.cudnn_version:type_name -> xla.CudnnVersion
	1, // 4: xla.gpu.AlgorithmDenylistEntry.algos:type_name -> xla.gpu.DenylistedAlgorithm
	2, // 5: xla.gpu.AlgorithmDenylist.entries:type_name -> xla.gpu.AlgorithmDenylistEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_init() }
func file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_init() {
	if File_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvInstructionLog); i {
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
		file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenylistedAlgorithm); i {
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
		file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmDenylistEntry); i {
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
		file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlgorithmDenylist); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto = out.File
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_rawDesc = nil
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_goTypes = nil
	file_tensorflow_compiler_xla_service_gpu_gpu_autotuning_proto_depIdxs = nil
}
