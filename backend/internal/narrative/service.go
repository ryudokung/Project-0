package narrative

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	EventDiscovery EventType = "DISCOVERY"
	EventCombat    EventType = "COMBAT"
	EventAccident  EventType = "ACCIDENT"
	EventLore      EventType = "LORE"
)

type NarrativeEvent struct {
	ID          uuid.UUID `json:"id"`
	Type        EventType `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Impact      string    `json:"impact"` // e.g., "LOSE_FUEL", "GAIN_SCRAP"
	CreatedAt   time.Time `json:"created_at"`
}

type Service struct {
	// Repository for narrative events
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateEvent(ctx context.Context, pilotID uuid.UUID, context string) (*NarrativeEvent, error) {
	// This would ideally call the AI service to generate a context-aware event
	return &NarrativeEvent{
		ID:          uuid.New(),
		Type:        EventLore,
		Title:       "The Echo of Sector 7",
		Description: "You find a derelict beacon emitting a signal from the Old War. It speaks of a 'God Machine' buried in the ice.",
		Impact:      "NONE",
		CreatedAt:   time.Now(),
	}, nil
}
