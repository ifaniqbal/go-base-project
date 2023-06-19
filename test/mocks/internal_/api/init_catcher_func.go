// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	catcher "github.com/ifaniqbal/go-base-project/pkg/catcher"
	mock "github.com/stretchr/testify/mock"
)

// initCatcherFunc is an autogenerated mock type for the initCatcherFunc type
type initCatcherFunc struct {
	mock.Mock
}

type initCatcherFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *initCatcherFunc) EXPECT() *initCatcherFunc_Expecter {
	return &initCatcherFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *initCatcherFunc) Execute() (catcher.Catcher, error) {
	ret := _m.Called()

	var r0 catcher.Catcher
	var r1 error
	if rf, ok := ret.Get(0).(func() (catcher.Catcher, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() catcher.Catcher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catcher.Catcher)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// initCatcherFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type initCatcherFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *initCatcherFunc_Expecter) Execute() *initCatcherFunc_Execute_Call {
	return &initCatcherFunc_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *initCatcherFunc_Execute_Call) Run(run func()) *initCatcherFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *initCatcherFunc_Execute_Call) Return(_a0 catcher.Catcher, _a1 error) *initCatcherFunc_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *initCatcherFunc_Execute_Call) RunAndReturn(run func() (catcher.Catcher, error)) *initCatcherFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTnewInitCatcherFunc interface {
	mock.TestingT
	Cleanup(func())
}

// newInitCatcherFunc creates a new instance of initCatcherFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newInitCatcherFunc(t mockConstructorTestingTnewInitCatcherFunc) *initCatcherFunc {
	mock := &initCatcherFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}