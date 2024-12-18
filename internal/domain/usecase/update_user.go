package usecase

import (
	"context"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserUpdater --outpkg=mocks --output=mocks --filename=user_updater.go --with-expecter

// UserUpdater defines functionality to update a user to the data layer.
type UserUpdater interface {
	UpdateUser(ctx context.Context, id model.UserID, info model.UserState) error
}

//go:generate mockery --name=UserUpdatedNotifier --outpkg=mocks --output=mocks --filename=user_updated_notifier.go --with-expecter

// UserUpdatedNotifier defines functionality to notify about a user updated.
type UserUpdatedNotifier interface {
	NotifyUserUpdated(ctx context.Context, id model.UserID, info model.UserState) error
}

// UpdateUser is a use case to update a user.
type UpdateUser struct {
	updater  UserUpdater
	notifier UserUpdatedNotifier

	logger ctxd.Logger
}

// NewUpdateUser creates a new UpdateUser use case.
func NewUpdateUser(userUpdater UserUpdater, notifier UserUpdatedNotifier, logger ctxd.Logger) *UpdateUser {
	return &UpdateUser{
		updater:  userUpdater,
		notifier: notifier,
		logger:   logger,
	}
}

// UpdateUser executes the update user use case.
func (a *UpdateUser) UpdateUser(ctx context.Context, id model.UserID, info model.UserState) error {
	ctx = ctxd.AddFields(ctx, "use_case", "UpdateUser", "user_id", id)

	if err := a.updater.UpdateUser(ctx, id, info); err != nil {
		return ctxd.WrapError(ctx, err, "update user") // error contains the context fields added
	}

	a.logger.Debug(ctx, "user updated")

	if err := a.notifier.NotifyUserUpdated(ctx, id, info); err != nil {
		return ctxd.WrapError(ctx, err, "notify user updated")
	}

	a.logger.Debug(ctx, "user updated notification sent")

	return nil
}
