// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	model "github.com/kzuabe/todolist-go-api/app/model"
	mock "github.com/stretchr/testify/mock"
)

// TaskUseCaseInterface is an autogenerated mock type for the TaskUseCaseInterface type
type TaskUseCaseInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *TaskUseCaseInterface) Create(_a0 model.Task) (model.Task, error) {
	ret := _m.Called(_a0)

	var r0 model.Task
	if rf, ok := ret.Get(0).(func(model.Task) model.Task); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *TaskUseCaseInterface) Delete(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: _a0
func (_m *TaskUseCaseInterface) Fetch(_a0 model.TaskFetchParam) ([]model.Task, error) {
	ret := _m.Called(_a0)

	var r0 []model.Task
	if rf, ok := ret.Get(0).(func(model.TaskFetchParam) []model.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.TaskFetchParam) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: _a0, _a1
func (_m *TaskUseCaseInterface) FetchByID(_a0 string, _a1 string) (model.Task, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.Task
	if rf, ok := ret.Get(0).(func(string, string) model.Task); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *TaskUseCaseInterface) Update(_a0 model.Task) (model.Task, error) {
	ret := _m.Called(_a0)

	var r0 model.Task
	if rf, ok := ret.Get(0).(func(model.Task) model.Task); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Task) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
