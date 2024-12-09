package model

import (
	"time"

	"github.com/google/uuid"
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
	UserCredentials

	UserInfo
}

// UserCredentials holds the user credentials.
type UserCredentials struct {
	PasswordHash string `db:"password_hash,omitempty"` // Hashed password
	Email        string `db:"email,omitempty"`         // Email address (unique)
}

// UserInfo holds the user information.
type UserInfo struct {
	FirstName string `db:"first_name,omitempty"` // First name of the user
	LastName  string `db:"last_name,omitempty"`  // Last name of the user
	Nickname  string `db:"nickname"`             // Optional nickname
	Country   string `db:"country,omitempty"`    // 2-character country code
}
