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

	RarityCommon    RarityTier = "COMMON"
	RarityRare      RarityTier = "RARE"
	RarityLegendary RarityTier = "LEGENDARY"

	StatusPending MechStatus = "PENDING"
	StatusMinted  MechStatus = "MINTED"
	StatusBurned  MechStatus = "BURNED"
)

type Mech struct {
	ID          uuid.UUID    `json:"id"`
	TokenID     *string      `json:"token_id,omitempty"` // Using string for uint256 compatibility
	OwnerID     uuid.UUID    `json:"owner_id"`
	VehicleType VehicleType  `json:"vehicle_type"`
	Class       VehicleClass `json:"class"`
	ImageURL    string       `json:"image_url"`
	Stats       MechStats    `json:"stats"`
	Rarity      RarityTier   `json:"rarity"`
	Season      string       `json:"season"`
	Status      MechStatus   `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
}

type MechStats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Speed   int `json:"speed"`
}

type MintRequest struct {
	UserID      uuid.UUID    `json:"user_id"`
	VehicleType VehicleType  `json:"vehicle_type"`
	Class       VehicleClass `json:"class"`
}
