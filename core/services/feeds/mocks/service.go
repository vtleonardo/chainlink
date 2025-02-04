// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	feeds "github.com/smartcontractkit/chainlink/core/services/feeds"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Service) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountManagers provides a mock function with given fields:
func (_m *Service) CountManagers() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJobProposal provides a mock function with given fields: jp
func (_m *Service) CreateJobProposal(jp *feeds.JobProposal) (int64, error) {
	ret := _m.Called(jp)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*feeds.JobProposal) int64); ok {
		r0 = rf(jp)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*feeds.JobProposal) error); ok {
		r1 = rf(jp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManager provides a mock function with given fields: id
func (_m *Service) GetManager(id int64) (*feeds.FeedsManager, error) {
	ret := _m.Called(id)

	var r0 *feeds.FeedsManager
	if rf, ok := ret.Get(0).(func(int64) *feeds.FeedsManager); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*feeds.FeedsManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManagers provides a mock function with given fields:
func (_m *Service) ListManagers() ([]feeds.FeedsManager, error) {
	ret := _m.Called()

	var r0 []feeds.FeedsManager
	if rf, ok := ret.Get(0).(func() []feeds.FeedsManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feeds.FeedsManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterManager provides a mock function with given fields: ms
func (_m *Service) RegisterManager(ms *feeds.FeedsManager) (int64, error) {
	ret := _m.Called(ms)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*feeds.FeedsManager) int64); ok {
		r0 = rf(ms)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*feeds.FeedsManager) error); ok {
		r1 = rf(ms)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Service) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
