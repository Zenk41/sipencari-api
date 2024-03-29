// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	categories "sipencari-api/businesses/categories"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: categoryDomain
func (_m *Usecase) Create(categoryDomain *categories.Domain) categories.Domain {
	ret := _m.Called(categoryDomain)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(*categories.Domain) categories.Domain); ok {
		r0 = rf(categoryDomain)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: category_id
func (_m *Usecase) Delete(category_id string) bool {
	ret := _m.Called(category_id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(category_id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Usecase) GetAll() []categories.Domain {
	ret := _m.Called()

	var r0 []categories.Domain
	if rf, ok := ret.Get(0).(func() []categories.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]categories.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: category_id
func (_m *Usecase) GetByID(category_id string) categories.Domain {
	ret := _m.Called(category_id)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(string) categories.Domain); ok {
		r0 = rf(category_id)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: category_id, categoryDomain
func (_m *Usecase) Update(category_id string, categoryDomain *categories.Domain) categories.Domain {
	ret := _m.Called(category_id, categoryDomain)

	var r0 categories.Domain
	if rf, ok := ret.Get(0).(func(string, *categories.Domain) categories.Domain); ok {
		r0 = rf(category_id, categoryDomain)
	} else {
		r0 = ret.Get(0).(categories.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
