package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

// Config contains the required parameters to configure a database connection
type Config struct {
	Name string
	Scheme string
	User string
	Password string
	Host string
	Port string
}

// Open opens a database connection with the given configuration parameters
func Open(cfg Config) (*sql.DB, error) {
	q := make(url.Values)
	q.Set("charset", "utf8")
	q.Set("parseTime", "True")
	q.Set("loc", "Local")

	u := url.URL{
		Scheme:      cfg.Scheme,
		User:        url.UserPassword(cfg.User, cfg.Password),
		Host:        fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Path:        cfg.Name,
		RawQuery:    q.Encode(),
	}

	db, err :=sql.Open(cfg.Name, u.String())
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ExecStmt is a helper function for executing queries that modify the current state
// pf the database.
func ExecStmt(ctx context.Context, db *sql.DB, query string, args ...string) error {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}

	return nil
}

// ExecQuery is a helper function for executing queries that return a collection
// of data to be unmarshalled into a slice.
func ExecQuery(ctx context.Context, db *sql.DB, target interface{}, query string, args ...string) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return errors.New("must provide a pointer to a slice")
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args)
	if err != nil {
		return err
	}

	slice := val.Elem()
	for rows.Next() {
		v := reflect.New(slice.Type().Elem())
		//TODO autopopulate struct collection with reflection
		slice.Set(reflect.Append(slice, v.Elem()))
	}
	return nil
}
