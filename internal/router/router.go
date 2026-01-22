package router
import (
	"github.com/gin-gonic/gin"
	"cloud-resource-auto-scaler-backend/internal/health"
	"cloud-resource-auto-scaler-backend/internal/metrics"
	"cloud-resource-auto-scaler-backend/internal/simulation"
)
func NewRouter(metricsHandler *metrics.Handler, simHandler *simulation.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.SetTrustedProxies(nil)
	health.RegisterRoutes(r)
	// Metrics
	r.GET("/metrics/stream", metricsHandler.StreamMetrics)
	// Simulation
	simGroup := r.Group("/simulation")
	{
		simGroup.POST("/start", simHandler.Start)
		simGroup.POST("/stop", simHandler.Stop)
		simGroup.GET("/status", simHandler.Status)
	}
	return r
}