package db

import (
	"context"
	"database/sql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) Add(ctx context.Context, u User) error {
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