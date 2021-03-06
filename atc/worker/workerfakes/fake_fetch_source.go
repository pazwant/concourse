// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc/worker"
)

type FakeFetchSource struct {
	CreateStub        func(context.Context) (worker.GetResult, worker.Volume, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
	}
	createReturns struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}
	createReturnsOnCall map[int]struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}
	FindStub        func() (worker.GetResult, worker.Volume, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
	}
	findReturns struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 bool
		result4 error
	}
	findReturnsOnCall map[int]struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 bool
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSource) Create(arg1 context.Context) (worker.GetResult, worker.Volume, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeFetchSource) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeFetchSource) CreateCalls(stub func(context.Context) (worker.GetResult, worker.Volume, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeFetchSource) CreateArgsForCall(i int) context.Context {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFetchSource) CreateReturns(result1 worker.GetResult, result2 worker.Volume, result3 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetchSource) CreateReturnsOnCall(i int, result1 worker.GetResult, result2 worker.Volume, result3 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 worker.GetResult
			result2 worker.Volume
			result3 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetchSource) Find() (worker.GetResult, worker.Volume, bool, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
	}{})
	fake.recordInvocation("Find", []interface{}{})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.findReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeFetchSource) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeFetchSource) FindCalls(stub func() (worker.GetResult, worker.Volume, bool, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeFetchSource) FindReturns(result1 worker.GetResult, result2 worker.Volume, result3 bool, result4 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeFetchSource) FindReturnsOnCall(i int, result1 worker.GetResult, result2 worker.Volume, result3 bool, result4 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 worker.GetResult
			result2 worker.Volume
			result3 bool
			result4 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 worker.GetResult
		result2 worker.Volume
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeFetchSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetchSource) recordInvocation(key string, args []interface{}) {
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

var _ worker.FetchSource = new(FakeFetchSource)
