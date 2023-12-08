// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/mlir/tfrt/analysis/analysis.proto

package analysis

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

type CompatibilityAnalysisReportProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnknownDialect        bool   `protobuf:"varint,1,opt,name=unknown_dialect,json=unknownDialect,proto3" json:"unknown_dialect,omitempty"`
	RefVariable           bool   `protobuf:"varint,2,opt,name=ref_variable,json=refVariable,proto3" json:"ref_variable,omitempty"`
	IncompatibleVariable  bool   `protobuf:"varint,3,opt,name=incompatible_variable,json=incompatibleVariable,proto3" json:"incompatible_variable,omitempty"`
	IncompatibleAttribute bool   `protobuf:"varint,4,opt,name=incompatible_attribute,json=incompatibleAttribute,proto3" json:"incompatible_attribute,omitempty"`
	ControlFlowV1         bool   `protobuf:"varint,5,opt,name=control_flow_v1,json=controlFlowV1,proto3" json:"control_flow_v1,omitempty"`
	MethodName            string `protobuf:"bytes,6,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (x *CompatibilityAnalysisReportProto) Reset() {
	*x = CompatibilityAnalysisReportProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompatibilityAnalysisReportProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompatibilityAnalysisReportProto) ProtoMessage() {}

func (x *CompatibilityAnalysisReportProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompatibilityAnalysisReportProto.ProtoReflect.Descriptor instead.
func (*CompatibilityAnalysisReportProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *CompatibilityAnalysisReportProto) GetUnknownDialect() bool {
	if x != nil {
		return x.UnknownDialect
	}
	return false
}

func (x *CompatibilityAnalysisReportProto) GetRefVariable() bool {
	if x != nil {
		return x.RefVariable
	}
	return false
}

func (x *CompatibilityAnalysisReportProto) GetIncompatibleVariable() bool {
	if x != nil {
		return x.IncompatibleVariable
	}
	return false
}

func (x *CompatibilityAnalysisReportProto) GetIncompatibleAttribute() bool {
	if x != nil {
		return x.IncompatibleAttribute
	}
	return false
}

func (x *CompatibilityAnalysisReportProto) GetControlFlowV1() bool {
	if x != nil {
		return x.ControlFlowV1
	}
	return false
}

func (x *CompatibilityAnalysisReportProto) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

type CompatibilityAnalysisProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary *CompatibilityAnalysisReportProto             `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Ops     map[string]*CompatibilityAnalysisProto_OpInfo `protobuf:"bytes,2,rep,name=ops,proto3" json:"ops,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CompatibilityAnalysisProto) Reset() {
	*x = CompatibilityAnalysisProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompatibilityAnalysisProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompatibilityAnalysisProto) ProtoMessage() {}

func (x *CompatibilityAnalysisProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompatibilityAnalysisProto.ProtoReflect.Descriptor instead.
func (*CompatibilityAnalysisProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *CompatibilityAnalysisProto) GetSummary() *CompatibilityAnalysisReportProto {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *CompatibilityAnalysisProto) GetOps() map[string]*CompatibilityAnalysisProto_OpInfo {
	if x != nil {
		return x.Ops
	}
	return nil
}

type CompatibilityAnalysisProto_OpInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int32                             `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Report *CompatibilityAnalysisReportProto `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *CompatibilityAnalysisProto_OpInfo) Reset() {
	*x = CompatibilityAnalysisProto_OpInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompatibilityAnalysisProto_OpInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompatibilityAnalysisProto_OpInfo) ProtoMessage() {}

func (x *CompatibilityAnalysisProto_OpInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompatibilityAnalysisProto_OpInfo.ProtoReflect.Descriptor instead.
func (*CompatibilityAnalysisProto_OpInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CompatibilityAnalysisProto_OpInfo) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CompatibilityAnalysisProto_OpInfo) GetReport() *CompatibilityAnalysisReportProto {
	if x != nil {
		return x.Report
	}
	return nil
}

var File_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x74, 0x66, 0x72, 0x74, 0x2f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x74, 0x66,
	0x72, 0x74, 0x22, 0xa3, 0x02, 0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x76, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x56, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf0, 0x02, 0x0a, 0x1a, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x45, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e,
	0x74, 0x66, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x40,
	0x0a, 0x03, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x6c,
	0x69, 0x72, 0x2e, 0x74, 0x66, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6f, 0x70, 0x73,
	0x1a, 0x63, 0x0a, 0x06, 0x4f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x74, 0x66, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x64, 0x0a, 0x08, 0x4f, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x74, 0x66, 0x72, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xb5, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x74, 0x66, 0x72, 0x74, 0x42, 0x0d, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e,
	0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c,
	0x69, 0x72, 0x2f, 0x74, 0x66, 0x72, 0x74, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0xa2, 0x02, 0x03, 0x4d, 0x54, 0x58, 0xaa, 0x02, 0x09, 0x4d, 0x6c, 0x69, 0x72, 0x2e, 0x54, 0x66,
	0x72, 0x74, 0xca, 0x02, 0x09, 0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x54, 0x66, 0x72, 0x74, 0xe2, 0x02,
	0x15, 0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x54, 0x66, 0x72, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x6c, 0x69, 0x72, 0x3a, 0x3a, 0x54,
	0x66, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescData = file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDesc
)

func file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescData)
	})
	return file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDescData
}

var file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_goTypes = []interface{}{
	(*CompatibilityAnalysisReportProto)(nil),  // 0: mlir.tfrt.CompatibilityAnalysisReportProto
	(*CompatibilityAnalysisProto)(nil),        // 1: mlir.tfrt.CompatibilityAnalysisProto
	(*CompatibilityAnalysisProto_OpInfo)(nil), // 2: mlir.tfrt.CompatibilityAnalysisProto.OpInfo
	nil, // 3: mlir.tfrt.CompatibilityAnalysisProto.OpsEntry
}
var file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_depIdxs = []int32{
	0, // 0: mlir.tfrt.CompatibilityAnalysisProto.summary:type_name -> mlir.tfrt.CompatibilityAnalysisReportProto
	3, // 1: mlir.tfrt.CompatibilityAnalysisProto.ops:type_name -> mlir.tfrt.CompatibilityAnalysisProto.OpsEntry
	0, // 2: mlir.tfrt.CompatibilityAnalysisProto.OpInfo.report:type_name -> mlir.tfrt.CompatibilityAnalysisReportProto
	2, // 3: mlir.tfrt.CompatibilityAnalysisProto.OpsEntry.value:type_name -> mlir.tfrt.CompatibilityAnalysisProto.OpInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_init() }
func file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_init() {
	if File_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompatibilityAnalysisReportProto); i {
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
		file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompatibilityAnalysisProto); i {
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
		file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompatibilityAnalysisProto_OpInfo); i {
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
			RawDescriptor: file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto = out.File
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_rawDesc = nil
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_goTypes = nil
	file_tensorflow_compiler_mlir_tfrt_analysis_analysis_proto_depIdxs = nil
}
