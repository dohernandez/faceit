package usecase

import (
	"context"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserAdder --outpkg=mocks --output=mocks --filename=user_adder.go --with-expecter

// UserAdder defines functionality to add a user to the data layer.
type UserAdder interface {
	AddUser(ctx context.Context, u *model.User) error
}

//go:generate mockery --name=UserAddedNotifier --outpkg=mocks --output=mocks --filename=user_added_notifier.go --with-expecter

// UserAddedNotifier defines functionality to notify about a user added.
type UserAddedNotifier interface {
	NotifyUserAdded(ctx context.Context, u *model.User) error
}

// AddUser is a use case to add a user.
type AddUser struct {
	adder    UserAdder
	notifier UserAddedNotifier

	logger ctxd.Logger
}

// NewAddUser creates a new AddUser use case.
func NewAddUser(userAdder UserAdder, notifier UserAddedNotifier, logger ctxd.Logger) *AddUser {
	return &AddUser{
		adder:    userAdder,
		notifier: notifier,
		logger:   logger,
	}
}

// AddUser executes the add user use case.
func (a *AddUser) AddUser(ctx context.Context, u *model.User) error {
	ctx = ctxd.AddFields(ctx, "use_case", "AddUser", "user_id", u.ID)

	err := a.adder.AddUser(ctx, u)
	if err != nil {
		return ctxd.WrapError(ctx, err, "add user") // error contains the context fields added
	}

	a.logger.Debug(ctx, "user added")

	if err = a.notifier.NotifyUserAdded(ctx, u); err != nil {
		return ctxd.WrapError(ctx, err, "notify user added")
	}

	a.logger.Debug(ctx, "user added notification sent")

	return nil
}
