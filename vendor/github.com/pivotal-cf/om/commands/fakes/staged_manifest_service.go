// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type StagedManifestService struct {
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
	GetStagedProductManifestStub        func(string) (string, error)
	getStagedProductManifestMutex       sync.RWMutex
	getStagedProductManifestArgsForCall []struct {
		arg1 string
	}
	getStagedProductManifestReturns struct {
		result1 string
		result2 error
	}
	getStagedProductManifestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedManifestService) GetStagedProductByName(arg1 string) (api.StagedProductsFindOutput, error) {
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

func (fake *StagedManifestService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *StagedManifestService) GetStagedProductByNameCalls(stub func(string) (api.StagedProductsFindOutput, error)) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = stub
}

func (fake *StagedManifestService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	argsForCall := fake.getStagedProductByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedManifestService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedManifestService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
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

func (fake *StagedManifestService) GetStagedProductManifest(arg1 string) (string, error) {
	fake.getStagedProductManifestMutex.Lock()
	ret, specificReturn := fake.getStagedProductManifestReturnsOnCall[len(fake.getStagedProductManifestArgsForCall)]
	fake.getStagedProductManifestArgsForCall = append(fake.getStagedProductManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductManifest", []interface{}{arg1})
	fake.getStagedProductManifestMutex.Unlock()
	if fake.GetStagedProductManifestStub != nil {
		return fake.GetStagedProductManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedManifestService) GetStagedProductManifestCallCount() int {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	return len(fake.getStagedProductManifestArgsForCall)
}

func (fake *StagedManifestService) GetStagedProductManifestCalls(stub func(string) (string, error)) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = stub
}

func (fake *StagedManifestService) GetStagedProductManifestArgsForCall(i int) string {
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	argsForCall := fake.getStagedProductManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedManifestService) GetStagedProductManifestReturns(result1 string, result2 error) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = nil
	fake.getStagedProductManifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *StagedManifestService) GetStagedProductManifestReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStagedProductManifestMutex.Lock()
	defer fake.getStagedProductManifestMutex.Unlock()
	fake.GetStagedProductManifestStub = nil
	if fake.getStagedProductManifestReturnsOnCall == nil {
		fake.getStagedProductManifestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStagedProductManifestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *StagedManifestService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductManifestMutex.RLock()
	defer fake.getStagedProductManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedManifestService) recordInvocation(key string, args []interface{}) {
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