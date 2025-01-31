// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"net/http"
	"sync"

	"github.com/deorbit/cf-mysql-plugin/cfmysql"
)

type FakeHttpClientFactory struct {
	NewClientStub        func(sslDisabled bool) *http.Client
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		sslDisabled bool
	}
	newClientReturns struct {
		result1 *http.Client
	}
	newClientReturnsOnCall map[int]struct {
		result1 *http.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHttpClientFactory) NewClient(sslDisabled bool) *http.Client {
	fake.newClientMutex.Lock()
	ret, specificReturn := fake.newClientReturnsOnCall[len(fake.newClientArgsForCall)]
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		sslDisabled bool
	}{sslDisabled})
	fake.recordInvocation("NewClient", []interface{}{sslDisabled})
	fake.newClientMutex.Unlock()
	if fake.NewClientStub != nil {
		return fake.NewClientStub(sslDisabled)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newClientReturns.result1
}

func (fake *FakeHttpClientFactory) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

func (fake *FakeHttpClientFactory) NewClientArgsForCall(i int) bool {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return fake.newClientArgsForCall[i].sslDisabled
}

func (fake *FakeHttpClientFactory) NewClientReturns(result1 *http.Client) {
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeHttpClientFactory) NewClientReturnsOnCall(i int, result1 *http.Client) {
	fake.NewClientStub = nil
	if fake.newClientReturnsOnCall == nil {
		fake.newClientReturnsOnCall = make(map[int]struct {
			result1 *http.Client
		})
	}
	fake.newClientReturnsOnCall[i] = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeHttpClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHttpClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.HttpClientFactory = new(FakeHttpClientFactory)
