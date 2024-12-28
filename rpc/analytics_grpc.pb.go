// Copyright 2024 Xtressials Corporation, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: rpc/analytics.proto

package rpc

import (
	context "context"
	wirtual "github.com/wirtual/protocol/wirtual"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AnalyticsRecorderService_IngestStats_FullMethodName          = "/wirtual.AnalyticsRecorderService/IngestStats"
	AnalyticsRecorderService_IngestEvents_FullMethodName         = "/wirtual.AnalyticsRecorderService/IngestEvents"
	AnalyticsRecorderService_IngestNodeRoomStates_FullMethodName = "/wirtual.AnalyticsRecorderService/IngestNodeRoomStates"
)

// AnalyticsRecorderServiceClient is the client API for AnalyticsRecorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsRecorderServiceClient interface {
	IngestStats(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsStats, emptypb.Empty], error)
	IngestEvents(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsEvents, emptypb.Empty], error)
	IngestNodeRoomStates(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsNodeRooms, emptypb.Empty], error)
}

type analyticsRecorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsRecorderServiceClient(cc grpc.ClientConnInterface) AnalyticsRecorderServiceClient {
	return &analyticsRecorderServiceClient{cc}
}

func (c *analyticsRecorderServiceClient) IngestStats(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsStats, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[0], AnalyticsRecorderService_IngestStats_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wirtual.AnalyticsStats, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestStatsClient = grpc.ClientStreamingClient[wirtual.AnalyticsStats, emptypb.Empty]

func (c *analyticsRecorderServiceClient) IngestEvents(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsEvents, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[1], AnalyticsRecorderService_IngestEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wirtual.AnalyticsEvents, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestEventsClient = grpc.ClientStreamingClient[wirtual.AnalyticsEvents, emptypb.Empty]

func (c *analyticsRecorderServiceClient) IngestNodeRoomStates(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wirtual.AnalyticsNodeRooms, emptypb.Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnalyticsRecorderService_ServiceDesc.Streams[2], AnalyticsRecorderService_IngestNodeRoomStates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wirtual.AnalyticsNodeRooms, emptypb.Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestNodeRoomStatesClient = grpc.ClientStreamingClient[wirtual.AnalyticsNodeRooms, emptypb.Empty]

// AnalyticsRecorderServiceServer is the server API for AnalyticsRecorderService service.
// All implementations must embed UnimplementedAnalyticsRecorderServiceServer
// for forward compatibility.
type AnalyticsRecorderServiceServer interface {
	IngestStats(grpc.ClientStreamingServer[wirtual.AnalyticsStats, emptypb.Empty]) error
	IngestEvents(grpc.ClientStreamingServer[wirtual.AnalyticsEvents, emptypb.Empty]) error
	IngestNodeRoomStates(grpc.ClientStreamingServer[wirtual.AnalyticsNodeRooms, emptypb.Empty]) error
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

// UnimplementedAnalyticsRecorderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnalyticsRecorderServiceServer struct{}

func (UnimplementedAnalyticsRecorderServiceServer) IngestStats(grpc.ClientStreamingServer[wirtual.AnalyticsStats, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method IngestStats not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) IngestEvents(grpc.ClientStreamingServer[wirtual.AnalyticsEvents, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method IngestEvents not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) IngestNodeRoomStates(grpc.ClientStreamingServer[wirtual.AnalyticsNodeRooms, emptypb.Empty]) error {
	return status.Errorf(codes.Unimplemented, "method IngestNodeRoomStates not implemented")
}
func (UnimplementedAnalyticsRecorderServiceServer) mustEmbedUnimplementedAnalyticsRecorderServiceServer() {
}
func (UnimplementedAnalyticsRecorderServiceServer) testEmbeddedByValue() {}

// UnsafeAnalyticsRecorderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsRecorderServiceServer will
// result in compilation errors.
type UnsafeAnalyticsRecorderServiceServer interface {
	mustEmbedUnimplementedAnalyticsRecorderServiceServer()
}

func RegisterAnalyticsRecorderServiceServer(s grpc.ServiceRegistrar, srv AnalyticsRecorderServiceServer) {
	// If the following call pancis, it indicates UnimplementedAnalyticsRecorderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnalyticsRecorderService_ServiceDesc, srv)
}

func _AnalyticsRecorderService_IngestStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestStats(&grpc.GenericServerStream[wirtual.AnalyticsStats, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestStatsServer = grpc.ClientStreamingServer[wirtual.AnalyticsStats, emptypb.Empty]

func _AnalyticsRecorderService_IngestEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestEvents(&grpc.GenericServerStream[wirtual.AnalyticsEvents, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestEventsServer = grpc.ClientStreamingServer[wirtual.AnalyticsEvents, emptypb.Empty]

func _AnalyticsRecorderService_IngestNodeRoomStates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnalyticsRecorderServiceServer).IngestNodeRoomStates(&grpc.GenericServerStream[wirtual.AnalyticsNodeRooms, emptypb.Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnalyticsRecorderService_IngestNodeRoomStatesServer = grpc.ClientStreamingServer[wirtual.AnalyticsNodeRooms, emptypb.Empty]

// AnalyticsRecorderService_ServiceDesc is the grpc.ServiceDesc for AnalyticsRecorderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsRecorderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wirtual.AnalyticsRecorderService",
	HandlerType: (*AnalyticsRecorderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IngestStats",
			Handler:       _AnalyticsRecorderService_IngestStats_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IngestEvents",
			Handler:       _AnalyticsRecorderService_IngestEvents_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "IngestNodeRoomStates",
			Handler:       _AnalyticsRecorderService_IngestNodeRoomStates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/analytics.proto",
}
