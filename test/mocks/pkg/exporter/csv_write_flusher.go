// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CSVWriteFlusher is an autogenerated mock type for the CSVWriteFlusher type
type CSVWriteFlusher struct {
	mock.Mock
}

type CSVWriteFlusher_Expecter struct {
	mock *mock.Mock
}

func (_m *CSVWriteFlusher) EXPECT() *CSVWriteFlusher_Expecter {
	return &CSVWriteFlusher_Expecter{mock: &_m.Mock}
}

// Flush provides a mock function with given fields:
func (_m *CSVWriteFlusher) Flush() {
	_m.Called()
}

// CSVWriteFlusher_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type CSVWriteFlusher_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
func (_e *CSVWriteFlusher_Expecter) Flush() *CSVWriteFlusher_Flush_Call {
	return &CSVWriteFlusher_Flush_Call{Call: _e.mock.On("Flush")}
}

func (_c *CSVWriteFlusher_Flush_Call) Run(run func()) *CSVWriteFlusher_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CSVWriteFlusher_Flush_Call) Return() *CSVWriteFlusher_Flush_Call {
	_c.Call.Return()
	return _c
}

func (_c *CSVWriteFlusher_Flush_Call) RunAndReturn(run func()) *CSVWriteFlusher_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: record
func (_m *CSVWriteFlusher) Write(record []string) error {
	ret := _m.Called(record)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSVWriteFlusher_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type CSVWriteFlusher_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - record []string
func (_e *CSVWriteFlusher_Expecter) Write(record interface{}) *CSVWriteFlusher_Write_Call {
	return &CSVWriteFlusher_Write_Call{Call: _e.mock.On("Write", record)}
}

func (_c *CSVWriteFlusher_Write_Call) Run(run func(record []string)) *CSVWriteFlusher_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *CSVWriteFlusher_Write_Call) Return(_a0 error) *CSVWriteFlusher_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CSVWriteFlusher_Write_Call) RunAndReturn(run func([]string) error) *CSVWriteFlusher_Write_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCSVWriteFlusher interface {
	mock.TestingT
	Cleanup(func())
}

// NewCSVWriteFlusher creates a new instance of CSVWriteFlusher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCSVWriteFlusher(t mockConstructorTestingTNewCSVWriteFlusher) *CSVWriteFlusher {
	mock := &CSVWriteFlusher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
