// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ValidatorFunc is an autogenerated mock type for the ValidatorFunc type
type ValidatorFunc struct {
	mock.Mock
}

type ValidatorFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidatorFunc) EXPECT() *ValidatorFunc_Expecter {
	return &ValidatorFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: payload
func (_m *ValidatorFunc) Execute(payload interface{}) bool {
	ret := _m.Called(payload)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidatorFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ValidatorFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - payload interface{}
func (_e *ValidatorFunc_Expecter) Execute(payload interface{}) *ValidatorFunc_Execute_Call {
	return &ValidatorFunc_Execute_Call{Call: _e.mock.On("Execute", payload)}
}

func (_c *ValidatorFunc_Execute_Call) Run(run func(payload interface{})) *ValidatorFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ValidatorFunc_Execute_Call) Return(_a0 bool) *ValidatorFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ValidatorFunc_Execute_Call) RunAndReturn(run func(interface{}) bool) *ValidatorFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewValidatorFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatorFunc creates a new instance of ValidatorFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatorFunc(t mockConstructorTestingTNewValidatorFunc) *ValidatorFunc {
	mock := &ValidatorFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
