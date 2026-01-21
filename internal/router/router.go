package router

import (
	"github.com/gin-gonic/gin"

	"cloud-resource-auto-scaler-backend/internal/health"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Register health routes
	health.RegisterRoutes(r)

	return r
}
