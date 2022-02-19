package user

import "time"

// User models an individual user
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	Roles        []string  `json:"roles"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
}

// NewUser contains the needed information for creating a new user entity.
type NewUser struct {
	Email           string   `json:"email" validate:"required"`
	Password        string   `json:"password" validate:"required"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"`
	Roles           []string `json:"roles"`
}
