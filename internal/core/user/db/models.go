package db

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrAlreadyExists = errors.New("email address already taken")
)

// Repository is the interface that needs to be implemented by any database
// implementation to allow dependency abstraction with the user.Core struct.
type Repository interface {
	Insert(context.Context, User) error
}

type Roles []string

// User represents the user.User model to be used for database operations.
type User struct {
	ID           string    `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash []byte    `json:"-" db:"password_hash"`
	Roles        Roles     `json:"roles" db:"roles"`
	DateCreated  time.Time `json:"date_created" db:"date_created"`
	DateUpdated  time.Time `json:"date_updated" db:"date_updated"`
}

// Scan implements the Scanner interface to allow the database driver to
// parse a custom type.
func (h *Roles) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), h)
	if err != nil {
		return err
	}
	return nil
}
