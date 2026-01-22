package metrics

import (
	"encoding/json"
	"fmt"
	"io"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// MetricsProvider is a consumer-defined interface.
// The handler only asks for what it specifically needs.
type MetricsProvider interface {
	GetBroadcast() chan *Metric
}

type Handler struct {
	service MetricsProvider
}

func NewHandler(service MetricsProvider) *Handler {
	return &Handler{service: service}
}

// StreamMetrics handles the Server-Sent Events connection
func (h *Handler) StreamMetrics(c *gin.Context) {
	// 1. Set SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// 2. Listen to the broadcast channel and send data to client
	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Request.Context().Done():
			// Client disconnected
			return false
		case metric := <-h.service.GetBroadcast():
			// Marshal metric to JSON
			data, err := json.Marshal(metric)
			if err != nil {
				return true // Continue loop, just skip this error
			}

			// Write in SSE format: "data: <json>\n\n"
			fmt.Fprintf(w, "data: %s\n\n", string(data))
			return true
		}
	})
}