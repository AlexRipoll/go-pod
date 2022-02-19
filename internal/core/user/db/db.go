package db

import (
	"context"
	"database/sql"
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

	stmt, err := mysql.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}