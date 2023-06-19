// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	utils "github.com/ifaniqbal/go-base-project/internal/utils"
	mock "github.com/stretchr/testify/mock"
)

// Catcher is an autogenerated mock type for the Catcher type
type Catcher struct {
	mock.Mock
}

type Catcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Catcher) EXPECT() *Catcher_Expecter {
	return &Catcher_Expecter{mock: &_m.Mock}
}

// CaptureMessage provides a mock function with given fields: msg
func (_m *Catcher) CaptureMessage(msg string) {
	_m.Called(msg)
}

// Catcher_CaptureMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CaptureMessage'
type Catcher_CaptureMessage_Call struct {
	*mock.Call
}

// CaptureMessage is a helper method to define mock.On call
//   - msg string
func (_e *Catcher_Expecter) CaptureMessage(msg interface{}) *Catcher_CaptureMessage_Call {
	return &Catcher_CaptureMessage_Call{Call: _e.mock.On("CaptureMessage", msg)}
}

func (_c *Catcher_CaptureMessage_Call) Run(run func(msg string)) *Catcher_CaptureMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Catcher_CaptureMessage_Call) Return() *Catcher_CaptureMessage_Call {
	_c.Call.Return()
	return _c
}

func (_c *Catcher_CaptureMessage_Call) RunAndReturn(run func(string)) *Catcher_CaptureMessage_Call {
	_c.Call.Return(run)
	return _c
}

// Catch provides a mock function with given fields: err
func (_m *Catcher) Catch(err error) {
	_m.Called(err)
}

// Catcher_Catch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Catch'
type Catcher_Catch_Call struct {
	*mock.Call
}

// Catch is a helper method to define mock.On call
//   - err error
func (_e *Catcher_Expecter) Catch(err interface{}) *Catcher_Catch_Call {
	return &Catcher_Catch_Call{Call: _e.mock.On("Catch", err)}
}

func (_c *Catcher_Catch_Call) Run(run func(err error)) *Catcher_Catch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *Catcher_Catch_Call) Return() *Catcher_Catch_Call {
	_c.Call.Return()
	return _c
}

func (_c *Catcher_Catch_Call) RunAndReturn(run func(error)) *Catcher_Catch_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields:
func (_m *Catcher) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Catcher_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Catcher_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *Catcher_Expecter) Init() *Catcher_Init_Call {
	return &Catcher_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *Catcher_Init_Call) Run(run func()) *Catcher_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Catcher_Init_Call) Return(_a0 error) *Catcher_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Catcher_Init_Call) RunAndReturn(run func() error) *Catcher_Init_Call {
	_c.Call.Return(run)
	return _c
}

// WithContext provides a mock function with given fields: ctx
func (_m *Catcher) WithContext(ctx utils.Context) utils.Catcher {
	ret := _m.Called(ctx)

	var r0 utils.Catcher
	if rf, ok := ret.Get(0).(func(utils.Context) utils.Catcher); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.Catcher)
		}
	}

	return r0
}

// Catcher_WithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithContext'
type Catcher_WithContext_Call struct {
	*mock.Call
}

// WithContext is a helper method to define mock.On call
//   - ctx utils.Context
func (_e *Catcher_Expecter) WithContext(ctx interface{}) *Catcher_WithContext_Call {
	return &Catcher_WithContext_Call{Call: _e.mock.On("WithContext", ctx)}
}

func (_c *Catcher_WithContext_Call) Run(run func(ctx utils.Context)) *Catcher_WithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(utils.Context))
	})
	return _c
}

func (_c *Catcher_WithContext_Call) Return(_a0 utils.Catcher) *Catcher_WithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Catcher_WithContext_Call) RunAndReturn(run func(utils.Context) utils.Catcher) *Catcher_WithContext_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCatcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewCatcher creates a new instance of Catcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCatcher(t mockConstructorTestingTNewCatcher) *Catcher {
	mock := &Catcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}