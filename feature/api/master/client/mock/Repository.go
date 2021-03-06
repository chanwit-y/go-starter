// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import (
	schema "go-stater/domain/schema"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 schema.Client) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(schema.Client) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *Repository) FindAll() []schema.Client {
	// ret := _m.Called()

	// var r0 []schema.Client
	// if rf, ok := ret.Get(0).(func() []schema.Client); ok {
	// 	r0 = rf()
	// } else {
	// 	if ret.Get(0) != nil {
	// 		r0 = ret.Get(0).([]schema.Client)
	// 	}
	// }

	// return r0

	mockData := []schema.Client{
		{
			Code: "1",
			Name: "Unit Test",
		},
	}
	return mockData
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
