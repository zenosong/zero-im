// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: rpc/zero.proto

package zero

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

// ZeroClient is the client API for Zero service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZeroClient interface {
	// Base service
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	// Admin service
	AdminList(ctx context.Context, in *AdminListReq, opts ...grpc.CallOption) (*AdminListResp, error)
	AdminCreate(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	AdminUpdate(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseResp, error)
	AdminDelete(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseResp, error)
	AdminDetail(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*AdminInfo, error)
}

type zeroClient struct {
	cc grpc.ClientConnInterface
}

func NewZeroClient(cc grpc.ClientConnInterface) ZeroClient {
	return &zeroClient{cc}
}

func (c *zeroClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/zero.Zero/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/zero.Zero/InitDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) AdminList(ctx context.Context, in *AdminListReq, opts ...grpc.CallOption) (*AdminListResp, error) {
	out := new(AdminListResp)
	err := c.cc.Invoke(ctx, "/zero.Zero/AdminList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) AdminCreate(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, "/zero.Zero/AdminCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) AdminUpdate(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/zero.Zero/AdminUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) AdminDelete(ctx context.Context, in *AdminInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/zero.Zero/AdminDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeroClient) AdminDetail(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*AdminInfo, error) {
	out := new(AdminInfo)
	err := c.cc.Invoke(ctx, "/zero.Zero/AdminDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZeroServer is the server API for Zero service.
// All implementations must embed UnimplementedZeroServer
// for forward compatibility
type ZeroServer interface {
	// Base service
	Ping(context.Context, *Request) (*Response, error)
	InitDatabase(context.Context, *Empty) (*BaseResp, error)
	// Admin service
	AdminList(context.Context, *AdminListReq) (*AdminListResp, error)
	AdminCreate(context.Context, *AdminInfo) (*BaseIDResp, error)
	AdminUpdate(context.Context, *AdminInfo) (*BaseResp, error)
	AdminDelete(context.Context, *AdminInfo) (*BaseResp, error)
	AdminDetail(context.Context, *IDReq) (*AdminInfo, error)
	mustEmbedUnimplementedZeroServer()
}

// UnimplementedZeroServer must be embedded to have forward compatible implementations.
type UnimplementedZeroServer struct {
}

func (UnimplementedZeroServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedZeroServer) InitDatabase(context.Context, *Empty) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDatabase not implemented")
}
func (UnimplementedZeroServer) AdminList(context.Context, *AdminListReq) (*AdminListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminList not implemented")
}
func (UnimplementedZeroServer) AdminCreate(context.Context, *AdminInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreate not implemented")
}
func (UnimplementedZeroServer) AdminUpdate(context.Context, *AdminInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdate not implemented")
}
func (UnimplementedZeroServer) AdminDelete(context.Context, *AdminInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDelete not implemented")
}
func (UnimplementedZeroServer) AdminDetail(context.Context, *IDReq) (*AdminInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDetail not implemented")
}
func (UnimplementedZeroServer) mustEmbedUnimplementedZeroServer() {}

// UnsafeZeroServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZeroServer will
// result in compilation errors.
type UnsafeZeroServer interface {
	mustEmbedUnimplementedZeroServer()
}

func RegisterZeroServer(s grpc.ServiceRegistrar, srv ZeroServer) {
	s.RegisterService(&Zero_ServiceDesc, srv)
}

func _Zero_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_InitDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).InitDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/InitDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).InitDatabase(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_AdminList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).AdminList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/AdminList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).AdminList(ctx, req.(*AdminListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_AdminCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).AdminCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/AdminCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).AdminCreate(ctx, req.(*AdminInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_AdminUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).AdminUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/AdminUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).AdminUpdate(ctx, req.(*AdminInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_AdminDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).AdminDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/AdminDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).AdminDelete(ctx, req.(*AdminInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zero_AdminDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeroServer).AdminDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zero.Zero/AdminDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeroServer).AdminDetail(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Zero_ServiceDesc is the grpc.ServiceDesc for Zero service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zero_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zero.Zero",
	HandlerType: (*ZeroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Zero_Ping_Handler,
		},
		{
			MethodName: "InitDatabase",
			Handler:    _Zero_InitDatabase_Handler,
		},
		{
			MethodName: "AdminList",
			Handler:    _Zero_AdminList_Handler,
		},
		{
			MethodName: "AdminCreate",
			Handler:    _Zero_AdminCreate_Handler,
		},
		{
			MethodName: "AdminUpdate",
			Handler:    _Zero_AdminUpdate_Handler,
		},
		{
			MethodName: "AdminDelete",
			Handler:    _Zero_AdminDelete_Handler,
		},
		{
			MethodName: "AdminDetail",
			Handler:    _Zero_AdminDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/zero.proto",
}
