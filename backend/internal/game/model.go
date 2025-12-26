package game

import (
	"time"
	"github.com/google/uuid"
)

type PilotStats struct {
	UserID            uuid.UUID              `json:"user_id"`
	CharacterID       uuid.UUID              `json:"character_id"`
	EquippedExosuitID *uuid.UUID             `json:"equipped_exosuit_id,omitempty"`
	ResonanceLevel    int                    `json:"resonance_level"`
	ResonanceExp      int                    `json:"resonance_exp"`
	Stress            int                    `json:"stress"`
	XP                int                    `json:"xp"`
	Rank              int                    `json:"rank"`
	CurrentO2         float64                `json:"current_o2"`
	CurrentFuel       float64                `json:"current_fuel"`
	ScrapMetal        int                    `json:"scrap_metal"`
	ResearchData      int                    `json:"research_data"`
	Metadata          map[string]interface{} `json:"metadata"`
	UpdatedAt         time.Time              `json:"updated_at"`
}

type GachaStats struct {
	UserID               uuid.UUID `json:"user_id"`
	PityRelicCount       int       `json:"pity_relic_count"`
	PitySingularityCount int       `json:"pity_singularity_count"`
	TotalPulls           int       `json:"total_pulls"`
	LastFreePullAt       *time.Time `json:"last_free_pull_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
