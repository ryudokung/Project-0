package game

import (
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"
)

type UseCase interface {
	InitializeNewCharacter(userID, charID uuid.UUID) error
	InitializeGachaStats(userID uuid.UUID) error
}

type gameUseCase struct {
	repo        Repository
	vehicleRepo vehicle.Repository
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
		VehicleType: vehicle.TypeShip,
		Class:       vehicle.ClassScout,
		Rarity:      vehicle.RarityCommon,
		Status:      vehicle.StatusPending,
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

	return nil
}

func (u *gameUseCase) InitializeGachaStats(userID uuid.UUID) error {
	return u.repo.InitializeGachaStats(userID)
}
