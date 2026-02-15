package main

import (
	"fmt"
	"log"
	"test-rest-api/internal/config"
	"test-rest-api/internal/database"
	"test-rest-api/internal/models"
	"test-rest-api/internal/server"
)

func main() {
	fmt.Println("Try starting")
	cfg := config.Load()
	fmt.Println("Loaded Configuration")

	db, err := database.NewPostgres(cfg.Database)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to PostgreSQL")

	err = db.AutoMigrate(models.User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.StartServer(db)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Started Server on port 8080")
}
