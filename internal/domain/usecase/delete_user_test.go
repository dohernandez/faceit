package usecase

import (
	"context"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/usecase/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteUser_DeleteUser(t *testing.T) {
	t.Parallel()

	uID := uuid.New()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		deleter := mocks.NewUserDeleter(t)
		deleter.EXPECT().DeleteUser(mock.Anything, uID).Return(nil)

		notifier := mocks.NewUserDeletedNotifier(t)
		notifier.EXPECT().NotifyUserDeleted(mock.Anything, uID).Return(nil)

		logger := &ctxd.LoggerMock{}

		uc := NewDeleteUser(deleter, notifier, logger)

		err := uc.DeleteUser(context.Background(), uID)
		require.NoError(t, err)
	})

	t.Run("error deleter", func(t *testing.T) {
		t.Parallel()

		deleter := mocks.NewUserDeleter(t)
		deleter.EXPECT().DeleteUser(mock.Anything, uID).Return(assert.AnError)

		notifier := mocks.NewUserDeletedNotifier(t)

		logger := &ctxd.LoggerMock{}

		uc := NewDeleteUser(deleter, notifier, logger)

		err := uc.DeleteUser(context.Background(), uID)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("error notifier", func(t *testing.T) {
		t.Parallel()

		deleter := mocks.NewUserDeleter(t)
		deleter.EXPECT().DeleteUser(mock.Anything, uID).Return(nil)

		notifier := mocks.NewUserDeletedNotifier(t)
		notifier.EXPECT().NotifyUserDeleted(mock.Anything, uID).Return(assert.AnError)

		logger := &ctxd.LoggerMock{}

		uc := NewDeleteUser(deleter, notifier, logger)

		err := uc.DeleteUser(context.Background(), uID)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
