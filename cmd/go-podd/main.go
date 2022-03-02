package main

import (
	"github.com/AlexRipoll/go-pod/cmd/go-podd/handlers"
	"github.com/AlexRipoll/go-pod/internal/sys/database"
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
		Name:     "pod_db",
	}
	// MySQL initialization
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalf("mysql setup error: %v", err.Error())
	}
	log.Printf(">> database connection established")

	handlers.MuxServer(db)
}