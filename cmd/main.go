package main

import (
	"log"

	"github.com/gabereiser/blog/data/database"
	"github.com/gabereiser/blog/service"
)

func main() {
	database.Init(false) // Initialize the database connection
	svc := service.NewWebService()
	svc.RegisterRoutes() // Register routes

	defer func() {
		if err := svc.Stop(); err != nil {
			log.Fatalf("Failed to stop service: %v", err)
		} else {
			log.Println("Service stopped gracefully")
		}
	}()

	log.Fatal(svc.Start()) // Start the service

	database.Close() // Close the database connection
}
