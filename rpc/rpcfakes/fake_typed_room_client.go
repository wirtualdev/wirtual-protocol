// Code generated by counterfeiter. DO NOT EDIT.
package rpcfakes

import (
	"context"
	"sync"

	"github.com/wirtualdev/wirtual-protocol/wirtual"
	"github.com/wirtualdev/wirtual-protocol/rpc"
	"github.com/livekit/psrpc"
)

type FakeTypedRoomClient struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	DeleteRoomStub        func(context.Context, rpc.RoomTopic, *wirtual.DeleteRoomRequest, ...psrpc.RequestOption) (*wirtual.DeleteRoomResponse, error)
	deleteRoomMutex       sync.RWMutex
	deleteRoomArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.DeleteRoomRequest
		arg4 []psrpc.RequestOption
	}
	deleteRoomReturns struct {
		result1 *wirtual.DeleteRoomResponse
		result2 error
	}
	deleteRoomReturnsOnCall map[int]struct {
		result1 *wirtual.DeleteRoomResponse
		result2 error
	}
	SendDataStub        func(context.Context, rpc.RoomTopic, *wirtual.SendDataRequest, ...psrpc.RequestOption) (*wirtual.SendDataResponse, error)
	sendDataMutex       sync.RWMutex
	sendDataArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.SendDataRequest
		arg4 []psrpc.RequestOption
	}
	sendDataReturns struct {
		result1 *wirtual.SendDataResponse
		result2 error
	}
	sendDataReturnsOnCall map[int]struct {
		result1 *wirtual.SendDataResponse
		result2 error
	}
	UpdateRoomMetadataStub        func(context.Context, rpc.RoomTopic, *wirtual.UpdateRoomMetadataRequest, ...psrpc.RequestOption) (*wirtual.Room, error)
	updateRoomMetadataMutex       sync.RWMutex
	updateRoomMetadataArgsForCall []struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.UpdateRoomMetadataRequest
		arg4 []psrpc.RequestOption
	}
	updateRoomMetadataReturns struct {
		result1 *wirtual.Room
		result2 error
	}
	updateRoomMetadataReturnsOnCall map[int]struct {
		result1 *wirtual.Room
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTypedRoomClient) Close() {
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

func (fake *FakeTypedRoomClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeTypedRoomClient) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeTypedRoomClient) DeleteRoom(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.DeleteRoomRequest, arg4 ...psrpc.RequestOption) (*wirtual.DeleteRoomResponse, error) {
	fake.deleteRoomMutex.Lock()
	ret, specificReturn := fake.deleteRoomReturnsOnCall[len(fake.deleteRoomArgsForCall)]
	fake.deleteRoomArgsForCall = append(fake.deleteRoomArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.DeleteRoomRequest
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.DeleteRoomStub
	fakeReturns := fake.deleteRoomReturns
	fake.recordInvocation("DeleteRoom", []interface{}{arg1, arg2, arg3, arg4})
	fake.deleteRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedRoomClient) DeleteRoomCallCount() int {
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	return len(fake.deleteRoomArgsForCall)
}

func (fake *FakeTypedRoomClient) DeleteRoomCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.DeleteRoomRequest, ...psrpc.RequestOption) (*wirtual.DeleteRoomResponse, error)) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = stub
}

func (fake *FakeTypedRoomClient) DeleteRoomArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.DeleteRoomRequest, []psrpc.RequestOption) {
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	argsForCall := fake.deleteRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedRoomClient) DeleteRoomReturns(result1 *wirtual.DeleteRoomResponse, result2 error) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = nil
	fake.deleteRoomReturns = struct {
		result1 *wirtual.DeleteRoomResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) DeleteRoomReturnsOnCall(i int, result1 *wirtual.DeleteRoomResponse, result2 error) {
	fake.deleteRoomMutex.Lock()
	defer fake.deleteRoomMutex.Unlock()
	fake.DeleteRoomStub = nil
	if fake.deleteRoomReturnsOnCall == nil {
		fake.deleteRoomReturnsOnCall = make(map[int]struct {
			result1 *wirtual.DeleteRoomResponse
			result2 error
		})
	}
	fake.deleteRoomReturnsOnCall[i] = struct {
		result1 *wirtual.DeleteRoomResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) SendData(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.SendDataRequest, arg4 ...psrpc.RequestOption) (*wirtual.SendDataResponse, error) {
	fake.sendDataMutex.Lock()
	ret, specificReturn := fake.sendDataReturnsOnCall[len(fake.sendDataArgsForCall)]
	fake.sendDataArgsForCall = append(fake.sendDataArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.SendDataRequest
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.SendDataStub
	fakeReturns := fake.sendDataReturns
	fake.recordInvocation("SendData", []interface{}{arg1, arg2, arg3, arg4})
	fake.sendDataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedRoomClient) SendDataCallCount() int {
	fake.sendDataMutex.RLock()
	defer fake.sendDataMutex.RUnlock()
	return len(fake.sendDataArgsForCall)
}

func (fake *FakeTypedRoomClient) SendDataCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.SendDataRequest, ...psrpc.RequestOption) (*wirtual.SendDataResponse, error)) {
	fake.sendDataMutex.Lock()
	defer fake.sendDataMutex.Unlock()
	fake.SendDataStub = stub
}

func (fake *FakeTypedRoomClient) SendDataArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.SendDataRequest, []psrpc.RequestOption) {
	fake.sendDataMutex.RLock()
	defer fake.sendDataMutex.RUnlock()
	argsForCall := fake.sendDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedRoomClient) SendDataReturns(result1 *wirtual.SendDataResponse, result2 error) {
	fake.sendDataMutex.Lock()
	defer fake.sendDataMutex.Unlock()
	fake.SendDataStub = nil
	fake.sendDataReturns = struct {
		result1 *wirtual.SendDataResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) SendDataReturnsOnCall(i int, result1 *wirtual.SendDataResponse, result2 error) {
	fake.sendDataMutex.Lock()
	defer fake.sendDataMutex.Unlock()
	fake.SendDataStub = nil
	if fake.sendDataReturnsOnCall == nil {
		fake.sendDataReturnsOnCall = make(map[int]struct {
			result1 *wirtual.SendDataResponse
			result2 error
		})
	}
	fake.sendDataReturnsOnCall[i] = struct {
		result1 *wirtual.SendDataResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadata(arg1 context.Context, arg2 rpc.RoomTopic, arg3 *wirtual.UpdateRoomMetadataRequest, arg4 ...psrpc.RequestOption) (*wirtual.Room, error) {
	fake.updateRoomMetadataMutex.Lock()
	ret, specificReturn := fake.updateRoomMetadataReturnsOnCall[len(fake.updateRoomMetadataArgsForCall)]
	fake.updateRoomMetadataArgsForCall = append(fake.updateRoomMetadataArgsForCall, struct {
		arg1 context.Context
		arg2 rpc.RoomTopic
		arg3 *wirtual.UpdateRoomMetadataRequest
		arg4 []psrpc.RequestOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.UpdateRoomMetadataStub
	fakeReturns := fake.updateRoomMetadataReturns
	fake.recordInvocation("UpdateRoomMetadata", []interface{}{arg1, arg2, arg3, arg4})
	fake.updateRoomMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadataCallCount() int {
	fake.updateRoomMetadataMutex.RLock()
	defer fake.updateRoomMetadataMutex.RUnlock()
	return len(fake.updateRoomMetadataArgsForCall)
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadataCalls(stub func(context.Context, rpc.RoomTopic, *wirtual.UpdateRoomMetadataRequest, ...psrpc.RequestOption) (*wirtual.Room, error)) {
	fake.updateRoomMetadataMutex.Lock()
	defer fake.updateRoomMetadataMutex.Unlock()
	fake.UpdateRoomMetadataStub = stub
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadataArgsForCall(i int) (context.Context, rpc.RoomTopic, *wirtual.UpdateRoomMetadataRequest, []psrpc.RequestOption) {
	fake.updateRoomMetadataMutex.RLock()
	defer fake.updateRoomMetadataMutex.RUnlock()
	argsForCall := fake.updateRoomMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadataReturns(result1 *wirtual.Room, result2 error) {
	fake.updateRoomMetadataMutex.Lock()
	defer fake.updateRoomMetadataMutex.Unlock()
	fake.UpdateRoomMetadataStub = nil
	fake.updateRoomMetadataReturns = struct {
		result1 *wirtual.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) UpdateRoomMetadataReturnsOnCall(i int, result1 *wirtual.Room, result2 error) {
	fake.updateRoomMetadataMutex.Lock()
	defer fake.updateRoomMetadataMutex.Unlock()
	fake.UpdateRoomMetadataStub = nil
	if fake.updateRoomMetadataReturnsOnCall == nil {
		fake.updateRoomMetadataReturnsOnCall = make(map[int]struct {
			result1 *wirtual.Room
			result2 error
		})
	}
	fake.updateRoomMetadataReturnsOnCall[i] = struct {
		result1 *wirtual.Room
		result2 error
	}{result1, result2}
}

func (fake *FakeTypedRoomClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.deleteRoomMutex.RLock()
	defer fake.deleteRoomMutex.RUnlock()
	fake.sendDataMutex.RLock()
	defer fake.sendDataMutex.RUnlock()
	fake.updateRoomMetadataMutex.RLock()
	defer fake.updateRoomMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTypedRoomClient) recordInvocation(key string, args []interface{}) {
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

var _ rpc.TypedRoomClient = new(FakeTypedRoomClient)
