// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_rebelion_grpc

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

// InformantesClient is the client API for Informantes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InformantesClient interface {
	AddCity(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Respuesta, error)
	UpdateName(ctx context.Context, in *InfoUpdateName, opts ...grpc.CallOption) (*Respuesta, error)
	UpdateNumber(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Respuesta, error)
	DeleteCity(ctx context.Context, in *InfoDelete, opts ...grpc.CallOption) (*Respuesta, error)
}

type informantesClient struct {
	cc grpc.ClientConnInterface
}

func NewInformantesClient(cc grpc.ClientConnInterface) InformantesClient {
	return &informantesClient{cc}
}

func (c *informantesClient) AddCity(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/informantes.Informantes/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informantesClient) UpdateName(ctx context.Context, in *InfoUpdateName, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/informantes.Informantes/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informantesClient) UpdateNumber(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/informantes.Informantes/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *informantesClient) DeleteCity(ctx context.Context, in *InfoDelete, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/informantes.Informantes/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InformantesServer is the server API for Informantes service.
// All implementations must embed UnimplementedInformantesServer
// for forward compatibility
type InformantesServer interface {
	AddCity(context.Context, *Info) (*Respuesta, error)
	UpdateName(context.Context, *InfoUpdateName) (*Respuesta, error)
	UpdateNumber(context.Context, *Info) (*Respuesta, error)
	DeleteCity(context.Context, *InfoDelete) (*Respuesta, error)
	mustEmbedUnimplementedInformantesServer()
}

// UnimplementedInformantesServer must be embedded to have forward compatible implementations.
type UnimplementedInformantesServer struct {
}

func (UnimplementedInformantesServer) AddCity(context.Context, *Info) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedInformantesServer) UpdateName(context.Context, *InfoUpdateName) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedInformantesServer) UpdateNumber(context.Context, *Info) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedInformantesServer) DeleteCity(context.Context, *InfoDelete) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedInformantesServer) mustEmbedUnimplementedInformantesServer() {}

// UnsafeInformantesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InformantesServer will
// result in compilation errors.
type UnsafeInformantesServer interface {
	mustEmbedUnimplementedInformantesServer()
}

func RegisterInformantesServer(s grpc.ServiceRegistrar, srv InformantesServer) {
	s.RegisterService(&Informantes_ServiceDesc, srv)
}

func _Informantes_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformantesServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informantes.Informantes/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformantesServer).AddCity(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

func _Informantes_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUpdateName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformantesServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informantes.Informantes/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformantesServer).UpdateName(ctx, req.(*InfoUpdateName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Informantes_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformantesServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informantes.Informantes/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformantesServer).UpdateNumber(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

func _Informantes_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformantesServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informantes.Informantes/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformantesServer).DeleteCity(ctx, req.(*InfoDelete))
	}
	return interceptor(ctx, in, info, handler)
}

// Informantes_ServiceDesc is the grpc.ServiceDesc for Informantes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Informantes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "informantes.Informantes",
	HandlerType: (*InformantesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _Informantes_AddCity_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _Informantes_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _Informantes_UpdateNumber_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _Informantes_DeleteCity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rebelion/informantes.proto",
}

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	SolicitarIP(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*IP, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) SolicitarIP(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*IP, error) {
	out := new(IP)
	err := c.cc.Invoke(ctx, "/informantes.Broker/SolicitarIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility
type BrokerServer interface {
	SolicitarIP(context.Context, *Comando) (*IP, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (UnimplementedBrokerServer) SolicitarIP(context.Context, *Comando) (*IP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolicitarIP not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_SolicitarIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).SolicitarIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/informantes.Broker/SolicitarIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).SolicitarIP(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "informantes.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolicitarIP",
			Handler:    _Broker_SolicitarIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rebelion/informantes.proto",
}