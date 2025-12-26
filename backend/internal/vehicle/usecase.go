package vehicle

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UseCase interface {
	InitializeStarterPack(ctx context.Context, userID uuid.UUID) (*Vehicle, error)
	GetUserVehicles(userID uuid.UUID) ([]Vehicle, error)
	GetCharacterVehicles(charID uuid.UUID) ([]Vehicle, error)
	GetVehicleByID(ctx context.Context, id uuid.UUID) (*Vehicle, error)

	// Item operations (DDS)
	ApplyDamage(ctx context.Context, itemID uuid.UUID, damage int) (*Item, error)
	RepairItem(ctx context.Context, itemID uuid.UUID, amount int) (*Item, error)
	GetItems(ctx context.Context, userID uuid.UUID) ([]Item, error)
	GetItemByID(ctx context.Context, itemID uuid.UUID) (*Item, error)
	GetVehicleCP(ctx context.Context, vehicleID uuid.UUID) (int, error)
	EquipItem(ctx context.Context, itemID uuid.UUID, vehicleID uuid.UUID) error
	UnequipItem(ctx context.Context, itemID uuid.UUID) error
}

type vehicleUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &vehicleUseCase{repo: repo}
}

func (u *vehicleUseCase) InitializeStarterPack(ctx context.Context, userID uuid.UUID) (*Vehicle, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Randomly pick a type and class for starter
	types := []VehicleType{TypeMech, TypeTank, TypeShip}
	classes := []VehicleClass{ClassStriker, ClassGuardian, ClassScout}
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	season := "Genesis"
	imageURL := "https://api.dicebear.com/7.x/bottts/svg?seed=" + uuid.New().String()

	v := &Vehicle{
		ID:          uuid.New(),
		OwnerID:     userID,
		VehicleType: types[r.Intn(len(types))],
		Class:       classes[r.Intn(len(classes))],
		Rarity:      RarityCommon,
		Season:      &season,
		Status:      StatusAvailable, // Manifested as an in-game asset
		Stats: VehicleStats{
			HP:      50 + r.Intn(50),
			Attack:  10 + r.Intn(10),
			Defense: 10 + r.Intn(10),
			Speed:   5 + r.Intn(15),
		},
		ImageURL: &imageURL, // Placeholder AI image
	}

	// Calculate initial CR
	v.CR = (v.Stats.Attack * 2) + (v.Stats.Defense * 2) + (v.Stats.HP / 10)
	
	// Add default suitability tags based on type
	switch v.VehicleType {
	case TypeMech:
		v.SuitabilityTags = []string{"urban", "forest"}
	case TypeTank:
		v.SuitabilityTags = []string{"desert", "plains"}
	case TypeShip:
		v.SuitabilityTags = []string{"ocean", "coastal"}
	}

	if err := u.repo.Create(v); err != nil {
		return nil, err
	}

	// Create Starter Items
	starterItems := []struct {
		name string
		slot string
		atk  int
		def  int
	}{
		{"Starter Kinetic Arm", "ARM_R", 5, 0},
		{"Starter Plating", "CORE", 0, 5},
	}

	for _, si := range starterItems {
		slot := si.slot
		item := &Item{
			ID:            uuid.New(),
			OwnerID:       userID,
			Name:          si.name,
			ItemType:      ItemTypePart,
			Rarity:        RarityCommon,
			Tier:          1,
			Slot:          &slot,
			Durability:    1000,
			MaxDurability: 1000,
			Condition:     ConditionPristine,
			Stats: ItemStats{
				BonusAttack:  si.atk,
				BonusDefense: si.def,
			},
			IsEquipped:   true,
			ParentItemID: &v.ID,
		}
		_ = u.repo.CreateItem(ctx, item)
	}

	return v, nil
}

func (u *vehicleUseCase) GetCharacterVehicles(charID uuid.UUID) ([]Vehicle, error) {
	return u.repo.GetByCharacterID(charID)
}

func (u *vehicleUseCase) GetUserVehicles(userID uuid.UUID) ([]Vehicle, error) {
	return u.repo.GetByOwnerID(userID)
}

func (u *vehicleUseCase) GetVehicleByID(ctx context.Context, id uuid.UUID) (*Vehicle, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return u.repo.GetByID(ctx, id)
}

// Item operations (DDS)

func (u *vehicleUseCase) GetItems(ctx context.Context, userID uuid.UUID) ([]Item, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return u.repo.GetItemsByOwnerID(ctx, userID)
}

func (u *vehicleUseCase) GetItemByID(ctx context.Context, itemID uuid.UUID) (*Item, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return u.repo.GetItemByID(ctx, itemID)
}

func (u *vehicleUseCase) GetVehicleCP(ctx context.Context, vehicleID uuid.UUID) (int, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// 1. Get Vehicle Base Stats
	v, err := u.repo.GetByID(ctx, vehicleID)
	if err != nil || v == nil {
		return 0, err
	}

	// 2. Get Equipped Parts (Items)
	// Note: We use the items table where parent_item_id = vehicleID
	items, err := u.repo.GetItemsByOwnerID(ctx, v.OwnerID)
	if err != nil {
		return 0, err
	}

	totalAttack := v.Stats.Attack
	totalDefense := v.Stats.Defense
	totalHP := v.Stats.HP

	for _, item := range items {
		if item.IsEquipped && item.ParentItemID != nil && *item.ParentItemID == vehicleID {
			totalAttack += item.Stats.Attack + item.Stats.BonusAttack
			totalDefense += item.Stats.Defense + item.Stats.BonusDefense
			totalHP += item.Stats.HP + item.Stats.BonusHP
		}
	}

	// CP Formula: (Attack * 3) + (Defense * 2) + (HP / 5)
	cp := (totalAttack * 3) + (totalDefense * 2) + (totalHP / 5)
	return cp, nil
}

func (u *vehicleUseCase) EquipItem(ctx context.Context, itemID uuid.UUID, vehicleID uuid.UUID) error {
	if ctx == nil {
		ctx = context.Background()
	}

	item, err := u.repo.GetItemByID(ctx, itemID)
	if err != nil || item == nil {
		return fmt.Errorf("item not found")
	}

	// Check if slot is already occupied on this vehicle
	if item.Slot != nil {
		items, _ := u.repo.GetItemsByOwnerID(ctx, item.OwnerID)
		for _, i := range items {
			if i.IsEquipped && i.ParentItemID != nil && *i.ParentItemID == vehicleID && i.Slot != nil && *i.Slot == *item.Slot {
				// Unequip the old item in this slot
				i.IsEquipped = false
				i.ParentItemID = nil
				_ = u.repo.UpdateItem(ctx, &i)
			}
		}
	}

	item.IsEquipped = true
	item.ParentItemID = &vehicleID
	return u.repo.UpdateItem(ctx, item)
}

func (u *vehicleUseCase) UnequipItem(ctx context.Context, itemID uuid.UUID) error {
	if ctx == nil {
		ctx = context.Background()
	}

	item, err := u.repo.GetItemByID(ctx, itemID)
	if err != nil || item == nil {
		return fmt.Errorf("item not found")
	}

	item.IsEquipped = false
	item.ParentItemID = nil
	return u.repo.UpdateItem(ctx, item)
}

func (u *vehicleUseCase) ApplyDamage(ctx context.Context, itemID uuid.UUID, damage int) (*Item, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	item, err := u.repo.GetItemByID(ctx, itemID)
	if err != nil || item == nil {
		return nil, err
	}

	item.Durability -= damage
	if item.Durability < 0 {
		item.Durability = 0
	}

	item.Condition = calculateCondition(item.Durability, item.MaxDurability)
	
	// Update Visual DNA based on condition
	updateVisualsByCondition(item)

	if err := u.repo.UpdateItem(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (u *vehicleUseCase) RepairItem(ctx context.Context, itemID uuid.UUID, amount int) (*Item, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	item, err := u.repo.GetItemByID(ctx, itemID)
	if err != nil || item == nil {
		return nil, err
	}

	item.Durability += amount
	if item.Durability > item.MaxDurability {
		item.Durability = item.MaxDurability
	}

	item.Condition = calculateCondition(item.Durability, item.MaxDurability)
	updateVisualsByCondition(item)

	if err := u.repo.UpdateItem(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func calculateCondition(durability, maxDurability int) ItemCondition {
	percentage := float64(durability) / float64(maxDurability) * 100

	switch {
	case percentage >= 80:
		return ConditionPristine
	case percentage >= 50:
		return ConditionWorn
	case percentage >= 20:
		return ConditionDamaged
	case percentage > 0:
		return ConditionCritical
	default:
		return ConditionBroken
	}
}

func updateVisualsByCondition(item *Item) {
	percentage := float64(item.Durability) / float64(item.MaxDurability) * 100

	if percentage < 50 {
		item.VisualDNA.SmokeLevel = (50 - percentage) / 50
	} else {
		item.VisualDNA.SmokeLevel = 0
	}

	if percentage < 20 {
		item.VisualDNA.GlitchIntensity = (20 - percentage) / 20
		item.VisualDNA.SparksEnabled = true
	} else {
		item.VisualDNA.GlitchIntensity = 0
		item.VisualDNA.SparksEnabled = false
	}
}
