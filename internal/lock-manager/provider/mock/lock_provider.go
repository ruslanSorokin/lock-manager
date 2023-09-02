// Code generated by mockery. DO NOT EDIT.

package mock

import (
	context "context"

	model "github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	mock "github.com/stretchr/testify/mock"
)

// LockProvider is an autogenerated mock type for the LockProviderI type
type LockProvider struct {
	mock.Mock
}

type LockProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *LockProvider) EXPECT() *LockProvider_Expecter {
	return &LockProvider_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, lock
func (_m *LockProvider) Create(ctx context.Context, lock *model.Lock) error {
	ret := _m.Called(ctx, lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Lock) error); ok {
		r0 = rf(ctx, lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LockProvider_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type LockProvider_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - lock *model.Lock
func (_e *LockProvider_Expecter) Create(ctx interface{}, lock interface{}) *LockProvider_Create_Call {
	return &LockProvider_Create_Call{Call: _e.mock.On("Create", ctx, lock)}
}

func (_c *LockProvider_Create_Call) Run(run func(ctx context.Context, lock *model.Lock)) *LockProvider_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Lock))
	})
	return _c
}

func (_c *LockProvider_Create_Call) Return(_a0 error) *LockProvider_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LockProvider_Create_Call) RunAndReturn(run func(context.Context, *model.Lock) error) *LockProvider_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, resourceID
func (_m *LockProvider) Delete(ctx context.Context, resourceID string) error {
	ret := _m.Called(ctx, resourceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, resourceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LockProvider_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type LockProvider_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - resourceID string
func (_e *LockProvider_Expecter) Delete(ctx interface{}, resourceID interface{}) *LockProvider_Delete_Call {
	return &LockProvider_Delete_Call{Call: _e.mock.On("Delete", ctx, resourceID)}
}

func (_c *LockProvider_Delete_Call) Run(run func(ctx context.Context, resourceID string)) *LockProvider_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LockProvider_Delete_Call) Return(_a0 error) *LockProvider_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LockProvider_Delete_Call) RunAndReturn(run func(context.Context, string) error) *LockProvider_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteIfTokenMatches provides a mock function with given fields: ctx, lock
func (_m *LockProvider) DeleteIfTokenMatches(ctx context.Context, lock *model.Lock) error {
	ret := _m.Called(ctx, lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Lock) error); ok {
		r0 = rf(ctx, lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LockProvider_DeleteIfTokenMatches_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteIfTokenMatches'
type LockProvider_DeleteIfTokenMatches_Call struct {
	*mock.Call
}

// DeleteIfTokenMatches is a helper method to define mock.On call
//   - ctx context.Context
//   - lock *model.Lock
func (_e *LockProvider_Expecter) DeleteIfTokenMatches(ctx interface{}, lock interface{}) *LockProvider_DeleteIfTokenMatches_Call {
	return &LockProvider_DeleteIfTokenMatches_Call{Call: _e.mock.On("DeleteIfTokenMatches", ctx, lock)}
}

func (_c *LockProvider_DeleteIfTokenMatches_Call) Run(run func(ctx context.Context, lock *model.Lock)) *LockProvider_DeleteIfTokenMatches_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Lock))
	})
	return _c
}

func (_c *LockProvider_DeleteIfTokenMatches_Call) Return(_a0 error) *LockProvider_DeleteIfTokenMatches_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LockProvider_DeleteIfTokenMatches_Call) RunAndReturn(run func(context.Context, *model.Lock) error) *LockProvider_DeleteIfTokenMatches_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, resourceID
func (_m *LockProvider) Get(ctx context.Context, resourceID string) (*model.Lock, error) {
	ret := _m.Called(ctx, resourceID)

	var r0 *model.Lock
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Lock, error)); ok {
		return rf(ctx, resourceID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Lock); ok {
		r0 = rf(ctx, resourceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Lock)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, resourceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockProvider_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type LockProvider_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - resourceID string
func (_e *LockProvider_Expecter) Get(ctx interface{}, resourceID interface{}) *LockProvider_Get_Call {
	return &LockProvider_Get_Call{Call: _e.mock.On("Get", ctx, resourceID)}
}

func (_c *LockProvider_Get_Call) Run(run func(ctx context.Context, resourceID string)) *LockProvider_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LockProvider_Get_Call) Return(_a0 *model.Lock, _a1 error) *LockProvider_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *LockProvider_Get_Call) RunAndReturn(run func(context.Context, string) (*model.Lock, error)) *LockProvider_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewLockProvider creates a new instance of LockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLockProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *LockProvider {
	mock := &LockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
