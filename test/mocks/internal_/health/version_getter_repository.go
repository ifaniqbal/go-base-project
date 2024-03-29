// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VersionGetterRepository is an autogenerated mock type for the VersionGetterRepository type
type VersionGetterRepository struct {
	mock.Mock
}

type VersionGetterRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *VersionGetterRepository) EXPECT() *VersionGetterRepository_Expecter {
	return &VersionGetterRepository_Expecter{mock: &_m.Mock}
}

// GetVersion provides a mock function with given fields:
func (_m *VersionGetterRepository) GetVersion() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VersionGetterRepository_GetVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVersion'
type VersionGetterRepository_GetVersion_Call struct {
	*mock.Call
}

// GetVersion is a helper method to define mock.On call
func (_e *VersionGetterRepository_Expecter) GetVersion() *VersionGetterRepository_GetVersion_Call {
	return &VersionGetterRepository_GetVersion_Call{Call: _e.mock.On("GetVersion")}
}

func (_c *VersionGetterRepository_GetVersion_Call) Run(run func()) *VersionGetterRepository_GetVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VersionGetterRepository_GetVersion_Call) Return(_a0 string, _a1 error) *VersionGetterRepository_GetVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VersionGetterRepository_GetVersion_Call) RunAndReturn(run func() (string, error)) *VersionGetterRepository_GetVersion_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewVersionGetterRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewVersionGetterRepository creates a new instance of VersionGetterRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVersionGetterRepository(t mockConstructorTestingTNewVersionGetterRepository) *VersionGetterRepository {
	mock := &VersionGetterRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
