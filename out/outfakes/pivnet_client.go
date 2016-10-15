// This file was generated by counterfeiter
package outfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
)

type PivnetClient struct {
	DeleteReleaseStub        func(productSlug string, release go_pivnet.Release) error
	deleteReleaseMutex       sync.RWMutex
	deleteReleaseArgsForCall []struct {
		productSlug string
		release     go_pivnet.Release
	}
	deleteReleaseReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PivnetClient) DeleteRelease(productSlug string, release go_pivnet.Release) error {
	fake.deleteReleaseMutex.Lock()
	fake.deleteReleaseArgsForCall = append(fake.deleteReleaseArgsForCall, struct {
		productSlug string
		release     go_pivnet.Release
	}{productSlug, release})
	fake.recordInvocation("DeleteRelease", []interface{}{productSlug, release})
	fake.deleteReleaseMutex.Unlock()
	if fake.DeleteReleaseStub != nil {
		return fake.DeleteReleaseStub(productSlug, release)
	} else {
		return fake.deleteReleaseReturns.result1
	}
}

func (fake *PivnetClient) DeleteReleaseCallCount() int {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return len(fake.deleteReleaseArgsForCall)
}

func (fake *PivnetClient) DeleteReleaseArgsForCall(i int) (string, go_pivnet.Release) {
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return fake.deleteReleaseArgsForCall[i].productSlug, fake.deleteReleaseArgsForCall[i].release
}

func (fake *PivnetClient) DeleteReleaseReturns(result1 error) {
	fake.DeleteReleaseStub = nil
	fake.deleteReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *PivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteReleaseMutex.RLock()
	defer fake.deleteReleaseMutex.RUnlock()
	return fake.invocations
}

func (fake *PivnetClient) recordInvocation(key string, args []interface{}) {
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
