package user

import (
	"context"
	"github.com/AlexRipoll/go-pod/internal/core/user/db"
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"github.com/AlexRipoll/go-pod/internal/sys/validate"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Core manages the access to the user's services.
type Core struct {
	db db.Repository
}

// NewCore initializes a new Core struct to access to the user's services.
func NewCore(db db.Repository) Core {
	return Core{db: db}
}

// Create creates a new user and inserts it to the database.
func (c *Core) Create(ctx context.Context, nu NewUser) (*User, error) {
	now := time.Now()

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errorFlag.New(err, errorFlag.Internal)
	}

	u := db.User{
		ID:           validate.GenerateID(),
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := c.db.Insert(ctx, u); err != nil {
		return nil, err
	}

	uconv := toUser(u)

	return &uconv, nil
}

func toUser(u db.User) User {
	return User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Roles:        u.Roles,
		DateCreated:  u.DateCreated,
		DateUpdated:  u.DateUpdated,
	}
}
