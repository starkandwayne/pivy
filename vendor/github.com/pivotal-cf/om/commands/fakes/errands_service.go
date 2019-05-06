// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type ErrandsService struct {
	GetStagedProductByNameStub        func(string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		arg1 string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	ListStagedProductErrandsStub        func(string) (api.ErrandsListOutput, error)
	listStagedProductErrandsMutex       sync.RWMutex
	listStagedProductErrandsArgsForCall []struct {
		arg1 string
	}
	listStagedProductErrandsReturns struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	listStagedProductErrandsReturnsOnCall map[int]struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ErrandsService) GetStagedProductByName(arg1 string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductByName", []interface{}{arg1})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ErrandsService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *ErrandsService) GetStagedProductByNameCalls(stub func(string) (api.StagedProductsFindOutput, error)) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = stub
}

func (fake *ErrandsService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	argsForCall := fake.getStagedProductByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ErrandsService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *ErrandsService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *ErrandsService) ListStagedProductErrands(arg1 string) (api.ErrandsListOutput, error) {
	fake.listStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.listStagedProductErrandsReturnsOnCall[len(fake.listStagedProductErrandsArgsForCall)]
	fake.listStagedProductErrandsArgsForCall = append(fake.listStagedProductErrandsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListStagedProductErrands", []interface{}{arg1})
	fake.listStagedProductErrandsMutex.Unlock()
	if fake.ListStagedProductErrandsStub != nil {
		return fake.ListStagedProductErrandsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedProductErrandsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ErrandsService) ListStagedProductErrandsCallCount() int {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return len(fake.listStagedProductErrandsArgsForCall)
}

func (fake *ErrandsService) ListStagedProductErrandsCalls(stub func(string) (api.ErrandsListOutput, error)) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = stub
}

func (fake *ErrandsService) ListStagedProductErrandsArgsForCall(i int) string {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	argsForCall := fake.listStagedProductErrandsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ErrandsService) ListStagedProductErrandsReturns(result1 api.ErrandsListOutput, result2 error) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = nil
	fake.listStagedProductErrandsReturns = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *ErrandsService) ListStagedProductErrandsReturnsOnCall(i int, result1 api.ErrandsListOutput, result2 error) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = nil
	if fake.listStagedProductErrandsReturnsOnCall == nil {
		fake.listStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 api.ErrandsListOutput
			result2 error
		})
	}
	fake.listStagedProductErrandsReturnsOnCall[i] = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *ErrandsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ErrandsService) recordInvocation(key string, args []interface{}) {
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