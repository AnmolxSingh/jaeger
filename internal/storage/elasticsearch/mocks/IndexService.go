// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	elasticsearch "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
	mock "github.com/stretchr/testify/mock"
)

// IndexService is an autogenerated mock type for the IndexService type
type IndexService struct {
	mock.Mock
}

type IndexService_Expecter struct {
	mock *mock.Mock
}

func (_m *IndexService) EXPECT() *IndexService_Expecter {
	return &IndexService_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with no fields
func (_m *IndexService) Add() {
	_m.Called()
}

// IndexService_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type IndexService_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
func (_e *IndexService_Expecter) Add() *IndexService_Add_Call {
	return &IndexService_Add_Call{Call: _e.mock.On("Add")}
}

func (_c *IndexService_Add_Call) Run(run func()) *IndexService_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IndexService_Add_Call) Return() *IndexService_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *IndexService_Add_Call) RunAndReturn(run func()) *IndexService_Add_Call {
	_c.Run(run)
	return _c
}

// BodyJson provides a mock function with given fields: body
func (_m *IndexService) BodyJson(body any) elasticsearch.IndexService {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for BodyJson")
	}

	var r0 elasticsearch.IndexService
	if rf, ok := ret.Get(0).(func(any) elasticsearch.IndexService); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(elasticsearch.IndexService)
		}
	}

	return r0
}

// IndexService_BodyJson_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BodyJson'
type IndexService_BodyJson_Call struct {
	*mock.Call
}

// BodyJson is a helper method to define mock.On call
//   - body any
func (_e *IndexService_Expecter) BodyJson(body interface{}) *IndexService_BodyJson_Call {
	return &IndexService_BodyJson_Call{Call: _e.mock.On("BodyJson", body)}
}

func (_c *IndexService_BodyJson_Call) Run(run func(body any)) *IndexService_BodyJson_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(any))
	})
	return _c
}

func (_c *IndexService_BodyJson_Call) Return(_a0 elasticsearch.IndexService) *IndexService_BodyJson_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexService_BodyJson_Call) RunAndReturn(run func(any) elasticsearch.IndexService) *IndexService_BodyJson_Call {
	_c.Call.Return(run)
	return _c
}

// Id provides a mock function with given fields: id
func (_m *IndexService) Id(id string) elasticsearch.IndexService {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Id")
	}

	var r0 elasticsearch.IndexService
	if rf, ok := ret.Get(0).(func(string) elasticsearch.IndexService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(elasticsearch.IndexService)
		}
	}

	return r0
}

// IndexService_Id_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Id'
type IndexService_Id_Call struct {
	*mock.Call
}

// Id is a helper method to define mock.On call
//   - id string
func (_e *IndexService_Expecter) Id(id interface{}) *IndexService_Id_Call {
	return &IndexService_Id_Call{Call: _e.mock.On("Id", id)}
}

func (_c *IndexService_Id_Call) Run(run func(id string)) *IndexService_Id_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IndexService_Id_Call) Return(_a0 elasticsearch.IndexService) *IndexService_Id_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexService_Id_Call) RunAndReturn(run func(string) elasticsearch.IndexService) *IndexService_Id_Call {
	_c.Call.Return(run)
	return _c
}

// Index provides a mock function with given fields: index
func (_m *IndexService) Index(index string) elasticsearch.IndexService {
	ret := _m.Called(index)

	if len(ret) == 0 {
		panic("no return value specified for Index")
	}

	var r0 elasticsearch.IndexService
	if rf, ok := ret.Get(0).(func(string) elasticsearch.IndexService); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(elasticsearch.IndexService)
		}
	}

	return r0
}

// IndexService_Index_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Index'
type IndexService_Index_Call struct {
	*mock.Call
}

// Index is a helper method to define mock.On call
//   - index string
func (_e *IndexService_Expecter) Index(index interface{}) *IndexService_Index_Call {
	return &IndexService_Index_Call{Call: _e.mock.On("Index", index)}
}

func (_c *IndexService_Index_Call) Run(run func(index string)) *IndexService_Index_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IndexService_Index_Call) Return(_a0 elasticsearch.IndexService) *IndexService_Index_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexService_Index_Call) RunAndReturn(run func(string) elasticsearch.IndexService) *IndexService_Index_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields: typ
func (_m *IndexService) Type(typ string) elasticsearch.IndexService {
	ret := _m.Called(typ)

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 elasticsearch.IndexService
	if rf, ok := ret.Get(0).(func(string) elasticsearch.IndexService); ok {
		r0 = rf(typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(elasticsearch.IndexService)
		}
	}

	return r0
}

// IndexService_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type IndexService_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
//   - typ string
func (_e *IndexService_Expecter) Type(typ interface{}) *IndexService_Type_Call {
	return &IndexService_Type_Call{Call: _e.mock.On("Type", typ)}
}

func (_c *IndexService_Type_Call) Run(run func(typ string)) *IndexService_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IndexService_Type_Call) Return(_a0 elasticsearch.IndexService) *IndexService_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IndexService_Type_Call) RunAndReturn(run func(string) elasticsearch.IndexService) *IndexService_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewIndexService creates a new instance of IndexService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndexService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndexService {
	mock := &IndexService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
