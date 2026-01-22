package main

import (
	"context"
	"log"

	"cloud-resource-auto-scaler-backend/internal/config"
	"cloud-resource-auto-scaler-backend/internal/db"
	"cloud-resource-auto-scaler-backend/internal/metrics"
	"cloud-resource-auto-scaler-backend/internal/router"
	"cloud-resource-auto-scaler-backend/internal/simulation"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	database, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	simService := simulation.NewService()
	simHandler := simulation.NewHandler(simService)

	metricsService := metrics.NewService(database, simService)
	metricsHandler := metrics.NewHandler(metricsService)


	ctx := context.Background()
	go metricsService.StartCollectionLoop(ctx)


	r := router.NewRouter(metricsHandler, simHandler)

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}