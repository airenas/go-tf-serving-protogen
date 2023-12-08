// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/profiler/protobuf/steps_db.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Breakdown of step-time on generic hardware. Note that these components are
// mutually exclusive so that adding them together is equal to the step time. If
// an execution time interval has multiple types of event happening, we need to
// pick one of the event type to attribute the time interval to.
type GenericStepBreakdown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map event type to the accumulated duration in
	// picoseconds of that type.
	TypePs map[int32]uint64 `protobuf:"bytes,1,rep,name=type_ps,json=typePs,proto3" json:"type_ps,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GenericStepBreakdown) Reset() {
	*x = GenericStepBreakdown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericStepBreakdown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericStepBreakdown) ProtoMessage() {}

func (x *GenericStepBreakdown) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericStepBreakdown.ProtoReflect.Descriptor instead.
func (*GenericStepBreakdown) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{0}
}

func (x *GenericStepBreakdown) GetTypePs() map[int32]uint64 {
	if x != nil {
		return x.TypePs
	}
	return nil
}

// Information about memory transfer to/from device memory.
type DeviceMemoryTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Occurrence       uint64  `protobuf:"varint,1,opt,name=occurrence,proto3" json:"occurrence,omitempty"`
	TimeUs           float64 `protobuf:"fixed64,2,opt,name=time_us,json=timeUs,proto3" json:"time_us,omitempty"`
	BytesTransferred uint64  `protobuf:"varint,3,opt,name=bytes_transferred,json=bytesTransferred,proto3" json:"bytes_transferred,omitempty"`
}

func (x *DeviceMemoryTransfer) Reset() {
	*x = DeviceMemoryTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceMemoryTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceMemoryTransfer) ProtoMessage() {}

func (x *DeviceMemoryTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceMemoryTransfer.ProtoReflect.Descriptor instead.
func (*DeviceMemoryTransfer) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceMemoryTransfer) GetOccurrence() uint64 {
	if x != nil {
		return x.Occurrence
	}
	return 0
}

func (x *DeviceMemoryTransfer) GetTimeUs() float64 {
	if x != nil {
		return x.TimeUs
	}
	return 0
}

func (x *DeviceMemoryTransfer) GetBytesTransferred() uint64 {
	if x != nil {
		return x.BytesTransferred
	}
	return 0
}

// Next ID: 6
// Result proto for StepInfo.
type StepInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The step number.
	StepNum uint32 `protobuf:"varint,1,opt,name=step_num,json=stepNum,proto3" json:"step_num,omitempty"`
	// The step name.
	StepName string `protobuf:"bytes,5,opt,name=step_name,json=stepName,proto3" json:"step_name,omitempty"`
	// The step duration in picoseconds.
	DurationPs uint64 `protobuf:"varint,2,opt,name=duration_ps,json=durationPs,proto3" json:"duration_ps,omitempty"`
	// The start time of this step in picoseconds.
	BeginPs uint64 `protobuf:"varint,3,opt,name=begin_ps,json=beginPs,proto3" json:"begin_ps,omitempty"`
	// Breakdown of the step-time. Can be unpacked into a GenericStepBreakdown.
	StepBreakdown *anypb.Any `protobuf:"bytes,4,opt,name=step_breakdown,json=stepBreakdown,proto3" json:"step_breakdown,omitempty"`
}

func (x *StepInfoResult) Reset() {
	*x = StepInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepInfoResult) ProtoMessage() {}

func (x *StepInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepInfoResult.ProtoReflect.Descriptor instead.
func (*StepInfoResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{2}
}

func (x *StepInfoResult) GetStepNum() uint32 {
	if x != nil {
		return x.StepNum
	}
	return 0
}

func (x *StepInfoResult) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (x *StepInfoResult) GetDurationPs() uint64 {
	if x != nil {
		return x.DurationPs
	}
	return 0
}

func (x *StepInfoResult) GetBeginPs() uint64 {
	if x != nil {
		return x.BeginPs
	}
	return 0
}

func (x *StepInfoResult) GetStepBreakdown() *anypb.Any {
	if x != nil {
		return x.StepBreakdown
	}
	return nil
}

// Result proto for all -educe ops.
type AllReduceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique id for all-reduce ops.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the hlo op. This field is no longer set by the profiler.
	//
	// Deprecated: Marked as deprecated in tensorflow/core/profiler/protobuf/steps_db.proto.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// For all-reduce nodes from different modules, if they have the same
	// all_reduce_id, they will be 'Allreduce'd'. If empty, AllReduce will not be
	// applied across modules.
	AllReduceId uint64 `protobuf:"varint,3,opt,name=all_reduce_id,json=allReduceId,proto3" json:"all_reduce_id,omitempty"`
	// The start time in picoseconds of the op event.
	StartTimePs uint64 `protobuf:"varint,4,opt,name=start_time_ps,json=startTimePs,proto3" json:"start_time_ps,omitempty"`
	// The end time in picoseconds of the op event.
	EndTimePs uint64 `protobuf:"varint,5,opt,name=end_time_ps,json=endTimePs,proto3" json:"end_time_ps,omitempty"`
	// The size of the op in bytes.
	ByteSize uint64 `protobuf:"varint,6,opt,name=byte_size,json=byteSize,proto3" json:"byte_size,omitempty"`
}

func (x *AllReduceInfo) Reset() {
	*x = AllReduceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReduceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReduceInfo) ProtoMessage() {}

func (x *AllReduceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReduceInfo.ProtoReflect.Descriptor instead.
func (*AllReduceInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{3}
}

func (x *AllReduceInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Deprecated: Marked as deprecated in tensorflow/core/profiler/protobuf/steps_db.proto.
func (x *AllReduceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AllReduceInfo) GetAllReduceId() uint64 {
	if x != nil {
		return x.AllReduceId
	}
	return 0
}

func (x *AllReduceInfo) GetStartTimePs() uint64 {
	if x != nil {
		return x.StartTimePs
	}
	return 0
}

func (x *AllReduceInfo) GetEndTimePs() uint64 {
	if x != nil {
		return x.EndTimePs
	}
	return 0
}

func (x *AllReduceInfo) GetByteSize() uint64 {
	if x != nil {
		return x.ByteSize
	}
	return 0
}

// Result database for all-reduce ops.
type AllReduceDbResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllReduceInfo []*AllReduceInfo `protobuf:"bytes,1,rep,name=all_reduce_info,json=allReduceInfo,proto3" json:"all_reduce_info,omitempty"`
}

func (x *AllReduceDbResult) Reset() {
	*x = AllReduceDbResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReduceDbResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReduceDbResult) ProtoMessage() {}

func (x *AllReduceDbResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReduceDbResult.ProtoReflect.Descriptor instead.
func (*AllReduceDbResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{4}
}

func (x *AllReduceDbResult) GetAllReduceInfo() []*AllReduceInfo {
	if x != nil {
		return x.AllReduceInfo
	}
	return nil
}

// Result proto for information in a step across all cores.
type PerCoreStepInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The step number.
	StepNum uint32 `protobuf:"varint,1,opt,name=step_num,json=stepNum,proto3" json:"step_num,omitempty"`
	// A map from core_id to StepInfo.
	StepInfoPerCore map[uint32]*StepInfoResult `protobuf:"bytes,2,rep,name=step_info_per_core,json=stepInfoPerCore,proto3" json:"step_info_per_core,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The result for the per-step HLO-metric database.
	HloMetricsDb *OpMetricsDb `protobuf:"bytes,3,opt,name=hlo_metrics_db,json=hloMetricsDb,proto3" json:"hlo_metrics_db,omitempty"`
	// A map from core ID to program replica id. Replica id map could change
	// during a profile session, but should stay stable within a step.
	CoreIdToReplicaIdMap map[uint32]uint32 `protobuf:"bytes,5,rep,name=core_id_to_replica_id_map,json=coreIdToReplicaIdMap,proto3" json:"core_id_to_replica_id_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// A map from core_id to all-reduce ops.
	AllReduceDbPerCore map[uint32]*AllReduceDbResult `protobuf:"bytes,6,rep,name=all_reduce_db_per_core,json=allReduceDbPerCore,proto3" json:"all_reduce_db_per_core,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Information about deivce memory transfers, categoried by source and
	// destination. Ordered by following categories:
	// 1. HostToDevice
	// 2. DeviceToHost
	// 3. DeviceToDevice
	DeviceMemoryTransfers []*DeviceMemoryTransfer `protobuf:"bytes,7,rep,name=device_memory_transfers,json=deviceMemoryTransfers,proto3" json:"device_memory_transfers,omitempty"`
}

func (x *PerCoreStepInfo) Reset() {
	*x = PerCoreStepInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerCoreStepInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerCoreStepInfo) ProtoMessage() {}

func (x *PerCoreStepInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerCoreStepInfo.ProtoReflect.Descriptor instead.
func (*PerCoreStepInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{5}
}

func (x *PerCoreStepInfo) GetStepNum() uint32 {
	if x != nil {
		return x.StepNum
	}
	return 0
}

func (x *PerCoreStepInfo) GetStepInfoPerCore() map[uint32]*StepInfoResult {
	if x != nil {
		return x.StepInfoPerCore
	}
	return nil
}

func (x *PerCoreStepInfo) GetHloMetricsDb() *OpMetricsDb {
	if x != nil {
		return x.HloMetricsDb
	}
	return nil
}

func (x *PerCoreStepInfo) GetCoreIdToReplicaIdMap() map[uint32]uint32 {
	if x != nil {
		return x.CoreIdToReplicaIdMap
	}
	return nil
}

func (x *PerCoreStepInfo) GetAllReduceDbPerCore() map[uint32]*AllReduceDbResult {
	if x != nil {
		return x.AllReduceDbPerCore
	}
	return nil
}

func (x *PerCoreStepInfo) GetDeviceMemoryTransfers() []*DeviceMemoryTransfer {
	if x != nil {
		return x.DeviceMemoryTransfers
	}
	return nil
}

// Result proto for a StepDatabase.
type StepDatabaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A sequence of PerCoreStepInfo.
	StepSequence []*PerCoreStepInfo `protobuf:"bytes,1,rep,name=step_sequence,json=stepSequence,proto3" json:"step_sequence,omitempty"`
	// Whether the step db uses incomplete step information.
	// This flag is set to true when:
	// 1) no step marker or annotation present.
	// 2) profiling duration is too short to cover a full step.
	// If this flag is false, we will group and breakdown the
	// profile by complete steps only and ignore incomplete steps.
	// If this flag is true, we will simply aggregate and breakdown over the total
	// profile as a single step.
	UseIncompleteStep bool `protobuf:"varint,2,opt,name=use_incomplete_step,json=useIncompleteStep,proto3" json:"use_incomplete_step,omitempty"`
	// Number of steps dropped during post processing.
	NumStepsDropped uint32 `protobuf:"varint,3,opt,name=num_steps_dropped,json=numStepsDropped,proto3" json:"num_steps_dropped,omitempty"`
	// If the step_sequence is empty because:
	//   - there is no step profiled on any host, then empty_intersect is false.
	//   - there are steps profiled on some host, but the intersection of steps
	//     over all hosts is empty, then empty_intersect is true.
	EmptyIntersect bool `protobuf:"varint,4,opt,name=empty_intersect,json=emptyIntersect,proto3" json:"empty_intersect,omitempty"`
}

func (x *StepDatabaseResult) Reset() {
	*x = StepDatabaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StepDatabaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StepDatabaseResult) ProtoMessage() {}

func (x *StepDatabaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StepDatabaseResult.ProtoReflect.Descriptor instead.
func (*StepDatabaseResult) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP(), []int{6}
}

func (x *StepDatabaseResult) GetStepSequence() []*PerCoreStepInfo {
	if x != nil {
		return x.StepSequence
	}
	return nil
}

func (x *StepDatabaseResult) GetUseIncompleteStep() bool {
	if x != nil {
		return x.UseIncompleteStep
	}
	return false
}

func (x *StepDatabaseResult) GetNumStepsDropped() uint32 {
	if x != nil {
		return x.NumStepsDropped
	}
	return 0
}

func (x *StepDatabaseResult) GetEmptyIntersect() bool {
	if x != nil {
		return x.EmptyIntersect
	}
	return false
}

var File_tensorflow_core_profiler_protobuf_steps_db_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x32, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x53, 0x74, 0x65, 0x70, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x12,
	0x4e, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x74,
	0x65, 0x70, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x50, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x50, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7c, 0x0a, 0x14, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x65,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x65, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x74, 0x65, 0x70, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x50, 0x73, 0x12,
	0x3b, 0x0a, 0x0e, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x73,
	0x74, 0x65, 0x70, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x22, 0xbc, 0x01, 0x0a,
	0x0d, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x73, 0x12, 0x1e,
	0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x11, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x4a, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xd1, 0x06, 0x0a,
	0x0f, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x65, 0x70, 0x4e, 0x75, 0x6d, 0x12, 0x66, 0x0a, 0x12, 0x73,
	0x74, 0x65, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x65,
	0x72, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x73, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x65, 0x72, 0x43,
	0x6f, 0x72, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x68, 0x6c, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x64, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x4f, 0x70, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x44, 0x62, 0x52, 0x0c, 0x68,
	0x6c, 0x6f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x44, 0x62, 0x12, 0x77, 0x0a, 0x19, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x65, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49,
	0x64, 0x4d, 0x61, 0x70, 0x12, 0x70, 0x0a, 0x16, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x5f, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x43, 0x6f,
	0x72, 0x65, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x65, 0x44, 0x62, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x62, 0x50,
	0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x61, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x15, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x1a, 0x67, 0x0a, 0x14, 0x53, 0x74, 0x65,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x47, 0x0a, 0x19, 0x43, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6d, 0x0a, 0x17, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x62, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x22, 0xe4, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x65, 0x70, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x49, 0x0a, 0x0d, 0x73, 0x74, 0x65, 0x70, 0x5f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x65, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x65, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x5f,
	0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e,
	0x75, 0x6d, 0x53, 0x74, 0x65, 0x70, 0x73, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x42, 0xe1, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x42, 0x0c, 0x53, 0x74, 0x65, 0x70, 0x73, 0x44, 0x62, 0x50, 0x72, 0x6f, 0x74,
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
	file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescOnce sync.Once
	file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescData = file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDesc
)

func file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescGZIP() []byte {
	file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescData)
	})
	return file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDescData
}

var file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tensorflow_core_profiler_protobuf_steps_db_proto_goTypes = []interface{}{
	(*GenericStepBreakdown)(nil), // 0: tensorflow.profiler.GenericStepBreakdown
	(*DeviceMemoryTransfer)(nil), // 1: tensorflow.profiler.DeviceMemoryTransfer
	(*StepInfoResult)(nil),       // 2: tensorflow.profiler.StepInfoResult
	(*AllReduceInfo)(nil),        // 3: tensorflow.profiler.AllReduceInfo
	(*AllReduceDbResult)(nil),    // 4: tensorflow.profiler.AllReduceDbResult
	(*PerCoreStepInfo)(nil),      // 5: tensorflow.profiler.PerCoreStepInfo
	(*StepDatabaseResult)(nil),   // 6: tensorflow.profiler.StepDatabaseResult
	nil,                          // 7: tensorflow.profiler.GenericStepBreakdown.TypePsEntry
	nil,                          // 8: tensorflow.profiler.PerCoreStepInfo.StepInfoPerCoreEntry
	nil,                          // 9: tensorflow.profiler.PerCoreStepInfo.CoreIdToReplicaIdMapEntry
	nil,                          // 10: tensorflow.profiler.PerCoreStepInfo.AllReduceDbPerCoreEntry
	(*anypb.Any)(nil),            // 11: google.protobuf.Any
	(*OpMetricsDb)(nil),          // 12: tensorflow.profiler.OpMetricsDb
}
var file_tensorflow_core_profiler_protobuf_steps_db_proto_depIdxs = []int32{
	7,  // 0: tensorflow.profiler.GenericStepBreakdown.type_ps:type_name -> tensorflow.profiler.GenericStepBreakdown.TypePsEntry
	11, // 1: tensorflow.profiler.StepInfoResult.step_breakdown:type_name -> google.protobuf.Any
	3,  // 2: tensorflow.profiler.AllReduceDbResult.all_reduce_info:type_name -> tensorflow.profiler.AllReduceInfo
	8,  // 3: tensorflow.profiler.PerCoreStepInfo.step_info_per_core:type_name -> tensorflow.profiler.PerCoreStepInfo.StepInfoPerCoreEntry
	12, // 4: tensorflow.profiler.PerCoreStepInfo.hlo_metrics_db:type_name -> tensorflow.profiler.OpMetricsDb
	9,  // 5: tensorflow.profiler.PerCoreStepInfo.core_id_to_replica_id_map:type_name -> tensorflow.profiler.PerCoreStepInfo.CoreIdToReplicaIdMapEntry
	10, // 6: tensorflow.profiler.PerCoreStepInfo.all_reduce_db_per_core:type_name -> tensorflow.profiler.PerCoreStepInfo.AllReduceDbPerCoreEntry
	1,  // 7: tensorflow.profiler.PerCoreStepInfo.device_memory_transfers:type_name -> tensorflow.profiler.DeviceMemoryTransfer
	5,  // 8: tensorflow.profiler.StepDatabaseResult.step_sequence:type_name -> tensorflow.profiler.PerCoreStepInfo
	2,  // 9: tensorflow.profiler.PerCoreStepInfo.StepInfoPerCoreEntry.value:type_name -> tensorflow.profiler.StepInfoResult
	4,  // 10: tensorflow.profiler.PerCoreStepInfo.AllReduceDbPerCoreEntry.value:type_name -> tensorflow.profiler.AllReduceDbResult
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_protobuf_steps_db_proto_init() }
func file_tensorflow_core_profiler_protobuf_steps_db_proto_init() {
	if File_tensorflow_core_profiler_protobuf_steps_db_proto != nil {
		return
	}
	file_tensorflow_core_profiler_protobuf_op_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericStepBreakdown); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceMemoryTransfer); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepInfoResult); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReduceInfo); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReduceDbResult); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerCoreStepInfo); i {
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
		file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StepDatabaseResult); i {
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
			RawDescriptor: file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_protobuf_steps_db_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_protobuf_steps_db_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_profiler_protobuf_steps_db_proto_msgTypes,
	}.Build()
	File_tensorflow_core_profiler_protobuf_steps_db_proto = out.File
	file_tensorflow_core_profiler_protobuf_steps_db_proto_rawDesc = nil
	file_tensorflow_core_profiler_protobuf_steps_db_proto_goTypes = nil
	file_tensorflow_core_profiler_protobuf_steps_db_proto_depIdxs = nil
}
