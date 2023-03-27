// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: data.proto

package grpc

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
	FetchData_FetchCpuUsage_FullMethodName = "/data.protobuf.FetchData/FetchCpuUsage"
	FetchData_InitCpuDetail_FullMethodName = "/data.protobuf.FetchData/InitCpuDetail"
)

// FetchDataClient is the client API for FetchData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchDataClient interface {
	FetchCpuUsage(ctx context.Context, in *CpuUsageRequest, opts ...grpc.CallOption) (*CpuUsage, error)
	InitCpuDetail(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*InitData, error)
}

type fetchDataClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchDataClient(cc grpc.ClientConnInterface) FetchDataClient {
	return &fetchDataClient{cc}
}

func (c *fetchDataClient) FetchCpuUsage(ctx context.Context, in *CpuUsageRequest, opts ...grpc.CallOption) (*CpuUsage, error) {
	out := new(CpuUsage)
	err := c.cc.Invoke(ctx, FetchData_FetchCpuUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fetchDataClient) InitCpuDetail(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*InitData, error) {
	out := new(InitData)
	err := c.cc.Invoke(ctx, FetchData_InitCpuDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchDataServer is the server API for FetchData service.
// All implementations must embed UnimplementedFetchDataServer
// for forward compatibility
type FetchDataServer interface {
	FetchCpuUsage(context.Context, *CpuUsageRequest) (*CpuUsage, error)
	InitCpuDetail(context.Context, *EmptyReq) (*InitData, error)
	mustEmbedUnimplementedFetchDataServer()
}

// UnimplementedFetchDataServer must be embedded to have forward compatible implementations.
type UnimplementedFetchDataServer struct {
}

func (UnimplementedFetchDataServer) FetchCpuUsage(context.Context, *CpuUsageRequest) (*CpuUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchCpuUsage not implemented")
}
func (UnimplementedFetchDataServer) InitCpuDetail(context.Context, *EmptyReq) (*InitData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitCpuDetail not implemented")
}
func (UnimplementedFetchDataServer) mustEmbedUnimplementedFetchDataServer() {}

// UnsafeFetchDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchDataServer will
// result in compilation errors.
type UnsafeFetchDataServer interface {
	mustEmbedUnimplementedFetchDataServer()
}

func RegisterFetchDataServer(s grpc.ServiceRegistrar, srv FetchDataServer) {
	s.RegisterService(&FetchData_ServiceDesc, srv)
}

func _FetchData_FetchCpuUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CpuUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchDataServer).FetchCpuUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FetchData_FetchCpuUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchDataServer).FetchCpuUsage(ctx, req.(*CpuUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FetchData_InitCpuDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchDataServer).InitCpuDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FetchData_InitCpuDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchDataServer).InitCpuDetail(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FetchData_ServiceDesc is the grpc.ServiceDesc for FetchData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.protobuf.FetchData",
	HandlerType: (*FetchDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCpuUsage",
			Handler:    _FetchData_FetchCpuUsage_Handler,
		},
		{
			MethodName: "InitCpuDetail",
			Handler:    _FetchData_InitCpuDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

const (
	FetchDataMem_FetchMemUsage_FullMethodName = "/data.protobuf.FetchDataMem/FetchMemUsage"
)

// FetchDataMemClient is the client API for FetchDataMem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchDataMemClient interface {
	FetchMemUsage(ctx context.Context, in *MemUsageRequest, opts ...grpc.CallOption) (*MemUsage, error)
}

type fetchDataMemClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchDataMemClient(cc grpc.ClientConnInterface) FetchDataMemClient {
	return &fetchDataMemClient{cc}
}

func (c *fetchDataMemClient) FetchMemUsage(ctx context.Context, in *MemUsageRequest, opts ...grpc.CallOption) (*MemUsage, error) {
	out := new(MemUsage)
	err := c.cc.Invoke(ctx, FetchDataMem_FetchMemUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchDataMemServer is the server API for FetchDataMem service.
// All implementations must embed UnimplementedFetchDataMemServer
// for forward compatibility
type FetchDataMemServer interface {
	FetchMemUsage(context.Context, *MemUsageRequest) (*MemUsage, error)
	mustEmbedUnimplementedFetchDataMemServer()
}

// UnimplementedFetchDataMemServer must be embedded to have forward compatible implementations.
type UnimplementedFetchDataMemServer struct {
}

func (UnimplementedFetchDataMemServer) FetchMemUsage(context.Context, *MemUsageRequest) (*MemUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMemUsage not implemented")
}
func (UnimplementedFetchDataMemServer) mustEmbedUnimplementedFetchDataMemServer() {}

// UnsafeFetchDataMemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchDataMemServer will
// result in compilation errors.
type UnsafeFetchDataMemServer interface {
	mustEmbedUnimplementedFetchDataMemServer()
}

func RegisterFetchDataMemServer(s grpc.ServiceRegistrar, srv FetchDataMemServer) {
	s.RegisterService(&FetchDataMem_ServiceDesc, srv)
}

func _FetchDataMem_FetchMemUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchDataMemServer).FetchMemUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FetchDataMem_FetchMemUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchDataMemServer).FetchMemUsage(ctx, req.(*MemUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FetchDataMem_ServiceDesc is the grpc.ServiceDesc for FetchDataMem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchDataMem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.protobuf.FetchDataMem",
	HandlerType: (*FetchDataMemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMemUsage",
			Handler:    _FetchDataMem_FetchMemUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

const (
	FetchDataBat_FetchBatUsage_FullMethodName = "/data.protobuf.FetchDataBat/FetchBatUsage"
)

// FetchDataBatClient is the client API for FetchDataBat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchDataBatClient interface {
	FetchBatUsage(ctx context.Context, in *BatUsageRequest, opts ...grpc.CallOption) (*BatUsage, error)
}

type fetchDataBatClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchDataBatClient(cc grpc.ClientConnInterface) FetchDataBatClient {
	return &fetchDataBatClient{cc}
}

func (c *fetchDataBatClient) FetchBatUsage(ctx context.Context, in *BatUsageRequest, opts ...grpc.CallOption) (*BatUsage, error) {
	out := new(BatUsage)
	err := c.cc.Invoke(ctx, FetchDataBat_FetchBatUsage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchDataBatServer is the server API for FetchDataBat service.
// All implementations must embed UnimplementedFetchDataBatServer
// for forward compatibility
type FetchDataBatServer interface {
	FetchBatUsage(context.Context, *BatUsageRequest) (*BatUsage, error)
	mustEmbedUnimplementedFetchDataBatServer()
}

// UnimplementedFetchDataBatServer must be embedded to have forward compatible implementations.
type UnimplementedFetchDataBatServer struct {
}

func (UnimplementedFetchDataBatServer) FetchBatUsage(context.Context, *BatUsageRequest) (*BatUsage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBatUsage not implemented")
}
func (UnimplementedFetchDataBatServer) mustEmbedUnimplementedFetchDataBatServer() {}

// UnsafeFetchDataBatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchDataBatServer will
// result in compilation errors.
type UnsafeFetchDataBatServer interface {
	mustEmbedUnimplementedFetchDataBatServer()
}

func RegisterFetchDataBatServer(s grpc.ServiceRegistrar, srv FetchDataBatServer) {
	s.RegisterService(&FetchDataBat_ServiceDesc, srv)
}

func _FetchDataBat_FetchBatUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchDataBatServer).FetchBatUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FetchDataBat_FetchBatUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchDataBatServer).FetchBatUsage(ctx, req.(*BatUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FetchDataBat_ServiceDesc is the grpc.ServiceDesc for FetchDataBat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchDataBat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.protobuf.FetchDataBat",
	HandlerType: (*FetchDataBatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchBatUsage",
			Handler:    _FetchDataBat_FetchBatUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
