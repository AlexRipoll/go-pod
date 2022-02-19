package db

import (
	"context"
	"time"
)

type Repository interface {
	Add(context.Context, User) error
}


type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Roles        []string  `json:"roles"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
}