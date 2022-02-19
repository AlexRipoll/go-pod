package db

import (
	"context"
	"time"
)

// Repository is the interface that needs to be implemented by any database
// implementation to allow dependency abstraction with the user.Core struct.
type Repository interface {
	Insert(context.Context, User) error
}

// User represents the user.User model to be used for database operations.
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Roles        []string  `json:"roles"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
}