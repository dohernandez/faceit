package usecase

import (
	"context"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserDeleter --outpkg=mocks --output=mocks --filename=user_deleter.go --with-expecter

// UserDeleter defines functionality to delete a user from the data layer.
type UserDeleter interface {
	DeleteUser(ctx context.Context, id model.UserID) error
}

//go:generate mockery --name=UserDeletedNotifier --outpkg=mocks --output=mocks --filename=user_deleted_notifier.go --with-expecter

// UserDeletedNotifier defines functionality to notify about a user deleted.
type UserDeletedNotifier interface {
	NotifyUserDeleted(ctx context.Context, id model.UserID) error
}

// DeleteUser is a use case to delete a user.
type DeleteUser struct {
	deleter  UserDeleter
	notifier UserDeletedNotifier

	logger ctxd.Logger
}

// NewDeleteUser creates a new DeleteUser use case.
func NewDeleteUser(userDeleter UserDeleter, notifier UserDeletedNotifier, logger ctxd.Logger) *DeleteUser {
	return &DeleteUser{
		deleter:  userDeleter,
		notifier: notifier,
		logger:   logger,
	}
}

// DeleteUser executes the delete user use case.
func (a *DeleteUser) DeleteUser(ctx context.Context, id model.UserID) error {
	ctx = ctxd.AddFields(ctx, "use_case", "DeleteUser", "user_id", id)

	if err := a.deleter.DeleteUser(ctx, id); err != nil {
		return ctxd.WrapError(ctx, err, "delete user") // error contains the context fields added
	}

	a.logger.Debug(ctx, "user deleted")

	if err := a.notifier.NotifyUserDeleted(ctx, id); err != nil {
		return ctxd.WrapError(ctx, err, "notify user deleted")
	}

	a.logger.Debug(ctx, "user deleted notification sent", "user_id", id)

	return nil
}
