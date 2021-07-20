// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/aofiee/barroth/models"
	mock "github.com/stretchr/testify/mock"
)

// SystemRepository is an autogenerated mock type for the SystemRepository type
type SystemRepository struct {
	mock.Mock
}

// CreateRole provides a mock function with given fields: m
func (_m *SystemRepository) CreateRole(m *models.RoleItems) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.RoleItems) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSystem provides a mock function with given fields: s
func (_m *SystemRepository) CreateSystem(s *models.System) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.System) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: m
func (_m *SystemRepository) CreateUser(m *models.Users) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Users) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFirstSystemInstallation provides a mock function with given fields: s
func (_m *SystemRepository) GetFirstSystemInstallation(s *models.System) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.System) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSystem provides a mock function with given fields: s, id
func (_m *SystemRepository) GetSystem(s *models.System, id string) error {
	ret := _m.Called(s, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.System, string) error); ok {
		r0 = rf(s, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSystem provides a mock function with given fields: s, id
func (_m *SystemRepository) UpdateSystem(s *models.System, id string) error {
	ret := _m.Called(s, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.System, string) error); ok {
		r0 = rf(s, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
