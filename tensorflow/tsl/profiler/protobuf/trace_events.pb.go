// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/tsl/profiler/protobuf/trace_events.proto

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

// A 'Trace' contains metadata for the individual traces of a system.
type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The devices that this trace has information about. Maps from device_id to
	// more data about the specific device.
	Devices map[uint32]*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// All trace events capturing in the profiling period.
	TraceEvents []*TraceEvent `protobuf:"bytes,4,rep,name=trace_events,json=traceEvents,proto3" json:"trace_events,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescGZIP(), []int{0}
}

func (x *Trace) GetDevices() map[uint32]*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *Trace) GetTraceEvents() []*TraceEvent {
	if x != nil {
		return x.TraceEvents
	}
	return nil
}

// A 'device' is a physical entity in the system and is comprised of several
// resources.
type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the device.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id of this device, unique in a single trace.
	DeviceId uint32 `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// The resources on this device, keyed by resource_id;
	Resources map[uint32]*Resource `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Device) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *Device) GetResources() map[uint32]*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// A 'resource' generally is a specific computation component on a device. These
// can range from threads on CPUs to specific arithmetic units on hardware
// devices.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id of the resource. Unique within a device.
	ResourceId uint32 `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The sort index of the resource. Resources within a device are ordered by
	// this value. if absent, use resource id as sort index.
	SortIndex uint32 `protobuf:"varint,3,opt,name=sort_index,json=sortIndex,proto3" json:"sort_index,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *Resource) GetSortIndex() uint32 {
	if x != nil {
		return x.SortIndex
	}
	return 0
}

type TraceEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the device that this event occurred on. The full dataset should
	// have this device present in the Trace object.
	DeviceId uint32 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// The id of the resource that this event occurred on. The full dataset should
	// have this resource present in the Device object of the Trace object. A
	// resource_id is unique on a specific device, but not necessarily within the
	// trace.
	ResourceId uint32 `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The name of this trace event.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp that this event occurred at (in picos since tracing started).
	TimestampPs uint64 `protobuf:"varint,9,opt,name=timestamp_ps,json=timestampPs,proto3" json:"timestamp_ps,omitempty"`
	// The duration of the event in picoseconds if applicable.
	// Events without duration are called instant events.
	DurationPs uint64 `protobuf:"varint,10,opt,name=duration_ps,json=durationPs,proto3" json:"duration_ps,omitempty"`
	// Extra arguments that will be displayed in trace view.
	Args map[string]string `protobuf:"bytes,11,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TraceEvent) Reset() {
	*x = TraceEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceEvent) ProtoMessage() {}

func (x *TraceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceEvent.ProtoReflect.Descriptor instead.
func (*TraceEvent) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescGZIP(), []int{3}
}

func (x *TraceEvent) GetDeviceId() uint32 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *TraceEvent) GetResourceId() uint32 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *TraceEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TraceEvent) GetTimestampPs() uint64 {
	if x != nil {
		return x.TimestampPs
	}
	return 0
}

func (x *TraceEvent) GetDurationPs() uint64 {
	if x != nil {
		return x.DurationPs
	}
	return 0
}

func (x *TraceEvent) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_tensorflow_tsl_profiler_protobuf_trace_events_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x22, 0xd2, 0x01, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x54, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x73, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x93, 0x02,
	0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x70, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x50,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x41, 0x72,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0xc4, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x73, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x42, 0x10, 0x54, 0x72, 0x61, 0x63, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x65, 0x6e, 0x61, 0x73,
	0x2f, 0x67, 0x6f, 0x2d, 0x74, 0x66, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x50,
	0x58, 0xaa, 0x02, 0x0c, 0x54, 0x73, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0xca, 0x02, 0x0c, 0x54, 0x73, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0xe2,
	0x02, 0x18, 0x54, 0x73, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x54, 0x73, 0x6c,
	0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_tsl_profiler_protobuf_trace_events_proto_goTypes = []interface{}{
	(*Trace)(nil),      // 0: tsl.profiler.Trace
	(*Device)(nil),     // 1: tsl.profiler.Device
	(*Resource)(nil),   // 2: tsl.profiler.Resource
	(*TraceEvent)(nil), // 3: tsl.profiler.TraceEvent
	nil,                // 4: tsl.profiler.Trace.DevicesEntry
	nil,                // 5: tsl.profiler.Device.ResourcesEntry
	nil,                // 6: tsl.profiler.TraceEvent.ArgsEntry
}
var file_tensorflow_tsl_profiler_protobuf_trace_events_proto_depIdxs = []int32{
	4, // 0: tsl.profiler.Trace.devices:type_name -> tsl.profiler.Trace.DevicesEntry
	3, // 1: tsl.profiler.Trace.trace_events:type_name -> tsl.profiler.TraceEvent
	5, // 2: tsl.profiler.Device.resources:type_name -> tsl.profiler.Device.ResourcesEntry
	6, // 3: tsl.profiler.TraceEvent.args:type_name -> tsl.profiler.TraceEvent.ArgsEntry
	1, // 4: tsl.profiler.Trace.DevicesEntry.value:type_name -> tsl.profiler.Device
	2, // 5: tsl.profiler.Device.ResourcesEntry.value:type_name -> tsl.profiler.Resource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_trace_events_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_trace_events_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_trace_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
		file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceEvent); i {
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
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_trace_events_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_trace_events_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_trace_events_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_trace_events_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_trace_events_proto_depIdxs = nil
}
