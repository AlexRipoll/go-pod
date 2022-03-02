package handlers

import (
	"database/sql"
	"github.com/AlexRipoll/go-pod/cmd/go-skeleton/handlers/usergrp"
	"github.com/AlexRipoll/go-pod/internal/core/user"
	"github.com/AlexRipoll/go-pod/internal/core/user/db"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// MuxServer constructs and runs a new http.Handler with all the accessible routes.
func MuxServer(dbConn *sql.DB) {

	mux := http.NewServeMux()

	mysql := db.NewMySQL(dbConn)

	ugrp := usergrp.Handler{
		User: user.NewCore(mysql),
	}

	mux.Handle("/users", http.HandlerFunc(ugrp.Create))


	log.Fatal(http.ListenAndServe(":8080", mux))
}