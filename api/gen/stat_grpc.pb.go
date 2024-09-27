// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: api/stat.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StatService_GetStats_FullMethodName = "/stat.StatService/GetStats"
)

// StatServiceClient is the client API for StatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatServiceClient interface {
	GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StatsResponse], error)
}

type statServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatServiceClient(cc grpc.ClientConnInterface) StatServiceClient {
	return &statServiceClient{cc}
}

func (c *statServiceClient) GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StatsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StatService_ServiceDesc.Streams[0], StatService_GetStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StatsRequest, StatsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StatService_GetStatsClient = grpc.ServerStreamingClient[StatsResponse]

// StatServiceServer is the server API for StatService service.
// All implementations must embed UnimplementedStatServiceServer
// for forward compatibility.
type StatServiceServer interface {
	GetStats(*StatsRequest, grpc.ServerStreamingServer[StatsResponse]) error
	mustEmbedUnimplementedStatServiceServer()
}

// UnimplementedStatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatServiceServer struct{}

func (UnimplementedStatServiceServer) GetStats(*StatsRequest, grpc.ServerStreamingServer[StatsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedStatServiceServer) mustEmbedUnimplementedStatServiceServer() {}
func (UnimplementedStatServiceServer) testEmbeddedByValue()                     {}

// UnsafeStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatServiceServer will
// result in compilation errors.
type UnsafeStatServiceServer interface {
	mustEmbedUnimplementedStatServiceServer()
}

func RegisterStatServiceServer(s grpc.ServiceRegistrar, srv StatServiceServer) {
	// If the following call pancis, it indicates UnimplementedStatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StatService_ServiceDesc, srv)
}

func _StatService_GetStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatServiceServer).GetStats(m, &grpc.GenericServerStream[StatsRequest, StatsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StatService_GetStatsServer = grpc.ServerStreamingServer[StatsResponse]

// StatService_ServiceDesc is the grpc.ServiceDesc for StatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stat.StatService",
	HandlerType: (*StatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStats",
			Handler:       _StatService_GetStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/stat.proto",
}
