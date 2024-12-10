package notifier

import (
	"context"

	"github.com/dohernandez/faceit/internal/domain/model"
)

// NoopNotifier is a notifier that does nothing.
type NoopNotifier struct{}

// NewNoopNotifier creates a new NoopNotifier.
func NewNoopNotifier() *NoopNotifier {
	return &NoopNotifier{}
}

// NotifyUserAdded does nothing.
func (n *NoopNotifier) NotifyUserAdded(_ context.Context, _ *model.User) error {
	return nil
}

// NotifyUserUpdated does nothing.
func (n *NoopNotifier) NotifyUserUpdated(_ context.Context, _ model.UserID, _ model.UserState) error {
	return nil
}

// NotifyUserDeleted does nothing.
func (n *NoopNotifier) NotifyUserDeleted(_ context.Context, _ model.UserID) error {
	return nil
}
