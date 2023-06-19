// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	sentry "github.com/getsentry/sentry-go"
	mock "github.com/stretchr/testify/mock"
)

// Hub is an autogenerated mock type for the Hub type
type Hub struct {
	mock.Mock
}

type Hub_Expecter struct {
	mock *mock.Mock
}

func (_m *Hub) EXPECT() *Hub_Expecter {
	return &Hub_Expecter{mock: &_m.Mock}
}

// CaptureException provides a mock function with given fields: exception
func (_m *Hub) CaptureException(exception error) *sentry.EventID {
	ret := _m.Called(exception)

	var r0 *sentry.EventID
	if rf, ok := ret.Get(0).(func(error) *sentry.EventID); ok {
		r0 = rf(exception)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sentry.EventID)
		}
	}

	return r0
}

// Hub_CaptureException_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CaptureException'
type Hub_CaptureException_Call struct {
	*mock.Call
}

// CaptureException is a helper method to define mock.On call
//   - exception error
func (_e *Hub_Expecter) CaptureException(exception interface{}) *Hub_CaptureException_Call {
	return &Hub_CaptureException_Call{Call: _e.mock.On("CaptureException", exception)}
}

func (_c *Hub_CaptureException_Call) Run(run func(exception error)) *Hub_CaptureException_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *Hub_CaptureException_Call) Return(_a0 *sentry.EventID) *Hub_CaptureException_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hub_CaptureException_Call) RunAndReturn(run func(error) *sentry.EventID) *Hub_CaptureException_Call {
	_c.Call.Return(run)
	return _c
}

// CaptureMessage provides a mock function with given fields: message
func (_m *Hub) CaptureMessage(message string) *sentry.EventID {
	ret := _m.Called(message)

	var r0 *sentry.EventID
	if rf, ok := ret.Get(0).(func(string) *sentry.EventID); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sentry.EventID)
		}
	}

	return r0
}

// Hub_CaptureMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CaptureMessage'
type Hub_CaptureMessage_Call struct {
	*mock.Call
}

// CaptureMessage is a helper method to define mock.On call
//   - message string
func (_e *Hub_Expecter) CaptureMessage(message interface{}) *Hub_CaptureMessage_Call {
	return &Hub_CaptureMessage_Call{Call: _e.mock.On("CaptureMessage", message)}
}

func (_c *Hub_CaptureMessage_Call) Run(run func(message string)) *Hub_CaptureMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Hub_CaptureMessage_Call) Return(_a0 *sentry.EventID) *Hub_CaptureMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Hub_CaptureMessage_Call) RunAndReturn(run func(string) *sentry.EventID) *Hub_CaptureMessage_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewHub interface {
	mock.TestingT
	Cleanup(func())
}

// NewHub creates a new instance of Hub. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHub(t mockConstructorTestingTNewHub) *Hub {
	mock := &Hub{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}