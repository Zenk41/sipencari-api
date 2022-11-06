// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	likesmissing "sipencari-api/businesses/likes_missing"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: idMissing
func (_m *Usecase) GetAll(idMissing string) []likesmissing.Domain {
	ret := _m.Called(idMissing)

	var r0 []likesmissing.Domain
	if rf, ok := ret.Get(0).(func(string) []likesmissing.Domain); ok {
		r0 = rf(idMissing)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]likesmissing.Domain)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: idLikes, idMissing
func (_m *Usecase) GetByID(idLikes string, idMissing string) likesmissing.Domain {
	ret := _m.Called(idLikes, idMissing)

	var r0 likesmissing.Domain
	if rf, ok := ret.Get(0).(func(string, string) likesmissing.Domain); ok {
		r0 = rf(idLikes, idMissing)
	} else {
		r0 = ret.Get(0).(likesmissing.Domain)
	}

	return r0
}

// Like provides a mock function with given fields: idUser, idMissing, likeDomain
func (_m *Usecase) Like(idUser string, idMissing string, likeDomain *likesmissing.Domain) likesmissing.Domain {
	ret := _m.Called(idUser, idMissing, likeDomain)

	var r0 likesmissing.Domain
	if rf, ok := ret.Get(0).(func(string, string, *likesmissing.Domain) likesmissing.Domain); ok {
		r0 = rf(idUser, idMissing, likeDomain)
	} else {
		r0 = ret.Get(0).(likesmissing.Domain)
	}

	return r0
}

// Unlike provides a mock function with given fields: idUser, idMissing
func (_m *Usecase) Unlike(idUser string, idMissing string) bool {
	ret := _m.Called(idUser, idMissing)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(idUser, idMissing)
	} else {
		r0 = ret.Get(0).(bool)
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