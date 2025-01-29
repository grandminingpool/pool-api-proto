// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pool_miners.proto

package poolMinersProto

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

// PoolMinersServiceClient is the client API for PoolMinersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolMinersServiceClient interface {
	ValidateAddress(ctx context.Context, in *MinerAddressRequest, opts ...grpc.CallOption) (*ValidateAddressResponse, error)
	GetMiners(ctx context.Context, in *GetMinersRequest, opts ...grpc.CallOption) (*MinersList, error)
	GetMiner(ctx context.Context, in *MinerAddressRequest, opts ...grpc.CallOption) (*Miner, error)
	GetMinerWorkers(ctx context.Context, in *GetMinerWorkersRequest, opts ...grpc.CallOption) (*MinerWorkersList, error)
	GetMinersWorkersFromList(ctx context.Context, in *MinerAddressesRequest, opts ...grpc.CallOption) (*MinersWorkersMap, error)
}

type poolMinersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolMinersServiceClient(cc grpc.ClientConnInterface) PoolMinersServiceClient {
	return &poolMinersServiceClient{cc}
}

func (c *poolMinersServiceClient) ValidateAddress(ctx context.Context, in *MinerAddressRequest, opts ...grpc.CallOption) (*ValidateAddressResponse, error) {
	out := new(ValidateAddressResponse)
	err := c.cc.Invoke(ctx, "/pool_miners.PoolMinersService/ValidateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolMinersServiceClient) GetMiners(ctx context.Context, in *GetMinersRequest, opts ...grpc.CallOption) (*MinersList, error) {
	out := new(MinersList)
	err := c.cc.Invoke(ctx, "/pool_miners.PoolMinersService/GetMiners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolMinersServiceClient) GetMiner(ctx context.Context, in *MinerAddressRequest, opts ...grpc.CallOption) (*Miner, error) {
	out := new(Miner)
	err := c.cc.Invoke(ctx, "/pool_miners.PoolMinersService/GetMiner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolMinersServiceClient) GetMinerWorkers(ctx context.Context, in *GetMinerWorkersRequest, opts ...grpc.CallOption) (*MinerWorkersList, error) {
	out := new(MinerWorkersList)
	err := c.cc.Invoke(ctx, "/pool_miners.PoolMinersService/GetMinerWorkers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolMinersServiceClient) GetMinersWorkersFromList(ctx context.Context, in *MinerAddressesRequest, opts ...grpc.CallOption) (*MinersWorkersMap, error) {
	out := new(MinersWorkersMap)
	err := c.cc.Invoke(ctx, "/pool_miners.PoolMinersService/GetMinersWorkersFromList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolMinersServiceServer is the server API for PoolMinersService service.
// All implementations must embed UnimplementedPoolMinersServiceServer
// for forward compatibility
type PoolMinersServiceServer interface {
	ValidateAddress(context.Context, *MinerAddressRequest) (*ValidateAddressResponse, error)
	GetMiners(context.Context, *GetMinersRequest) (*MinersList, error)
	GetMiner(context.Context, *MinerAddressRequest) (*Miner, error)
	GetMinerWorkers(context.Context, *GetMinerWorkersRequest) (*MinerWorkersList, error)
	GetMinersWorkersFromList(context.Context, *MinerAddressesRequest) (*MinersWorkersMap, error)
	mustEmbedUnimplementedPoolMinersServiceServer()
}

// UnimplementedPoolMinersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoolMinersServiceServer struct {
}

func (UnimplementedPoolMinersServiceServer) ValidateAddress(context.Context, *MinerAddressRequest) (*ValidateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAddress not implemented")
}
func (UnimplementedPoolMinersServiceServer) GetMiners(context.Context, *GetMinersRequest) (*MinersList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMiners not implemented")
}
func (UnimplementedPoolMinersServiceServer) GetMiner(context.Context, *MinerAddressRequest) (*Miner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMiner not implemented")
}
func (UnimplementedPoolMinersServiceServer) GetMinerWorkers(context.Context, *GetMinerWorkersRequest) (*MinerWorkersList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinerWorkers not implemented")
}
func (UnimplementedPoolMinersServiceServer) GetMinersWorkersFromList(context.Context, *MinerAddressesRequest) (*MinersWorkersMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinersWorkersFromList not implemented")
}
func (UnimplementedPoolMinersServiceServer) mustEmbedUnimplementedPoolMinersServiceServer() {}

// UnsafePoolMinersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolMinersServiceServer will
// result in compilation errors.
type UnsafePoolMinersServiceServer interface {
	mustEmbedUnimplementedPoolMinersServiceServer()
}

func RegisterPoolMinersServiceServer(s grpc.ServiceRegistrar, srv PoolMinersServiceServer) {
	s.RegisterService(&PoolMinersService_ServiceDesc, srv)
}

func _PoolMinersService_ValidateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolMinersServiceServer).ValidateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pool_miners.PoolMinersService/ValidateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolMinersServiceServer).ValidateAddress(ctx, req.(*MinerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolMinersService_GetMiners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolMinersServiceServer).GetMiners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pool_miners.PoolMinersService/GetMiners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolMinersServiceServer).GetMiners(ctx, req.(*GetMinersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolMinersService_GetMiner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolMinersServiceServer).GetMiner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pool_miners.PoolMinersService/GetMiner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolMinersServiceServer).GetMiner(ctx, req.(*MinerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolMinersService_GetMinerWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMinerWorkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolMinersServiceServer).GetMinerWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pool_miners.PoolMinersService/GetMinerWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolMinersServiceServer).GetMinerWorkers(ctx, req.(*GetMinerWorkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolMinersService_GetMinersWorkersFromList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolMinersServiceServer).GetMinersWorkersFromList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pool_miners.PoolMinersService/GetMinersWorkersFromList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolMinersServiceServer).GetMinersWorkersFromList(ctx, req.(*MinerAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoolMinersService_ServiceDesc is the grpc.ServiceDesc for PoolMinersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoolMinersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pool_miners.PoolMinersService",
	HandlerType: (*PoolMinersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateAddress",
			Handler:    _PoolMinersService_ValidateAddress_Handler,
		},
		{
			MethodName: "GetMiners",
			Handler:    _PoolMinersService_GetMiners_Handler,
		},
		{
			MethodName: "GetMiner",
			Handler:    _PoolMinersService_GetMiner_Handler,
		},
		{
			MethodName: "GetMinerWorkers",
			Handler:    _PoolMinersService_GetMinerWorkers_Handler,
		},
		{
			MethodName: "GetMinersWorkersFromList",
			Handler:    _PoolMinersService_GetMinersWorkersFromList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool_miners.proto",
}
