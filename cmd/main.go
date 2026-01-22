package main

import (
	"context"
	"log"

	"cloud-resource-auto-scaler-backend/internal/config"
	"cloud-resource-auto-scaler-backend/internal/db"
	"cloud-resource-auto-scaler-backend/internal/metrics"
	"cloud-resource-auto-scaler-backend/internal/router"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize database
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	// 3. Initialize Metrics Service & Handler
	metricsService := metrics.NewService(database)
	metricsHandler := metrics.NewHandler(metricsService)

	// 4. Start the background metrics collection loop (in a goroutine)
	ctx := context.Background()
	go metricsService.StartCollectionLoop(ctx)

	// 5. Initialize router with the metrics handler
	r := router.NewRouter(metricsHandler)

	// 6. Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}