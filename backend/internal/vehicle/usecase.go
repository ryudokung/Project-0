package vehicle

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UseCase interface {
	MintStarterVehicle(userID uuid.UUID) (*Vehicle, error)
	GetUserVehicles(userID uuid.UUID) ([]Vehicle, error)
	GetCharacterVehicles(charID uuid.UUID) ([]Vehicle, error)
	GetVehicleByID(ctx context.Context, id uuid.UUID) (*Vehicle, error)

	// Item operations (DDS)
	ApplyDamage(ctx context.Context, itemID uuid.UUID, damage int) (*Item, error)
	RepairItem(ctx context.Context, itemID uuid.UUID, amount int) (*Item, error)
	GetItems(ctx context.Context, userID uuid.UUID) ([]Item, error)
	GetItemByID(ctx context.Context, itemID uuid.UUID) (*Item, error)
}

type vehicleUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &vehicleUseCase{repo: repo}
}

func (u *vehicleUseCase) MintStarterVehicle(userID uuid.UUID) (*Vehicle, error) {
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
		Status:      StatusMinted, // For starter, we mark as minted immediately in DB
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
