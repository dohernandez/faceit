package usecase

import (
	"context"
	"fmt"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
)

//go:generate mockery --name=UserUpdater --outpkg=mocks --output=mocks --filename=user_updater.go --with-expecter

// UserUpdater defines functionality to update a user to the data layer.
type UserUpdater interface {
	UpdateUser(ctx context.Context, id model.UserID, info model.UserInfo) error
}

//go:generate mockery --name=UserUpdatedNotifier --outpkg=mocks --output=mocks --filename=user_updated_notifier.go --with-expecter

// UserUpdatedNotifier defines functionality to notify about a user updated.
type UserUpdatedNotifier interface {
	NotifyUserAdded(ctx context.Context, id model.UserID, info model.UserInfo) error
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
func (a *UpdateUser) UpdateUser(ctx context.Context, id model.UserID, info model.UserInfo) error {
	ctx = ctxd.AddFields(ctx, "use_case", "UpdateUser")

	if err := a.updater.UpdateUser(ctx, id, info); err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	a.logger.Debug(ctx, "user updated", "user_id", id)

	if err := a.notifier.NotifyUserAdded(ctx, id, info); err != nil {
		return fmt.Errorf("notify user updated: %w", err)
	}

	a.logger.Debug(ctx, "user updated notification sent", "user_id", id)

	return nil
}
