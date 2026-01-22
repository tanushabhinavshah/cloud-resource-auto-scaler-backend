package simulation
import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	service Service
}
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
// Start handles POST /simulation/start
func (h *Handler) Start(c *gin.Context) {
	var config SimulationConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid simulation config"})
		return
	}
	h.service.SetConfig(config)
	c.JSON(http.StatusOK, gin.H{"status": "Simulation started", "config": config})
}
// Stop handles POST /simulation/stop
func (h *Handler) Stop(c *gin.Context) {
	h.service.Stop()
	c.JSON(http.StatusOK, gin.H{"status": "Simulation stopped"})
}
// Status handles GET /simulation/status
func (h *Handler) Status(c *gin.Context) {
	config := h.service.GetConfig()
	c.JSON(http.StatusOK, config)
}