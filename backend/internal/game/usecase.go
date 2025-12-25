package game

import (
	"github.com/google/uuid"
	"github.com/ryudokung/Project-0/backend/internal/mech"
)

type UseCase interface {
	InitializeNewCharacter(userID, charID uuid.UUID) error
	InitializeGachaStats(userID uuid.UUID) error
}

type gameUseCase struct {
	repo     Repository
	mechRepo mech.Repository
}

func NewUseCase(repo Repository, mechRepo mech.Repository) UseCase {
	return &gameUseCase{
		repo:     repo,
		mechRepo: mechRepo,
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
	starterShip := mech.Mech{
		ID:          uuid.New(),
		OwnerID:     userID,
		CharacterID: &charID,
		VehicleType: mech.TypeShip,
		Class:       mech.ClassScout,
		Rarity:      mech.RarityCommon,
		Status:      mech.StatusPending,
		Stats: mech.MechStats{
			HP:      100,
			Attack:  10,
			Defense: 10,
			Speed:   20,
		},
	}

	if err := u.mechRepo.Create(&starterShip); err != nil {
		return err
	}

	return nil
}

func (u *gameUseCase) InitializeGachaStats(userID uuid.UUID) error {
	return u.repo.InitializeGachaStats(userID)
}
