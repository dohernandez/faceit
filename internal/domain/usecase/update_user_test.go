package usecase

import (
	"context"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/faceit/internal/domain/usecase/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUser_UpdateUser(t *testing.T) {
	t.Parallel()

	uID := uuid.New()

	userState := model.UserState{
		FirstName: "Alice",
		LastName:  "Bob",
		Nickname:  "AB123",
		Country:   "UK",
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		updater := mocks.NewUserUpdater(t)
		updater.EXPECT().UpdateUser(mock.Anything, uID, userState).Return(nil)

		notifier := mocks.NewUserUpdatedNotifier(t)
		notifier.EXPECT().NotifyUserUpdated(mock.Anything, uID, userState).Return(nil)

		logger := &ctxd.LoggerMock{}

		uc := NewUpdateUser(updater, notifier, logger)

		err := uc.UpdateUser(context.Background(), uID, userState)
		require.NoError(t, err)
	})

	t.Run("error updater", func(t *testing.T) {
		t.Parallel()

		updater := mocks.NewUserUpdater(t)
		updater.EXPECT().UpdateUser(mock.Anything, uID, userState).Return(assert.AnError)

		notifier := mocks.NewUserUpdatedNotifier(t)

		logger := &ctxd.LoggerMock{}

		uc := NewUpdateUser(updater, notifier, logger)

		err := uc.UpdateUser(context.Background(), uID, userState)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("error notifier", func(t *testing.T) {
		t.Parallel()

		updater := mocks.NewUserUpdater(t)
		updater.EXPECT().UpdateUser(mock.Anything, uID, userState).Return(nil)

		notifier := mocks.NewUserUpdatedNotifier(t)
		notifier.EXPECT().NotifyUserUpdated(mock.Anything, uID, userState).Return(assert.AnError)

		logger := &ctxd.LoggerMock{}

		uc := NewUpdateUser(updater, notifier, logger)

		err := uc.UpdateUser(context.Background(), uID, userState)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
