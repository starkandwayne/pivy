// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type DeleteInstallationService struct {
	DeleteInstallationAssetCollectionStub        func() (api.InstallationsServiceOutput, error)
	deleteInstallationAssetCollectionMutex       sync.RWMutex
	deleteInstallationAssetCollectionArgsForCall []struct {
	}
	deleteInstallationAssetCollectionReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	deleteInstallationAssetCollectionReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	GetInstallationStub        func(int) (api.InstallationsServiceOutput, error)
	getInstallationMutex       sync.RWMutex
	getInstallationArgsForCall []struct {
		arg1 int
	}
	getInstallationReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	getInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	GetInstallationLogsStub        func(int) (api.InstallationsServiceOutput, error)
	getInstallationLogsMutex       sync.RWMutex
	getInstallationLogsArgsForCall []struct {
		arg1 int
	}
	getInstallationLogsReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	getInstallationLogsReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	RunningInstallationStub        func() (api.InstallationsServiceOutput, error)
	runningInstallationMutex       sync.RWMutex
	runningInstallationArgsForCall []struct {
	}
	runningInstallationReturns struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	runningInstallationReturnsOnCall map[int]struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeleteInstallationService) DeleteInstallationAssetCollection() (api.InstallationsServiceOutput, error) {
	fake.deleteInstallationAssetCollectionMutex.Lock()
	ret, specificReturn := fake.deleteInstallationAssetCollectionReturnsOnCall[len(fake.deleteInstallationAssetCollectionArgsForCall)]
	fake.deleteInstallationAssetCollectionArgsForCall = append(fake.deleteInstallationAssetCollectionArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteInstallationAssetCollection", []interface{}{})
	fake.deleteInstallationAssetCollectionMutex.Unlock()
	if fake.DeleteInstallationAssetCollectionStub != nil {
		return fake.DeleteInstallationAssetCollectionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteInstallationAssetCollectionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeleteInstallationService) DeleteInstallationAssetCollectionCallCount() int {
	fake.deleteInstallationAssetCollectionMutex.RLock()
	defer fake.deleteInstallationAssetCollectionMutex.RUnlock()
	return len(fake.deleteInstallationAssetCollectionArgsForCall)
}

func (fake *DeleteInstallationService) DeleteInstallationAssetCollectionCalls(stub func() (api.InstallationsServiceOutput, error)) {
	fake.deleteInstallationAssetCollectionMutex.Lock()
	defer fake.deleteInstallationAssetCollectionMutex.Unlock()
	fake.DeleteInstallationAssetCollectionStub = stub
}

func (fake *DeleteInstallationService) DeleteInstallationAssetCollectionReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.deleteInstallationAssetCollectionMutex.Lock()
	defer fake.deleteInstallationAssetCollectionMutex.Unlock()
	fake.DeleteInstallationAssetCollectionStub = nil
	fake.deleteInstallationAssetCollectionReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) DeleteInstallationAssetCollectionReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.deleteInstallationAssetCollectionMutex.Lock()
	defer fake.deleteInstallationAssetCollectionMutex.Unlock()
	fake.DeleteInstallationAssetCollectionStub = nil
	if fake.deleteInstallationAssetCollectionReturnsOnCall == nil {
		fake.deleteInstallationAssetCollectionReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.deleteInstallationAssetCollectionReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) GetInstallation(arg1 int) (api.InstallationsServiceOutput, error) {
	fake.getInstallationMutex.Lock()
	ret, specificReturn := fake.getInstallationReturnsOnCall[len(fake.getInstallationArgsForCall)]
	fake.getInstallationArgsForCall = append(fake.getInstallationArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("GetInstallation", []interface{}{arg1})
	fake.getInstallationMutex.Unlock()
	if fake.GetInstallationStub != nil {
		return fake.GetInstallationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstallationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeleteInstallationService) GetInstallationCallCount() int {
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	return len(fake.getInstallationArgsForCall)
}

func (fake *DeleteInstallationService) GetInstallationCalls(stub func(int) (api.InstallationsServiceOutput, error)) {
	fake.getInstallationMutex.Lock()
	defer fake.getInstallationMutex.Unlock()
	fake.GetInstallationStub = stub
}

func (fake *DeleteInstallationService) GetInstallationArgsForCall(i int) int {
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	argsForCall := fake.getInstallationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeleteInstallationService) GetInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.getInstallationMutex.Lock()
	defer fake.getInstallationMutex.Unlock()
	fake.GetInstallationStub = nil
	fake.getInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) GetInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.getInstallationMutex.Lock()
	defer fake.getInstallationMutex.Unlock()
	fake.GetInstallationStub = nil
	if fake.getInstallationReturnsOnCall == nil {
		fake.getInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.getInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) GetInstallationLogs(arg1 int) (api.InstallationsServiceOutput, error) {
	fake.getInstallationLogsMutex.Lock()
	ret, specificReturn := fake.getInstallationLogsReturnsOnCall[len(fake.getInstallationLogsArgsForCall)]
	fake.getInstallationLogsArgsForCall = append(fake.getInstallationLogsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("GetInstallationLogs", []interface{}{arg1})
	fake.getInstallationLogsMutex.Unlock()
	if fake.GetInstallationLogsStub != nil {
		return fake.GetInstallationLogsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstallationLogsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeleteInstallationService) GetInstallationLogsCallCount() int {
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	return len(fake.getInstallationLogsArgsForCall)
}

func (fake *DeleteInstallationService) GetInstallationLogsCalls(stub func(int) (api.InstallationsServiceOutput, error)) {
	fake.getInstallationLogsMutex.Lock()
	defer fake.getInstallationLogsMutex.Unlock()
	fake.GetInstallationLogsStub = stub
}

func (fake *DeleteInstallationService) GetInstallationLogsArgsForCall(i int) int {
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	argsForCall := fake.getInstallationLogsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeleteInstallationService) GetInstallationLogsReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.getInstallationLogsMutex.Lock()
	defer fake.getInstallationLogsMutex.Unlock()
	fake.GetInstallationLogsStub = nil
	fake.getInstallationLogsReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) GetInstallationLogsReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.getInstallationLogsMutex.Lock()
	defer fake.getInstallationLogsMutex.Unlock()
	fake.GetInstallationLogsStub = nil
	if fake.getInstallationLogsReturnsOnCall == nil {
		fake.getInstallationLogsReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.getInstallationLogsReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) RunningInstallation() (api.InstallationsServiceOutput, error) {
	fake.runningInstallationMutex.Lock()
	ret, specificReturn := fake.runningInstallationReturnsOnCall[len(fake.runningInstallationArgsForCall)]
	fake.runningInstallationArgsForCall = append(fake.runningInstallationArgsForCall, struct {
	}{})
	fake.recordInvocation("RunningInstallation", []interface{}{})
	fake.runningInstallationMutex.Unlock()
	if fake.RunningInstallationStub != nil {
		return fake.RunningInstallationStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runningInstallationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeleteInstallationService) RunningInstallationCallCount() int {
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	return len(fake.runningInstallationArgsForCall)
}

func (fake *DeleteInstallationService) RunningInstallationCalls(stub func() (api.InstallationsServiceOutput, error)) {
	fake.runningInstallationMutex.Lock()
	defer fake.runningInstallationMutex.Unlock()
	fake.RunningInstallationStub = stub
}

func (fake *DeleteInstallationService) RunningInstallationReturns(result1 api.InstallationsServiceOutput, result2 error) {
	fake.runningInstallationMutex.Lock()
	defer fake.runningInstallationMutex.Unlock()
	fake.RunningInstallationStub = nil
	fake.runningInstallationReturns = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) RunningInstallationReturnsOnCall(i int, result1 api.InstallationsServiceOutput, result2 error) {
	fake.runningInstallationMutex.Lock()
	defer fake.runningInstallationMutex.Unlock()
	fake.RunningInstallationStub = nil
	if fake.runningInstallationReturnsOnCall == nil {
		fake.runningInstallationReturnsOnCall = make(map[int]struct {
			result1 api.InstallationsServiceOutput
			result2 error
		})
	}
	fake.runningInstallationReturnsOnCall[i] = struct {
		result1 api.InstallationsServiceOutput
		result2 error
	}{result1, result2}
}

func (fake *DeleteInstallationService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteInstallationAssetCollectionMutex.RLock()
	defer fake.deleteInstallationAssetCollectionMutex.RUnlock()
	fake.getInstallationMutex.RLock()
	defer fake.getInstallationMutex.RUnlock()
	fake.getInstallationLogsMutex.RLock()
	defer fake.getInstallationLogsMutex.RUnlock()
	fake.runningInstallationMutex.RLock()
	defer fake.runningInstallationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeleteInstallationService) recordInvocation(key string, args []interface{}) {
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
