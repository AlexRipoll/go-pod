package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AlexRipoll/go-pod/internal/sys/database"
)

// mysql manages the access to the database's methods.
type mysql struct {
	db *sql.DB
}

// NewMySQL initializes a new mysql struct to access to the database's methods.
func NewMySQL(db *sql.DB) *mysql {
	return &mysql{db: db}
}

// Insert inserts a new User instance to the database
func (mysql *mysql) Insert(ctx context.Context, u User) error {
	const q = `
	INSERT INTO users
		(id, email, password_hash, roles, date_created, date_updated)
	VALUES
		(?, ?, ?, ?, ?, ?);`

	roles, err := json.Marshal(u.Roles)
	if err != nil {
		return fmt.Errorf("error json marshal: %s", err.Error())
	}
	// TODO improve
	err = database.ExecStmt(ctx, mysql.db, q, u.ID, u.Email, string(u.PasswordHash),string(roles), u.DateCreated.String(), u.DateUpdated.String())
	if err != nil {
		return err
	}

	return nil
}