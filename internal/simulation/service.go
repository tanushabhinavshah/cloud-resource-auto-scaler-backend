package simulation
import (
	"sync"
)
type Service interface {
	GetConfig() SimulationConfig
	SetConfig(config SimulationConfig)
	Stop()
}
type simulationService struct {
	mu     sync.RWMutex
	config SimulationConfig
}
func NewService() Service {
	return &simulationService{
		config: SimulationConfig{IsActive: false},
	}
}
func (s *simulationService) GetConfig() SimulationConfig {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.config
}
func (s *simulationService) SetConfig(config SimulationConfig) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.config = config
	s.config.IsActive = true
}
func (s *simulationService) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.config.IsActive = false
}