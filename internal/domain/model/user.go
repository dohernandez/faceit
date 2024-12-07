package model

import (
	"github.com/google/uuid"
	"time"
)

// UserID represents the User id.
type UserID = uuid.UUID

// User represents the user.
type User struct {
	ID UserID `db:"id"` // User ID

	UserState

	CreatedAt time.Time `db:"created_at"` // Creation timestamp
	UpdatedAt time.Time `db:"updated_at"` // Last update timestamp
}

// UserState holds the user state.
type UserState struct {
	FirstName    string `db:"first_name"`    // First name of the user
	LastName     string `db:"last_name"`     // Last name of the user
	Nickname     string `db:"nickname"`      // Optional nickname
	PasswordHash string `db:"password_hash"` // Hashed password
	Email        string `db:"email"`         // Email address (unique)
	Country      string `db:"country"`       // 2-character country code
}
