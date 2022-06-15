// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package globalRPCProto

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

// GlobalRPCClient is the client API for GlobalRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalRPCClient interface {
	// Asset
	GetLatestAccountLp(ctx context.Context, in *ReqGetLatestAccountLp, opts ...grpc.CallOption) (*RespGetLatestAccountLp, error)
	GetLatestAssetsListByAccountIndex(ctx context.Context, in *ReqGetLatestAssetsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestAssetsListByAccountIndex, error)
	GetLatestAssetInfoByAccountIndexAndAssetId(ctx context.Context, in *ReqGetLatestAssetInfoByAccountIndexAndAssetId, opts ...grpc.CallOption) (*RespGetLatestAssetInfoByAccountIndexAndAssetId, error)
	// Liquidity
	GetLatestPairInfo(ctx context.Context, in *ReqGetLatestPairInfo, opts ...grpc.CallOption) (*RespGetLatestPairInfo, error)
	GetSwapAmount(ctx context.Context, in *ReqGetSwapAmount, opts ...grpc.CallOption) (*RespGetSwapAmount, error)
	GetLpValue(ctx context.Context, in *ReqGetLpValue, opts ...grpc.CallOption) (*RespGetLpValue, error)
	// Transaction
	GetLatestTxsListByAccountIndex(ctx context.Context, in *ReqGetLatestTxsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestTxsListByAccountIndex, error)
	GetLatestTxsListByAccountIndexAndTxType(ctx context.Context, in *ReqGetLatestTxsListByAccountIndexAndTxType, opts ...grpc.CallOption) (*RespGetLatestTxsListByAccountIndexAndTxType, error)
	SendTx(ctx context.Context, in *ReqSendTx, opts ...grpc.CallOption) (*RespSendTx, error)
}

type globalRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalRPCClient(cc grpc.ClientConnInterface) GlobalRPCClient {
	return &globalRPCClient{cc}
}

func (c *globalRPCClient) GetLatestAccountLp(ctx context.Context, in *ReqGetLatestAccountLp, opts ...grpc.CallOption) (*RespGetLatestAccountLp, error) {
	out := new(RespGetLatestAccountLp)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestAccountLp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetLatestAssetsListByAccountIndex(ctx context.Context, in *ReqGetLatestAssetsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestAssetsListByAccountIndex, error) {
	out := new(RespGetLatestAssetsListByAccountIndex)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestAssetsListByAccountIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetLatestAssetInfoByAccountIndexAndAssetId(ctx context.Context, in *ReqGetLatestAssetInfoByAccountIndexAndAssetId, opts ...grpc.CallOption) (*RespGetLatestAssetInfoByAccountIndexAndAssetId, error) {
	out := new(RespGetLatestAssetInfoByAccountIndexAndAssetId)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestAssetInfoByAccountIndexAndAssetId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetLatestPairInfo(ctx context.Context, in *ReqGetLatestPairInfo, opts ...grpc.CallOption) (*RespGetLatestPairInfo, error) {
	out := new(RespGetLatestPairInfo)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestPairInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetSwapAmount(ctx context.Context, in *ReqGetSwapAmount, opts ...grpc.CallOption) (*RespGetSwapAmount, error) {
	out := new(RespGetSwapAmount)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getSwapAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *globalRPCClient) GetLpValue(ctx context.Context, in *ReqGetLpValue, opts ...grpc.CallOption) (*RespGetLpValue, error) {
	out := new(RespGetLpValue)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLpValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetLatestTxsListByAccountIndex(ctx context.Context, in *ReqGetLatestTxsListByAccountIndex, opts ...grpc.CallOption) (*RespGetLatestTxsListByAccountIndex, error) {
	out := new(RespGetLatestTxsListByAccountIndex)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestTxsListByAccountIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) GetLatestTxsListByAccountIndexAndTxType(ctx context.Context, in *ReqGetLatestTxsListByAccountIndexAndTxType, opts ...grpc.CallOption) (*RespGetLatestTxsListByAccountIndexAndTxType, error) {
	out := new(RespGetLatestTxsListByAccountIndexAndTxType)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/getLatestTxsListByAccountIndexAndTxType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *globalRPCClient) SendTx(ctx context.Context, in *ReqSendTx, opts ...grpc.CallOption) (*RespSendTx, error) {
	out := new(RespSendTx)
	err := c.cc.Invoke(ctx, "/globalRPCProto.globalRPC/sendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalRPCServer is the server API for GlobalRPC service.
// All implementations must embed UnimplementedGlobalRPCServer
// for forward compatibility
type GlobalRPCServer interface {
	// Asset
	GetLatestAccountLp(context.Context, *ReqGetLatestAccountLp) (*RespGetLatestAccountLp, error)
	GetLatestAssetsListByAccountIndex(context.Context, *ReqGetLatestAssetsListByAccountIndex) (*RespGetLatestAssetsListByAccountIndex, error)
	GetLatestAssetInfoByAccountIndexAndAssetId(context.Context, *ReqGetLatestAssetInfoByAccountIndexAndAssetId) (*RespGetLatestAssetInfoByAccountIndexAndAssetId, error)
	// Liquidity
	GetLatestPairInfo(context.Context, *ReqGetLatestPairInfo) (*RespGetLatestPairInfo, error)
	GetSwapAmount(context.Context, *ReqGetSwapAmount) (*RespGetSwapAmount, error)
	GetLpValue(context.Context, *ReqGetLpValue) (*RespGetLpValue, error)
	// Transaction
	GetLatestTxsListByAccountIndex(context.Context, *ReqGetLatestTxsListByAccountIndex) (*RespGetLatestTxsListByAccountIndex, error)
	GetLatestTxsListByAccountIndexAndTxType(context.Context, *ReqGetLatestTxsListByAccountIndexAndTxType) (*RespGetLatestTxsListByAccountIndexAndTxType, error)
	SendTx(context.Context, *ReqSendTx) (*RespSendTx, error)
	mustEmbedUnimplementedGlobalRPCServer()
}

// UnimplementedGlobalRPCServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalRPCServer struct {
}

func (UnimplementedGlobalRPCServer) GetLatestAccountLp(context.Context, *ReqGetLatestAccountLp) (*RespGetLatestAccountLp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestAccountLp not implemented")
}
func (UnimplementedGlobalRPCServer) GetLatestAssetsListByAccountIndex(context.Context, *ReqGetLatestAssetsListByAccountIndex) (*RespGetLatestAssetsListByAccountIndex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestAssetsListByAccountIndex not implemented")
}
func (UnimplementedGlobalRPCServer) GetLatestAssetInfoByAccountIndexAndAssetId(context.Context, *ReqGetLatestAssetInfoByAccountIndexAndAssetId) (*RespGetLatestAssetInfoByAccountIndexAndAssetId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestAssetInfoByAccountIndexAndAssetId not implemented")
}
func (UnimplementedGlobalRPCServer) GetLatestPairInfo(context.Context, *ReqGetLatestPairInfo) (*RespGetLatestPairInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestPairInfo not implemented")
}
func (UnimplementedGlobalRPCServer) GetSwapAmount(context.Context, *ReqGetSwapAmount) (*RespGetSwapAmount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwapAmount not implemented")
}
func (UnimplementedGlobalRPCServer) GetLpValue(context.Context, *ReqGetLpValue) (*RespGetLpValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLpValue not implemented")
}
func (UnimplementedGlobalRPCServer) GetLatestTxsListByAccountIndex(context.Context, *ReqGetLatestTxsListByAccountIndex) (*RespGetLatestTxsListByAccountIndex, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTxsListByAccountIndex not implemented")
}
func (UnimplementedGlobalRPCServer) GetLatestTxsListByAccountIndexAndTxType(context.Context, *ReqGetLatestTxsListByAccountIndexAndTxType) (*RespGetLatestTxsListByAccountIndexAndTxType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTxsListByAccountIndexAndTxType not implemented")
}
func (UnimplementedGlobalRPCServer) SendTx(context.Context, *ReqSendTx) (*RespSendTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedGlobalRPCServer) mustEmbedUnimplementedGlobalRPCServer() {}

// UnsafeGlobalRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalRPCServer will
// result in compilation errors.
type UnsafeGlobalRPCServer interface {
	mustEmbedUnimplementedGlobalRPCServer()
}

func RegisterGlobalRPCServer(s grpc.ServiceRegistrar, srv GlobalRPCServer) {
	s.RegisterService(&GlobalRPC_ServiceDesc, srv)
}

func _GlobalRPC_GetLatestAccountLp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestAccountLp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestAccountLp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestAccountLp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestAccountLp(ctx, req.(*ReqGetLatestAccountLp))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLatestAssetsListByAccountIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestAssetsListByAccountIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestAssetsListByAccountIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestAssetsListByAccountIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestAssetsListByAccountIndex(ctx, req.(*ReqGetLatestAssetsListByAccountIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLatestAssetInfoByAccountIndexAndAssetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestAssetInfoByAccountIndexAndAssetId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestAssetInfoByAccountIndexAndAssetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestAssetInfoByAccountIndexAndAssetId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestAssetInfoByAccountIndexAndAssetId(ctx, req.(*ReqGetLatestAssetInfoByAccountIndexAndAssetId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLatestPairInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestPairInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestPairInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestPairInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestPairInfo(ctx, req.(*ReqGetLatestPairInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetSwapAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetSwapAmount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetSwapAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getSwapAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetSwapAmount(ctx, req.(*ReqGetSwapAmount))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLpValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLpValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLpValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLpValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLpValue(ctx, req.(*ReqGetLpValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLatestTxsListByAccountIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestTxsListByAccountIndex)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestTxsListByAccountIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestTxsListByAccountIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestTxsListByAccountIndex(ctx, req.(*ReqGetLatestTxsListByAccountIndex))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_GetLatestTxsListByAccountIndexAndTxType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetLatestTxsListByAccountIndexAndTxType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).GetLatestTxsListByAccountIndexAndTxType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/getLatestTxsListByAccountIndexAndTxType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).GetLatestTxsListByAccountIndexAndTxType(ctx, req.(*ReqGetLatestTxsListByAccountIndexAndTxType))
	}
	return interceptor(ctx, in, info, handler)
}

func _GlobalRPC_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSendTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalRPCServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/globalRPCProto.globalRPC/sendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalRPCServer).SendTx(ctx, req.(*ReqSendTx))
	}
	return interceptor(ctx, in, info, handler)
}

// GlobalRPC_ServiceDesc is the grpc.ServiceDesc for GlobalRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "globalRPCProto.globalRPC",
	HandlerType: (*GlobalRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getLatestAccountLp",
			Handler:    _GlobalRPC_GetLatestAccountLp_Handler,
		},
		{
			MethodName: "getLatestAssetsListByAccountIndex",
			Handler:    _GlobalRPC_GetLatestAssetsListByAccountIndex_Handler,
		},
		{
			MethodName: "getLatestAssetInfoByAccountIndexAndAssetId",
			Handler:    _GlobalRPC_GetLatestAssetInfoByAccountIndexAndAssetId_Handler,
		},
		{
			MethodName: "getLatestPairInfo",
			Handler:    _GlobalRPC_GetLatestPairInfo_Handler,
		},
		{
			MethodName: "getSwapAmount",
			Handler:    _GlobalRPC_GetSwapAmount_Handler,
		},
		{
			MethodName: "getLpValue",
			Handler:    _GlobalRPC_GetLpValue_Handler,
		},
		{
			MethodName: "getLatestTxsListByAccountIndex",
			Handler:    _GlobalRPC_GetLatestTxsListByAccountIndex_Handler,
		},
		{
			MethodName: "getLatestTxsListByAccountIndexAndTxType",
			Handler:    _GlobalRPC_GetLatestTxsListByAccountIndexAndTxType_Handler,
		},
		{
			MethodName: "sendTx",
			Handler:    _GlobalRPC_SendTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "globalRPC.proto",
}
