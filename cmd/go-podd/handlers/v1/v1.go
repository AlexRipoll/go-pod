package v1

import (
	"database/sql"
	"fmt"
	"github.com/AlexRipoll/go-pod/cmd/go-podd/handlers/v1/usergrp"
	"github.com/AlexRipoll/go-pod/internal/core/user"
	"github.com/AlexRipoll/go-pod/internal/core/user/db"
	"github.com/AlexRipoll/go-pod/logger"
	"net/http"
)

const (
	version = "v1"
)

// Routes binds all the version 1 routes.
func Routes(mux *http.ServeMux, dbConn *sql.DB, log *logger.Logger)  {

	mysql := db.NewMySQL(dbConn)

	ugrp := usergrp.Handler{
		Log: log,
		User: user.NewCore(mysql),
	}

	// list of accessible routes
	// TODO add HTTP method validation
	mux.Handle(endpoint("/users"), http.HandlerFunc(ugrp.Create))

}

//endpoint builds the route including the version to it.
func endpoint(pattern string ) string {
	return fmt.Sprintf("/%s%s", version, pattern)
}
