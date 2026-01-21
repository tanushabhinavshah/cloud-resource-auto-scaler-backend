package main

import (
	"log"

	"cloud-resource-auto-scaler-backend/internal/config"
	"cloud-resource-auto-scaler-backend/internal/db"
	"cloud-resource-auto-scaler-backend/internal/router"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// For Feature 1, we just ensure we can connect.
	// Future features will use 'database' for migrations and queries.
	_ = database

	// Initialize router
	r := router.NewRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
