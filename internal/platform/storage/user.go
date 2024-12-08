package storage

import (
	"context"
	"fmt"
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
	colCountry string
}

// NewUser returns instance of User repository.
func NewUser(storage *sqluct.Storage) *User {
	var user model.User

	return &User{
		storage:    storage,
		colCountry: storage.Mapper.Col(&user, &user.Country),
	}
}

// AddUser store the user data.
func (s *User) AddUser(ctx context.Context, u model.UserState) (*model.User, error) {
	errMsg := "storage.User: add user"

	q := s.storage.InsertStmt(UserTable, u).Suffix("RETURNING *")

	var user model.User

	err := s.storage.Select(ctx, q, &user)
	if err == nil {
		return &user, nil
	}

	if pgx.IsUniqueViolation(err) {
		return nil, fmt.Errorf("%s: %w: %w", errMsg, database.ErrAlreadyExists, err)
	}

	return nil, fmt.Errorf("%s: %w", errMsg, err)
}
