package usecase

import (
	"context"
	"fmt"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserAdder --outpkg=mocks --output=mocks --filename=user_adder.go --with-expecter

// UserAdder defines functionality to add a user to the data layer.
type UserAdder interface {
	AddUser(ctx context.Context, u model.UserState) (*model.User, error)
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
func (a *AddUser) AddUser(ctx context.Context, us model.UserState) (model.UserID, error) {
	ctx = ctxd.AddFields(ctx, "use_case", "AddUser")

	u, err := a.adder.AddUser(ctx, us)
	if err != nil {
		return model.UserID{}, fmt.Errorf("add user: %w", err)
	}

	a.logger.Debug(ctx, "user added", "user_id", u.ID)

	if err := a.notifier.NotifyUserAdded(ctx, u); err != nil {
		return model.UserID{}, fmt.Errorf("notify user added: %w", err)
	}

	a.logger.Debug(ctx, "user added notification sent", "user_id", u.ID)

	return u.ID, nil
}
