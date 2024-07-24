// Code generated by mockery v2.42.1. DO NOT EDIT.

package sessionManager

import (
	context "context"

	errs "github.com/raffops/chat/pkg/errs"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// ExpireAt provides a mock function with given fields: ctx, key, at
func (_m *Repository) ExpireAt(ctx context.Context, key string, at time.Time) errs.ChatError {
	ret := _m.Called(ctx, key, at)

	if len(ret) == 0 {
		panic("no return value specified for ExpireAt")
	}

	var r0 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) errs.ChatError); ok {
		r0 = rf(ctx, key, at)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errs.ChatError)
		}
	}

	return r0
}

// Repository_ExpireAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpireAt'
type Repository_ExpireAt_Call struct {
	*mock.Call
}

// ExpireAt is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - at time.Time
func (_e *Repository_Expecter) ExpireAt(ctx interface{}, key interface{}, at interface{}) *Repository_ExpireAt_Call {
	return &Repository_ExpireAt_Call{Call: _e.mock.On("ExpireAt", ctx, key, at)}
}

func (_c *Repository_ExpireAt_Call) Run(run func(ctx context.Context, key string, at time.Time)) *Repository_ExpireAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Time))
	})
	return _c
}

func (_c *Repository_ExpireAt_Call) Return(_a0 errs.ChatError) *Repository_ExpireAt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_ExpireAt_Call) RunAndReturn(run func(context.Context, string, time.Time) errs.ChatError) *Repository_ExpireAt_Call {
	_c.Call.Return(run)
	return _c
}

// GetEncryptedHashmap provides a mock function with given fields: ctx, key, secretKey
func (_m *Repository) GetEncryptedHashmap(ctx context.Context, key string, secretKey string) (map[string]interface{}, errs.ChatError) {
	ret := _m.Called(ctx, key, secretKey)

	if len(ret) == 0 {
		panic("no return value specified for GetEncryptedHashmap")
	}

	var r0 map[string]interface{}
	var r1 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (map[string]interface{}, errs.ChatError)); ok {
		return rf(ctx, key, secretKey)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) map[string]interface{}); ok {
		r0 = rf(ctx, key, secretKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) errs.ChatError); ok {
		r1 = rf(ctx, key, secretKey)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errs.ChatError)
		}
	}

	return r0, r1
}

// Repository_GetEncryptedHashmap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEncryptedHashmap'
type Repository_GetEncryptedHashmap_Call struct {
	*mock.Call
}

// GetEncryptedHashmap is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - secretKey string
func (_e *Repository_Expecter) GetEncryptedHashmap(ctx interface{}, key interface{}, secretKey interface{}) *Repository_GetEncryptedHashmap_Call {
	return &Repository_GetEncryptedHashmap_Call{Call: _e.mock.On("GetEncryptedHashmap", ctx, key, secretKey)}
}

func (_c *Repository_GetEncryptedHashmap_Call) Run(run func(ctx context.Context, key string, secretKey string)) *Repository_GetEncryptedHashmap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Repository_GetEncryptedHashmap_Call) Return(_a0 map[string]interface{}, _a1 errs.ChatError) *Repository_GetEncryptedHashmap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetEncryptedHashmap_Call) RunAndReturn(run func(context.Context, string, string) (map[string]interface{}, errs.ChatError)) *Repository_GetEncryptedHashmap_Call {
	_c.Call.Return(run)
	return _c
}

// Hashget provides a mock function with given fields: ctx, key, field
func (_m *Repository) Hashget(ctx context.Context, key string, field string) (string, errs.ChatError) {
	ret := _m.Called(ctx, key, field)

	if len(ret) == 0 {
		panic("no return value specified for Hashget")
	}

	var r0 string
	var r1 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, errs.ChatError)); ok {
		return rf(ctx, key, field)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) errs.ChatError); ok {
		r1 = rf(ctx, key, field)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errs.ChatError)
		}
	}

	return r0, r1
}

// Repository_Hashget_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hashget'
type Repository_Hashget_Call struct {
	*mock.Call
}

// Hashget is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - field string
func (_e *Repository_Expecter) Hashget(ctx interface{}, key interface{}, field interface{}) *Repository_Hashget_Call {
	return &Repository_Hashget_Call{Call: _e.mock.On("Hashget", ctx, key, field)}
}

func (_c *Repository_Hashget_Call) Run(run func(ctx context.Context, key string, field string)) *Repository_Hashget_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Repository_Hashget_Call) Return(_a0 string, _a1 errs.ChatError) *Repository_Hashget_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Hashget_Call) RunAndReturn(run func(context.Context, string, string) (string, errs.ChatError)) *Repository_Hashget_Call {
	_c.Call.Return(run)
	return _c
}

// Hashset provides a mock function with given fields: ctx, key, value
func (_m *Repository) Hashset(ctx context.Context, key string, value map[string]interface{}) errs.ChatError {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Hashset")
	}

	var r0 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) errs.ChatError); ok {
		r0 = rf(ctx, key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errs.ChatError)
		}
	}

	return r0
}

// Repository_Hashset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hashset'
type Repository_Hashset_Call struct {
	*mock.Call
}

// Hashset is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value map[string]interface{}
func (_e *Repository_Expecter) Hashset(ctx interface{}, key interface{}, value interface{}) *Repository_Hashset_Call {
	return &Repository_Hashset_Call{Call: _e.mock.On("Hashset", ctx, key, value)}
}

func (_c *Repository_Hashset_Call) Run(run func(ctx context.Context, key string, value map[string]interface{})) *Repository_Hashset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *Repository_Hashset_Call) Return(_a0 errs.ChatError) *Repository_Hashset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Hashset_Call) RunAndReturn(run func(context.Context, string, map[string]interface{}) errs.ChatError) *Repository_Hashset_Call {
	_c.Call.Return(run)
	return _c
}

// SetAppend provides a mock function with given fields: ctx, key, value
func (_m *Repository) SetAppend(ctx context.Context, key string, value string) errs.ChatError {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for SetAppend")
	}

	var r0 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, string) errs.ChatError); ok {
		r0 = rf(ctx, key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errs.ChatError)
		}
	}

	return r0
}

// Repository_SetAppend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAppend'
type Repository_SetAppend_Call struct {
	*mock.Call
}

// SetAppend is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value string
func (_e *Repository_Expecter) SetAppend(ctx interface{}, key interface{}, value interface{}) *Repository_SetAppend_Call {
	return &Repository_SetAppend_Call{Call: _e.mock.On("SetAppend", ctx, key, value)}
}

func (_c *Repository_SetAppend_Call) Run(run func(ctx context.Context, key string, value string)) *Repository_SetAppend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Repository_SetAppend_Call) Return(_a0 errs.ChatError) *Repository_SetAppend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_SetAppend_Call) RunAndReturn(run func(context.Context, string, string) errs.ChatError) *Repository_SetAppend_Call {
	_c.Call.Return(run)
	return _c
}

// SetEncryptedHashmap provides a mock function with given fields: ctx, key, secretKey, value
func (_m *Repository) SetEncryptedHashmap(ctx context.Context, key string, secretKey string, value map[string]interface{}) errs.ChatError {
	ret := _m.Called(ctx, key, secretKey, value)

	if len(ret) == 0 {
		panic("no return value specified for SetEncryptedHashmap")
	}

	var r0 errs.ChatError
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]interface{}) errs.ChatError); ok {
		r0 = rf(ctx, key, secretKey, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errs.ChatError)
		}
	}

	return r0
}

// Repository_SetEncryptedHashmap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEncryptedHashmap'
type Repository_SetEncryptedHashmap_Call struct {
	*mock.Call
}

// SetEncryptedHashmap is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - secretKey string
//   - value map[string]interface{}
func (_e *Repository_Expecter) SetEncryptedHashmap(ctx interface{}, key interface{}, secretKey interface{}, value interface{}) *Repository_SetEncryptedHashmap_Call {
	return &Repository_SetEncryptedHashmap_Call{Call: _e.mock.On("SetEncryptedHashmap", ctx, key, secretKey, value)}
}

func (_c *Repository_SetEncryptedHashmap_Call) Run(run func(ctx context.Context, key string, secretKey string, value map[string]interface{})) *Repository_SetEncryptedHashmap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(map[string]interface{}))
	})
	return _c
}

func (_c *Repository_SetEncryptedHashmap_Call) Return(_a0 errs.ChatError) *Repository_SetEncryptedHashmap_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_SetEncryptedHashmap_Call) RunAndReturn(run func(context.Context, string, string, map[string]interface{}) errs.ChatError) *Repository_SetEncryptedHashmap_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
