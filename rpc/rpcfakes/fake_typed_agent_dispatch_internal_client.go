// Code generated by counterfeiter. DO NOT EDIT.
package rpcfakes

import (
	"context"
	"sync"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/wirtualdev/wirtual-protocol/rpc"
	"github.com/wirtual/psrpc"
)

type FakeTypedAgentDispatchInternalClient struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	CreateDispatchStub        func(context.Context, rpc.RoomTopic, *wirtual.AgentDispatch, ...psrpc.RequestOption) (*wirtual.AgentDispatch, error)
	createDispatchMutex       sync.RWMutex
	createDispatchArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.AgentDispatch
		arg4 []psrpc.RequestOption
	}
	createDispatchReturns struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}
	createDispatchReturnsOnCall map[int]struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}
	DeleteDispatchStub        func(context.Context, rpc.RoomTopic, *wirtual.DeleteAgentDispatchRequest, ...psrpc.RequestOption) (*wirtual.AgentDispatch, error)
	deleteDispatchMutex       sync.RWMutex
	deleteDispatchArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.DeleteAgentDispatchRequest
		arg4 []psrpc.RequestOption
	}
	deleteDispatchReturns struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}
	deleteDispatchReturnsOnCall map[int]struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}
	ListDispatchStub        func(context.Context, rpc.RoomTopic, *wirtual.ListAgentDispatchRequest, ...psrpc.RequestOption) (*wirtual.ListAgentDispatchResponse, error)
	listDispatchMutex       sync.RWMutex
	listDispatchArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.ListAgentDispatchRequest
		arg4 []psrpc.RequestOption
	}
	listDispatchReturns struct {
		result1 *wirtual.ListAgentDispatchResponse
		result2 error
	}
	listDispatchReturnsOnCall map[int]struct {
		result1 *wirtual.ListAgentDispatchResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypedAgentDispatchInternalClient) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeTypedAgentDispatchInternalClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeTypedAgentDispatchInternalClient) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatch(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.AgentDispatch, arg4 ...psrpc.RequestOption) (*wirtual.AgentDispatch, error) {
	fake.createDispatchMutex.Lock()
	ret, specificReturn := fake.createDispatchReturnsOnCall[len(fake.createDispatchArgsForCall)]
	fake.createDispatchArgsForCall = append(fake.createDispatchArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.AgentDispatch
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateDispatchStub
	fakeReturns := fake.createDispatchReturns
	fake.recordInvocation("CreateDispatch", []interface{}{arg1, arg2, arg3, arg4})
	fake.createDispatchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatchCallCount() int {
	fake.createDispatchMutex.RLock()
	defer fake.createDispatchMutex.RUnlock()
	return len(fake.createDispatchArgsForCall)
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatchCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.AgentDispatch, ...psrpc.RequestOption) (*wirtual.AgentDispatch, error)) {
	fake.createDispatchMutex.Lock()
	defer fake.createDispatchMutex.Unlock()
	fake.CreateDispatchStub = stub
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatchArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.AgentDispatch, []psrpc.RequestOption) {
	fake.createDispatchMutex.RLock()
	defer fake.createDispatchMutex.RUnlock()
	argsForCall := fake.createDispatchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatchReturns(result1 *wirtual.AgentDispatch, result2 error) {
	fake.createDispatchMutex.Lock()
	defer fake.createDispatchMutex.Unlock()
	fake.CreateDispatchStub = nil
	fake.createDispatchReturns = struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) CreateDispatchReturnsOnCall(i int, result1 *wirtual.AgentDispatch, result2 error) {
	fake.createDispatchMutex.Lock()
	defer fake.createDispatchMutex.Unlock()
	fake.CreateDispatchStub = nil
	if fake.createDispatchReturnsOnCall == nil {
		fake.createDispatchReturnsOnCall = make(map[int]struct {
			result1 *wirtual.AgentDispatch
			result2 error
		})
	}
	fake.createDispatchReturnsOnCall[i] = struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatch(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.DeleteAgentDispatchRequest, arg4 ...psrpc.RequestOption) (*wirtual.AgentDispatch, error) {
	fake.deleteDispatchMutex.Lock()
	ret, specificReturn := fake.deleteDispatchReturnsOnCall[len(fake.deleteDispatchArgsForCall)]
	fake.deleteDispatchArgsForCall = append(fake.deleteDispatchArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.DeleteAgentDispatchRequest
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.DeleteDispatchStub
	fakeReturns := fake.deleteDispatchReturns
	fake.recordInvocation("DeleteDispatch", []interface{}{arg1, arg2, arg3, arg4})
	fake.deleteDispatchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatchCallCount() int {
	fake.deleteDispatchMutex.RLock()
	defer fake.deleteDispatchMutex.RUnlock()
	return len(fake.deleteDispatchArgsForCall)
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatchCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.DeleteAgentDispatchRequest, ...psrpc.RequestOption) (*wirtual.AgentDispatch, error)) {
	fake.deleteDispatchMutex.Lock()
	defer fake.deleteDispatchMutex.Unlock()
	fake.DeleteDispatchStub = stub
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatchArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.DeleteAgentDispatchRequest, []psrpc.RequestOption) {
	fake.deleteDispatchMutex.RLock()
	defer fake.deleteDispatchMutex.RUnlock()
	argsForCall := fake.deleteDispatchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatchReturns(result1 *wirtual.AgentDispatch, result2 error) {
	fake.deleteDispatchMutex.Lock()
	defer fake.deleteDispatchMutex.Unlock()
	fake.DeleteDispatchStub = nil
	fake.deleteDispatchReturns = struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) DeleteDispatchReturnsOnCall(i int, result1 *wirtual.AgentDispatch, result2 error) {
	fake.deleteDispatchMutex.Lock()
	defer fake.deleteDispatchMutex.Unlock()
	fake.DeleteDispatchStub = nil
	if fake.deleteDispatchReturnsOnCall == nil {
		fake.deleteDispatchReturnsOnCall = make(map[int]struct {
			result1 *wirtual.AgentDispatch
			result2 error
		})
	}
	fake.deleteDispatchReturnsOnCall[i] = struct {
		result1 *wirtual.AgentDispatch
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatch(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.ListAgentDispatchRequest, arg4 ...psrpc.RequestOption) (*wirtual.ListAgentDispatchResponse, error) {
	fake.listDispatchMutex.Lock()
	ret, specificReturn := fake.listDispatchReturnsOnCall[len(fake.listDispatchArgsForCall)]
	fake.listDispatchArgsForCall = append(fake.listDispatchArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.ListAgentDispatchRequest
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.ListDispatchStub
	fakeReturns := fake.listDispatchReturns
	fake.recordInvocation("ListDispatch", []interface{}{arg1, arg2, arg3, arg4})
	fake.listDispatchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatchCallCount() int {
	fake.listDispatchMutex.RLock()
	defer fake.listDispatchMutex.RUnlock()
	return len(fake.listDispatchArgsForCall)
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatchCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.ListAgentDispatchRequest, ...psrpc.RequestOption) (*wirtual.ListAgentDispatchResponse, error)) {
	fake.listDispatchMutex.Lock()
	defer fake.listDispatchMutex.Unlock()
	fake.ListDispatchStub = stub
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatchArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.ListAgentDispatchRequest, []psrpc.RequestOption) {
	fake.listDispatchMutex.RLock()
	defer fake.listDispatchMutex.RUnlock()
	argsForCall := fake.listDispatchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatchReturns(result1 *wirtual.ListAgentDispatchResponse, result2 error) {
	fake.listDispatchMutex.Lock()
	defer fake.listDispatchMutex.Unlock()
	fake.ListDispatchStub = nil
	fake.listDispatchReturns = struct {
		result1 *wirtual.ListAgentDispatchResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) ListDispatchReturnsOnCall(i int, result1 *wirtual.ListAgentDispatchResponse, result2 error) {
	fake.listDispatchMutex.Lock()
	defer fake.listDispatchMutex.Unlock()
	fake.ListDispatchStub = nil
	if fake.listDispatchReturnsOnCall == nil {
		fake.listDispatchReturnsOnCall = make(map[int]struct {
			result1 *wirtual.ListAgentDispatchResponse
			result2 error
		})
	}
	fake.listDispatchReturnsOnCall[i] = struct {
		result1 *wirtual.ListAgentDispatchResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedAgentDispatchInternalClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createDispatchMutex.RLock()
	defer fake.createDispatchMutex.RUnlock()
	fake.deleteDispatchMutex.RLock()
	defer fake.deleteDispatchMutex.RUnlock()
	fake.listDispatchMutex.RLock()
	defer fake.listDispatchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypedAgentDispatchInternalClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rpc.TypedAgentDispatchInternalClient = new(FakeTypedAgentDispatchInternalClient)
