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
	ResonanceGauge    float64                `json:"resonance_gauge"` // 0-100
	Stress            int                    `json:"stress"`
	XP                int                    `json:"xp"`
	SyncLevel         int                    `json:"sync_level"`
	CurrentO2         float64                `json:"current_o2"`
	CurrentFuel       float64                `json:"current_fuel"`
	CurrentNE         float64                `json:"current_ne"`
	MaxNE             float64                `json:"max_ne"`
	ExpeditionsCompleted int                 `json:"expeditions_completed"`
	CharacterAttributes  map[string]int      `json:"character_attributes"`
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

type BastionModule struct {
	ID          uuid.UUID              `json:"id"`
	UserID      uuid.UUID              `json:"user_id"`
	ModuleType  string                 `json:"module_type"` // RADAR, LAB, WARP_DRIVE
	Level       int                    `json:"level"`
	IsActive    bool                   `json:"is_active"`
	Metadata    map[string]interface{} `json:"metadata"`
	UnlockedAt  time.Time              `json:"unlocked_at"`
}

type ScriptEvent struct {
	Trigger  string `json:"trigger" yaml:"trigger"`
	Action   string `json:"action" yaml:"action"`
	Dialogue string `json:"dialogue" yaml:"dialogue"`
}
