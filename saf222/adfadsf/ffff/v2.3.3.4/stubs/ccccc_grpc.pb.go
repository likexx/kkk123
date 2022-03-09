// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ffff

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

// CccccClient is the client API for Ccccc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CccccClient interface {
	TestHello12345678(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Test223(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type cccccClient struct {
	cc grpc.ClientConnInterface
}

func NewCccccClient(cc grpc.ClientConnInterface) CccccClient {
	return &cccccClient{cc}
}

func (c *cccccClient) TestHello12345678(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/adfadsf.ffff.ccccc/TestHello12345678", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cccccClient) Test223(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/adfadsf.ffff.ccccc/Test223", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CccccServer is the server API for Ccccc service.
// All implementations must embed UnimplementedCccccServer
// for forward compatibility
type CccccServer interface {
	TestHello12345678(context.Context, *HelloRequest) (*HelloReply, error)
	Test223(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedCccccServer()
}

// UnimplementedCccccServer must be embedded to have forward compatible implementations.
type UnimplementedCccccServer struct {
}

func (UnimplementedCccccServer) TestHello12345678(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestHello12345678 not implemented")
}
func (UnimplementedCccccServer) Test223(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test223 not implemented")
}
func (UnimplementedCccccServer) mustEmbedUnimplementedCccccServer() {}

// UnsafeCccccServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CccccServer will
// result in compilation errors.
type UnsafeCccccServer interface {
	mustEmbedUnimplementedCccccServer()
}

func RegisterCccccServer(s grpc.ServiceRegistrar, srv CccccServer) {
	s.RegisterService(&Ccccc_ServiceDesc, srv)
}

func _Ccccc_TestHello12345678_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CccccServer).TestHello12345678(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adfadsf.ffff.ccccc/TestHello12345678",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CccccServer).TestHello12345678(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ccccc_Test223_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CccccServer).Test223(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adfadsf.ffff.ccccc/Test223",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CccccServer).Test223(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ccccc_ServiceDesc is the grpc.ServiceDesc for Ccccc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ccccc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adfadsf.ffff.ccccc",
	HandlerType: (*CccccServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestHello12345678",
			Handler:    _Ccccc_TestHello12345678_Handler,
		},
		{
			MethodName: "Test223",
			Handler:    _Ccccc_Test223_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ccccc.proto",
}
