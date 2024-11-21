// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: charts.proto

package chartsProto

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

// ChartsServiceClient is the client API for ChartsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChartsServiceClient interface {
	GetPoolStats(ctx context.Context, in *GetPoolStatsRequest, opts ...grpc.CallOption) (*PoolStatsPoints, error)
	GetPoolDifficulties(ctx context.Context, in *GetPoolDifficultiesRequest, opts ...grpc.CallOption) (*PoolDifficultiesPoints, error)
	GetRounds(ctx context.Context, in *GetRoundsRequest, opts ...grpc.CallOption) (*RoundsPoints, error)
	GetMinerProfitabilities(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerProfitabilitiesPoints, error)
	GetMinerHashrates(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerHashratesPoints, error)
	GetMinerWorkerHashrates(ctx context.Context, in *GetMinerWorkerChartRequest, opts ...grpc.CallOption) (*MinerHashratesPoints, error)
	GetMinerShares(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerSharesPoints, error)
	GetMinerWorkerShares(ctx context.Context, in *GetMinerWorkerChartRequest, opts ...grpc.CallOption) (*MinerSharesPoints, error)
}

type chartsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChartsServiceClient(cc grpc.ClientConnInterface) ChartsServiceClient {
	return &chartsServiceClient{cc}
}

func (c *chartsServiceClient) GetPoolStats(ctx context.Context, in *GetPoolStatsRequest, opts ...grpc.CallOption) (*PoolStatsPoints, error) {
	out := new(PoolStatsPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetPoolStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetPoolDifficulties(ctx context.Context, in *GetPoolDifficultiesRequest, opts ...grpc.CallOption) (*PoolDifficultiesPoints, error) {
	out := new(PoolDifficultiesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetPoolDifficulties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetRounds(ctx context.Context, in *GetRoundsRequest, opts ...grpc.CallOption) (*RoundsPoints, error) {
	out := new(RoundsPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetRounds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetMinerProfitabilities(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerProfitabilitiesPoints, error) {
	out := new(MinerProfitabilitiesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetMinerProfitabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetMinerHashrates(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerHashratesPoints, error) {
	out := new(MinerHashratesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetMinerHashrates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetMinerWorkerHashrates(ctx context.Context, in *GetMinerWorkerChartRequest, opts ...grpc.CallOption) (*MinerHashratesPoints, error) {
	out := new(MinerHashratesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetMinerWorkerHashrates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetMinerShares(ctx context.Context, in *GetMinerChartRequest, opts ...grpc.CallOption) (*MinerSharesPoints, error) {
	out := new(MinerSharesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetMinerShares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chartsServiceClient) GetMinerWorkerShares(ctx context.Context, in *GetMinerWorkerChartRequest, opts ...grpc.CallOption) (*MinerSharesPoints, error) {
	out := new(MinerSharesPoints)
	err := c.cc.Invoke(ctx, "/charts.ChartsService/GetMinerWorkerShares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChartsServiceServer is the server API for ChartsService service.
// All implementations must embed UnimplementedChartsServiceServer
// for forward compatibility
type ChartsServiceServer interface {
	GetPoolStats(context.Context, *GetPoolStatsRequest) (*PoolStatsPoints, error)
	GetPoolDifficulties(context.Context, *GetPoolDifficultiesRequest) (*PoolDifficultiesPoints, error)
	GetRounds(context.Context, *GetRoundsRequest) (*RoundsPoints, error)
	GetMinerProfitabilities(context.Context, *GetMinerChartRequest) (*MinerProfitabilitiesPoints, error)
	GetMinerHashrates(context.Context, *GetMinerChartRequest) (*MinerHashratesPoints, error)
	GetMinerWorkerHashrates(context.Context, *GetMinerWorkerChartRequest) (*MinerHashratesPoints, error)
	GetMinerShares(context.Context, *GetMinerChartRequest) (*MinerSharesPoints, error)
	GetMinerWorkerShares(context.Context, *GetMinerWorkerChartRequest) (*MinerSharesPoints, error)
	mustEmbedUnimplementedChartsServiceServer()
}

// UnimplementedChartsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChartsServiceServer struct {
}

func (UnimplementedChartsServiceServer) GetPoolStats(context.Context, *GetPoolStatsRequest) (*PoolStatsPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolStats not implemented")
}
func (UnimplementedChartsServiceServer) GetPoolDifficulties(context.Context, *GetPoolDifficultiesRequest) (*PoolDifficultiesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolDifficulties not implemented")
}
func (UnimplementedChartsServiceServer) GetRounds(context.Context, *GetRoundsRequest) (*RoundsPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRounds not implemented")
}
func (UnimplementedChartsServiceServer) GetMinerProfitabilities(context.Context, *GetMinerChartRequest) (*MinerProfitabilitiesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerProfitabilities not implemented")
}
func (UnimplementedChartsServiceServer) GetMinerHashrates(context.Context, *GetMinerChartRequest) (*MinerHashratesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerHashrates not implemented")
}
func (UnimplementedChartsServiceServer) GetMinerWorkerHashrates(context.Context, *GetMinerWorkerChartRequest) (*MinerHashratesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerWorkerHashrates not implemented")
}
func (UnimplementedChartsServiceServer) GetMinerShares(context.Context, *GetMinerChartRequest) (*MinerSharesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerShares not implemented")
}
func (UnimplementedChartsServiceServer) GetMinerWorkerShares(context.Context, *GetMinerWorkerChartRequest) (*MinerSharesPoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerWorkerShares not implemented")
}
func (UnimplementedChartsServiceServer) mustEmbedUnimplementedChartsServiceServer() {}

// UnsafeChartsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChartsServiceServer will
// result in compilation errors.
type UnsafeChartsServiceServer interface {
	mustEmbedUnimplementedChartsServiceServer()
}

func RegisterChartsServiceServer(s grpc.ServiceRegistrar, srv ChartsServiceServer) {
	s.RegisterService(&ChartsService_ServiceDesc, srv)
}

func _ChartsService_GetPoolStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetPoolStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetPoolStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetPoolStats(ctx, req.(*GetPoolStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetPoolDifficulties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolDifficultiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetPoolDifficulties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetPoolDifficulties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetPoolDifficulties(ctx, req.(*GetPoolDifficultiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetRounds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetRounds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetRounds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetRounds(ctx, req.(*GetRoundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetMinerProfitabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetMinerProfitabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetMinerProfitabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetMinerProfitabilities(ctx, req.(*GetMinerChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetMinerHashrates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetMinerHashrates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetMinerHashrates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetMinerHashrates(ctx, req.(*GetMinerChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetMinerWorkerHashrates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerWorkerChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetMinerWorkerHashrates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetMinerWorkerHashrates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetMinerWorkerHashrates(ctx, req.(*GetMinerWorkerChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetMinerShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetMinerShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetMinerShares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetMinerShares(ctx, req.(*GetMinerChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChartsService_GetMinerWorkerShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerWorkerChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChartsServiceServer).GetMinerWorkerShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/charts.ChartsService/GetMinerWorkerShares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChartsServiceServer).GetMinerWorkerShares(ctx, req.(*GetMinerWorkerChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChartsService_ServiceDesc is the grpc.ServiceDesc for ChartsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChartsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "charts.ChartsService",
	HandlerType: (*ChartsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoolStats",
			Handler:    _ChartsService_GetPoolStats_Handler,
		},
		{
			MethodName: "GetPoolDifficulties",
			Handler:    _ChartsService_GetPoolDifficulties_Handler,
		},
		{
			MethodName: "GetRounds",
			Handler:    _ChartsService_GetRounds_Handler,
		},
		{
			MethodName: "GetMinerProfitabilities",
			Handler:    _ChartsService_GetMinerProfitabilities_Handler,
		},
		{
			MethodName: "GetMinerHashrates",
			Handler:    _ChartsService_GetMinerHashrates_Handler,
		},
		{
			MethodName: "GetMinerWorkerHashrates",
			Handler:    _ChartsService_GetMinerWorkerHashrates_Handler,
		},
		{
			MethodName: "GetMinerShares",
			Handler:    _ChartsService_GetMinerShares_Handler,
		},
		{
			MethodName: "GetMinerWorkerShares",
			Handler:    _ChartsService_GetMinerWorkerShares_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "charts.proto",
}
