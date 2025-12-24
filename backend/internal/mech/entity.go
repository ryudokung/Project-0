package mech

import (
	"time"

	"github.com/google/uuid"
)

type VehicleType string
type VehicleClass string
type RarityTier string
type MechStatus string

const (
	TypeMech VehicleType = "MECH"
	TypeTank VehicleType = "TANK"
	TypeShip VehicleType = "SHIP"

	ClassStriker   VehicleClass = "STRIKER"
	ClassGuardian  VehicleClass = "GUARDIAN"
	ClassScout     VehicleClass = "SCOUT"
	ClassArtillery VehicleClass = "ARTILLERY"

	RarityCommon      RarityTier = "COMMON"
	RarityRare        RarityTier = "RARE"
	RarityLegendary   RarityTier = "LEGENDARY"
	RarityRefined     RarityTier = "REFINED"
	RarityPrototype   RarityTier = "PROTOTYPE"
	RarityRelic       RarityTier = "RELIC"
	RaritySingularity RarityTier = "SINGULARITY"

	StatusPending MechStatus = "PENDING"
	StatusMinted  MechStatus = "MINTED"
	StatusBurned  MechStatus = "BURNED"
)

type Mech struct {
	ID            uuid.UUID    `json:"id"`
	TokenID       *string      `json:"token_id,omitempty"` // Using string for uint256 compatibility
	OwnerID       uuid.UUID    `json:"owner_id"`
	VehicleType   VehicleType  `json:"vehicle_type"`
	Class         VehicleClass `json:"class"`
	ImageURL      *string      `json:"image_url,omitempty"`
	Stats         MechStats    `json:"stats"`
	Rarity        RarityTier   `json:"rarity"`
	Tier          int          `json:"tier"`
	IsVoidTouched bool         `json:"is_void_touched"`
	Season        *string      `json:"season,omitempty"`
	Status        MechStatus   `json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
}

type MechStats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Speed   int `json:"speed"`
}

type Part struct {
	ID            uuid.UUID  `json:"id"`
	OwnerID       uuid.UUID  `json:"owner_id"`
	MechID        *uuid.UUID `json:"mech_id,omitempty"`
	Slot          string     `json:"slot"`
	Name          string     `json:"name"`
	Rarity        RarityTier `json:"rarity"`
	Tier          int        `json:"tier"`
	IsVoidTouched bool       `json:"is_void_touched"`
	IsMinted      bool       `json:"is_minted"`
	Stats         PartStats  `json:"stats"`
	VisualDNA     VisualDNA  `json:"visual_dna"`
	CreatedAt     time.Time  `json:"created_at"`
}

type PartStats struct {
	BonusHP      int `json:"bonus_hp"`
	BonusAttack  int `json:"bonus_attack"`
	BonusDefense int `json:"bonus_defense"`
}

type VisualDNA struct {
	Keywords []string `json:"keywords"`
	Faction  string   `json:"faction"`
	Style    string   `json:"style"`
}

type MintRequest struct {
	UserID      uuid.UUID    `json:"user_id"`
	VehicleType VehicleType  `json:"vehicle_type"`
	Class       VehicleClass `json:"class"`
}
