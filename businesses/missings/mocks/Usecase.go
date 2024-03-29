// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	missings "sipencari-api/businesses/missings"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: idUser, missingDomain
func (_m *Usecase) Create(idUser string, missingDomain *missings.Domain) missings.Domain {
	ret := _m.Called(idUser, missingDomain)

	var r0 missings.Domain
	if rf, ok := ret.Get(0).(func(string, *missings.Domain) missings.Domain); ok {
		r0 = rf(idUser, missingDomain)
	} else {
		r0 = ret.Get(0).(missings.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Usecase) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Usecase) GetAll() []missings.Domain {
	ret := _m.Called()

	var r0 []missings.Domain
	if rf, ok := ret.Get(0).(func() []missings.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]missings.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Usecase) GetByID(id string) missings.Domain {
	ret := _m.Called(id)

	var r0 missings.Domain
	if rf, ok := ret.Get(0).(func(string) missings.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(missings.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: idUser, id, missingDomain
func (_m *Usecase) Update(idUser string, id string, missingDomain *missings.Domain) missings.Domain {
	ret := _m.Called(idUser, id, missingDomain)

	var r0 missings.Domain
	if rf, ok := ret.Get(0).(func(string, string, *missings.Domain) missings.Domain); ok {
		r0 = rf(idUser, id, missingDomain)
	} else {
		r0 = ret.Get(0).(missings.Domain)
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
