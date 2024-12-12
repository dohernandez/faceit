// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/dohernandez/faceit/internal/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// UserByCountryFinder is an autogenerated mock type for the UserByCountryFinder type
type UserByCountryFinder struct {
	mock.Mock
}

type UserByCountryFinder_Expecter struct {
	mock *mock.Mock
}

func (_m *UserByCountryFinder) EXPECT() *UserByCountryFinder_Expecter {
	return &UserByCountryFinder_Expecter{mock: &_m.Mock}
}

// ListByCountry provides a mock function with given fields: ctx, country, limit, offset
func (_m *UserByCountryFinder) ListByCountry(ctx context.Context, country string, limit uint64, offset uint64) ([]*model.User, error) {
	ret := _m.Called(ctx, country, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for ListByCountry")
	}

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64) ([]*model.User, error)); ok {
		return rf(ctx, country, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, uint64) []*model.User); ok {
		r0 = rf(ctx, country, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, uint64) error); ok {
		r1 = rf(ctx, country, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserByCountryFinder_ListByCountry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByCountry'
type UserByCountryFinder_ListByCountry_Call struct {
	*mock.Call
}

// ListByCountry is a helper method to define mock.On call
//   - ctx context.Context
//   - country string
//   - limit uint64
//   - offset uint64
func (_e *UserByCountryFinder_Expecter) ListByCountry(ctx interface{}, country interface{}, limit interface{}, offset interface{}) *UserByCountryFinder_ListByCountry_Call {
	return &UserByCountryFinder_ListByCountry_Call{Call: _e.mock.On("ListByCountry", ctx, country, limit, offset)}
}

func (_c *UserByCountryFinder_ListByCountry_Call) Run(run func(ctx context.Context, country string, limit uint64, offset uint64)) *UserByCountryFinder_ListByCountry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uint64), args[3].(uint64))
	})
	return _c
}

func (_c *UserByCountryFinder_ListByCountry_Call) Return(_a0 []*model.User, _a1 error) *UserByCountryFinder_ListByCountry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserByCountryFinder_ListByCountry_Call) RunAndReturn(run func(context.Context, string, uint64, uint64) ([]*model.User, error)) *UserByCountryFinder_ListByCountry_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserByCountryFinder creates a new instance of UserByCountryFinder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserByCountryFinder(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserByCountryFinder {
	mock := &UserByCountryFinder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}