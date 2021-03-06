package executor

import mock "github.com/stretchr/testify/mock"
import os "os"
import time "time"

// MockTaskHandle is an autogenerated mock type for the TaskHandle type
type MockTaskHandle struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *MockTaskHandle) Address() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EraseOutput provides a mock function with given fields:
func (_m *MockTaskHandle) EraseOutput() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExitCode provides a mock function with given fields:
func (_m *MockTaskHandle) ExitCode() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *MockTaskHandle) Status() TaskState {
	ret := _m.Called()

	var r0 TaskState
	if rf, ok := ret.Get(0).(func() TaskState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(TaskState)
	}

	return r0
}

// StderrFile provides a mock function with given fields:
func (_m *MockTaskHandle) StderrFile() (*os.File, error) {
	ret := _m.Called()

	var r0 *os.File
	if rf, ok := ret.Get(0).(func() *os.File); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
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

// StdoutFile provides a mock function with given fields:
func (_m *MockTaskHandle) StdoutFile() (*os.File, error) {
	ret := _m.Called()

	var r0 *os.File
	if rf, ok := ret.Get(0).(func() *os.File); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
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

// Stop provides a mock function with given fields:
func (_m *MockTaskHandle) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *MockTaskHandle) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Wait provides a mock function with given fields: timeout
func (_m *MockTaskHandle) Wait(timeout time.Duration) (bool, error) {
	ret := _m.Called(timeout)

	var r0 bool
	if rf, ok := ret.Get(0).(func(time.Duration) bool); ok {
		r0 = rf(timeout)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Duration) error); ok {
		r1 = rf(timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
