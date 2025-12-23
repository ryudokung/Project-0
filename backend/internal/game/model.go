package game

import (
	"time"
	"github.com/google/uuid"
)

type PilotStats struct {
	UserID         uuid.UUID `json:"user_id"`
	ResonanceLevel int       `json:"resonance_level"`
	ResonanceExp   int       `json:"resonance_exp"`
	CurrentO2      float64   `json:"current_o2"`
	CurrentFuel    float64   `json:"current_fuel"`
	UpdatedAt      time.Time `json:"updated_at"`
}
