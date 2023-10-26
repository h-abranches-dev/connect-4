// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: game-engine/game_engine.proto

package gameengine

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
	Route_Ping_FullMethodName = "/gameengine.Route/Ping"
)

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteClient interface {
	Ping(ctx context.Context, in *PingPayload, opts ...grpc.CallOption) (*PingResponse, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) Ping(ctx context.Context, in *PingPayload, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Route_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServer is the server API for Route service.
// All implementations must embed UnimplementedRouteServer
// for forward compatibility
type RouteServer interface {
	Ping(context.Context, *PingPayload) (*PingResponse, error)
	mustEmbedUnimplementedRouteServer()
}

// UnimplementedRouteServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (UnimplementedRouteServer) Ping(context.Context, *PingPayload) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRouteServer) mustEmbedUnimplementedRouteServer() {}

// UnsafeRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServer will
// result in compilation errors.
type UnsafeRouteServer interface {
	mustEmbedUnimplementedRouteServer()
}

func RegisterRouteServer(s grpc.ServiceRegistrar, srv RouteServer) {
	s.RegisterService(&Route_ServiceDesc, srv)
}

func _Route_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).Ping(ctx, req.(*PingPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// Route_ServiceDesc is the grpc.ServiceDesc for Route service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gameengine.Route",
	HandlerType: (*RouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Route_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game-engine/game_engine.proto",
}
