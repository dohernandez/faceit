package storage

import (
	"context"

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
func (s *User) AddUser(ctx context.Context, u model.UserState) (*model.User, error) {
	q := s.storage.InsertStmt(UserTable, u).Suffix("RETURNING *")

	var user model.User

	err := s.storage.Select(ctx, q, &user)
	if err == nil {
		return &user, nil
	}

	if pgx.IsUniqueViolation(err) {
		return nil, ctxd.LabeledError(database.ErrAlreadyExists, err)
	}

	return nil, err
}

// UpdateUser updates the user data.
func (s *User) UpdateUser(ctx context.Context, id model.UserID, info model.UserState) error {
	q := s.storage.UpdateStmt(UserTable, info).Where(squirrel.Eq{s.colID: id})

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
