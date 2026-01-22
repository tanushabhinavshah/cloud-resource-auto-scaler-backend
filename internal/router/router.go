package router

import (
	"github.com/gin-gonic/gin"

	"cloud-resource-auto-scaler-backend/internal/health"
	"cloud-resource-auto-scaler-backend/internal/metrics"
)

func NewRouter(metricsHandler *metrics.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.SetTrustedProxies(nil)

	// Health check
	health.RegisterRoutes(r)

	// Metrics routes
	r.GET("/metrics/stream", metricsHandler.StreamMetrics)

	return r
}