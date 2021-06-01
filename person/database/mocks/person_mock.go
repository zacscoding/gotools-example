// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/zacscoding/gotools-example/person/model"
)

// PersonDB is an autogenerated mock type for the PersonDB type
type PersonDB struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *PersonDB) FindByEmail(ctx context.Context, email string) (*model.Person, error) {
	ret := _m.Called(ctx, email)

	var r0 *model.Person
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Person); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, person
func (_m *PersonDB) Save(ctx context.Context, person *model.Person) error {
	ret := _m.Called(ctx, person)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Person) error); ok {
		r0 = rf(ctx, person)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}