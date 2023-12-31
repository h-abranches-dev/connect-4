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
	Route_VerifyCompatibility_FullMethodName = "/gameengine.Route/VerifyCompatibility"
	Route_Connect_FullMethodName             = "/gameengine.Route/Connect"
	Route_POL_FullMethodName                 = "/gameengine.Route/POL"
	Route_ServeBoard_FullMethodName          = "/gameengine.Route/ServeBoard"
	Route_Play_FullMethodName                = "/gameengine.Route/Play"
)

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteClient interface {
	VerifyCompatibility(ctx context.Context, in *VerifyCompatibilityPayload, opts ...grpc.CallOption) (*VerifyCompatibilityResponse, error)
	Connect(ctx context.Context, in *ConnectPayload, opts ...grpc.CallOption) (*ConnectResponse, error)
	POL(ctx context.Context, in *POLPayload, opts ...grpc.CallOption) (*POLResponse, error)
	ServeBoard(ctx context.Context, in *ServeBoardPayload, opts ...grpc.CallOption) (*ServeBoardResponse, error)
	Play(ctx context.Context, in *PlayPayload, opts ...grpc.CallOption) (*PlayResponse, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) VerifyCompatibility(ctx context.Context, in *VerifyCompatibilityPayload, opts ...grpc.CallOption) (*VerifyCompatibilityResponse, error) {
	out := new(VerifyCompatibilityResponse)
	err := c.cc.Invoke(ctx, Route_VerifyCompatibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeClient) Connect(ctx context.Context, in *ConnectPayload, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, Route_Connect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeClient) POL(ctx context.Context, in *POLPayload, opts ...grpc.CallOption) (*POLResponse, error) {
	out := new(POLResponse)
	err := c.cc.Invoke(ctx, Route_POL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeClient) ServeBoard(ctx context.Context, in *ServeBoardPayload, opts ...grpc.CallOption) (*ServeBoardResponse, error) {
	out := new(ServeBoardResponse)
	err := c.cc.Invoke(ctx, Route_ServeBoard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeClient) Play(ctx context.Context, in *PlayPayload, opts ...grpc.CallOption) (*PlayResponse, error) {
	out := new(PlayResponse)
	err := c.cc.Invoke(ctx, Route_Play_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServer is the server API for Route service.
// All implementations must embed UnimplementedRouteServer
// for forward compatibility
type RouteServer interface {
	VerifyCompatibility(context.Context, *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error)
	Connect(context.Context, *ConnectPayload) (*ConnectResponse, error)
	POL(context.Context, *POLPayload) (*POLResponse, error)
	ServeBoard(context.Context, *ServeBoardPayload) (*ServeBoardResponse, error)
	Play(context.Context, *PlayPayload) (*PlayResponse, error)
	mustEmbedUnimplementedRouteServer()
}

// UnimplementedRouteServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (UnimplementedRouteServer) VerifyCompatibility(context.Context, *VerifyCompatibilityPayload) (*VerifyCompatibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCompatibility not implemented")
}
func (UnimplementedRouteServer) Connect(context.Context, *ConnectPayload) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedRouteServer) POL(context.Context, *POLPayload) (*POLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method POL not implemented")
}
func (UnimplementedRouteServer) ServeBoard(context.Context, *ServeBoardPayload) (*ServeBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeBoard not implemented")
}
func (UnimplementedRouteServer) Play(context.Context, *PlayPayload) (*PlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
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

func _Route_VerifyCompatibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCompatibilityPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).VerifyCompatibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_VerifyCompatibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).VerifyCompatibility(ctx, req.(*VerifyCompatibilityPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).Connect(ctx, req.(*ConnectPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route_POL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(POLPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).POL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_POL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).POL(ctx, req.(*POLPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route_ServeBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServeBoardPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).ServeBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_ServeBoard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).ServeBoard(ctx, req.(*ServeBoardPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_Play_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).Play(ctx, req.(*PlayPayload))
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
			MethodName: "VerifyCompatibility",
			Handler:    _Route_VerifyCompatibility_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Route_Connect_Handler,
		},
		{
			MethodName: "POL",
			Handler:    _Route_POL_Handler,
		},
		{
			MethodName: "ServeBoard",
			Handler:    _Route_ServeBoard_Handler,
		},
		{
			MethodName: "Play",
			Handler:    _Route_Play_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game-engine/game_engine.proto",
}
