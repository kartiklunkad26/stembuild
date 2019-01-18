// Code generated by counterfeiter. DO NOT EDIT.
package iaas_clientsfakes

import (
	sync "sync"

	iaas_clients "github.com/cloudfoundry-incubator/stembuild/package_stemcell/iaas_clients"
)

type FakeIaasClient struct {
	FindVMStub        func(string) error
	findVMMutex       sync.RWMutex
	findVMArgsForCall []struct {
		arg1 string
	}
	findVMReturns struct {
		result1 error
	}
	findVMReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateCredentialsStub        func() error
	validateCredentialsMutex       sync.RWMutex
	validateCredentialsArgsForCall []struct {
	}
	validateCredentialsReturns struct {
		result1 error
	}
	validateCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateUrlStub        func() error
	validateUrlMutex       sync.RWMutex
	validateUrlArgsForCall []struct {
	}
	validateUrlReturns struct {
		result1 error
	}
	validateUrlReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIaasClient) FindVM(arg1 string) error {
	fake.findVMMutex.Lock()
	ret, specificReturn := fake.findVMReturnsOnCall[len(fake.findVMArgsForCall)]
	fake.findVMArgsForCall = append(fake.findVMArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindVM", []interface{}{arg1})
	fake.findVMMutex.Unlock()
	if fake.FindVMStub != nil {
		return fake.FindVMStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.findVMReturns
	return fakeReturns.result1
}

func (fake *FakeIaasClient) FindVMCallCount() int {
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	return len(fake.findVMArgsForCall)
}

func (fake *FakeIaasClient) FindVMCalls(stub func(string) error) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = stub
}

func (fake *FakeIaasClient) FindVMArgsForCall(i int) string {
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	argsForCall := fake.findVMArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIaasClient) FindVMReturns(result1 error) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = nil
	fake.findVMReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) FindVMReturnsOnCall(i int, result1 error) {
	fake.findVMMutex.Lock()
	defer fake.findVMMutex.Unlock()
	fake.FindVMStub = nil
	if fake.findVMReturnsOnCall == nil {
		fake.findVMReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.findVMReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) ValidateCredentials() error {
	fake.validateCredentialsMutex.Lock()
	ret, specificReturn := fake.validateCredentialsReturnsOnCall[len(fake.validateCredentialsArgsForCall)]
	fake.validateCredentialsArgsForCall = append(fake.validateCredentialsArgsForCall, struct {
	}{})
	fake.recordInvocation("ValidateCredentials", []interface{}{})
	fake.validateCredentialsMutex.Unlock()
	if fake.ValidateCredentialsStub != nil {
		return fake.ValidateCredentialsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeIaasClient) ValidateCredentialsCallCount() int {
	fake.validateCredentialsMutex.RLock()
	defer fake.validateCredentialsMutex.RUnlock()
	return len(fake.validateCredentialsArgsForCall)
}

func (fake *FakeIaasClient) ValidateCredentialsCalls(stub func() error) {
	fake.validateCredentialsMutex.Lock()
	defer fake.validateCredentialsMutex.Unlock()
	fake.ValidateCredentialsStub = stub
}

func (fake *FakeIaasClient) ValidateCredentialsReturns(result1 error) {
	fake.validateCredentialsMutex.Lock()
	defer fake.validateCredentialsMutex.Unlock()
	fake.ValidateCredentialsStub = nil
	fake.validateCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) ValidateCredentialsReturnsOnCall(i int, result1 error) {
	fake.validateCredentialsMutex.Lock()
	defer fake.validateCredentialsMutex.Unlock()
	fake.ValidateCredentialsStub = nil
	if fake.validateCredentialsReturnsOnCall == nil {
		fake.validateCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) ValidateUrl() error {
	fake.validateUrlMutex.Lock()
	ret, specificReturn := fake.validateUrlReturnsOnCall[len(fake.validateUrlArgsForCall)]
	fake.validateUrlArgsForCall = append(fake.validateUrlArgsForCall, struct {
	}{})
	fake.recordInvocation("ValidateUrl", []interface{}{})
	fake.validateUrlMutex.Unlock()
	if fake.ValidateUrlStub != nil {
		return fake.ValidateUrlStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateUrlReturns
	return fakeReturns.result1
}

func (fake *FakeIaasClient) ValidateUrlCallCount() int {
	fake.validateUrlMutex.RLock()
	defer fake.validateUrlMutex.RUnlock()
	return len(fake.validateUrlArgsForCall)
}

func (fake *FakeIaasClient) ValidateUrlCalls(stub func() error) {
	fake.validateUrlMutex.Lock()
	defer fake.validateUrlMutex.Unlock()
	fake.ValidateUrlStub = stub
}

func (fake *FakeIaasClient) ValidateUrlReturns(result1 error) {
	fake.validateUrlMutex.Lock()
	defer fake.validateUrlMutex.Unlock()
	fake.ValidateUrlStub = nil
	fake.validateUrlReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) ValidateUrlReturnsOnCall(i int, result1 error) {
	fake.validateUrlMutex.Lock()
	defer fake.validateUrlMutex.Unlock()
	fake.ValidateUrlStub = nil
	if fake.validateUrlReturnsOnCall == nil {
		fake.validateUrlReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateUrlReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIaasClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findVMMutex.RLock()
	defer fake.findVMMutex.RUnlock()
	fake.validateCredentialsMutex.RLock()
	defer fake.validateCredentialsMutex.RUnlock()
	fake.validateUrlMutex.RLock()
	defer fake.validateUrlMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIaasClient) recordInvocation(key string, args []interface{}) {
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

var _ iaas_clients.IaasClient = new(FakeIaasClient)