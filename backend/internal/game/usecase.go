package game

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

type UseCase interface {
	InitializeNewCharacter(userID, charID uuid.UUID) error
	InitializeGachaStats(userID uuid.UUID) error
	UnlockResearch(ctx context.Context, userID uuid.UUID, researchID string) (*PilotStats, error)
}

type gameUseCase struct {
	repo        Repository
	vehicleRepo vehicle.Repository
}

var ResearchCosts = map[string]int{
	"atmosphericEntry": 100,
	"quantumGate":      500,
	"miningDrill":      200,
	"hackingModule":    300,
}

func NewUseCase(repo Repository, vehicleRepo vehicle.Repository) UseCase {
	return &gameUseCase{
		repo:        repo,
		vehicleRepo: vehicleRepo,
	}
}

func (u *gameUseCase) InitializeNewCharacter(userID, charID uuid.UUID) error {
	// 1. Initialize Stats for the character
	if err := u.repo.InitializePilot(charID); err != nil {
		return err
	}

	// 2. Initialize Gacha Stats for the user (if not already)
	if err := u.repo.InitializeGachaStats(userID); err != nil {
		return err
	}

	// 3. Assign Starter Gear (Ship)
	starterShip := vehicle.Vehicle{
		ID:          uuid.New(),
		OwnerID:     userID,
		CharacterID: &charID,
		Name:        "Vanguard Scout",
		VehicleType: vehicle.TypeShip,
		Class:       vehicle.ClassScout,
		Rarity:      vehicle.RarityCommon,
		Status:      vehicle.StatusAvailable,
		Stats: vehicle.VehicleStats{
			HP:      100,
			Attack:  10,
			Defense: 10,
			Speed:   20,
		},
		CR:              50, // Initial CR for starter ship
		SuitabilityTags: []string{"ocean", "coastal"},
	}

	if err := u.vehicleRepo.Create(&starterShip); err != nil {
		return err
	}

	// 4. Assign Starter Items (Modules)
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
		item := vehicle.Item{
			ID:            uuid.New(),
			OwnerID:       userID,
			CharacterID:   &charID,
			Name:          si.name,
			ItemType:      vehicle.ItemTypePart,
			Rarity:        vehicle.RarityCommon,
			Tier:          1,
			Slot:          &slot,
			Durability:    100,
			MaxDurability: 100,
			Condition:     vehicle.ConditionPristine,
			Stats: vehicle.ItemStats{
				BonusAttack:  si.atk,
				BonusDefense: si.def,
			},
			IsEquipped:   true,
			ParentItemID: &starterShip.ID,
		}
		_ = u.vehicleRepo.CreateItem(context.Background(), &item)
	}

	return nil
}

func (uc *gameUseCase) UnlockResearch(ctx context.Context, userID uuid.UUID, researchID string) (*PilotStats, error) {
	stats, err := uc.repo.GetActivePilotStats(userID)
	if err != nil {
		return nil, err
	}

	if stats == nil {
		return nil, errors.New("pilot stats not found")
	}

	cost, ok := ResearchCosts[researchID]
	if !ok {
		return nil, errors.New("invalid research ID")
	}

	// Check if already unlocked
	unlocked, ok := stats.Metadata["unlocked_research"].([]interface{})
	if ok {
		for _, id := range unlocked {
			if id == researchID {
				return nil, errors.New("research already unlocked")
			}
		}
	} else {
		// Initialize if nil
		stats.Metadata["unlocked_research"] = []interface{}{}
	}

	if stats.ResearchData < cost {
		return nil, errors.New("insufficient research data")
	}

	// Deduct cost
	stats.ResearchData -= cost

	// Add to unlocked list
	currentUnlocked := stats.Metadata["unlocked_research"].([]interface{})
	stats.Metadata["unlocked_research"] = append(currentUnlocked, researchID)

	if err := uc.repo.UpdatePilotStats(stats); err != nil {
		return nil, err
	}

	return stats, nil
}

func (u *gameUseCase) InitializeGachaStats(userID uuid.UUID) error {
	return u.repo.InitializeGachaStats(userID)
}
