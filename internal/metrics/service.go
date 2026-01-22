package metrics

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	db           *gorm.DB
	broadcast    chan *Metric
	lastNetBytes uint64
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:        db,
		broadcast: make(chan *Metric, 10),
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
			metric, currentBytes, err := Collect(s.lastNetBytes)
			if err != nil {
				log.Printf("Collection error: %v", err)
				continue
			}
			s.lastNetBytes = currentBytes

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