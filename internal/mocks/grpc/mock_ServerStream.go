// Code generated by mockery v2.42.1. DO NOT EDIT.

package grpc

import (
	context "context"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// ServerStream is an autogenerated mock type for the ServerStream type
type ServerStream struct {
	mock.Mock
}

type ServerStream_Expecter struct {
	mock *mock.Mock
}

func (_m *ServerStream) EXPECT() *ServerStream_Expecter {
	return &ServerStream_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *ServerStream) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// ServerStream_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type ServerStream_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *ServerStream_Expecter) Context() *ServerStream_Context_Call {
	return &ServerStream_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *ServerStream_Context_Call) Run(run func()) *ServerStream_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServerStream_Context_Call) Return(_a0 context.Context) *ServerStream_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerStream_Context_Call) RunAndReturn(run func() context.Context) *ServerStream_Context_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *ServerStream) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerStream_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type ServerStream_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ServerStream_Expecter) RecvMsg(m interface{}) *ServerStream_RecvMsg_Call {
	return &ServerStream_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *ServerStream_RecvMsg_Call) Run(run func(m interface{})) *ServerStream_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ServerStream_RecvMsg_Call) Return(_a0 error) *ServerStream_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerStream_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *ServerStream_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *ServerStream) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerStream_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type ServerStream_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerStream_Expecter) SendHeader(_a0 interface{}) *ServerStream_SendHeader_Call {
	return &ServerStream_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *ServerStream_SendHeader_Call) Run(run func(_a0 metadata.MD)) *ServerStream_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerStream_SendHeader_Call) Return(_a0 error) *ServerStream_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerStream_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *ServerStream_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *ServerStream) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerStream_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type ServerStream_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ServerStream_Expecter) SendMsg(m interface{}) *ServerStream_SendMsg_Call {
	return &ServerStream_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *ServerStream_SendMsg_Call) Run(run func(m interface{})) *ServerStream_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ServerStream_SendMsg_Call) Return(_a0 error) *ServerStream_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerStream_SendMsg_Call) RunAndReturn(run func(interface{}) error) *ServerStream_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *ServerStream) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServerStream_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type ServerStream_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerStream_Expecter) SetHeader(_a0 interface{}) *ServerStream_SetHeader_Call {
	return &ServerStream_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *ServerStream_SetHeader_Call) Run(run func(_a0 metadata.MD)) *ServerStream_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerStream_SetHeader_Call) Return(_a0 error) *ServerStream_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServerStream_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *ServerStream_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *ServerStream) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// ServerStream_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type ServerStream_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *ServerStream_Expecter) SetTrailer(_a0 interface{}) *ServerStream_SetTrailer_Call {
	return &ServerStream_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *ServerStream_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *ServerStream_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *ServerStream_SetTrailer_Call) Return() *ServerStream_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *ServerStream_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *ServerStream_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewServerStream creates a new instance of ServerStream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerStream(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerStream {
	mock := &ServerStream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
