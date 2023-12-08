// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/tsl/distributed_runtime/coordination/test_device.proto

package coordination

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

type TestDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LocalId  int64  `protobuf:"varint,2,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	GlobalId int64  `protobuf:"varint,3,opt,name=global_id,json=globalId,proto3" json:"global_id,omitempty"`
}

func (x *TestDevice) Reset() {
	*x = TestDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDevice) ProtoMessage() {}

func (x *TestDevice) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDevice.ProtoReflect.Descriptor instead.
func (*TestDevice) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescGZIP(), []int{0}
}

func (x *TestDevice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestDevice) GetLocalId() int64 {
	if x != nil {
		return x.LocalId
	}
	return 0
}

func (x *TestDevice) GetGlobalId() int64 {
	if x != nil {
		return x.GlobalId
	}
	return 0
}

type TestDeviceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device []*TestDevice `protobuf:"bytes,1,rep,name=device,proto3" json:"device,omitempty"`
}

func (x *TestDeviceList) Reset() {
	*x = TestDeviceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDeviceList) ProtoMessage() {}

func (x *TestDeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDeviceList.ProtoReflect.Descriptor instead.
func (*TestDeviceList) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescGZIP(), []int{1}
}

func (x *TestDeviceList) GetDevice() []*TestDevice {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_tensorflow_tsl_distributed_runtime_coordination_test_device_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDesc = []byte{
	0x0a, 0x41, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22,
	0x58, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x0e, 0x54, 0x65, 0x73,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0xc4, 0x01, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0f,
	0x54, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69,
	0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x54,
	0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca,
	0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescData = file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDesc
)

func file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescData)
	})
	return file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDescData
}

var file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_goTypes = []interface{}{
	(*TestDevice)(nil),     // 0: tensorflow.TestDevice
	(*TestDeviceList)(nil), // 1: tensorflow.TestDeviceList
}
var file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_depIdxs = []int32{
	0, // 0: tensorflow.TestDeviceList.device:type_name -> tensorflow.TestDevice
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_init() }
func file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_init() {
	if File_tensorflow_tsl_distributed_runtime_coordination_test_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDevice); i {
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
		file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDeviceList); i {
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
			RawDescriptor: file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_distributed_runtime_coordination_test_device_proto = out.File
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_rawDesc = nil
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_goTypes = nil
	file_tensorflow_tsl_distributed_runtime_coordination_test_device_proto_depIdxs = nil
}
