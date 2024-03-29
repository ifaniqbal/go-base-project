// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	authenticator "github.com/ifaniqbal/go-base-project/pkg/ginwrapper/middleware/authenticator"

	mock "github.com/stretchr/testify/mock"

	utils "github.com/ifaniqbal/go-base-project/internal/utils"
)

// Authenticator is an autogenerated mock type for the Authenticator type
type Authenticator struct {
	mock.Mock
}

type Authenticator_Expecter struct {
	mock *mock.Mock
}

func (_m *Authenticator) EXPECT() *Authenticator_Expecter {
	return &Authenticator_Expecter{mock: &_m.Mock}
}

// Authenticate provides a mock function with given fields: secret
func (_m *Authenticator) Authenticate(secret string) utils.HandlerFunc {
	ret := _m.Called(secret)

	var r0 utils.HandlerFunc
	if rf, ok := ret.Get(0).(func(string) utils.HandlerFunc); ok {
		r0 = rf(secret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.HandlerFunc)
		}
	}

	return r0
}

// Authenticator_Authenticate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authenticate'
type Authenticator_Authenticate_Call struct {
	*mock.Call
}

// Authenticate is a helper method to define mock.On call
//   - secret string
func (_e *Authenticator_Expecter) Authenticate(secret interface{}) *Authenticator_Authenticate_Call {
	return &Authenticator_Authenticate_Call{Call: _e.mock.On("Authenticate", secret)}
}

func (_c *Authenticator_Authenticate_Call) Run(run func(secret string)) *Authenticator_Authenticate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Authenticator_Authenticate_Call) Return(_a0 utils.HandlerFunc) *Authenticator_Authenticate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Authenticator_Authenticate_Call) RunAndReturn(run func(string) utils.HandlerFunc) *Authenticator_Authenticate_Call {
	_c.Call.Return(run)
	return _c
}

// Authorize provides a mock function with given fields: secret, isValid
func (_m *Authenticator) Authorize(secret string, isValid authenticator.ValidatorFunc) utils.HandlerFunc {
	ret := _m.Called(secret, isValid)

	var r0 utils.HandlerFunc
	if rf, ok := ret.Get(0).(func(string, authenticator.ValidatorFunc) utils.HandlerFunc); ok {
		r0 = rf(secret, isValid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.HandlerFunc)
		}
	}

	return r0
}

// Authenticator_Authorize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authorize'
type Authenticator_Authorize_Call struct {
	*mock.Call
}

// Authorize is a helper method to define mock.On call
//   - secret string
//   - isValid authenticator.ValidatorFunc
func (_e *Authenticator_Expecter) Authorize(secret interface{}, isValid interface{}) *Authenticator_Authorize_Call {
	return &Authenticator_Authorize_Call{Call: _e.mock.On("Authorize", secret, isValid)}
}

func (_c *Authenticator_Authorize_Call) Run(run func(secret string, isValid authenticator.ValidatorFunc)) *Authenticator_Authorize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(authenticator.ValidatorFunc))
	})
	return _c
}

func (_c *Authenticator_Authorize_Call) Return(_a0 utils.HandlerFunc) *Authenticator_Authorize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Authenticator_Authorize_Call) RunAndReturn(run func(string, authenticator.ValidatorFunc) utils.HandlerFunc) *Authenticator_Authorize_Call {
	_c.Call.Return(run)
	return _c
}

// GetPayloadFromContext provides a mock function with given fields: ctx
func (_m *Authenticator) GetPayloadFromContext(ctx utils.Context) (interface{}, bool) {
	ret := _m.Called(ctx)

	var r0 interface{}
	var r1 bool
	if rf, ok := ret.Get(0).(func(utils.Context) (interface{}, bool)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(utils.Context) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(utils.Context) bool); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Authenticator_GetPayloadFromContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPayloadFromContext'
type Authenticator_GetPayloadFromContext_Call struct {
	*mock.Call
}

// GetPayloadFromContext is a helper method to define mock.On call
//   - ctx utils.Context
func (_e *Authenticator_Expecter) GetPayloadFromContext(ctx interface{}) *Authenticator_GetPayloadFromContext_Call {
	return &Authenticator_GetPayloadFromContext_Call{Call: _e.mock.On("GetPayloadFromContext", ctx)}
}

func (_c *Authenticator_GetPayloadFromContext_Call) Run(run func(ctx utils.Context)) *Authenticator_GetPayloadFromContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(utils.Context))
	})
	return _c
}

func (_c *Authenticator_GetPayloadFromContext_Call) Return(_a0 interface{}, _a1 bool) *Authenticator_GetPayloadFromContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Authenticator_GetPayloadFromContext_Call) RunAndReturn(run func(utils.Context) (interface{}, bool)) *Authenticator_GetPayloadFromContext_Call {
	_c.Call.Return(run)
	return _c
}

// SendUnauthorizedResponse provides a mock function with given fields: ctx
func (_m *Authenticator) SendUnauthorizedResponse(ctx utils.Context) {
	_m.Called(ctx)
}

// Authenticator_SendUnauthorizedResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendUnauthorizedResponse'
type Authenticator_SendUnauthorizedResponse_Call struct {
	*mock.Call
}

// SendUnauthorizedResponse is a helper method to define mock.On call
//   - ctx utils.Context
func (_e *Authenticator_Expecter) SendUnauthorizedResponse(ctx interface{}) *Authenticator_SendUnauthorizedResponse_Call {
	return &Authenticator_SendUnauthorizedResponse_Call{Call: _e.mock.On("SendUnauthorizedResponse", ctx)}
}

func (_c *Authenticator_SendUnauthorizedResponse_Call) Run(run func(ctx utils.Context)) *Authenticator_SendUnauthorizedResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(utils.Context))
	})
	return _c
}

func (_c *Authenticator_SendUnauthorizedResponse_Call) Return() *Authenticator_SendUnauthorizedResponse_Call {
	_c.Call.Return()
	return _c
}

func (_c *Authenticator_SendUnauthorizedResponse_Call) RunAndReturn(run func(utils.Context)) *Authenticator_SendUnauthorizedResponse_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAuthenticator interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticator creates a new instance of Authenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticator(t mockConstructorTestingTNewAuthenticator) *Authenticator {
	mock := &Authenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
