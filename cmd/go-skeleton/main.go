package main

import (
	"github.com/AlexRipoll/go-skeleton/cmd/go-skeleton/handlers"
	"github.com/AlexRipoll/go-skeleton/internal/sys/database"
	"log"
)

func main() {

	cfg := database.Config{
		Driver:   "mysql",
		User:     "root",
		Password: "root",
		Protocol: "tcp",
		Host:     "127.0.0.1",
		Port:     "3306",
		Name:     "sample_project",
	}
	// MySQL initialization
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalf("mysql setup error: %v", err.Error())
	}
	log.Printf(">> database connection established")

	handlers.MuxServer(db)
}