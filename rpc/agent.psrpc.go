// Code generated by protoc-gen-psrpc v0.6.0, DO NOT EDIT.
// source: rpc/agent.proto

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
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"
import livekit2 "github.com/wirtual/protocol/wirtual"

var _ = version.PsrpcVersion_0_6

// ==============================
// AgentInternal Client Interface
// ==============================

type AgentInternalClient interface {
	CheckEnabled(ctx context.Context, req *CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*CheckEnabledResponse], error)

	JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error)

	JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error)

	SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error)

	// Close immediately, without waiting for pending RPCs
	Close()
}

// ==================================
// AgentInternal ServerImpl Interface
// ==================================

type AgentInternalServerImpl interface {
	CheckEnabled(context.Context, *CheckEnabledRequest) (*CheckEnabledResponse, error)

	JobRequest(context.Context, *livekit2.Job) (*JobRequestResponse, error)
	JobRequestAffinity(context.Context, *livekit2.Job) float32

	JobTerminate(context.Context, *JobTerminateRequest) (*JobTerminateResponse, error)
}

// ==============================
// AgentInternal Server Interface
// ==============================

type AgentInternalServer interface {
	RegisterJobRequestTopic(namespace string, jobType string) error
	DeregisterJobRequestTopic(namespace string, jobType string)
	RegisterJobTerminateTopic(jobId string) error
	DeregisterJobTerminateTopic(jobId string)
	PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ====================
// AgentInternal Client
// ====================

type agentInternalClient struct {
	client *client.RPCClient
}

// NewAgentInternalClient creates a psrpc client that implements the AgentInternalClient interface.
func NewAgentInternalClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (AgentInternalClient, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("JobTerminate", false, false, true, true)
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &agentInternalClient{
		client: rpcClient,
	}, nil
}

func (c *agentInternalClient) CheckEnabled(ctx context.Context, req *CheckEnabledRequest, opts ...psrpc.RequestOption) (<-chan *psrpc.Response[*CheckEnabledResponse], error) {
	return client.RequestMulti[*CheckEnabledResponse](ctx, c.client, "CheckEnabled", nil, req, opts...)
}

func (c *agentInternalClient) JobRequest(ctx context.Context, namespace string, jobType string, req *livekit2.Job, opts ...psrpc.RequestOption) (*JobRequestResponse, error) {
	return client.RequestSingle[*JobRequestResponse](ctx, c.client, "JobRequest", []string{namespace, jobType}, req, opts...)
}

func (c *agentInternalClient) JobTerminate(ctx context.Context, jobId string, req *JobTerminateRequest, opts ...psrpc.RequestOption) (*JobTerminateResponse, error) {
	return client.RequestSingle[*JobTerminateResponse](ctx, c.client, "JobTerminate", []string{jobId}, req, opts...)
}

func (c *agentInternalClient) SubscribeWorkerRegistered(ctx context.Context, handlerNamespace string) (psrpc.Subscription[*google_protobuf.Empty], error) {
	return client.Join[*google_protobuf.Empty](ctx, c.client, "WorkerRegistered", []string{handlerNamespace})
}

func (s *agentInternalClient) Close() {
	s.client.Close()
}

// ====================
// AgentInternal Server
// ====================

type agentInternalServer struct {
	svc AgentInternalServerImpl
	rpc *server.RPCServer
}

// NewAgentInternalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewAgentInternalServer(svc AgentInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (AgentInternalServer, error) {
	sd := &info.ServiceDefinition{
		Name: "AgentInternal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("CheckEnabled", false, true, false, false)
	var err error
	err = server.RegisterHandler(s, "CheckEnabled", nil, svc.CheckEnabled, nil)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	sd.RegisterMethod("JobRequest", true, false, true, false)
	sd.RegisterMethod("JobTerminate", false, false, true, true)
	sd.RegisterMethod("WorkerRegistered", false, true, false, false)
	return &agentInternalServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *agentInternalServer) RegisterJobRequestTopic(namespace string, jobType string) error {
	return server.RegisterHandler(s.rpc, "JobRequest", []string{namespace, jobType}, s.svc.JobRequest, s.svc.JobRequestAffinity)
}

func (s *agentInternalServer) DeregisterJobRequestTopic(namespace string, jobType string) {
	s.rpc.DeregisterHandler("JobRequest", []string{namespace, jobType})
}

func (s *agentInternalServer) RegisterJobTerminateTopic(jobId string) error {
	return server.RegisterHandler(s.rpc, "JobTerminate", []string{jobId}, s.svc.JobTerminate, nil)
}

func (s *agentInternalServer) DeregisterJobTerminateTopic(jobId string) {
	s.rpc.DeregisterHandler("JobTerminate", []string{jobId})
}

func (s *agentInternalServer) PublishWorkerRegistered(ctx context.Context, handlerNamespace string, msg *google_protobuf.Empty) error {
	return s.rpc.Publish(ctx, "WorkerRegistered", []string{handlerNamespace}, msg)
}

func (s *agentInternalServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *agentInternalServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor0 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x8f, 0xd2, 0x4c,
	0x14, 0x7e, 0x87, 0xfd, 0x08, 0x1c, 0xd8, 0x77, 0xcb, 0xb0, 0x28, 0x5b, 0x63, 0x76, 0xe9, 0x8d,
	0x44, 0x93, 0xd6, 0xe0, 0xb5, 0x71, 0x59, 0xb7, 0x1a, 0x88, 0x40, 0xec, 0x76, 0x35, 0x31, 0x31,
	0x4d, 0x5b, 0x8e, 0xd0, 0xa5, 0x74, 0xc6, 0xe9, 0x60, 0xb2, 0x3f, 0x81, 0xdf, 0xe2, 0x1d, 0xbf,
	0xcd, 0x1f, 0x60, 0xfa, 0xb1, 0x15, 0x36, 0x78, 0xe1, 0xe5, 0x3c, 0xe7, 0xc9, 0xc9, 0xf3, 0x31,
	0x07, 0x8e, 0x05, 0xf7, 0x0d, 0x77, 0x8a, 0x91, 0xd4, 0xb9, 0x60, 0x92, 0xd1, 0x3d, 0xc1, 0x7d,
	0xf5, 0xc9, 0x94, 0xb1, 0x69, 0x88, 0x46, 0x0a, 0x79, 0xcb, 0x6f, 0x06, 0x2e, 0xb8, 0xbc, 0xcb,
	0x18, 0xea, 0x11, 0xe3, 0x32, 0x60, 0x51, 0x9c, 0x3f, 0x1b, 0x61, 0xf0, 0x03, 0xe7, 0x81, 0x74,
	0x36, 0xb6, 0x68, 0x4d, 0x68, 0xbc, 0x9d, 0xa1, 0x3f, 0x37, 0x23, 0xd7, 0x0b, 0x71, 0x62, 0xe1,
	0xf7, 0x25, 0xc6, 0x52, 0xfb, 0x49, 0xe0, 0x64, 0x1b, 0x8f, 0x39, 0x8b, 0x62, 0xa4, 0x6d, 0xa8,
	0x09, 0xc6, 0x16, 0x0e, 0x66, 0x78, 0x8b, 0x9c, 0x93, 0x4e, 0xd9, 0xaa, 0x26, 0x58, 0x4e, 0xa5,
	0x2f, 0xa0, 0xce, 0x97, 0x5e, 0x18, 0xc4, 0x33, 0x14, 0x05, 0xaf, 0x94, 0xf2, 0x94, 0x62, 0x70,
	0x4f, 0xd6, 0x00, 0x22, 0x77, 0x81, 0x31, 0x77, 0x7d, 0x8c, 0x5b, 0x7b, 0xe7, 0x7b, 0x9d, 0xca,
	0x65, 0xa9, 0x45, 0xac, 0x0d, 0x94, 0x9e, 0x41, 0x35, 0x95, 0xec, 0xa4, 0x58, 0x6b, 0x3f, 0x21,
	0x59, 0x90, 0x42, 0xa3, 0x04, 0xd1, 0x5e, 0x03, 0x1d, 0x30, 0x2f, 0xd7, 0x5e, 0x48, 0x7d, 0x06,
	0x07, 0xb1, 0x74, 0x25, 0xa6, 0x1a, 0xab, 0xdd, 0xba, 0x9e, 0xfb, 0xd7, 0x07, 0xcc, 0xbb, 0x4e,
	0x06, 0x56, 0x36, 0xd7, 0xbe, 0x42, 0x63, 0xc0, 0x3c, 0x1b, 0xc5, 0x22, 0x88, 0x12, 0x38, 0xdb,
	0x43, 0x9b, 0x70, 0x78, 0xcb, 0x3c, 0x27, 0xc8, 0x4c, 0x56, 0xac, 0x83, 0x5b, 0xe6, 0xf5, 0x27,
	0xd4, 0x80, 0x43, 0x81, 0x6e, 0xcc, 0xa2, 0xd4, 0xd3, 0xff, 0xdd, 0xc7, 0xba, 0xe0, 0xbe, 0xbe,
	0xbd, 0x20, 0x19, 0x5b, 0x39, 0x4d, 0x7b, 0x03, 0x27, 0xdb, 0xd3, 0x7f, 0xd4, 0xf7, 0xfc, 0x2a,
	0xb5, 0xf7, 0x60, 0x3d, 0x3d, 0x85, 0xa6, 0x6d, 0x5a, 0xc3, 0xfe, 0xa8, 0x67, 0xf7, 0xc7, 0x23,
	0xc7, 0x32, 0x3f, 0xde, 0x98, 0xd7, 0xb6, 0x79, 0xa5, 0xfc, 0x47, 0x1b, 0x70, 0xdc, 0x7b, 0x6f,
	0x8e, 0x6c, 0xe7, 0x83, 0xf9, 0xce, 0x76, 0xac, 0xf1, 0x78, 0xa8, 0x90, 0xee, 0xaf, 0x12, 0x1c,
	0xf5, 0x92, 0xcc, 0xfa, 0x91, 0x44, 0x11, 0xb9, 0x21, 0x1d, 0x42, 0x6d, 0xb3, 0x63, 0xda, 0x4a,
	0x9d, 0xec, 0xf8, 0x0e, 0xea, 0xe9, 0x8e, 0x49, 0xe6, 0x42, 0x2b, 0xaf, 0x57, 0x64, 0xff, 0xa2,
	0xd4, 0x21, 0xf4, 0x13, 0xc0, 0x9f, 0x16, 0x68, 0x6d, 0xd3, 0x8e, 0x5a, 0x84, 0xf4, 0xa0, 0x24,
	0xad, 0xbd, 0x5e, 0x91, 0xa7, 0x0a, 0x51, 0x9b, 0xb4, 0x52, 0x34, 0x4e, 0xcb, 0x49, 0xee, 0xf2,
	0x8e, 0xe3, 0x05, 0x79, 0x49, 0xe8, 0x0d, 0xd4, 0x36, 0xed, 0xe7, 0x32, 0x77, 0x34, 0x96, 0xcb,
	0xdc, 0x15, 0xb6, 0xa6, 0xac, 0x57, 0xa4, 0xa6, 0x10, 0xb5, 0x4c, 0xf3, 0x52, 0x29, 0x82, 0xf2,
	0x99, 0x89, 0x39, 0x0a, 0x0b, 0xa7, 0x41, 0x2c, 0x51, 0xe0, 0x84, 0x3e, 0xd2, 0xb3, 0x7b, 0xd2,
	0xef, 0xef, 0x49, 0x37, 0x93, 0x7b, 0x52, 0xff, 0x82, 0x67, 0xea, 0xcb, 0x44, 0x21, 0x6a, 0x83,
	0xd6, 0x67, 0x6e, 0x34, 0x09, 0x51, 0x38, 0x85, 0x8f, 0x24, 0x95, 0xcb, 0xf6, 0x97, 0xb3, 0x69,
	0x20, 0x67, 0x4b, 0x4f, 0xf7, 0xd9, 0xc2, 0xc8, 0x33, 0xc9, 0xee, 0xd5, 0x67, 0xa1, 0x21, 0xb8,
	0xef, 0x1d, 0xa6, 0xaf, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0xf7, 0x34, 0xd6, 0xe3,
	0x03, 0x00, 0x00,
}
