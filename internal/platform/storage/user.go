package storage

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/bool64/ctxd"
	"github.com/bool64/sqluct"
	"github.com/dohernandez/faceit/internal/domain/model"
	"github.com/dohernandez/go-grpc-service/database"
	"github.com/dohernandez/go-grpc-service/database/pgx"
)

// UserTable is the table name for users.
const UserTable = "users"

// User represents a User repository.
type User struct {
	storage *sqluct.Storage

	// col names for users table search
	colID      string
	colCountry string
}

// NewUser returns instance of User repository.
func NewUser(storage *sqluct.Storage) *User {
	var user model.User

	return &User{
		storage:    storage,
		colID:      storage.Mapper.Col(&user, &user.ID),
		colCountry: storage.Mapper.Col(&user, &user.Country),
	}
}

// AddUser store the user data.
func (s *User) AddUser(ctx context.Context, u *model.User) error {
	q := s.storage.InsertStmt(UserTable, u, sqluct.SkipZeroValues)

	res, err := s.storage.Exec(ctx, q)
	if err != nil {
		if pgx.IsUniqueViolation(err) {
			return ctxd.LabeledError(database.ErrAlreadyExists, err)
		}

		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return err
}

// UpdateUser updates the user data.
func (s *User) UpdateUser(ctx context.Context, id model.UserID, state model.UserState) error {
	q := s.storage.UpdateStmt(UserTable, state).Where(squirrel.Eq{s.colID: id})

	res, err := s.storage.Exec(ctx, q)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return database.ErrNotFound
	}

	return nil
}

// DeleteUser deletes the user data.
func (s *User) DeleteUser(ctx context.Context, id model.UserID) error {
	q := s.storage.DeleteStmt(UserTable).Where(squirrel.Eq{s.colID: id})

	res, err := s.storage.Exec(ctx, q)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return database.ErrNotFound
	}

	return nil
}
