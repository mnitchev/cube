// Code generated by counterfeiter. DO NOT EDIT.
package cubefakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/runtimeschema/cc_messages"
	"github.com/julz/cube"
	"github.com/julz/cube/opi"
)

type FakeBackend struct {
	CreateStagingTaskStub        func(string, cc_messages.StagingRequestFromCC) (opi.Task, error)
	createStagingTaskMutex       sync.RWMutex
	createStagingTaskArgsForCall []struct {
		arg1 string
		arg2 cc_messages.StagingRequestFromCC
	}
	createStagingTaskReturns struct {
		result1 opi.Task
		result2 error
	}
	createStagingTaskReturnsOnCall map[int]struct {
		result1 opi.Task
		result2 error
	}
	BuildStagingResponseStub        func(*models.TaskCallbackResponse) (cc_messages.StagingResponseForCC, error)
	buildStagingResponseMutex       sync.RWMutex
	buildStagingResponseArgsForCall []struct {
		arg1 *models.TaskCallbackResponse
	}
	buildStagingResponseReturns struct {
		result1 cc_messages.StagingResponseForCC
		result2 error
	}
	buildStagingResponseReturnsOnCall map[int]struct {
		result1 cc_messages.StagingResponseForCC
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackend) CreateStagingTask(arg1 string, arg2 cc_messages.StagingRequestFromCC) (opi.Task, error) {
	fake.createStagingTaskMutex.Lock()
	ret, specificReturn := fake.createStagingTaskReturnsOnCall[len(fake.createStagingTaskArgsForCall)]
	fake.createStagingTaskArgsForCall = append(fake.createStagingTaskArgsForCall, struct {
		arg1 string
		arg2 cc_messages.StagingRequestFromCC
	}{arg1, arg2})
	fake.recordInvocation("CreateStagingTask", []interface{}{arg1, arg2})
	fake.createStagingTaskMutex.Unlock()
	if fake.CreateStagingTaskStub != nil {
		return fake.CreateStagingTaskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createStagingTaskReturns.result1, fake.createStagingTaskReturns.result2
}

func (fake *FakeBackend) CreateStagingTaskCallCount() int {
	fake.createStagingTaskMutex.RLock()
	defer fake.createStagingTaskMutex.RUnlock()
	return len(fake.createStagingTaskArgsForCall)
}

func (fake *FakeBackend) CreateStagingTaskArgsForCall(i int) (string, cc_messages.StagingRequestFromCC) {
	fake.createStagingTaskMutex.RLock()
	defer fake.createStagingTaskMutex.RUnlock()
	return fake.createStagingTaskArgsForCall[i].arg1, fake.createStagingTaskArgsForCall[i].arg2
}

func (fake *FakeBackend) CreateStagingTaskReturns(result1 opi.Task, result2 error) {
	fake.CreateStagingTaskStub = nil
	fake.createStagingTaskReturns = struct {
		result1 opi.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) CreateStagingTaskReturnsOnCall(i int, result1 opi.Task, result2 error) {
	fake.CreateStagingTaskStub = nil
	if fake.createStagingTaskReturnsOnCall == nil {
		fake.createStagingTaskReturnsOnCall = make(map[int]struct {
			result1 opi.Task
			result2 error
		})
	}
	fake.createStagingTaskReturnsOnCall[i] = struct {
		result1 opi.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) BuildStagingResponse(arg1 *models.TaskCallbackResponse) (cc_messages.StagingResponseForCC, error) {
	fake.buildStagingResponseMutex.Lock()
	ret, specificReturn := fake.buildStagingResponseReturnsOnCall[len(fake.buildStagingResponseArgsForCall)]
	fake.buildStagingResponseArgsForCall = append(fake.buildStagingResponseArgsForCall, struct {
		arg1 *models.TaskCallbackResponse
	}{arg1})
	fake.recordInvocation("BuildStagingResponse", []interface{}{arg1})
	fake.buildStagingResponseMutex.Unlock()
	if fake.BuildStagingResponseStub != nil {
		return fake.BuildStagingResponseStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildStagingResponseReturns.result1, fake.buildStagingResponseReturns.result2
}

func (fake *FakeBackend) BuildStagingResponseCallCount() int {
	fake.buildStagingResponseMutex.RLock()
	defer fake.buildStagingResponseMutex.RUnlock()
	return len(fake.buildStagingResponseArgsForCall)
}

func (fake *FakeBackend) BuildStagingResponseArgsForCall(i int) *models.TaskCallbackResponse {
	fake.buildStagingResponseMutex.RLock()
	defer fake.buildStagingResponseMutex.RUnlock()
	return fake.buildStagingResponseArgsForCall[i].arg1
}

func (fake *FakeBackend) BuildStagingResponseReturns(result1 cc_messages.StagingResponseForCC, result2 error) {
	fake.BuildStagingResponseStub = nil
	fake.buildStagingResponseReturns = struct {
		result1 cc_messages.StagingResponseForCC
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) BuildStagingResponseReturnsOnCall(i int, result1 cc_messages.StagingResponseForCC, result2 error) {
	fake.BuildStagingResponseStub = nil
	if fake.buildStagingResponseReturnsOnCall == nil {
		fake.buildStagingResponseReturnsOnCall = make(map[int]struct {
			result1 cc_messages.StagingResponseForCC
			result2 error
		})
	}
	fake.buildStagingResponseReturnsOnCall[i] = struct {
		result1 cc_messages.StagingResponseForCC
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createStagingTaskMutex.RLock()
	defer fake.createStagingTaskMutex.RUnlock()
	fake.buildStagingResponseMutex.RLock()
	defer fake.buildStagingResponseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackend) recordInvocation(key string, args []interface{}) {
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

var _ cube.Backend = new(FakeBackend)
