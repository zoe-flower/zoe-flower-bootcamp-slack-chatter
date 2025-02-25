// Code generated by counterfeiter. DO NOT EDIT.
package eventbusfakes

import (
	"context"
	"sync"
	"time"

	"github.com/flypay/go-kit/v4/pkg/eventbus"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FakeDelayedEmitter struct {
	EmitAtStub        func(context.Context, time.Time, protoreflect.ProtoMessage, ...eventbus.Header) error
	emitAtMutex       sync.RWMutex
	emitAtArgsForCall []struct {
		arg1 context.Context
		arg2 time.Time
		arg3 protoreflect.ProtoMessage
		arg4 []eventbus.Header
	}
	emitAtReturns struct {
		result1 error
	}
	emitAtReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDelayedEmitter) EmitAt(arg1 context.Context, arg2 time.Time, arg3 protoreflect.ProtoMessage, arg4 ...eventbus.Header) error {
	fake.emitAtMutex.Lock()
	ret, specificReturn := fake.emitAtReturnsOnCall[len(fake.emitAtArgsForCall)]
	fake.emitAtArgsForCall = append(fake.emitAtArgsForCall, struct {
		arg1 context.Context
		arg2 time.Time
		arg3 protoreflect.ProtoMessage
		arg4 []eventbus.Header
	}{arg1, arg2, arg3, arg4})
	stub := fake.EmitAtStub
	fakeReturns := fake.emitAtReturns
	fake.recordInvocation("EmitAt", []interface{}{arg1, arg2, arg3, arg4})
	fake.emitAtMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDelayedEmitter) EmitAtCallCount() int {
	fake.emitAtMutex.RLock()
	defer fake.emitAtMutex.RUnlock()
	return len(fake.emitAtArgsForCall)
}

func (fake *FakeDelayedEmitter) EmitAtCalls(stub func(context.Context, time.Time, protoreflect.ProtoMessage, ...eventbus.Header) error) {
	fake.emitAtMutex.Lock()
	defer fake.emitAtMutex.Unlock()
	fake.EmitAtStub = stub
}

func (fake *FakeDelayedEmitter) EmitAtArgsForCall(i int) (context.Context, time.Time, protoreflect.ProtoMessage, []eventbus.Header) {
	fake.emitAtMutex.RLock()
	defer fake.emitAtMutex.RUnlock()
	argsForCall := fake.emitAtArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDelayedEmitter) EmitAtReturns(result1 error) {
	fake.emitAtMutex.Lock()
	defer fake.emitAtMutex.Unlock()
	fake.EmitAtStub = nil
	fake.emitAtReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDelayedEmitter) EmitAtReturnsOnCall(i int, result1 error) {
	fake.emitAtMutex.Lock()
	defer fake.emitAtMutex.Unlock()
	fake.EmitAtStub = nil
	if fake.emitAtReturnsOnCall == nil {
		fake.emitAtReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.emitAtReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDelayedEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emitAtMutex.RLock()
	defer fake.emitAtMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDelayedEmitter) recordInvocation(key string, args []interface{}) {
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

var _ eventbus.DelayedEmitter = new(FakeDelayedEmitter)
