package metrics
import (
	"context"
	"log"
	"time"
	"cloud-resource-auto-scaler-backend/internal/simulation"
	"gorm.io/gorm"
)
type Service struct {
	db           *gorm.DB
	broadcast    chan *Metric
	simService   simulation.Service // Added simulation service
	lastNetBytes uint64
}
// Update constructor to accept simService
func NewService(db *gorm.DB, simService simulation.Service) *Service {
	return &Service{
		db:         db,
		broadcast:  make(chan *Metric, 10),
		simService: simService,
	}
}
func (s *Service) GetBroadcast() chan *Metric {
	return s.broadcast
}
func (s *Service) StartCollectionLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var metric *Metric
			var currentBytes uint64
			var err error
			// 1. Check if simulation is active
			simConfig := s.simService.GetConfig()
			if simConfig.IsActive {
				// Use fake data
				metric = &Metric{
					CPUUsagePercent:    simConfig.CPUUsagePercent,
					MemoryUsagePercent: simConfig.MemoryUsagePercent,
					NetworkUsageKbps:   simConfig.NetworkUsageKbps,
					IsOutlier:          false,
					CreatedAt:          time.Now(),
				}
				// We don't update lastNetBytes during simulation to avoid jumps
			} else {
				// Use real data
				metric, currentBytes, err = Collect(s.lastNetBytes)
				if err != nil {
					log.Printf("Collection error: %v", err)
					continue
				}
				s.lastNetBytes = currentBytes
			}
			// 2. Save and Broadcast (Remaining logic is the same)
			if err := s.db.Create(metric).Error; err != nil {
				log.Printf("DB Save error: %v", err)
			}
			select {
			case s.broadcast <- metric:
			default:
			}
		}
	}
}