// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: data_exchange.proto

package data_exchange

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

// GrpcServiceClient is the client API for GrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcServiceClient interface {
	UploadData(ctx context.Context, opts ...grpc.CallOption) (GrpcService_UploadDataClient, error)
	GetData(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (GrpcService_GetDataClient, error)
}

type grpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServiceClient(cc grpc.ClientConnInterface) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) UploadData(ctx context.Context, opts ...grpc.CallOption) (GrpcService_UploadDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &GrpcService_ServiceDesc.Streams[0], "/data_exchange.GrpcService/UploadData", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServiceUploadDataClient{stream}
	return x, nil
}

type GrpcService_UploadDataClient interface {
	Send(*Document) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type grpcServiceUploadDataClient struct {
	grpc.ClientStream
}

func (x *grpcServiceUploadDataClient) Send(m *Document) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcServiceUploadDataClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcServiceClient) GetData(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (GrpcService_GetDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &GrpcService_ServiceDesc.Streams[1], "/data_exchange.GrpcService/GetData", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServiceGetDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcService_GetDataClient interface {
	Recv() (*Document, error)
	grpc.ClientStream
}

type grpcServiceGetDataClient struct {
	grpc.ClientStream
}

func (x *grpcServiceGetDataClient) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GrpcServiceServer is the server API for GrpcService service.
// All implementations must embed UnimplementedGrpcServiceServer
// for forward compatibility
type GrpcServiceServer interface {
	UploadData(GrpcService_UploadDataServer) error
	GetData(*Pagination, GrpcService_GetDataServer) error
	mustEmbedUnimplementedGrpcServiceServer()
}

// UnimplementedGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServiceServer struct {
}

func (UnimplementedGrpcServiceServer) UploadData(GrpcService_UploadDataServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadData not implemented")
}
func (UnimplementedGrpcServiceServer) GetData(*Pagination, GrpcService_GetDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGrpcServiceServer) mustEmbedUnimplementedGrpcServiceServer() {}

// UnsafeGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServiceServer will
// result in compilation errors.
type UnsafeGrpcServiceServer interface {
	mustEmbedUnimplementedGrpcServiceServer()
}

func RegisterGrpcServiceServer(s grpc.ServiceRegistrar, srv GrpcServiceServer) {
	s.RegisterService(&GrpcService_ServiceDesc, srv)
}

func _GrpcService_UploadData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcServiceServer).UploadData(&grpcServiceUploadDataServer{stream})
}

type GrpcService_UploadDataServer interface {
	SendAndClose(*Status) error
	Recv() (*Document, error)
	grpc.ServerStream
}

type grpcServiceUploadDataServer struct {
	grpc.ServerStream
}

func (x *grpcServiceUploadDataServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcServiceUploadDataServer) Recv() (*Document, error) {
	m := new(Document)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GrpcService_GetData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Pagination)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcServiceServer).GetData(m, &grpcServiceGetDataServer{stream})
}

type GrpcService_GetDataServer interface {
	Send(*Document) error
	grpc.ServerStream
}

type grpcServiceGetDataServer struct {
	grpc.ServerStream
}

func (x *grpcServiceGetDataServer) Send(m *Document) error {
	return x.ServerStream.SendMsg(m)
}

// GrpcService_ServiceDesc is the grpc.ServiceDesc for GrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data_exchange.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadData",
			Handler:       _GrpcService_UploadData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetData",
			Handler:       _GrpcService_GetData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data_exchange.proto",
}
