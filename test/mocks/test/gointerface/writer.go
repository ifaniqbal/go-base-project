// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

type Writer_Expecter struct {
	mock *mock.Mock
}

func (_m *Writer) EXPECT() *Writer_Expecter {
	return &Writer_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: p
func (_m *Writer) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Writer_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type Writer_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *Writer_Expecter) Write(p interface{}) *Writer_Write_Call {
	return &Writer_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *Writer_Write_Call) Run(run func(p []byte)) *Writer_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *Writer_Write_Call) Return(n int, err error) *Writer_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *Writer_Write_Call) RunAndReturn(run func([]byte) (int, error)) *Writer_Write_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriter(t mockConstructorTestingTNewWriter) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
