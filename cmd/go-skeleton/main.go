package main

import (
	"github.com/AlexRipoll/go-skeleton/cmd/go-skeleton/handlers"
	"github.com/AlexRipoll/go-skeleton/internal/sys/database"
	"log"
)

func main() {

	cfg := database.Config{
		Name:     "mysql",
		Scheme:   "sample-project",
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     "3306",
	}
	// MySQL initialization
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalf("mysql setup error: %v", err.Error())
	}

	handlers.MuxServer(db)
}