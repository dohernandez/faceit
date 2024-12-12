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

func TestListUserByCountry_ListByCountry(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		finder := mocks.NewUserByCountryFinder(t)
		finder.EXPECT().ListByCountry(mock.Anything, "UK", uint64(100), uint64(0)).Return([]*model.User{
			{
				ID: uuid.New(),
			},
			{
				ID: uuid.New(),
			},
		}, nil)

		logger := &ctxd.LoggerMock{}

		uc := NewListUsersByCountry(finder, logger)

		users, err := uc.ListUsersByCountry(context.Background(), "UK", 100, 0)
		require.NoError(t, err)

		require.Len(t, users, 2)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		finder := mocks.NewUserByCountryFinder(t)
		finder.EXPECT().ListByCountry(mock.Anything, "UK", uint64(100), uint64(0)).Return(nil, assert.AnError)

		logger := &ctxd.LoggerMock{}

		uc := NewListUsersByCountry(finder, logger)

		users, err := uc.ListUsersByCountry(context.Background(), "UK", 100, 0)
		require.Error(t, err)
		require.Nil(t, users)
		require.ErrorIs(t, err, assert.AnError)
	})
}
