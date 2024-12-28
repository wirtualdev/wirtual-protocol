// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/roommanager.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit1 "github.com/wirtual/protocol/wirtual"
import livekit6 "github.com/wirtual/protocol/wirtual"

var _ = version.PsrpcVersion_0_6

// ============================
// RoomManager Client Interface
// ============================

type RoomManagerClient[NodeIdTopicType ~string] interface {
	CreateRoom(ctx context.Context, nodeId NodeIdTopicType, req *livekit6.CreateRoomRequest, opts ...psrpc.RequestOption) (*livekit1.Room, error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ================================
// RoomManager ServerImpl Interface
// ================================

type RoomManagerServerImpl interface {
	CreateRoom(context.Context, *livekit6.CreateRoomRequest) (*livekit1.Room, error)
}

// ============================
// RoomManager Server Interface
// ============================

type RoomManagerServer[NodeIdTopicType ~string] interface {
	RegisterCreateRoomTopic(nodeId NodeIdTopicType) error
	DeregisterCreateRoomTopic(nodeId NodeIdTopicType)
	RegisterAllNodeTopics(nodeId NodeIdTopicType) error
	DeregisterAllNodeTopics(nodeId NodeIdTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ==================
// RoomManager Client
// ==================

type roomManagerClient[NodeIdTopicType ~string] struct {
	client *client.RPCClient
}

// NewRoomManagerClient creates a psrpc client that implements the RoomManagerClient interface.
func NewRoomManagerClient[NodeIdTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (RoomManagerClient[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "RoomManager",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CreateRoom", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &roomManagerClient[NodeIdTopicType]{
		client: rpcClient,
	}, nil
}

func (c *roomManagerClient[NodeIdTopicType]) CreateRoom(ctx context.Context, nodeId NodeIdTopicType, req *livekit6.CreateRoomRequest, opts ...psrpc.RequestOption) (*livekit1.Room, error) {
	return client.RequestSingle[*livekit1.Room](ctx, c.client, "CreateRoom", []string{string(nodeId)}, req, opts...)
}

func (s *roomManagerClient[NodeIdTopicType]) Close() {
	s.client.Close()
}

// ==================
// RoomManager Server
// ==================

type roomManagerServer[NodeIdTopicType ~string] struct {
	svc RoomManagerServerImpl
	rpc *server.RPCServer
}

// NewRoomManagerServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewRoomManagerServer[NodeIdTopicType ~string](svc RoomManagerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (RoomManagerServer[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "RoomManager",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CreateRoom", false, false, true, true)
	return &roomManagerServer[NodeIdTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *roomManagerServer[NodeIdTopicType]) RegisterCreateRoomTopic(nodeId NodeIdTopicType) error {
	return server.RegisterHandler(s.rpc, "CreateRoom", []string{string(nodeId)}, s.svc.CreateRoom, nil)
}

func (s *roomManagerServer[NodeIdTopicType]) DeregisterCreateRoomTopic(nodeId NodeIdTopicType) {
	s.rpc.DeregisterHandler("CreateRoom", []string{string(nodeId)})
}

func (s *roomManagerServer[NodeIdTopicType]) allNodeTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterCreateRoomTopic, s.DeregisterCreateRoomTopic),
	}
}

func (s *roomManagerServer[NodeIdTopicType]) RegisterAllNodeTopics(nodeId NodeIdTopicType) error {
	return s.allNodeTopicRegisterers().Register(nodeId)
}

func (s *roomManagerServer[NodeIdTopicType]) DeregisterAllNodeTopics(nodeId NodeIdTopicType) {
	s.allNodeTopicRegisterers().Deregister(nodeId)
}

func (s *roomManagerServer[NodeIdTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *roomManagerServer[NodeIdTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor8 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcd, 0x4d, 0x0a, 0xc2, 0x30,
	0x10, 0x05, 0x60, 0x8a, 0xa2, 0x10, 0x29, 0x68, 0xb0, 0xa0, 0xd9, 0x88, 0x1e, 0x20, 0x01, 0xbd,
	0x81, 0xae, 0xdd, 0x74, 0xe9, 0xa6, 0xa4, 0xe9, 0x50, 0x83, 0x4d, 0x27, 0xa6, 0xa9, 0x77, 0xf0,
	0x3a, 0x9e, 0x50, 0xd2, 0x1f, 0xdd, 0x0c, 0xcc, 0x37, 0xbc, 0x37, 0x24, 0x71, 0x56, 0x09, 0x87,
	0x68, 0x8c, 0xac, 0x65, 0x09, 0x8e, 0x5b, 0x87, 0x1e, 0xe9, 0xc4, 0x59, 0xc5, 0x62, 0xb4, 0x5e,
	0x63, 0xdd, 0xf4, 0xc6, 0xd6, 0x95, 0x7e, 0xc1, 0x43, 0xfb, 0xcc, 0x60, 0x01, 0xd5, 0xa8, 0x74,
	0xd4, 0x50, 0xd2, 0xdb, 0x51, 0x92, 0x45, 0x8a, 0x68, 0xae, 0x7d, 0x25, 0x4d, 0x09, 0xb9, 0x38,
	0x90, 0x1e, 0x02, 0x52, 0xc6, 0x87, 0x04, 0xff, 0x63, 0x0a, 0xcf, 0x16, 0x1a, 0xcf, 0xe2, 0xdf,
	0x2d, 0xe8, 0x61, 0xfb, 0x79, 0x47, 0xc9, 0x32, 0x62, 0x2b, 0x32, 0xad, 0xb1, 0x00, 0x3a, 0x0f,
	0x33, 0xd3, 0xc5, 0x26, 0x3a, 0xef, 0x6f, 0xbb, 0x52, 0xfb, 0x7b, 0x9b, 0x73, 0x85, 0x46, 0x0c,
	0x29, 0xd1, 0xbd, 0x57, 0x58, 0x09, 0x67, 0x55, 0x3e, 0xeb, 0xb6, 0xd3, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0xa9, 0xd5, 0x75, 0xe3, 0x00, 0x00, 0x00,
}
