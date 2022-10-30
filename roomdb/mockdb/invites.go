// Code generated by counterfeiter. DO NOT EDIT.
package mockdb

import (
	"context"
	"sync"

	"github.com/ssb-ngi-pointer/go-ssb-room/v2/roomdb"
	refs "go.mindeco.de/ssb-refs"
)

type FakeInvitesService struct {
	ConsumeStub        func(context.Context, string, refs.FeedRef) (roomdb.Invite, error)
	consumeMutex       sync.RWMutex
	consumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 refs.FeedRef
	}
	consumeReturns struct {
		result1 roomdb.Invite
		result2 error
	}
	consumeReturnsOnCall map[int]struct {
		result1 roomdb.Invite
		result2 error
	}
	CountStub        func(context.Context, bool) (uint, error)
	countMutex       sync.RWMutex
	countArgsForCall []struct {
		arg1 context.Context
		arg2 bool
	}
	countReturns struct {
		result1 uint
		result2 error
	}
	countReturnsOnCall map[int]struct {
		result1 uint
		result2 error
	}
	CreateStub        func(context.Context, int64) (string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	createReturns struct {
		result1 string
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetByIDStub        func(context.Context, int64) (roomdb.Invite, error)
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getByIDReturns struct {
		result1 roomdb.Invite
		result2 error
	}
	getByIDReturnsOnCall map[int]struct {
		result1 roomdb.Invite
		result2 error
	}
	GetByTokenStub        func(context.Context, string) (roomdb.Invite, error)
	getByTokenMutex       sync.RWMutex
	getByTokenArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getByTokenReturns struct {
		result1 roomdb.Invite
		result2 error
	}
	getByTokenReturnsOnCall map[int]struct {
		result1 roomdb.Invite
		result2 error
	}
	ListStub        func(context.Context) ([]roomdb.Invite, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 []roomdb.Invite
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []roomdb.Invite
		result2 error
	}
	RevokeStub        func(context.Context, int64) error
	revokeMutex       sync.RWMutex
	revokeArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	revokeReturns struct {
		result1 error
	}
	revokeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInvitesService) Consume(arg1 context.Context, arg2 string, arg3 refs.FeedRef) (roomdb.Invite, error) {
	fake.consumeMutex.Lock()
	ret, specificReturn := fake.consumeReturnsOnCall[len(fake.consumeArgsForCall)]
	fake.consumeArgsForCall = append(fake.consumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 refs.FeedRef
	}{arg1, arg2, arg3})
	stub := fake.ConsumeStub
	fakeReturns := fake.consumeReturns
	fake.recordInvocation("Consume", []interface{}{arg1, arg2, arg3})
	fake.consumeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) ConsumeCallCount() int {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	return len(fake.consumeArgsForCall)
}

func (fake *FakeInvitesService) ConsumeCalls(stub func(context.Context, string, refs.FeedRef) (roomdb.Invite, error)) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = stub
}

func (fake *FakeInvitesService) ConsumeArgsForCall(i int) (context.Context, string, refs.FeedRef) {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	argsForCall := fake.consumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInvitesService) ConsumeReturns(result1 roomdb.Invite, result2 error) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = nil
	fake.consumeReturns = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) ConsumeReturnsOnCall(i int, result1 roomdb.Invite, result2 error) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = nil
	if fake.consumeReturnsOnCall == nil {
		fake.consumeReturnsOnCall = make(map[int]struct {
			result1 roomdb.Invite
			result2 error
		})
	}
	fake.consumeReturnsOnCall[i] = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) Count(arg1 context.Context, arg2 bool) (uint, error) {
	fake.countMutex.Lock()
	ret, specificReturn := fake.countReturnsOnCall[len(fake.countArgsForCall)]
	fake.countArgsForCall = append(fake.countArgsForCall, struct {
		arg1 context.Context
		arg2 bool
	}{arg1, arg2})
	stub := fake.CountStub
	fakeReturns := fake.countReturns
	fake.recordInvocation("Count", []interface{}{arg1, arg2})
	fake.countMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) CountCallCount() int {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return len(fake.countArgsForCall)
}

func (fake *FakeInvitesService) CountCalls(stub func(context.Context, bool) (uint, error)) {
	fake.countMutex.Lock()
	defer fake.countMutex.Unlock()
	fake.CountStub = stub
}

func (fake *FakeInvitesService) CountArgsForCall(i int) (context.Context, bool) {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	argsForCall := fake.countArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInvitesService) CountReturns(result1 uint, result2 error) {
	fake.countMutex.Lock()
	defer fake.countMutex.Unlock()
	fake.CountStub = nil
	fake.countReturns = struct {
		result1 uint
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) CountReturnsOnCall(i int, result1 uint, result2 error) {
	fake.countMutex.Lock()
	defer fake.countMutex.Unlock()
	fake.CountStub = nil
	if fake.countReturnsOnCall == nil {
		fake.countReturnsOnCall = make(map[int]struct {
			result1 uint
			result2 error
		})
	}
	fake.countReturnsOnCall[i] = struct {
		result1 uint
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) Create(arg1 context.Context, arg2 int64) (string, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeInvitesService) CreateCalls(stub func(context.Context, int64) (string, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeInvitesService) CreateArgsForCall(i int) (context.Context, int64) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInvitesService) CreateReturns(result1 string, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) CreateReturnsOnCall(i int, result1 string, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) GetByID(arg1 context.Context, arg2 int64) (roomdb.Invite, error) {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.GetByIDStub
	fakeReturns := fake.getByIDReturns
	fake.recordInvocation("GetByID", []interface{}{arg1, arg2})
	fake.getByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeInvitesService) GetByIDCalls(stub func(context.Context, int64) (roomdb.Invite, error)) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = stub
}

func (fake *FakeInvitesService) GetByIDArgsForCall(i int) (context.Context, int64) {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	argsForCall := fake.getByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInvitesService) GetByIDReturns(result1 roomdb.Invite, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) GetByIDReturnsOnCall(i int, result1 roomdb.Invite, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 roomdb.Invite
			result2 error
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) GetByToken(arg1 context.Context, arg2 string) (roomdb.Invite, error) {
	fake.getByTokenMutex.Lock()
	ret, specificReturn := fake.getByTokenReturnsOnCall[len(fake.getByTokenArgsForCall)]
	fake.getByTokenArgsForCall = append(fake.getByTokenArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetByTokenStub
	fakeReturns := fake.getByTokenReturns
	fake.recordInvocation("GetByToken", []interface{}{arg1, arg2})
	fake.getByTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) GetByTokenCallCount() int {
	fake.getByTokenMutex.RLock()
	defer fake.getByTokenMutex.RUnlock()
	return len(fake.getByTokenArgsForCall)
}

func (fake *FakeInvitesService) GetByTokenCalls(stub func(context.Context, string) (roomdb.Invite, error)) {
	fake.getByTokenMutex.Lock()
	defer fake.getByTokenMutex.Unlock()
	fake.GetByTokenStub = stub
}

func (fake *FakeInvitesService) GetByTokenArgsForCall(i int) (context.Context, string) {
	fake.getByTokenMutex.RLock()
	defer fake.getByTokenMutex.RUnlock()
	argsForCall := fake.getByTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInvitesService) GetByTokenReturns(result1 roomdb.Invite, result2 error) {
	fake.getByTokenMutex.Lock()
	defer fake.getByTokenMutex.Unlock()
	fake.GetByTokenStub = nil
	fake.getByTokenReturns = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) GetByTokenReturnsOnCall(i int, result1 roomdb.Invite, result2 error) {
	fake.getByTokenMutex.Lock()
	defer fake.getByTokenMutex.Unlock()
	fake.GetByTokenStub = nil
	if fake.getByTokenReturnsOnCall == nil {
		fake.getByTokenReturnsOnCall = make(map[int]struct {
			result1 roomdb.Invite
			result2 error
		})
	}
	fake.getByTokenReturnsOnCall[i] = struct {
		result1 roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) List(arg1 context.Context) ([]roomdb.Invite, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInvitesService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeInvitesService) ListCalls(stub func(context.Context) ([]roomdb.Invite, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeInvitesService) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInvitesService) ListReturns(result1 []roomdb.Invite, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) ListReturnsOnCall(i int, result1 []roomdb.Invite, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []roomdb.Invite
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []roomdb.Invite
		result2 error
	}{result1, result2}
}

func (fake *FakeInvitesService) Revoke(arg1 context.Context, arg2 int64) error {
	fake.revokeMutex.Lock()
	ret, specificReturn := fake.revokeReturnsOnCall[len(fake.revokeArgsForCall)]
	fake.revokeArgsForCall = append(fake.revokeArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.RevokeStub
	fakeReturns := fake.revokeReturns
	fake.recordInvocation("Revoke", []interface{}{arg1, arg2})
	fake.revokeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInvitesService) RevokeCallCount() int {
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	return len(fake.revokeArgsForCall)
}

func (fake *FakeInvitesService) RevokeCalls(stub func(context.Context, int64) error) {
	fake.revokeMutex.Lock()
	defer fake.revokeMutex.Unlock()
	fake.RevokeStub = stub
}

func (fake *FakeInvitesService) RevokeArgsForCall(i int) (context.Context, int64) {
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	argsForCall := fake.revokeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInvitesService) RevokeReturns(result1 error) {
	fake.revokeMutex.Lock()
	defer fake.revokeMutex.Unlock()
	fake.RevokeStub = nil
	fake.revokeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInvitesService) RevokeReturnsOnCall(i int, result1 error) {
	fake.revokeMutex.Lock()
	defer fake.revokeMutex.Unlock()
	fake.RevokeStub = nil
	if fake.revokeReturnsOnCall == nil {
		fake.revokeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.revokeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInvitesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.getByTokenMutex.RLock()
	defer fake.getByTokenMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.revokeMutex.RLock()
	defer fake.revokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInvitesService) recordInvocation(key string, args []interface{}) {
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

var _ roomdb.InvitesService = new(FakeInvitesService)
