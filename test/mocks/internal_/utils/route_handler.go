// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	utils "github.com/ifaniqbal/go-base-project/internal/utils"
	mock "github.com/stretchr/testify/mock"
)

// RouteHandler is an autogenerated mock type for the RouteHandler type
type RouteHandler struct {
	mock.Mock
}

type RouteHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *RouteHandler) EXPECT() *RouteHandler_Expecter {
	return &RouteHandler_Expecter{mock: &_m.Mock}
}

// BasePath provides a mock function with given fields:
func (_m *RouteHandler) BasePath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RouteHandler_BasePath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BasePath'
type RouteHandler_BasePath_Call struct {
	*mock.Call
}

// BasePath is a helper method to define mock.On call
func (_e *RouteHandler_Expecter) BasePath() *RouteHandler_BasePath_Call {
	return &RouteHandler_BasePath_Call{Call: _e.mock.On("BasePath")}
}

func (_c *RouteHandler_BasePath_Call) Run(run func()) *RouteHandler_BasePath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RouteHandler_BasePath_Call) Return(_a0 string) *RouteHandler_BasePath_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_BasePath_Call) RunAndReturn(run func() string) *RouteHandler_BasePath_Call {
	_c.Call.Return(run)
	return _c
}

// DELETE provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) DELETE(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_DELETE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DELETE'
type RouteHandler_DELETE_Call struct {
	*mock.Call
}

// DELETE is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) DELETE(relativePath interface{}, handles ...interface{}) *RouteHandler_DELETE_Call {
	return &RouteHandler_DELETE_Call{Call: _e.mock.On("DELETE",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_DELETE_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_DELETE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_DELETE_Call) Return(_a0 utils.RouteHandler) *RouteHandler_DELETE_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_DELETE_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_DELETE_Call {
	_c.Call.Return(run)
	return _c
}

// GET provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) GET(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_GET_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GET'
type RouteHandler_GET_Call struct {
	*mock.Call
}

// GET is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) GET(relativePath interface{}, handles ...interface{}) *RouteHandler_GET_Call {
	return &RouteHandler_GET_Call{Call: _e.mock.On("GET",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_GET_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_GET_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_GET_Call) Return(_a0 utils.RouteHandler) *RouteHandler_GET_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_GET_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_GET_Call {
	_c.Call.Return(run)
	return _c
}

// Group provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) Group(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_Group_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Group'
type RouteHandler_Group_Call struct {
	*mock.Call
}

// Group is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) Group(relativePath interface{}, handles ...interface{}) *RouteHandler_Group_Call {
	return &RouteHandler_Group_Call{Call: _e.mock.On("Group",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_Group_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_Group_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_Group_Call) Return(_a0 utils.RouteHandler) *RouteHandler_Group_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_Group_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_Group_Call {
	_c.Call.Return(run)
	return _c
}

// HEAD provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) HEAD(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_HEAD_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HEAD'
type RouteHandler_HEAD_Call struct {
	*mock.Call
}

// HEAD is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) HEAD(relativePath interface{}, handles ...interface{}) *RouteHandler_HEAD_Call {
	return &RouteHandler_HEAD_Call{Call: _e.mock.On("HEAD",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_HEAD_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_HEAD_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_HEAD_Call) Return(_a0 utils.RouteHandler) *RouteHandler_HEAD_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_HEAD_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_HEAD_Call {
	_c.Call.Return(run)
	return _c
}

// Handle provides a mock function with given fields: httpMethod, relativePath, handles
func (_m *RouteHandler) Handle(httpMethod string, relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, httpMethod, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(httpMethod, relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type RouteHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - httpMethod string
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) Handle(httpMethod interface{}, relativePath interface{}, handles ...interface{}) *RouteHandler_Handle_Call {
	return &RouteHandler_Handle_Call{Call: _e.mock.On("Handle",
		append([]interface{}{httpMethod, relativePath}, handles...)...)}
}

func (_c *RouteHandler_Handle_Call) Run(run func(httpMethod string, relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_Handle_Call) Return(_a0 utils.RouteHandler) *RouteHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_Handle_Call) RunAndReturn(run func(string, string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// OPTIONS provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) OPTIONS(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_OPTIONS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OPTIONS'
type RouteHandler_OPTIONS_Call struct {
	*mock.Call
}

// OPTIONS is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) OPTIONS(relativePath interface{}, handles ...interface{}) *RouteHandler_OPTIONS_Call {
	return &RouteHandler_OPTIONS_Call{Call: _e.mock.On("OPTIONS",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_OPTIONS_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_OPTIONS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_OPTIONS_Call) Return(_a0 utils.RouteHandler) *RouteHandler_OPTIONS_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_OPTIONS_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_OPTIONS_Call {
	_c.Call.Return(run)
	return _c
}

// PATCH provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) PATCH(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_PATCH_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PATCH'
type RouteHandler_PATCH_Call struct {
	*mock.Call
}

// PATCH is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) PATCH(relativePath interface{}, handles ...interface{}) *RouteHandler_PATCH_Call {
	return &RouteHandler_PATCH_Call{Call: _e.mock.On("PATCH",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_PATCH_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_PATCH_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_PATCH_Call) Return(_a0 utils.RouteHandler) *RouteHandler_PATCH_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_PATCH_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_PATCH_Call {
	_c.Call.Return(run)
	return _c
}

// POST provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) POST(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_POST_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'POST'
type RouteHandler_POST_Call struct {
	*mock.Call
}

// POST is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) POST(relativePath interface{}, handles ...interface{}) *RouteHandler_POST_Call {
	return &RouteHandler_POST_Call{Call: _e.mock.On("POST",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_POST_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_POST_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_POST_Call) Return(_a0 utils.RouteHandler) *RouteHandler_POST_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_POST_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_POST_Call {
	_c.Call.Return(run)
	return _c
}

// PUT provides a mock function with given fields: relativePath, handles
func (_m *RouteHandler) PUT(relativePath string, handles ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(handles))
	for _i := range handles {
		_va[_i] = handles[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, relativePath)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(string, ...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(relativePath, handles...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_PUT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PUT'
type RouteHandler_PUT_Call struct {
	*mock.Call
}

// PUT is a helper method to define mock.On call
//   - relativePath string
//   - handles ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) PUT(relativePath interface{}, handles ...interface{}) *RouteHandler_PUT_Call {
	return &RouteHandler_PUT_Call{Call: _e.mock.On("PUT",
		append([]interface{}{relativePath}, handles...)...)}
}

func (_c *RouteHandler_PUT_Call) Run(run func(relativePath string, handles ...utils.HandlerFunc)) *RouteHandler_PUT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_PUT_Call) Return(_a0 utils.RouteHandler) *RouteHandler_PUT_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_PUT_Call) RunAndReturn(run func(string, ...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_PUT_Call {
	_c.Call.Return(run)
	return _c
}

// Use provides a mock function with given fields: middleware
func (_m *RouteHandler) Use(middleware ...utils.HandlerFunc) utils.RouteHandler {
	_va := make([]interface{}, len(middleware))
	for _i := range middleware {
		_va[_i] = middleware[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 utils.RouteHandler
	if rf, ok := ret.Get(0).(func(...utils.HandlerFunc) utils.RouteHandler); ok {
		r0 = rf(middleware...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(utils.RouteHandler)
		}
	}

	return r0
}

// RouteHandler_Use_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Use'
type RouteHandler_Use_Call struct {
	*mock.Call
}

// Use is a helper method to define mock.On call
//   - middleware ...utils.HandlerFunc
func (_e *RouteHandler_Expecter) Use(middleware ...interface{}) *RouteHandler_Use_Call {
	return &RouteHandler_Use_Call{Call: _e.mock.On("Use",
		append([]interface{}{}, middleware...)...)}
}

func (_c *RouteHandler_Use_Call) Run(run func(middleware ...utils.HandlerFunc)) *RouteHandler_Use_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]utils.HandlerFunc, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(utils.HandlerFunc)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *RouteHandler_Use_Call) Return(_a0 utils.RouteHandler) *RouteHandler_Use_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RouteHandler_Use_Call) RunAndReturn(run func(...utils.HandlerFunc) utils.RouteHandler) *RouteHandler_Use_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRouteHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewRouteHandler creates a new instance of RouteHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRouteHandler(t mockConstructorTestingTNewRouteHandler) *RouteHandler {
	mock := &RouteHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
