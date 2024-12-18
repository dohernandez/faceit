package usecase_test

import (
	"context"
	"testing"

	"github.com/bool64/ctxd"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/faceit/internal/domain/usecase"
	"github.com/dohernandez/faceit/internal/domain/usecase/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddUser_AddUser(t *testing.T) {
	t.Parallel()

	user := &model.User{
		ID: uuid.New(),
		UserState: model.UserState{
			PasswordHash: "supersecurepassword",
			Email:        "alice@bob.com",
			FirstName:    "Alice",
			LastName:     "Bob",
			Nickname:     "AB123",
			Country:      "UK",
		},
	}

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		adder := mocks.NewUserAdder(t)
		adder.EXPECT().AddUser(mock.Anything, user).Return(nil)

		notifier := mocks.NewUserAddedNotifier(t)
		notifier.EXPECT().NotifyUserAdded(mock.Anything, user).Return(nil)

		logger := &ctxd.LoggerMock{}

		uc := usecase.NewAddUser(adder, notifier, logger)

		err := uc.AddUser(context.Background(), user)
		require.NoError(t, err)
	})

	t.Run("error adder", func(t *testing.T) {
		t.Parallel()

		adder := mocks.NewUserAdder(t)
		adder.EXPECT().AddUser(mock.Anything, user).Return(assert.AnError)

		notifier := mocks.NewUserAddedNotifier(t)

		logger := &ctxd.LoggerMock{}

		uc := usecase.NewAddUser(adder, notifier, logger)

		err := uc.AddUser(context.Background(), user)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})

	t.Run("error notifier", func(t *testing.T) {
		t.Parallel()

		adder := mocks.NewUserAdder(t)
		adder.EXPECT().AddUser(mock.Anything, user).Return(nil)

		notifier := mocks.NewUserAddedNotifier(t)
		notifier.EXPECT().NotifyUserAdded(mock.Anything, user).Return(assert.AnError)

		logger := &ctxd.LoggerMock{}

		uc := usecase.NewAddUser(adder, notifier, logger)

		err := uc.AddUser(context.Background(), user)
		require.Error(t, err)
		require.ErrorIs(t, err, assert.AnError)
	})
}
