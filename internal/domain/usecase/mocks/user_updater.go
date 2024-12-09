// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/dohernandez/faceit/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserUpdater is an autogenerated mock type for the UserUpdater type
type UserUpdater struct {
	mock.Mock
}

type UserUpdater_Expecter struct {
	mock *mock.Mock
}

func (_m *UserUpdater) EXPECT() *UserUpdater_Expecter {
	return &UserUpdater_Expecter{mock: &_m.Mock}
}

// UpdateUser provides a mock function with given fields: ctx, id, info
func (_m *UserUpdater) UpdateUser(ctx context.Context, id uuid.UUID, info model.UserState) error {
	ret := _m.Called(ctx, id, info)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, model.UserState) error); ok {
		r0 = rf(ctx, id, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserUpdater_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserUpdater_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - info model.UserState
func (_e *UserUpdater_Expecter) UpdateUser(ctx interface{}, id interface{}, info interface{}) *UserUpdater_UpdateUser_Call {
	return &UserUpdater_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, id, info)}
}

func (_c *UserUpdater_UpdateUser_Call) Run(run func(ctx context.Context, id uuid.UUID, info model.UserState)) *UserUpdater_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(model.UserState))
	})
	return _c
}

func (_c *UserUpdater_UpdateUser_Call) Return(_a0 error) *UserUpdater_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserUpdater_UpdateUser_Call) RunAndReturn(run func(context.Context, uuid.UUID, model.UserState) error) *UserUpdater_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserUpdater creates a new instance of UserUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUpdater {
	mock := &UserUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
