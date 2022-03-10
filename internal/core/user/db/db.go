package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AlexRipoll/go-pod/internal/sys/database"
	"github.com/AlexRipoll/go-pod/internal/sys/errorFlag"
	"time"
)

const (
	timeFormat = time.RFC3339
)

// mysql manages the access to the database's methods.
type mysql struct {
	db *sql.DB
}

// NewMySQL initializes a new mysql struct to access to the database's methods.
func NewMySQL(db *sql.DB) *mysql {
	return &mysql{db: db}
}

// Insert inserts a new User instance to the database.
func (mysql *mysql) Insert(ctx context.Context, u User) error {

	usr, err := mysql.QueryByEmail(ctx, u.Email)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return err
	}
	if usr != nil {
		return errorFlag.New(ErrAlreadyExists, errorFlag.AlreadyExists)
	}

	const q = `
	INSERT INTO users
		(id, email, password_hash, roles, date_created, date_updated)
	VALUES
		(?, ?, ?, ?, ?, ?);`

	roles, err := json.Marshal(u.Roles)
	if err != nil {
		return errorFlag.New(fmt.Errorf("error json marshal: %w", err), errorFlag.Internal)
	}
	// TODO improve
	err = database.ExecStmt(ctx, mysql.db, q, u.ID, u.Email, string(u.PasswordHash), roles, u.DateCreated.Format(timeFormat), u.DateUpdated.Format(timeFormat))
	if err != nil {
		return errorFlag.New(err, errorFlag.Internal)
	}

	return nil
}

// QueryByEmail gets the specified user from the database.
func (mysql *mysql) QueryByEmail(ctx context.Context, email string) (*User, error) {
	const q = `
	SELECT
		id, email, password_hash, roles, date_created, date_updated
	FROM
		users
	WHERE 
		email = ?;`

	var user User
	err := database.ExecQueryStruct(ctx, mysql.db, &user, q, email)
	if err != nil {
		return nil, errorFlag.New(err, errorFlag.Internal)
	}

	return &user, nil
}
