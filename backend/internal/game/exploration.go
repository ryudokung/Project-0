package game

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ExplorationStatus string

const (
	StatusInTransit ExplorationStatus = "IN_TRANSIT"
	StatusScanning  ExplorationStatus = "SCANNING"
	StatusEVA       ExplorationStatus = "EVA"
	StatusCompleted ExplorationStatus = "COMPLETED"
	StatusFailed    ExplorationStatus = "FAILED"
)

type ExplorationSession struct {
	ID           uuid.UUID         `json:"id"`
	PilotID      uuid.UUID         `json:"pilot_id"`
	VehicleID    uuid.UUID         `json:"vehicle_id"`
	Status       ExplorationStatus `json:"status"`
	Fuel         int               `json:"fuel"`
	ScannerRange int               `json:"scanner_range"`
	CurrentSector string            `json:"current_sector"`
	StartTime    time.Time         `json:"start_time"`
}

type ExplorationService struct {
	// Repositories
}

func NewExplorationService() *ExplorationService {
	return &ExplorationService{}
}

func (s *ExplorationService) StartSession(ctx context.Context, pilotID, vehicleID uuid.UUID) (*ExplorationSession, error) {
	return &ExplorationSession{
		ID:           uuid.New(),
		PilotID:      pilotID,
		VehicleID:    vehicleID,
		Status:       StatusInTransit,
		Fuel:         100,
		ScannerRange: 50,
		CurrentSector: "SOL-GATE",
		StartTime:    time.Now(),
	}, nil
}

func (s *ExplorationService) MoveToSector(ctx context.Context, session *ExplorationSession, targetSector string) error {
	fuelCost := 10 // Simplified
	if session.Fuel < fuelCost {
		return errors.New("insufficient fuel for transit")
	}

	session.Fuel -= fuelCost
	session.CurrentSector = targetSector
	return nil
}
