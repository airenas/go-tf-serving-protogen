// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow_serving/servables/tensorflow/saved_model_bundle_source_adapter.proto

package tensorflow

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

// Config proto for SavedModelBundleSourceAdapter.
type SavedModelBundleSourceAdapterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A SessionBundleConfig.
	// FOR INTERNAL USE ONLY DURING TRANSITION TO SAVED_MODEL. WILL BE DEPRECATED.
	// TODO(b/32248363): Replace this field with the "real" field(s).
	LegacyConfig *SessionBundleConfig `protobuf:"bytes,1000,opt,name=legacy_config,json=legacyConfig,proto3" json:"legacy_config,omitempty"`
}

func (x *SavedModelBundleSourceAdapterConfig) Reset() {
	*x = SavedModelBundleSourceAdapterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedModelBundleSourceAdapterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedModelBundleSourceAdapterConfig) ProtoMessage() {}

func (x *SavedModelBundleSourceAdapterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedModelBundleSourceAdapterConfig.ProtoReflect.Descriptor instead.
func (*SavedModelBundleSourceAdapterConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *SavedModelBundleSourceAdapterConfig) GetLegacyConfig() *SessionBundleConfig {
	if x != nil {
		return x.LegacyConfig
	}
	return nil
}

var File_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto protoreflect.FileDescriptor

var file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x43, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x23, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x4d, 0x0a, 0x0d, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0c, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x42, 0xf8, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x22, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69,
	0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca,
	0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescOnce sync.Once
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescData = file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDesc
)

func file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescData)
	})
	return file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDescData
}

var file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_goTypes = []interface{}{
	(*SavedModelBundleSourceAdapterConfig)(nil), // 0: tensorflow.serving.SavedModelBundleSourceAdapterConfig
	(*SessionBundleConfig)(nil),                 // 1: tensorflow.serving.SessionBundleConfig
}
var file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_depIdxs = []int32{
	1, // 0: tensorflow.serving.SavedModelBundleSourceAdapterConfig.legacy_config:type_name -> tensorflow.serving.SessionBundleConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_init()
}
func file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_init() {
	if File_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto != nil {
		return
	}
	file_tensorflow_serving_servables_tensorflow_session_bundle_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedModelBundleSourceAdapterConfig); i {
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
			RawDescriptor: file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto = out.File
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_rawDesc = nil
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_goTypes = nil
	file_tensorflow_serving_servables_tensorflow_saved_model_bundle_source_adapter_proto_depIdxs = nil
}
