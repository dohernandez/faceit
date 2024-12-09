package storage_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bool64/sqluct"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/faceit/internal/platform/storage"
	"github.com/dohernandez/go-grpc-service/database"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUser_AddUser(t *testing.T) {
	t.Parallel()

	userState := model.UserState{
		PasswordHash: "supersecurepassword",
		Email:        "alice@bob.com",
		FirstName:    "Alice",
		LastName:     "Bob",
		Nickname:     "AB123",
		Country:      "UK",
	}

	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				INSERT INTO users (password_hash,email,first_name,last_name,nickname,country) 
				VALUES ($1,$2,$3,$4,$5,$6) RETURNING *
			`).
			WithArgs(
				userState.PasswordHash,
				userState.Email,
				userState.FirstName,
				userState.LastName,
				userState.Nickname,
				userState.Country,
			)

		rows := sqlmock.NewRows([]string{
			"id", "first_name", "last_name", "nickname", "password_hash", "email", "country", "created_at", "updated_at",
		})

		rows.AddRow(
			userID,
			userState.FirstName,
			userState.LastName,
			userState.Nickname,
			userState.PasswordHash,
			userState.Email,
			userState.Country,
			time.Now(),
			time.Now(),
		)

		meQuery.WillReturnRows(rows)

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		user, err := repo.AddUser(context.Background(), userState)
		require.NoError(t, err)

		require.NotNil(t, user)
		require.NotEmpty(t, user.ID)
		require.Equal(t, userState.FirstName, user.FirstName)
		require.Equal(t, userState.LastName, user.LastName)
		require.Equal(t, userState.Nickname, user.Nickname)
		require.Equal(t, userState.PasswordHash, user.PasswordHash)
		require.Equal(t, userState.Email, user.Email)
		require.Equal(t, userState.Country, user.Country)
		require.NotEmpty(t, user.CreatedAt)
		require.NotEmpty(t, user.UpdatedAt)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error exists", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				INSERT INTO users (password_hash,email,first_name,last_name,nickname,country)
				VALUES ($1,$2,$3,$4,$5,$6) RETURNING *
			`).
			WithArgs(
				userState.PasswordHash,
				userState.Email,
				userState.FirstName,
				userState.LastName,
				userState.Nickname,
				userState.Country,
			)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.UniqueViolation})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		_, err = repo.AddUser(context.Background(), userState)
		require.Error(t, err)
		require.ErrorIs(t, err, database.ErrAlreadyExists)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				INSERT INTO users (password_hash,email,first_name,last_name,nickname,country)
				VALUES ($1,$2,$3,$4,$5,$6) RETURNING *
			`).
			WithArgs(
				userState.PasswordHash,
				userState.Email,
				userState.FirstName,
				userState.LastName,
				userState.Nickname,
				userState.Country,
			)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.InternalError})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		_, err = repo.AddUser(context.Background(), userState)
		require.Error(t, err)
	})
}

func TestUser_UpdateUser(t *testing.T) {
	t.Parallel()

	userIfo := model.UserState{
		FirstName: "Alice",
		LastName:  "Bob",
		Nickname:  "AB123",
		Country:   "UK",
	}

	userID := uuid.New()

	t.Run("success, partial update 4", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				UPDATE users SET first_name = $1, last_name = $2, nickname = $3, country = $4 WHERE id = $5
			`).
			WithArgs(
				userIfo.FirstName,
				userIfo.LastName,
				userIfo.Nickname,
				userIfo.Country,
				userID,
			)

		meQuery.WillReturnResult(sqlmock.NewResult(0, 1))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.UpdateUser(context.Background(), userID, userIfo)
		require.NoError(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("success, partial update 2", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				UPDATE users SET first_name = $1, last_name = $2, nickname = $3 WHERE id = $4
			`).
			WithArgs(
				userIfo.FirstName,
				userIfo.LastName,
				"",
				userID,
			)

		meQuery.WillReturnResult(sqlmock.NewResult(0, 1))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.UpdateUser(context.Background(), userID, model.UserState{
			FirstName: userIfo.FirstName,
			LastName:  userIfo.LastName,
		})
		require.NoError(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				UPDATE users SET first_name = $1, last_name = $2, nickname = $3, country = $4 WHERE id = $5
			`).
			WithArgs(
				userIfo.FirstName,
				userIfo.LastName,
				userIfo.Nickname,
				userIfo.Country,
				userID,
			)

		meQuery.WillReturnResult(sqlmock.NewResult(0, 0))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.UpdateUser(context.Background(), userID, userIfo)
		require.Error(t, err)
		require.ErrorIs(t, err, database.ErrNotFound)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				UPDATE users SET first_name = $1, last_name = $2, nickname = $3, country = $4 WHERE id = $5
			`).
			WithArgs(
				userIfo.FirstName,
				userIfo.LastName,
				userIfo.Nickname,
				userIfo.Country,
				userID,
			)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.InternalError})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.UpdateUser(context.Background(), userID, userIfo)
		require.Error(t, err)
	})
}
