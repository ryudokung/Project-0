package gacha

import (
	"time"
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

type PullType string

const (
	StandardSignal PullType = "STANDARD_SIGNAL"
	VoidSignal     PullType = "VOID_SIGNAL"
)

type GachaPullRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	PullType PullType  `json:"pull_type"`
	Count    int       `json:"count"` // 1 or 10
}

type GachaPullResponse struct {
	Results []GachaResult `json:"results"`
}

type GachaResult struct {
	Item      interface{}     `json:"item"` // Can be mech.Mech or mech.Part
	ItemType  string          `json:"item_type"`
	Rarity    mech.RarityTier `json:"rarity"`
	IsNew     bool            `json:"is_new"`
	PityCount int             `json:"pity_count"`
}

type GachaHistory struct {
	ID        uuid.UUID       `json:"id"`
	UserID    uuid.UUID       `json:"user_id"`
	ItemID    uuid.UUID       `json:"item_id"`
	ItemType  string          `json:"item_type"`
	PullType  PullType        `json:"pull_type"`
	Rarity    mech.RarityTier `json:"rarity"`
	CreatedAt time.Time       `json:"created_at"`
}
