// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/tf2xla/support.proto

package tf2xla

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

type OpSupport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the op (in the OpDef).
	GraphOpName string `protobuf:"bytes,1,opt,name=graph_op_name,json=graphOpName,proto3" json:"graph_op_name,omitempty"`
	// Is this TF op supported by XLA compilation?
	// i.e. JIT compiled
	SupportsXlaCompilation bool `protobuf:"varint,7,opt,name=supports_xla_compilation,json=supportsXlaCompilation,proto3" json:"supports_xla_compilation,omitempty"`
	// Is this TF op supported by the old bridge?
	SupportsOldBridge bool `protobuf:"varint,2,opt,name=supports_old_bridge,json=supportsOldBridge,proto3" json:"supports_old_bridge,omitempty"`
	// Is this TF op supported by the new bridge?
	SupportsNewBridge bool `protobuf:"varint,3,opt,name=supports_new_bridge,json=supportsNewBridge,proto3" json:"supports_new_bridge,omitempty"`
	// Does this TF op support bounded, dynamic, ranked shapes?
	// Note: Unavailable in MLIR
	SupportsBoundedDynamicRanked bool `protobuf:"varint,4,opt,name=supports_bounded_dynamic_ranked,json=supportsBoundedDynamicRanked,proto3" json:"supports_bounded_dynamic_ranked,omitempty"`
	// Does this TF op support unbounded, dynamic, ranked shapes?
	// i.e. tensor<?x3xf32>
	SupportsDynamicRanked bool `protobuf:"varint,5,opt,name=supports_dynamic_ranked,json=supportsDynamicRanked,proto3" json:"supports_dynamic_ranked,omitempty"`
	// Does this TF op support unbounded, dynamic, unranked shapes?
	// i.e. tensor<*xf32>
	SupportsDynamicUnranked bool `protobuf:"varint,6,opt,name=supports_dynamic_unranked,json=supportsDynamicUnranked,proto3" json:"supports_dynamic_unranked,omitempty"`
}

func (x *OpSupport) Reset() {
	*x = OpSupport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_support_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpSupport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpSupport) ProtoMessage() {}

func (x *OpSupport) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_support_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpSupport.ProtoReflect.Descriptor instead.
func (*OpSupport) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_support_proto_rawDescGZIP(), []int{0}
}

func (x *OpSupport) GetGraphOpName() string {
	if x != nil {
		return x.GraphOpName
	}
	return ""
}

func (x *OpSupport) GetSupportsXlaCompilation() bool {
	if x != nil {
		return x.SupportsXlaCompilation
	}
	return false
}

func (x *OpSupport) GetSupportsOldBridge() bool {
	if x != nil {
		return x.SupportsOldBridge
	}
	return false
}

func (x *OpSupport) GetSupportsNewBridge() bool {
	if x != nil {
		return x.SupportsNewBridge
	}
	return false
}

func (x *OpSupport) GetSupportsBoundedDynamicRanked() bool {
	if x != nil {
		return x.SupportsBoundedDynamicRanked
	}
	return false
}

func (x *OpSupport) GetSupportsDynamicRanked() bool {
	if x != nil {
		return x.SupportsDynamicRanked
	}
	return false
}

func (x *OpSupport) GetSupportsDynamicUnranked() bool {
	if x != nil {
		return x.SupportsDynamicUnranked
	}
	return false
}

type OpSupports struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Support []*OpSupport `protobuf:"bytes,1,rep,name=support,proto3" json:"support,omitempty"`
}

func (x *OpSupports) Reset() {
	*x = OpSupports{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_tf2xla_support_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpSupports) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpSupports) ProtoMessage() {}

func (x *OpSupports) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_tf2xla_support_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpSupports.ProtoReflect.Descriptor instead.
func (*OpSupports) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_tf2xla_support_proto_rawDescGZIP(), []int{1}
}

func (x *OpSupports) GetSupport() []*OpSupport {
	if x != nil {
		return x.Support
	}
	return nil
}

var File_tensorflow_compiler_tf2xla_support_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_tf2xla_support_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x22, 0x84, 0x03,
	0x0a, 0x09, 0x4f, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x18, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x78, 0x6c, 0x61, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x16, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x58, 0x6c, 0x61, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6f, 0x6c, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4f, 0x6c, 0x64, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4e, 0x65, 0x77, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x1f, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1c, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x65, 0x64, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x6b, 0x65, 0x64,
	0x12, 0x36, 0x0a, 0x17, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x52, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x55, 0x6e, 0x72, 0x61,
	0x6e, 0x6b, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x0a, 0x4f, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x2e, 0x4f, 0x70, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x07, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x42, 0xd3, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x66,
	0x32, 0x78, 0x6c, 0x61, 0x42, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x72, 0x2f, 0x74, 0x66, 0x32, 0x78, 0x6c, 0x61, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x54, 0x54, 0x58, 0xaa, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x54, 0x66, 0x32, 0x78, 0x6c, 0x61, 0xca, 0x02, 0x11, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x54, 0x66, 0x32, 0x78, 0x6c, 0x61, 0xe2, 0x02, 0x1d, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x54, 0x66, 0x32, 0x78, 0x6c, 0x61, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x54, 0x66, 0x32, 0x78, 0x6c, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_tf2xla_support_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_tf2xla_support_proto_rawDescData = file_tensorflow_compiler_tf2xla_support_proto_rawDesc
)

func file_tensorflow_compiler_tf2xla_support_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_tf2xla_support_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_tf2xla_support_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_tf2xla_support_proto_rawDescData)
	})
	return file_tensorflow_compiler_tf2xla_support_proto_rawDescData
}

var file_tensorflow_compiler_tf2xla_support_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_compiler_tf2xla_support_proto_goTypes = []interface{}{
	(*OpSupport)(nil),  // 0: tensorflow.tf2xla.OpSupport
	(*OpSupports)(nil), // 1: tensorflow.tf2xla.OpSupports
}
var file_tensorflow_compiler_tf2xla_support_proto_depIdxs = []int32{
	0, // 0: tensorflow.tf2xla.OpSupports.support:type_name -> tensorflow.tf2xla.OpSupport
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_tf2xla_support_proto_init() }
func file_tensorflow_compiler_tf2xla_support_proto_init() {
	if File_tensorflow_compiler_tf2xla_support_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_tf2xla_support_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpSupport); i {
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
		file_tensorflow_compiler_tf2xla_support_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpSupports); i {
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
			RawDescriptor: file_tensorflow_compiler_tf2xla_support_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_tf2xla_support_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_tf2xla_support_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_tf2xla_support_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_tf2xla_support_proto = out.File
	file_tensorflow_compiler_tf2xla_support_proto_rawDesc = nil
	file_tensorflow_compiler_tf2xla_support_proto_goTypes = nil
	file_tensorflow_compiler_tf2xla_support_proto_depIdxs = nil
}
