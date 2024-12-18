package storage_test

import (
	"context"
	"testing"

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

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		mock.ExpectExec(`
				INSERT INTO users (id,password_hash,email,first_name,last_name,nickname,country) 
				VALUES ($1,$2,$3,$4,$5,$6,$7)
			`).
			WithArgs(
				user.ID,
				user.PasswordHash,
				user.Email,
				user.FirstName,
				user.LastName,
				user.Nickname,
				user.Country,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.AddUser(context.Background(), user)
		require.NoError(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error exists", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				INSERT INTO users (id,password_hash,email,first_name,last_name,nickname,country)
				VALUES ($1,$2,$3,$4,$5,$6,$7)
			`).
			WithArgs(
				user.ID,
				user.PasswordHash,
				user.Email,
				user.FirstName,
				user.LastName,
				user.Nickname,
				user.Country,
			)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.UniqueViolation})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.AddUser(context.Background(), user)
		require.Error(t, err)
		require.ErrorIs(t, err, database.ErrAlreadyExists)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				INSERT INTO users (id,password_hash,email,first_name,last_name,nickname,country)
				VALUES ($1,$2,$3,$4,$5,$6,$7)
			`).
			WithArgs(
				user.ID,
				user.PasswordHash,
				user.Email,
				user.FirstName,
				user.LastName,
				user.Nickname,
				user.Country,
			)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.InternalError})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.AddUser(context.Background(), user)
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

func TestUser_DeleteUser(t *testing.T) {
	t.Parallel()

	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				DELETE FROM users WHERE id = $1
			`).
			WithArgs(userID)

		meQuery.WillReturnResult(sqlmock.NewResult(0, 1))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.DeleteUser(context.Background(), userID)
		require.NoError(t, err)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				DELETE FROM users WHERE id = $1
			`).
			WithArgs(userID)

		meQuery.WillReturnResult(sqlmock.NewResult(0, 0))

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.DeleteUser(context.Background(), userID)
		require.Error(t, err)
		require.ErrorIs(t, err, database.ErrNotFound)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectExec(`
				DELETE FROM users WHERE id = $1
			`).
			WithArgs(userID)

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.InternalError})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		err = repo.DeleteUser(context.Background(), userID)
		require.Error(t, err)
	})
}

func TestUser_ListByCountry(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				SELECT id, created_at, updated_at, password_hash, email, first_name, last_name, nickname, country FROM users WHERE country = $1 LIMIT 100 OFFSET 0
			`).
			WithArgs("UK")

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(uuid.New()).
			AddRow(uuid.New())

		meQuery.WillReturnRows(rows)

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		users, err := repo.ListByCountry(context.Background(), "UK", 100, 0)
		require.NoError(t, err)

		require.Len(t, users, 2)

		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		defer db.Close() //nolint:errcheck

		meQuery := mock.ExpectQuery(`
				SELECT id, created_at, updated_at, password_hash, email, first_name, last_name, nickname, country FROM users WHERE country = $1 LIMIT 100 OFFSET 0
			`).
			WithArgs("UK")

		meQuery.WillReturnError(&pgconn.PgError{Code: pgerrcode.InternalError})

		st := sqluct.NewStorage(sqlx.NewDb(db, "sqlmock"))

		repo := storage.NewUser(st)

		users, err := repo.ListByCountry(context.Background(), "UK", 100, 0)
		require.Error(t, err)
		require.Nil(t, users)
	})
}
