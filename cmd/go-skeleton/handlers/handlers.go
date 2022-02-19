package handlers

import (
	"database/sql"
	"github.com/AlexRipoll/go-skeleton/cmd/go-skeleton/handlers/usergrp"
	"github.com/AlexRipoll/go-skeleton/internal/core/user"
	"github.com/AlexRipoll/go-skeleton/internal/core/user/db"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {

	mux := http.NewServeMux()

	// MySQL initialization
	sqlDb, err := sql.Open("mysql", "admin:root@tcp(localhost:3306)/sample-project")
	if err != nil {
		log.Fatalf("mysql setup error: %v", err.Error())
	}
	mysql := db.NewMySQL(sqlDb)

	ugrp := usergrp.Handler{
		User: user.NewCore(mysql),
	}

	mux.Handle("/users", http.HandlerFunc(ugrp.Create))


	log.Fatal(http.ListenAndServe(":8080", mux))
}