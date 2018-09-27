// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"os/exec"
	"sync"

	"github.com/andreasf/cf-mysql-plugin/cfmysql"
)

type FakeExecWrapper struct {
	LookPathStub        func(file string) (string, error)
	lookPathMutex       sync.RWMutex
	lookPathArgsForCall []struct {
		file string
	}
	lookPathReturns struct {
		result1 string
		result2 error
	}
	lookPathReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RunStub        func(*exec.Cmd) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 *exec.Cmd
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecWrapper) LookPath(file string) (string, error) {
	fake.lookPathMutex.Lock()
	ret, specificReturn := fake.lookPathReturnsOnCall[len(fake.lookPathArgsForCall)]
	fake.lookPathArgsForCall = append(fake.lookPathArgsForCall, struct {
		file string
	}{file})
	fake.recordInvocation("LookPath", []interface{}{file})
	fake.lookPathMutex.Unlock()
	if fake.LookPathStub != nil {
		return fake.LookPathStub(file)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lookPathReturns.result1, fake.lookPathReturns.result2
}

func (fake *FakeExecWrapper) LookPathCallCount() int {
	fake.lookPathMutex.RLock()
	defer fake.lookPathMutex.RUnlock()
	return len(fake.lookPathArgsForCall)
}

func (fake *FakeExecWrapper) LookPathArgsForCall(i int) string {
	fake.lookPathMutex.RLock()
	defer fake.lookPathMutex.RUnlock()
	return fake.lookPathArgsForCall[i].file
}

func (fake *FakeExecWrapper) LookPathReturns(result1 string, result2 error) {
	fake.LookPathStub = nil
	fake.lookPathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecWrapper) LookPathReturnsOnCall(i int, result1 string, result2 error) {
	fake.LookPathStub = nil
	if fake.lookPathReturnsOnCall == nil {
		fake.lookPathReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lookPathReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecWrapper) Run(arg1 *exec.Cmd) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	fake.recordInvocation("Run", []interface{}{arg1})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeExecWrapper) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeExecWrapper) RunArgsForCall(i int) *exec.Cmd {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1
}

func (fake *FakeExecWrapper) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecWrapper) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookPathMutex.RLock()
	defer fake.lookPathMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecWrapper) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.ExecWrapper = new(FakeExecWrapper)
