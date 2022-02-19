package database

import (
	"database/sql"
	"fmt"
	"net/url"
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
