// Copyright 2018 The TensorFlow Authors. All Rights Reserved.
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

// XLA service API.
//
// Users 1) build up computations and 2) create allocations via this API.
// Computations are composed of data flowing between arbitrarily-sized
// vector-oriented operations.
//
// Users build up computations using a ComputationHandle, and talk about
// allocations using GlobalDataHandles.
//
// There are currently no checkpointing capabilities or distribution/replication
// guarantees. The service runs on a single machine (e.g. one task) and that is
// its failure domain.
//
// Canonical example of "alpha * X + Y":
// * Make a computation.
// * Add alpha and X and Y as parameters.
// * Request the multiplication of alpha and X.
// * Request the addition of that result and Y.
//
// Then, pass the computation and appropriately shaped inputs to the XLA
// service's Execute method, which provides a result as a GlobalDataHandle.
//
// All data in XLA computations are conceptually immutable.
//
// Note: this API is subject to change / refinement over time -- use the
// provided client libraries to insulate code from changes to this service API.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tensorflow/compiler/xla/rpc/xla_service.proto

package rpc

import (
	context "context"
	xla "github.com/airenas/go-tf-serving-protogen/tensorflow/compiler/xla"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	XlaService_Unregister_FullMethodName               = "/xla.XlaService/Unregister"
	XlaService_DeconstructTuple_FullMethodName         = "/xla.XlaService/DeconstructTuple"
	XlaService_Unpack_FullMethodName                   = "/xla.XlaService/Unpack"
	XlaService_GetShape_FullMethodName                 = "/xla.XlaService/GetShape"
	XlaService_GetComputationGraphStats_FullMethodName = "/xla.XlaService/GetComputationGraphStats"
	XlaService_LoadData_FullMethodName                 = "/xla.XlaService/LoadData"
	XlaService_TransferToClient_FullMethodName         = "/xla.XlaService/TransferToClient"
	XlaService_TransferToServer_FullMethodName         = "/xla.XlaService/TransferToServer"
	XlaService_TransferToInfeed_FullMethodName         = "/xla.XlaService/TransferToInfeed"
	XlaService_TransferFromOutfeed_FullMethodName      = "/xla.XlaService/TransferFromOutfeed"
	XlaService_ResetDevice_FullMethodName              = "/xla.XlaService/ResetDevice"
	XlaService_ComputeConstantGraph_FullMethodName     = "/xla.XlaService/ComputeConstantGraph"
	XlaService_GetDeviceHandles_FullMethodName         = "/xla.XlaService/GetDeviceHandles"
	XlaService_CreateChannelHandle_FullMethodName      = "/xla.XlaService/CreateChannelHandle"
	XlaService_Compile_FullMethodName                  = "/xla.XlaService/Compile"
	XlaService_Execute_FullMethodName                  = "/xla.XlaService/Execute"
	XlaService_ExecuteGraphParallel_FullMethodName     = "/xla.XlaService/ExecuteGraphParallel"
	XlaService_WaitForExecution_FullMethodName         = "/xla.XlaService/WaitForExecution"
)

// XlaServiceClient is the client API for XlaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XlaServiceClient interface {
	// Unregisters a global allocation.
	//
	// If the handle given is not currently allocated, a NOT_FOUND status is
	// returned.
	Unregister(ctx context.Context, in *xla.UnregisterRequest, opts ...grpc.CallOption) (*xla.UnregisterResponse, error)
	// Deconstructs a tuple. Returns a newly created GlobalDataHandle for each
	// element in the tuple.
	DeconstructTuple(ctx context.Context, in *xla.DeconstructTupleRequest, opts ...grpc.CallOption) (*xla.DeconstructTupleResponse, error)
	// Unpack requests that a global data handle, with a tuple shape, has global
	// data handles created for each of its constituent members. This is the
	// equivalent of the "destructuring assignment" present in various programming
	// languages.
	Unpack(ctx context.Context, in *xla.UnpackRequest, opts ...grpc.CallOption) (*xla.UnpackResponse, error)
	// Requests the shape of the referenced global data.
	GetShape(ctx context.Context, in *xla.GetShapeRequest, opts ...grpc.CallOption) (*xla.GetShapeResponse, error)
	// Requests the statistics of the given computation.
	GetComputationGraphStats(ctx context.Context, in *xla.ComputationGraphStatsRequest, opts ...grpc.CallOption) (*xla.ComputationStatsResponse, error)
	// Loads a variable number of values with a given element type from ColumnIO.
	LoadData(ctx context.Context, in *xla.LoadDataRequest, opts ...grpc.CallOption) (*xla.LoadDataResponse, error)
	// Transfers the given global data to the client in the form of a Literal.
	TransferToClient(ctx context.Context, in *xla.TransferToClientRequest, opts ...grpc.CallOption) (*xla.TransferToClientResponse, error)
	// Transfers the given literal to the server to be stored in a global
	// allocation, which is returned.
	TransferToServer(ctx context.Context, in *xla.TransferToServerRequest, opts ...grpc.CallOption) (*xla.TransferToServerResponse, error)
	// Transfers the given literal to the Infeed buffer of the device.
	TransferToInfeed(ctx context.Context, in *xla.TransferToInfeedRequest, opts ...grpc.CallOption) (*xla.TransferToInfeedResponse, error)
	// Transferred literal from the Outfeed buffer of the device.
	TransferFromOutfeed(ctx context.Context, in *xla.TransferFromOutfeedRequest, opts ...grpc.CallOption) (*xla.TransferFromOutfeedResponse, error)
	// Resets the device, clearing all existing state on the device.
	ResetDevice(ctx context.Context, in *xla.ResetDeviceRequest, opts ...grpc.CallOption) (*xla.ResetDeviceResponse, error)
	// Computes the value of a constant expression. The request contains the
	// computation graph for the constant expression.
	ComputeConstantGraph(ctx context.Context, in *xla.ComputeConstantGraphRequest, opts ...grpc.CallOption) (*xla.ComputeConstantResponse, error)
	// Requests one or more device handles from the target. The returned device
	// handles can be used to specify the device on which to execute computations
	// or transfer data.
	GetDeviceHandles(ctx context.Context, in *xla.GetDeviceHandlesRequest, opts ...grpc.CallOption) (*xla.GetDeviceHandlesResponse, error)
	// Creates a channel handle that can be used to transfer data between
	// two computations via a pair of Send and Recv instructions.
	CreateChannelHandle(ctx context.Context, in *xla.CreateChannelHandleRequest, opts ...grpc.CallOption) (*xla.CreateChannelHandleResponse, error)
	// Compiles the provided computation into executable. Returns the handle of
	// the executable.
	Compile(ctx context.Context, in *xla.CompileRequest, opts ...grpc.CallOption) (*xla.CompileResponse, error)
	// Invokes the provided executable with the provided global data passed as
	// immutable arguments. The request contains the handle to the executable.
	// Returns global data output and execution timing.
	Execute(ctx context.Context, in *xla.ExecuteRequest, opts ...grpc.CallOption) (*xla.ExecuteResponse, error)
	// Invokes the provided list of computations in parallel with the provided
	// global data for each computation. Returns a list of global data output and
	// execution timing.
	ExecuteGraphParallel(ctx context.Context, in *xla.ExecuteGraphParallelRequest, opts ...grpc.CallOption) (*xla.ExecuteParallelResponse, error)
	// Waits until the given execution (aysnchronously launched) is complete, and
	// returns the global data output.
	WaitForExecution(ctx context.Context, in *xla.WaitForExecutionRequest, opts ...grpc.CallOption) (*xla.WaitForExecutionResponse, error)
}

type xlaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewXlaServiceClient(cc grpc.ClientConnInterface) XlaServiceClient {
	return &xlaServiceClient{cc}
}

func (c *xlaServiceClient) Unregister(ctx context.Context, in *xla.UnregisterRequest, opts ...grpc.CallOption) (*xla.UnregisterResponse, error) {
	out := new(xla.UnregisterResponse)
	err := c.cc.Invoke(ctx, XlaService_Unregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) DeconstructTuple(ctx context.Context, in *xla.DeconstructTupleRequest, opts ...grpc.CallOption) (*xla.DeconstructTupleResponse, error) {
	out := new(xla.DeconstructTupleResponse)
	err := c.cc.Invoke(ctx, XlaService_DeconstructTuple_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Unpack(ctx context.Context, in *xla.UnpackRequest, opts ...grpc.CallOption) (*xla.UnpackResponse, error) {
	out := new(xla.UnpackResponse)
	err := c.cc.Invoke(ctx, XlaService_Unpack_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetShape(ctx context.Context, in *xla.GetShapeRequest, opts ...grpc.CallOption) (*xla.GetShapeResponse, error) {
	out := new(xla.GetShapeResponse)
	err := c.cc.Invoke(ctx, XlaService_GetShape_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetComputationGraphStats(ctx context.Context, in *xla.ComputationGraphStatsRequest, opts ...grpc.CallOption) (*xla.ComputationStatsResponse, error) {
	out := new(xla.ComputationStatsResponse)
	err := c.cc.Invoke(ctx, XlaService_GetComputationGraphStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) LoadData(ctx context.Context, in *xla.LoadDataRequest, opts ...grpc.CallOption) (*xla.LoadDataResponse, error) {
	out := new(xla.LoadDataResponse)
	err := c.cc.Invoke(ctx, XlaService_LoadData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToClient(ctx context.Context, in *xla.TransferToClientRequest, opts ...grpc.CallOption) (*xla.TransferToClientResponse, error) {
	out := new(xla.TransferToClientResponse)
	err := c.cc.Invoke(ctx, XlaService_TransferToClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToServer(ctx context.Context, in *xla.TransferToServerRequest, opts ...grpc.CallOption) (*xla.TransferToServerResponse, error) {
	out := new(xla.TransferToServerResponse)
	err := c.cc.Invoke(ctx, XlaService_TransferToServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToInfeed(ctx context.Context, in *xla.TransferToInfeedRequest, opts ...grpc.CallOption) (*xla.TransferToInfeedResponse, error) {
	out := new(xla.TransferToInfeedResponse)
	err := c.cc.Invoke(ctx, XlaService_TransferToInfeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferFromOutfeed(ctx context.Context, in *xla.TransferFromOutfeedRequest, opts ...grpc.CallOption) (*xla.TransferFromOutfeedResponse, error) {
	out := new(xla.TransferFromOutfeedResponse)
	err := c.cc.Invoke(ctx, XlaService_TransferFromOutfeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ResetDevice(ctx context.Context, in *xla.ResetDeviceRequest, opts ...grpc.CallOption) (*xla.ResetDeviceResponse, error) {
	out := new(xla.ResetDeviceResponse)
	err := c.cc.Invoke(ctx, XlaService_ResetDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ComputeConstantGraph(ctx context.Context, in *xla.ComputeConstantGraphRequest, opts ...grpc.CallOption) (*xla.ComputeConstantResponse, error) {
	out := new(xla.ComputeConstantResponse)
	err := c.cc.Invoke(ctx, XlaService_ComputeConstantGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetDeviceHandles(ctx context.Context, in *xla.GetDeviceHandlesRequest, opts ...grpc.CallOption) (*xla.GetDeviceHandlesResponse, error) {
	out := new(xla.GetDeviceHandlesResponse)
	err := c.cc.Invoke(ctx, XlaService_GetDeviceHandles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) CreateChannelHandle(ctx context.Context, in *xla.CreateChannelHandleRequest, opts ...grpc.CallOption) (*xla.CreateChannelHandleResponse, error) {
	out := new(xla.CreateChannelHandleResponse)
	err := c.cc.Invoke(ctx, XlaService_CreateChannelHandle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Compile(ctx context.Context, in *xla.CompileRequest, opts ...grpc.CallOption) (*xla.CompileResponse, error) {
	out := new(xla.CompileResponse)
	err := c.cc.Invoke(ctx, XlaService_Compile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Execute(ctx context.Context, in *xla.ExecuteRequest, opts ...grpc.CallOption) (*xla.ExecuteResponse, error) {
	out := new(xla.ExecuteResponse)
	err := c.cc.Invoke(ctx, XlaService_Execute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ExecuteGraphParallel(ctx context.Context, in *xla.ExecuteGraphParallelRequest, opts ...grpc.CallOption) (*xla.ExecuteParallelResponse, error) {
	out := new(xla.ExecuteParallelResponse)
	err := c.cc.Invoke(ctx, XlaService_ExecuteGraphParallel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) WaitForExecution(ctx context.Context, in *xla.WaitForExecutionRequest, opts ...grpc.CallOption) (*xla.WaitForExecutionResponse, error) {
	out := new(xla.WaitForExecutionResponse)
	err := c.cc.Invoke(ctx, XlaService_WaitForExecution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XlaServiceServer is the server API for XlaService service.
// All implementations must embed UnimplementedXlaServiceServer
// for forward compatibility
type XlaServiceServer interface {
	// Unregisters a global allocation.
	//
	// If the handle given is not currently allocated, a NOT_FOUND status is
	// returned.
	Unregister(context.Context, *xla.UnregisterRequest) (*xla.UnregisterResponse, error)
	// Deconstructs a tuple. Returns a newly created GlobalDataHandle for each
	// element in the tuple.
	DeconstructTuple(context.Context, *xla.DeconstructTupleRequest) (*xla.DeconstructTupleResponse, error)
	// Unpack requests that a global data handle, with a tuple shape, has global
	// data handles created for each of its constituent members. This is the
	// equivalent of the "destructuring assignment" present in various programming
	// languages.
	Unpack(context.Context, *xla.UnpackRequest) (*xla.UnpackResponse, error)
	// Requests the shape of the referenced global data.
	GetShape(context.Context, *xla.GetShapeRequest) (*xla.GetShapeResponse, error)
	// Requests the statistics of the given computation.
	GetComputationGraphStats(context.Context, *xla.ComputationGraphStatsRequest) (*xla.ComputationStatsResponse, error)
	// Loads a variable number of values with a given element type from ColumnIO.
	LoadData(context.Context, *xla.LoadDataRequest) (*xla.LoadDataResponse, error)
	// Transfers the given global data to the client in the form of a Literal.
	TransferToClient(context.Context, *xla.TransferToClientRequest) (*xla.TransferToClientResponse, error)
	// Transfers the given literal to the server to be stored in a global
	// allocation, which is returned.
	TransferToServer(context.Context, *xla.TransferToServerRequest) (*xla.TransferToServerResponse, error)
	// Transfers the given literal to the Infeed buffer of the device.
	TransferToInfeed(context.Context, *xla.TransferToInfeedRequest) (*xla.TransferToInfeedResponse, error)
	// Transferred literal from the Outfeed buffer of the device.
	TransferFromOutfeed(context.Context, *xla.TransferFromOutfeedRequest) (*xla.TransferFromOutfeedResponse, error)
	// Resets the device, clearing all existing state on the device.
	ResetDevice(context.Context, *xla.ResetDeviceRequest) (*xla.ResetDeviceResponse, error)
	// Computes the value of a constant expression. The request contains the
	// computation graph for the constant expression.
	ComputeConstantGraph(context.Context, *xla.ComputeConstantGraphRequest) (*xla.ComputeConstantResponse, error)
	// Requests one or more device handles from the target. The returned device
	// handles can be used to specify the device on which to execute computations
	// or transfer data.
	GetDeviceHandles(context.Context, *xla.GetDeviceHandlesRequest) (*xla.GetDeviceHandlesResponse, error)
	// Creates a channel handle that can be used to transfer data between
	// two computations via a pair of Send and Recv instructions.
	CreateChannelHandle(context.Context, *xla.CreateChannelHandleRequest) (*xla.CreateChannelHandleResponse, error)
	// Compiles the provided computation into executable. Returns the handle of
	// the executable.
	Compile(context.Context, *xla.CompileRequest) (*xla.CompileResponse, error)
	// Invokes the provided executable with the provided global data passed as
	// immutable arguments. The request contains the handle to the executable.
	// Returns global data output and execution timing.
	Execute(context.Context, *xla.ExecuteRequest) (*xla.ExecuteResponse, error)
	// Invokes the provided list of computations in parallel with the provided
	// global data for each computation. Returns a list of global data output and
	// execution timing.
	ExecuteGraphParallel(context.Context, *xla.ExecuteGraphParallelRequest) (*xla.ExecuteParallelResponse, error)
	// Waits until the given execution (aysnchronously launched) is complete, and
	// returns the global data output.
	WaitForExecution(context.Context, *xla.WaitForExecutionRequest) (*xla.WaitForExecutionResponse, error)
	mustEmbedUnimplementedXlaServiceServer()
}

// UnimplementedXlaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedXlaServiceServer struct {
}

func (UnimplementedXlaServiceServer) Unregister(context.Context, *xla.UnregisterRequest) (*xla.UnregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (UnimplementedXlaServiceServer) DeconstructTuple(context.Context, *xla.DeconstructTupleRequest) (*xla.DeconstructTupleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeconstructTuple not implemented")
}
func (UnimplementedXlaServiceServer) Unpack(context.Context, *xla.UnpackRequest) (*xla.UnpackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpack not implemented")
}
func (UnimplementedXlaServiceServer) GetShape(context.Context, *xla.GetShapeRequest) (*xla.GetShapeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShape not implemented")
}
func (UnimplementedXlaServiceServer) GetComputationGraphStats(context.Context, *xla.ComputationGraphStatsRequest) (*xla.ComputationStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputationGraphStats not implemented")
}
func (UnimplementedXlaServiceServer) LoadData(context.Context, *xla.LoadDataRequest) (*xla.LoadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadData not implemented")
}
func (UnimplementedXlaServiceServer) TransferToClient(context.Context, *xla.TransferToClientRequest) (*xla.TransferToClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToClient not implemented")
}
func (UnimplementedXlaServiceServer) TransferToServer(context.Context, *xla.TransferToServerRequest) (*xla.TransferToServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToServer not implemented")
}
func (UnimplementedXlaServiceServer) TransferToInfeed(context.Context, *xla.TransferToInfeedRequest) (*xla.TransferToInfeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToInfeed not implemented")
}
func (UnimplementedXlaServiceServer) TransferFromOutfeed(context.Context, *xla.TransferFromOutfeedRequest) (*xla.TransferFromOutfeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferFromOutfeed not implemented")
}
func (UnimplementedXlaServiceServer) ResetDevice(context.Context, *xla.ResetDeviceRequest) (*xla.ResetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDevice not implemented")
}
func (UnimplementedXlaServiceServer) ComputeConstantGraph(context.Context, *xla.ComputeConstantGraphRequest) (*xla.ComputeConstantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeConstantGraph not implemented")
}
func (UnimplementedXlaServiceServer) GetDeviceHandles(context.Context, *xla.GetDeviceHandlesRequest) (*xla.GetDeviceHandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceHandles not implemented")
}
func (UnimplementedXlaServiceServer) CreateChannelHandle(context.Context, *xla.CreateChannelHandleRequest) (*xla.CreateChannelHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannelHandle not implemented")
}
func (UnimplementedXlaServiceServer) Compile(context.Context, *xla.CompileRequest) (*xla.CompileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (UnimplementedXlaServiceServer) Execute(context.Context, *xla.ExecuteRequest) (*xla.ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedXlaServiceServer) ExecuteGraphParallel(context.Context, *xla.ExecuteGraphParallelRequest) (*xla.ExecuteParallelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteGraphParallel not implemented")
}
func (UnimplementedXlaServiceServer) WaitForExecution(context.Context, *xla.WaitForExecutionRequest) (*xla.WaitForExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForExecution not implemented")
}
func (UnimplementedXlaServiceServer) mustEmbedUnimplementedXlaServiceServer() {}

// UnsafeXlaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XlaServiceServer will
// result in compilation errors.
type UnsafeXlaServiceServer interface {
	mustEmbedUnimplementedXlaServiceServer()
}

func RegisterXlaServiceServer(s grpc.ServiceRegistrar, srv XlaServiceServer) {
	s.RegisterService(&XlaService_ServiceDesc, srv)
}

func _XlaService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_Unregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Unregister(ctx, req.(*xla.UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_DeconstructTuple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.DeconstructTupleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).DeconstructTuple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_DeconstructTuple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).DeconstructTuple(ctx, req.(*xla.DeconstructTupleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Unpack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.UnpackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Unpack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_Unpack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Unpack(ctx, req.(*xla.UnpackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetShape_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.GetShapeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetShape(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_GetShape_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetShape(ctx, req.(*xla.GetShapeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetComputationGraphStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.ComputationGraphStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetComputationGraphStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_GetComputationGraphStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetComputationGraphStats(ctx, req.(*xla.ComputationGraphStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_LoadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.LoadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).LoadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_LoadData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).LoadData(ctx, req.(*xla.LoadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.TransferToClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_TransferToClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToClient(ctx, req.(*xla.TransferToClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.TransferToServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_TransferToServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToServer(ctx, req.(*xla.TransferToServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToInfeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.TransferToInfeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToInfeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_TransferToInfeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToInfeed(ctx, req.(*xla.TransferToInfeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferFromOutfeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.TransferFromOutfeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferFromOutfeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_TransferFromOutfeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferFromOutfeed(ctx, req.(*xla.TransferFromOutfeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ResetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.ResetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ResetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_ResetDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ResetDevice(ctx, req.(*xla.ResetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ComputeConstantGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.ComputeConstantGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ComputeConstantGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_ComputeConstantGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ComputeConstantGraph(ctx, req.(*xla.ComputeConstantGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetDeviceHandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.GetDeviceHandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetDeviceHandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_GetDeviceHandles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetDeviceHandles(ctx, req.(*xla.GetDeviceHandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_CreateChannelHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.CreateChannelHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).CreateChannelHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_CreateChannelHandle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).CreateChannelHandle(ctx, req.(*xla.CreateChannelHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_Compile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Compile(ctx, req.(*xla.CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Execute(ctx, req.(*xla.ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ExecuteGraphParallel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.ExecuteGraphParallelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ExecuteGraphParallel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_ExecuteGraphParallel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ExecuteGraphParallel(ctx, req.(*xla.ExecuteGraphParallelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_WaitForExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(xla.WaitForExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).WaitForExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: XlaService_WaitForExecution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).WaitForExecution(ctx, req.(*xla.WaitForExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XlaService_ServiceDesc is the grpc.ServiceDesc for XlaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XlaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xla.XlaService",
	HandlerType: (*XlaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unregister",
			Handler:    _XlaService_Unregister_Handler,
		},
		{
			MethodName: "DeconstructTuple",
			Handler:    _XlaService_DeconstructTuple_Handler,
		},
		{
			MethodName: "Unpack",
			Handler:    _XlaService_Unpack_Handler,
		},
		{
			MethodName: "GetShape",
			Handler:    _XlaService_GetShape_Handler,
		},
		{
			MethodName: "GetComputationGraphStats",
			Handler:    _XlaService_GetComputationGraphStats_Handler,
		},
		{
			MethodName: "LoadData",
			Handler:    _XlaService_LoadData_Handler,
		},
		{
			MethodName: "TransferToClient",
			Handler:    _XlaService_TransferToClient_Handler,
		},
		{
			MethodName: "TransferToServer",
			Handler:    _XlaService_TransferToServer_Handler,
		},
		{
			MethodName: "TransferToInfeed",
			Handler:    _XlaService_TransferToInfeed_Handler,
		},
		{
			MethodName: "TransferFromOutfeed",
			Handler:    _XlaService_TransferFromOutfeed_Handler,
		},
		{
			MethodName: "ResetDevice",
			Handler:    _XlaService_ResetDevice_Handler,
		},
		{
			MethodName: "ComputeConstantGraph",
			Handler:    _XlaService_ComputeConstantGraph_Handler,
		},
		{
			MethodName: "GetDeviceHandles",
			Handler:    _XlaService_GetDeviceHandles_Handler,
		},
		{
			MethodName: "CreateChannelHandle",
			Handler:    _XlaService_CreateChannelHandle_Handler,
		},
		{
			MethodName: "Compile",
			Handler:    _XlaService_Compile_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _XlaService_Execute_Handler,
		},
		{
			MethodName: "ExecuteGraphParallel",
			Handler:    _XlaService_ExecuteGraphParallel_Handler,
		},
		{
			MethodName: "WaitForExecution",
			Handler:    _XlaService_WaitForExecution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/compiler/xla/rpc/xla_service.proto",
}