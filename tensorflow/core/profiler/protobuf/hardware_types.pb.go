// This proto describes the types of hardware profiled by the TensorFlow
// profiler.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/profiler/protobuf/hardware_types.proto

package protobuf

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

// Types of hardware profiled.
type HardwareType int32

const (
	// Unknown hardware.
	HardwareType_UNKNOWN_HARDWARE HardwareType = 0
	// CPU only without any hardware accelerator.
	HardwareType_CPU_ONLY HardwareType = 1
	// GPU.
	HardwareType_GPU HardwareType = 2
	// TPU.
	HardwareType_TPU HardwareType = 3
)

// Enum value maps for HardwareType.
var (
	HardwareType_name = map[int32]string{
		0: "UNKNOWN_HARDWARE",
		1: "CPU_ONLY",
		2: "GPU",
		3: "TPU",
	}
	HardwareType_value = map[string]int32{
		"UNKNOWN_HARDWARE": 0,
		"CPU_ONLY":         1,
		"GPU":              2,
		"TPU":              3,
	}
)

func (x HardwareType) Enum() *HardwareType {
	p := new(HardwareType)
	*p = x
	return p
}

func (x HardwareType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HardwareType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_profiler_protobuf_hardware_types_proto_enumTypes[0].Descriptor()
}

func (HardwareType) Type() protoreflect.EnumType {
	return &file_tensorflow_core_profiler_protobuf_hardware_types_proto_enumTypes[0]
}

func (x HardwareType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HardwareType.Descriptor instead.
func (HardwareType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescGZIP(), []int{0}
}

type GPUComputeCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
}

func (x *GPUComputeCapability) Reset() {
	*x = GPUComputeCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GPUComputeCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GPUComputeCapability) ProtoMessage() {}

func (x *GPUComputeCapability) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GPUComputeCapability.ProtoReflect.Descriptor instead.
func (*GPUComputeCapability) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescGZIP(), []int{0}
}

func (x *GPUComputeCapability) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *GPUComputeCapability) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

type DeviceCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClockRateInGhz    float64               `protobuf:"fixed64,1,opt,name=clock_rate_in_ghz,json=clockRateInGhz,proto3" json:"clock_rate_in_ghz,omitempty"`
	NumCores          uint32                `protobuf:"varint,2,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	MemorySizeInBytes uint64                `protobuf:"varint,3,opt,name=memory_size_in_bytes,json=memorySizeInBytes,proto3" json:"memory_size_in_bytes,omitempty"`
	MemoryBandwidth   uint64                `protobuf:"varint,4,opt,name=memory_bandwidth,json=memoryBandwidth,proto3" json:"memory_bandwidth,omitempty"` // Bytes/s.
	ComputeCapability *GPUComputeCapability `protobuf:"bytes,5,opt,name=compute_capability,json=computeCapability,proto3" json:"compute_capability,omitempty"`
	DeviceVendor      string                `protobuf:"bytes,6,opt,name=device_vendor,json=deviceVendor,proto3" json:"device_vendor,omitempty"`
}

func (x *DeviceCapabilities) Reset() {
	*x = DeviceCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCapabilities) ProtoMessage() {}

func (x *DeviceCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCapabilities.ProtoReflect.Descriptor instead.
func (*DeviceCapabilities) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceCapabilities) GetClockRateInGhz() float64 {
	if x != nil {
		return x.ClockRateInGhz
	}
	return 0
}

func (x *DeviceCapabilities) GetNumCores() uint32 {
	if x != nil {
		return x.NumCores
	}
	return 0
}

func (x *DeviceCapabilities) GetMemorySizeInBytes() uint64 {
	if x != nil {
		return x.MemorySizeInBytes
	}
	return 0
}

func (x *DeviceCapabilities) GetMemoryBandwidth() uint64 {
	if x != nil {
		return x.MemoryBandwidth
	}
	return 0
}

func (x *DeviceCapabilities) GetComputeCapability() *GPUComputeCapability {
	if x != nil {
		return x.ComputeCapability
	}
	return nil
}

func (x *DeviceCapabilities) GetDeviceVendor() string {
	if x != nil {
		return x.DeviceVendor
	}
	return ""
}

var File_tensorflow_core_profiler_protobuf_hardware_types_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDesc = []byte{
	0x0a, 0x36, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x42, 0x0a,
	0x14, 0x47, 0x50, 0x55, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x22, 0xb7, 0x02, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x47, 0x68, 0x7a, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x47,
	0x50, 0x55, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2a, 0x44, 0x0a, 0x0c, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x57, 0x41, 0x52, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x50, 0x55, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x50, 0x55, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x50, 0x55, 0x10,
	0x03, 0x42, 0xe7, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x12, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0xa2, 0x02, 0x03, 0x54, 0x50, 0x58, 0xaa, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xca, 0x02, 0x13, 0x54,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x72, 0xe2, 0x02, 0x1f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescOnce sync.Once
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescData = file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDesc
)

func file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescGZIP() []byte {
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescData)
	})
	return file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDescData
}

var file_tensorflow_core_profiler_protobuf_hardware_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_profiler_protobuf_hardware_types_proto_goTypes = []interface{}{
	(HardwareType)(0),            // 0: tensorflow.profiler.HardwareType
	(*GPUComputeCapability)(nil), // 1: tensorflow.profiler.GPUComputeCapability
	(*DeviceCapabilities)(nil),   // 2: tensorflow.profiler.DeviceCapabilities
}
var file_tensorflow_core_profiler_protobuf_hardware_types_proto_depIdxs = []int32{
	1, // 0: tensorflow.profiler.DeviceCapabilities.compute_capability:type_name -> tensorflow.profiler.GPUComputeCapability
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_hardware_types_proto_init() }
func file_tensorflow_core_profiler_protobuf_hardware_types_proto_init() {
	if File_tensorflow_core_profiler_protobuf_hardware_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GPUComputeCapability); i {
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
		file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceCapabilities); i {
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
			RawDescriptor: file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_hardware_types_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_hardware_types_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_profiler_protobuf_hardware_types_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_profiler_protobuf_hardware_types_proto_msgTypes,
	}.Build()
	File_tensorflow_core_profiler_protobuf_hardware_types_proto = out.File
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_hardware_types_proto_depIdxs = nil
}
