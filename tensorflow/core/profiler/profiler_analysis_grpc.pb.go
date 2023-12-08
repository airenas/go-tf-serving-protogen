// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tensorflow/core/profiler/profiler_analysis.proto

package profiler

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProfileAnalysis_NewSession_FullMethodName         = "/tensorflow.ProfileAnalysis/NewSession"
	ProfileAnalysis_EnumSessions_FullMethodName       = "/tensorflow.ProfileAnalysis/EnumSessions"
	ProfileAnalysis_GetSessionToolData_FullMethodName = "/tensorflow.ProfileAnalysis/GetSessionToolData"
)

// ProfileAnalysisClient is the client API for ProfileAnalysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileAnalysisClient interface {
	// Starts a profiling session, blocks until it completes.
	// TPUProfileAnalysis service delegate this to TPUProfiler service.
	// Populate the profiled data in repository, then return status to caller.
	NewSession(ctx context.Context, in *NewProfileSessionRequest, opts ...grpc.CallOption) (*NewProfileSessionResponse, error)
	// Enumerate existing sessions and return available profile tools.
	EnumSessions(ctx context.Context, in *EnumProfileSessionsAndToolsRequest, opts ...grpc.CallOption) (*EnumProfileSessionsAndToolsResponse, error)
	// Retrieve specific tool's data for specific session.
	GetSessionToolData(ctx context.Context, in *ProfileSessionDataRequest, opts ...grpc.CallOption) (*ProfileSessionDataResponse, error)
}

type profileAnalysisClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileAnalysisClient(cc grpc.ClientConnInterface) ProfileAnalysisClient {
	return &profileAnalysisClient{cc}
}

func (c *profileAnalysisClient) NewSession(ctx context.Context, in *NewProfileSessionRequest, opts ...grpc.CallOption) (*NewProfileSessionResponse, error) {
	out := new(NewProfileSessionResponse)
	err := c.cc.Invoke(ctx, ProfileAnalysis_NewSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileAnalysisClient) EnumSessions(ctx context.Context, in *EnumProfileSessionsAndToolsRequest, opts ...grpc.CallOption) (*EnumProfileSessionsAndToolsResponse, error) {
	out := new(EnumProfileSessionsAndToolsResponse)
	err := c.cc.Invoke(ctx, ProfileAnalysis_EnumSessions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileAnalysisClient) GetSessionToolData(ctx context.Context, in *ProfileSessionDataRequest, opts ...grpc.CallOption) (*ProfileSessionDataResponse, error) {
	out := new(ProfileSessionDataResponse)
	err := c.cc.Invoke(ctx, ProfileAnalysis_GetSessionToolData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileAnalysisServer is the server API for ProfileAnalysis service.
// All implementations must embed UnimplementedProfileAnalysisServer
// for forward compatibility
type ProfileAnalysisServer interface {
	// Starts a profiling session, blocks until it completes.
	// TPUProfileAnalysis service delegate this to TPUProfiler service.
	// Populate the profiled data in repository, then return status to caller.
	NewSession(context.Context, *NewProfileSessionRequest) (*NewProfileSessionResponse, error)
	// Enumerate existing sessions and return available profile tools.
	EnumSessions(context.Context, *EnumProfileSessionsAndToolsRequest) (*EnumProfileSessionsAndToolsResponse, error)
	// Retrieve specific tool's data for specific session.
	GetSessionToolData(context.Context, *ProfileSessionDataRequest) (*ProfileSessionDataResponse, error)
	mustEmbedUnimplementedProfileAnalysisServer()
}

// UnimplementedProfileAnalysisServer must be embedded to have forward compatible implementations.
type UnimplementedProfileAnalysisServer struct {
}

func (UnimplementedProfileAnalysisServer) NewSession(context.Context, *NewProfileSessionRequest) (*NewProfileSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSession not implemented")
}
func (UnimplementedProfileAnalysisServer) EnumSessions(context.Context, *EnumProfileSessionsAndToolsRequest) (*EnumProfileSessionsAndToolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnumSessions not implemented")
}
func (UnimplementedProfileAnalysisServer) GetSessionToolData(context.Context, *ProfileSessionDataRequest) (*ProfileSessionDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionToolData not implemented")
}
func (UnimplementedProfileAnalysisServer) mustEmbedUnimplementedProfileAnalysisServer() {}

// UnsafeProfileAnalysisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileAnalysisServer will
// result in compilation errors.
type UnsafeProfileAnalysisServer interface {
	mustEmbedUnimplementedProfileAnalysisServer()
}

func RegisterProfileAnalysisServer(s grpc.ServiceRegistrar, srv ProfileAnalysisServer) {
	s.RegisterService(&ProfileAnalysis_ServiceDesc, srv)
}

func _ProfileAnalysis_NewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProfileSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAnalysisServer).NewSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileAnalysis_NewSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAnalysisServer).NewSession(ctx, req.(*NewProfileSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileAnalysis_EnumSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnumProfileSessionsAndToolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAnalysisServer).EnumSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileAnalysis_EnumSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAnalysisServer).EnumSessions(ctx, req.(*EnumProfileSessionsAndToolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileAnalysis_GetSessionToolData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileSessionDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAnalysisServer).GetSessionToolData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileAnalysis_GetSessionToolData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAnalysisServer).GetSessionToolData(ctx, req.(*ProfileSessionDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileAnalysis_ServiceDesc is the grpc.ServiceDesc for ProfileAnalysis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileAnalysis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.ProfileAnalysis",
	HandlerType: (*ProfileAnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSession",
			Handler:    _ProfileAnalysis_NewSession_Handler,
		},
		{
			MethodName: "EnumSessions",
			Handler:    _ProfileAnalysis_EnumSessions_Handler,
		},
		{
			MethodName: "GetSessionToolData",
			Handler:    _ProfileAnalysis_GetSessionToolData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/core/profiler/profiler_analysis.proto",
}
