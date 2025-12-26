package vehicle

import (
	"time"

	"github.com/google/uuid"
)

type VehicleType string
type VehicleClass string
type RarityTier string
type VehicleStatus string
type ItemType string
type ItemCondition string

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

	StatusPending   VehicleStatus = "PENDING"
	StatusAvailable VehicleStatus = "AVAILABLE"
	StatusMinted    VehicleStatus = "MINTED"
	StatusBurned    VehicleStatus = "BURNED"

	ItemTypeVehicle       ItemType = "VEHICLE"
	ItemTypeExosuit       ItemType = "EXOSUIT"
	ItemTypePart          ItemType = "PART"
	ItemTypeBastionModule ItemType = "BASTION_MODULE"
	ItemTypeConsumable    ItemType = "CONSUMABLE"

	ConditionPristine ItemCondition = "PRISTINE"
	ConditionWorn     ItemCondition = "WORN"
	ConditionDamaged  ItemCondition = "DAMAGED"
	ConditionCritical ItemCondition = "CRITICAL"
	ConditionBroken   ItemCondition = "BROKEN"
)

type Item struct {
	ID            uuid.UUID     `json:"id"`
	OwnerID       uuid.UUID     `json:"owner_id"`
	CharacterID   *uuid.UUID    `json:"character_id,omitempty"`
	Name          string        `json:"name"`
	ItemType      ItemType      `json:"item_type"`
	Rarity        RarityTier    `json:"rarity"`
	Tier          int           `json:"tier"`
	Slot          *string       `json:"slot,omitempty"`
	DamageType    *string       `json:"damage_type,omitempty"` // KINETIC, ENERGY, VOID
	SeriesID      *string       `json:"series_id,omitempty"`   // For Set Synergy
	IsNFT         bool          `json:"is_nft"`
	TokenID       *string       `json:"token_id,omitempty"`
	Durability    int           `json:"durability"`
	MaxDurability int           `json:"max_durability"`
	Condition     ItemCondition `json:"condition"`
	Stats         ItemStats     `json:"stats"`
	VisualDNA     VisualDNA     `json:"visual_dna"`
	Metadata      interface{}   `json:"metadata"`
	IsEquipped    bool          `json:"is_equipped"`
	ParentItemID  *uuid.UUID    `json:"parent_item_id,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type ItemStats struct {
	HP               int `json:"hp,omitempty"`
	Attack           int `json:"attack,omitempty"`
	Defense          int `json:"defense,omitempty"`
	Speed            int `json:"speed,omitempty"`
	EnergyConsume    int `json:"energy_consume,omitempty"`
	BonusHP          int `json:"bonus_hp,omitempty"`
	BonusAttack      int `json:"bonus_attack,omitempty"`
	BonusDefense     int `json:"bonus_defense,omitempty"`
	ShieldGeneration int `json:"shield_generation,omitempty"`
	WarpRange        int `json:"warp_range,omitempty"`
}

type Vehicle struct {
	ID              uuid.UUID     `json:"id"`
	TokenID         *string       `json:"token_id,omitempty"` // Using string for uint256 compatibility
	OwnerID         uuid.UUID     `json:"owner_id"`
	CharacterID     *uuid.UUID    `json:"character_id,omitempty"`
	VehicleType     VehicleType   `json:"vehicle_type"`
	Class           VehicleClass  `json:"class"`
	ImageURL        *string       `json:"image_url,omitempty"`
	Stats           VehicleStats  `json:"stats"`
	CR              int           `json:"cr"`
	SuitabilityTags []string      `json:"suitability_tags"`
	Rarity          RarityTier    `json:"rarity"`
	Tier            int           `json:"tier"`
	IsVoidTouched   bool          `json:"is_void_touched"`
	Season          *string       `json:"season,omitempty"`
	Status          VehicleStatus `json:"status"`
	Metadata        interface{}   `json:"metadata"`
	CreatedAt       time.Time     `json:"created_at"`
}

type VehicleStats struct {
	HP      int `json:"hp"`
	Attack  int `json:"attack"`
	Defense int `json:"defense"`
	Speed   int `json:"speed"`
}

type Part struct {
	ID            uuid.UUID  `json:"id"`
	OwnerID       uuid.UUID  `json:"owner_id"`
	VehicleID     *uuid.UUID `json:"vehicle_id,omitempty"`
	CharacterID   *uuid.UUID `json:"character_id,omitempty"` // For Pilot Gear
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
	Keywords         []string `json:"keywords"`
	Faction          string   `json:"faction"`
	Style            string   `json:"style"`
	GlitchIntensity  float64  `json:"glitch_intensity"` // 0.0 - 1.0
	SmokeLevel       float64  `json:"smoke_level"`      // 0.0 - 1.0
	SparksEnabled    bool     `json:"sparks_enabled"`
}

type MintRequest struct {
	UserID      uuid.UUID    `json:"user_id"`
	VehicleType VehicleType  `json:"vehicle_type"`
	Class       VehicleClass `json:"class"`
}
