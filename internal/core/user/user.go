package user

import (
	"context"
	"github.com/AlexRipoll/go-skeleton/internal/core/user/db"
	"time"
)

type Core struct {
	db db.Repository
}

func NewCore(db db.Repository) Core {
	return Core{db: db}
}

func (c *Core) Create(ctx context.Context, nu NewUser) (*User, error) {
	now := time.Now()

	u := db.User{
		ID:           "uuid",
		Email:        nu.Email,
		PasswordHash: nil,
		Roles:        nil,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := c.db.Add(ctx, u); err != nil {
		return nil, err
	}

	// TODO type conversion from db.User to User
	uconv := User(u)

	return &uconv, nil
}