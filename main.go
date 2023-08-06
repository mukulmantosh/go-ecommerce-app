package main

import (
	"github.com/mukulmantosh/go-ecommerce-app/internal/database"
	"github.com/mukulmantosh/go-ecommerce-app/internal/server"
	"log"
)

func main() {
	db, err := database.NewDBClient()
	if err != nil {
		log.Fatalf("Failed DB Startup: %s\n", err)
		return
	}

	err = db.RunMigration()
	if err != nil {
		log.Fatalf("Migration Failed: %s\n", err)
		return
	}

	service := server.NewServer(db)
	log.Fatal(service.Start())

}
