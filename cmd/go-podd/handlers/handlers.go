package handlers

import (
	"database/sql"
	v1 "github.com/AlexRipoll/go-pod/cmd/go-podd/handlers/v1"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// MuxServer constructs and runs a new http.Handler with all the accessible routes.
func MuxServer(dbConn *sql.DB) {

	mux := http.NewServeMux()

	// Load the v1 routes.
	v1.Routes(mux, dbConn)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
