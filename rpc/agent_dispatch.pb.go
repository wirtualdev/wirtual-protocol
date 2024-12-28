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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.23.4
// source: rpc/agent_dispatch.proto

package rpc

import (
	wirtual "github.com/wirtualdev/wirtual-protocol/wirtual"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_agent_dispatch_proto protoreflect.FileDescriptor

var file_rpc_agent_dispatch_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc7, 0x02, 0x0a,
	0x15, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22, 0x16, 0xb2, 0x89, 0x01, 0x12, 0x10, 0x01,
	0x1a, 0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01,
	0x12, 0x65, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x22,
	0x16, 0xb2, 0x89, 0x01, 0x12, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x12, 0x6d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16,
	0xb2, 0x89, 0x01, 0x12, 0x10, 0x01, 0x1a, 0x0e, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_rpc_agent_dispatch_proto_goTypes = []any{
	(*wirtual.AgentDispatch)(nil),              // 0: wirtual.AgentDispatch
	(*wirtual.DeleteAgentDispatchRequest)(nil), // 1: wirtual.DeleteAgentDispatchRequest
	(*wirtual.ListAgentDispatchRequest)(nil),   // 2: wirtual.ListAgentDispatchRequest
	(*wirtual.ListAgentDispatchResponse)(nil),  // 3: wirtual.ListAgentDispatchResponse
}
var file_rpc_agent_dispatch_proto_depIdxs = []int32{
	0, // 0: rpc.AgentDispatchInternal.CreateDispatch:input_type -> wirtual.AgentDispatch
	1, // 1: rpc.AgentDispatchInternal.DeleteDispatch:input_type -> wirtual.DeleteAgentDispatchRequest
	2, // 2: rpc.AgentDispatchInternal.ListDispatch:input_type -> wirtual.ListAgentDispatchRequest
	0, // 3: rpc.AgentDispatchInternal.CreateDispatch:output_type -> wirtual.AgentDispatch
	0, // 4: rpc.AgentDispatchInternal.DeleteDispatch:output_type -> wirtual.AgentDispatch
	3, // 5: rpc.AgentDispatchInternal.ListDispatch:output_type -> wirtual.ListAgentDispatchResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_agent_dispatch_proto_init() }
func file_rpc_agent_dispatch_proto_init() {
	if File_rpc_agent_dispatch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_agent_dispatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_agent_dispatch_proto_goTypes,
		DependencyIndexes: file_rpc_agent_dispatch_proto_depIdxs,
	}.Build()
	File_rpc_agent_dispatch_proto = out.File
	file_rpc_agent_dispatch_proto_rawDesc = nil
	file_rpc_agent_dispatch_proto_goTypes = nil
	file_rpc_agent_dispatch_proto_depIdxs = nil
}
