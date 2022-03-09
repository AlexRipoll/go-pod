package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Config contains the required parameters to configure a database connection
type Config struct {
	Driver string
	Scheme string
	User   string
	Password string
	Protocol string
	Host string
	Port string
	Name   string
}

// Open opens a database connection with the given configuration parameters
func Open(cfg Config) (*sql.DB, error) {
	q := make(url.Values)
	q.Set("charset", "utf8")
	q.Set("parseTime", "true")
	q.Set("loc", "Local")

	u := url.URL{
		Scheme:      cfg.Scheme,
		User:        url.UserPassword(cfg.User, cfg.Password),
		Host:        fmt.Sprintf("%s(%s:%s)",cfg.Protocol, cfg.Host, cfg.Port),
		Path:        cfg.Name,
		RawQuery:    q.Encode(),
	}
	fmt.Println(strings.Trim(u.String(), "/"))

	db, err :=sql.Open(cfg.Driver, strings.Trim(u.String(), "/"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// ExecStmt is a helper function for executing queries that modify the current state
// pf the database.
func ExecStmt(ctx context.Context, db *sql.DB, query string, args ...interface{}) error {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

// ExecQueryStruct is a helper function for executing queries that return a single result
// of data to be unmarshalled into a slice.
func ExecQueryStruct(ctx context.Context, db *sql.DB, target interface{}, query string, args ...interface{}) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr && val.Elem().Kind() != reflect.Slice {
		return errors.New("must provide a pointer to a slice")
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)

	fields, err := argsScanner(stmt, target, args...)
	if err != nil {
		return err
	}

	if err := row.Scan(fields...); err != nil {
		return err
	}

	return nil
}

// ExecQueryCollection is a helper function for executing queries that return a collection
// of data to be unmarshalled into a slice.
func ExecQueryCollection(ctx context.Context, db *sql.DB, target interface{}, query string, args ...interface{}) error {
	val := reflect.ValueOf(target)
	if val.Kind() != reflect.Ptr && val.Elem().Kind() != reflect.Slice {
		return errors.New("must provide a pointer to a slice")
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		fields, err := argsScanner(stmt, target, args...)
		if err != nil {
			return err
		}

		if err := rows.Scan(fields...); err != nil {
			return err
		}
	}

	return nil
}

// argsScanner scans the struct and the query's fields and creates a slice of pointers containing
// only the required fields by the query.
func argsScanner(stmt *sql.Stmt, target interface{}, args ...interface{}) ([]interface{}, error) {
	n := reflect.TypeOf(target).Elem().NumField()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var fields []interface{}
	for i := 0; i < n; i++ {
		fieldAddr := reflect.ValueOf(target).Elem().Field(i).Addr().Interface()
		fieldTag := reflect.TypeOf(target).Elem().Field(i).Tag.Get("db")

		for _, col := range cols {
			if string(fieldTag) == col {
				fields = append(fields, fieldAddr)
				break
			}
		}
	}
	return fields, nil
}
