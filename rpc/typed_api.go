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

package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/wirtualdev/wirtual-protocol/logger"
	"github.com/wirtualdev/wirtual-protocol/utils/must"
	"github.com/wirtualdev/psrpc"
	"github.com/wirtualdev/psrpc/pkg/middleware"
)

type PSRPCConfig struct {
	MaxAttempts int           `yaml:"max_attempts,omitempty"`
	Timeout     time.Duration `yaml:"timeout,omitempty"`
	Backoff     time.Duration `yaml:"backoff,omitempty"`
	BufferSize  int           `yaml:"buffer_size,omitempty"`
}

var DefaultPSRPCConfig = PSRPCConfig{
	MaxAttempts: 3,
	Timeout:     3 * time.Second,
	Backoff:     2 * time.Second,
	BufferSize:  1000,
}

type ClientParams struct {
	PSRPCConfig
	Bus      psrpc.MessageBus
	Logger   logger.Logger
	Observer middleware.MetricsObserver
}

func NewClientParams(
	config PSRPCConfig,
	bus psrpc.MessageBus,
	logger logger.Logger,
	observer middleware.MetricsObserver,
) ClientParams {
	return ClientParams{
		PSRPCConfig: config,
		Bus:         bus,
		Logger:      logger,
		Observer:    observer,
	}
}

func (p *ClientParams) Options() []psrpc.ClientOption {
	opts := make([]psrpc.ClientOption, 0, 4)
	if p.BufferSize != 0 {
		opts = append(opts, psrpc.WithClientChannelSize(p.BufferSize))
	}
	if p.Observer != nil {
		opts = append(opts, middleware.WithClientMetrics(p.Observer))
	}
	if p.Logger != nil {
		opts = append(opts, WithClientLogger(p.Logger))
	}
	if p.MaxAttempts != 0 || p.Timeout != 0 || p.Backoff != 0 {
		opts = append(opts, middleware.WithRPCRetries(middleware.RetryOptions{
			MaxAttempts: p.MaxAttempts,
			Timeout:     p.Timeout,
			Backoff:     p.Backoff,
		}))
	}
	return opts
}

func (p *ClientParams) Args() (psrpc.MessageBus, psrpc.ClientOption) {
	return p.Bus, psrpc.WithClientOptions(p.Options()...)
}

func WithServerObservability(logger logger.Logger) psrpc.ServerOption {
	return psrpc.WithServerOptions(
		middleware.WithServerMetrics(PSRPCMetricsObserver{}),
		WithServerLogger(logger),
	)
}

func WithDefaultServerOptions(psrpcConfig PSRPCConfig, logger logger.Logger) psrpc.ServerOption {
	return psrpc.WithServerOptions(
		psrpc.WithServerChannelSize(psrpcConfig.BufferSize),
		WithServerObservability(logger),
	)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type TypedSignalClient = SignalClient[wirtual.NodeID]
type TypedSignalServer = SignalServer[wirtual.NodeID]

func NewTypedSignalClient(nodeID wirtual.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedSignalClient, error) {
	return NewSignalClient[wirtual.NodeID](bus, psrpc.WithClientOptions(opts...), psrpc.WithClientID(string(nodeID)))
}

func NewTypedSignalServer(nodeID wirtual.NodeID, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalServer, error) {
	return NewSignalServer[wirtual.NodeID](svc, bus, psrpc.WithServerOptions(opts...), psrpc.WithServerID(string(nodeID)))
}

type TypedRoomManagerClient = RoomManagerClient[wirtual.NodeID]
type TypedRoomManagerServer = RoomManagerServer[wirtual.NodeID]

func NewTypedRoomManagerClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedRoomManagerClient, error) {
	return NewRoomManagerClient[wirtual.NodeID](bus, opts...)
}

func NewTypedRoomManagerServer(svc RoomManagerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomManagerServer, error) {
	return NewRoomManagerServer[wirtual.NodeID](svc, bus, opts...)
}

type ParticipantTopic string
type RoomTopic string

func FormatParticipantTopic(roomName wirtual.RoomName, identity wirtual.ParticipantIdentity) ParticipantTopic {
	return ParticipantTopic(fmt.Sprintf("%s_%s", roomName, identity))
}

func FormatRoomTopic(roomName wirtual.RoomName) RoomTopic {
	return RoomTopic(roomName)
}

type topicFormatter struct{}

func NewTopicFormatter() TopicFormatter {
	return topicFormatter{}
}

func (f topicFormatter) ParticipantTopic(ctx context.Context, roomName wirtual.RoomName, identity wirtual.ParticipantIdentity) ParticipantTopic {
	return FormatParticipantTopic(roomName, identity)
}

func (f topicFormatter) RoomTopic(ctx context.Context, roomName wirtual.RoomName) RoomTopic {
	return FormatRoomTopic(roomName)
}

type TopicFormatter interface {
	ParticipantTopic(ctx context.Context, roomName wirtual.RoomName, identity wirtual.ParticipantIdentity) ParticipantTopic
	RoomTopic(ctx context.Context, roomName wirtual.RoomName) RoomTopic
}

//counterfeiter:generate . TypedRoomClient
type TypedRoomClient = RoomClient[RoomTopic]
type TypedRoomServer = RoomServer[RoomTopic]

func NewTypedRoomClient(params ClientParams) (TypedRoomClient, error) {
	return NewRoomClient[RoomTopic](params.Args())
}

func NewTypedRoomServer(svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomServer, error) {
	return NewRoomServer[RoomTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedParticipantClient
type TypedParticipantClient = ParticipantClient[ParticipantTopic]
type TypedParticipantServer = ParticipantServer[ParticipantTopic]

func NewTypedParticipantClient(params ClientParams) (TypedParticipantClient, error) {
	return NewParticipantClient[ParticipantTopic](params.Args())
}

func NewTypedParticipantServer(svc ParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedParticipantServer, error) {
	return NewParticipantServer[ParticipantTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedAgentDispatchInternalClient
type TypedAgentDispatchInternalClient = AgentDispatchInternalClient[RoomTopic]
type TypedAgentDispatchInternalServer = AgentDispatchInternalServer[RoomTopic]

func NewTypedAgentDispatchInternalClient(params ClientParams) (TypedAgentDispatchInternalClient, error) {
	return NewAgentDispatchInternalClient[RoomTopic](params.Args())
}

func NewTypedAgentDispatchInternalServer(svc AgentDispatchInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedAgentDispatchInternalServer, error) {
	return NewAgentDispatchInternalServer[RoomTopic](svc, bus, opts...)
}

//counterfeiter:generate . KeepalivePubSub
type KeepalivePubSub interface {
	KeepaliveClient[wirtual.NodeID]
	KeepaliveServer[wirtual.NodeID]
}

type keepalivePubSub struct {
	KeepaliveClient[wirtual.NodeID]
	KeepaliveServer[wirtual.NodeID]
}

func NewKeepalivePubSub(params ClientParams) (KeepalivePubSub, error) {
	client, err := NewKeepaliveClient[wirtual.NodeID](params.Args())
	if err != nil {
		return nil, err
	}
	server := must.Get(NewKeepaliveServer[wirtual.NodeID](nil, params.Bus))
	return &keepalivePubSub{client, server}, nil
}
