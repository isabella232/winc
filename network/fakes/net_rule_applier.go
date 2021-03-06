// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/network"
	"code.cloudfoundry.org/winc/network/netrules"
	"github.com/Microsoft/hcsshim"
)

type NetRuleApplier struct {
	InStub        func(netrules.NetIn, string) (*hcsshim.NatPolicy, *hcsshim.ACLPolicy, error)
	inMutex       sync.RWMutex
	inArgsForCall []struct {
		arg1 netrules.NetIn
		arg2 string
	}
	inReturns struct {
		result1 *hcsshim.NatPolicy
		result2 *hcsshim.ACLPolicy
		result3 error
	}
	inReturnsOnCall map[int]struct {
		result1 *hcsshim.NatPolicy
		result2 *hcsshim.ACLPolicy
		result3 error
	}
	OutStub        func(netrules.NetOut, string) (*hcsshim.ACLPolicy, error)
	outMutex       sync.RWMutex
	outArgsForCall []struct {
		arg1 netrules.NetOut
		arg2 string
	}
	outReturns struct {
		result1 *hcsshim.ACLPolicy
		result2 error
	}
	outReturnsOnCall map[int]struct {
		result1 *hcsshim.ACLPolicy
		result2 error
	}
	CleanupStub        func() error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct{}
	cleanupReturns     struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	OpenPortStub        func(port uint32) error
	openPortMutex       sync.RWMutex
	openPortArgsForCall []struct {
		port uint32
	}
	openPortReturns struct {
		result1 error
	}
	openPortReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetRuleApplier) In(arg1 netrules.NetIn, arg2 string) (*hcsshim.NatPolicy, *hcsshim.ACLPolicy, error) {
	fake.inMutex.Lock()
	ret, specificReturn := fake.inReturnsOnCall[len(fake.inArgsForCall)]
	fake.inArgsForCall = append(fake.inArgsForCall, struct {
		arg1 netrules.NetIn
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("In", []interface{}{arg1, arg2})
	fake.inMutex.Unlock()
	if fake.InStub != nil {
		return fake.InStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.inReturns.result1, fake.inReturns.result2, fake.inReturns.result3
}

func (fake *NetRuleApplier) InCallCount() int {
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	return len(fake.inArgsForCall)
}

func (fake *NetRuleApplier) InArgsForCall(i int) (netrules.NetIn, string) {
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	return fake.inArgsForCall[i].arg1, fake.inArgsForCall[i].arg2
}

func (fake *NetRuleApplier) InReturns(result1 *hcsshim.NatPolicy, result2 *hcsshim.ACLPolicy, result3 error) {
	fake.InStub = nil
	fake.inReturns = struct {
		result1 *hcsshim.NatPolicy
		result2 *hcsshim.ACLPolicy
		result3 error
	}{result1, result2, result3}
}

func (fake *NetRuleApplier) InReturnsOnCall(i int, result1 *hcsshim.NatPolicy, result2 *hcsshim.ACLPolicy, result3 error) {
	fake.InStub = nil
	if fake.inReturnsOnCall == nil {
		fake.inReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.NatPolicy
			result2 *hcsshim.ACLPolicy
			result3 error
		})
	}
	fake.inReturnsOnCall[i] = struct {
		result1 *hcsshim.NatPolicy
		result2 *hcsshim.ACLPolicy
		result3 error
	}{result1, result2, result3}
}

func (fake *NetRuleApplier) Out(arg1 netrules.NetOut, arg2 string) (*hcsshim.ACLPolicy, error) {
	fake.outMutex.Lock()
	ret, specificReturn := fake.outReturnsOnCall[len(fake.outArgsForCall)]
	fake.outArgsForCall = append(fake.outArgsForCall, struct {
		arg1 netrules.NetOut
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Out", []interface{}{arg1, arg2})
	fake.outMutex.Unlock()
	if fake.OutStub != nil {
		return fake.OutStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.outReturns.result1, fake.outReturns.result2
}

func (fake *NetRuleApplier) OutCallCount() int {
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	return len(fake.outArgsForCall)
}

func (fake *NetRuleApplier) OutArgsForCall(i int) (netrules.NetOut, string) {
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	return fake.outArgsForCall[i].arg1, fake.outArgsForCall[i].arg2
}

func (fake *NetRuleApplier) OutReturns(result1 *hcsshim.ACLPolicy, result2 error) {
	fake.OutStub = nil
	fake.outReturns = struct {
		result1 *hcsshim.ACLPolicy
		result2 error
	}{result1, result2}
}

func (fake *NetRuleApplier) OutReturnsOnCall(i int, result1 *hcsshim.ACLPolicy, result2 error) {
	fake.OutStub = nil
	if fake.outReturnsOnCall == nil {
		fake.outReturnsOnCall = make(map[int]struct {
			result1 *hcsshim.ACLPolicy
			result2 error
		})
	}
	fake.outReturnsOnCall[i] = struct {
		result1 *hcsshim.ACLPolicy
		result2 error
	}{result1, result2}
}

func (fake *NetRuleApplier) Cleanup() error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct{}{})
	fake.recordInvocation("Cleanup", []interface{}{})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *NetRuleApplier) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *NetRuleApplier) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetRuleApplier) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetRuleApplier) OpenPort(port uint32) error {
	fake.openPortMutex.Lock()
	ret, specificReturn := fake.openPortReturnsOnCall[len(fake.openPortArgsForCall)]
	fake.openPortArgsForCall = append(fake.openPortArgsForCall, struct {
		port uint32
	}{port})
	fake.recordInvocation("OpenPort", []interface{}{port})
	fake.openPortMutex.Unlock()
	if fake.OpenPortStub != nil {
		return fake.OpenPortStub(port)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openPortReturns.result1
}

func (fake *NetRuleApplier) OpenPortCallCount() int {
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	return len(fake.openPortArgsForCall)
}

func (fake *NetRuleApplier) OpenPortArgsForCall(i int) uint32 {
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	return fake.openPortArgsForCall[i].port
}

func (fake *NetRuleApplier) OpenPortReturns(result1 error) {
	fake.OpenPortStub = nil
	fake.openPortReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetRuleApplier) OpenPortReturnsOnCall(i int, result1 error) {
	fake.OpenPortStub = nil
	if fake.openPortReturnsOnCall == nil {
		fake.openPortReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openPortReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetRuleApplier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inMutex.RLock()
	defer fake.inMutex.RUnlock()
	fake.outMutex.RLock()
	defer fake.outMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.openPortMutex.RLock()
	defer fake.openPortMutex.RUnlock()
	return fake.invocations
}

func (fake *NetRuleApplier) recordInvocation(key string, args []interface{}) {
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

var _ network.NetRuleApplier = new(NetRuleApplier)
